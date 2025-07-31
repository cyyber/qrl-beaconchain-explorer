package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/theQRL/qrl-beaconchain-explorer/db"
	"github.com/theQRL/qrl-beaconchain-explorer/eth1data"
	"github.com/theQRL/qrl-beaconchain-explorer/templates"
	"github.com/theQRL/qrl-beaconchain-explorer/types"
	"github.com/theQRL/qrl-beaconchain-explorer/utils"

	"github.com/gorilla/mux"
	"github.com/theQRL/go-zond/common"
	"golang.org/x/sync/errgroup"
)

func Eth1Address(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "execution/address.html")
	var eth1AddressTemplate = templates.GetTemplate(templateFiles...)

	w.Header().Set("Content-Type", "text/html")
	vars := mux.Vars(r)
	address := template.HTMLEscapeString(vars["address"])
	// TODO(now.youtrack.cloud/issue/TZB-1)
	// znsData, err := GetZnsDomain(address)
	// if err != nil && utils.IsValidZnsDomain(address) {
	// 	handleNotFoundHtml(w, r)
	// 	return
	// }
	// if len(znsData.Address) > 0 {
	// 	address = znsData.Address
	// }

	isValid := utils.IsAddress(address)
	if !isValid {
		handleNotFoundHtml(w, r)
		return
	}

	addressHex := strings.Replace(address, "Z", "", -1)
	addressHex = strings.ToLower(addressHex)

	addressBytes := common.FromHex(addressHex)
	data := InitPageData(w, r, "blockchain", "/address", fmt.Sprintf("Address Z%x", addressBytes), templateFiles)

	metadata, err := db.BigtableClient.GetMetadataForAddress(addressBytes, 0, db.ZRC20TokensPerAddressLimit)
	if err != nil {
		logger.Errorf("error retrieving balances for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	g := new(errgroup.Group)
	g.SetLimit(11)

	isContract := false
	txns := &types.DataTableResponse{}
	internal := &types.DataTableResponse{}
	zrc20 := &types.DataTableResponse{}
	zrc721 := &types.DataTableResponse{}
	zrc1155 := &types.DataTableResponse{}
	blocksMined := &types.DataTableResponse{}
	withdrawals := &types.DataTableResponse{}
	withdrawalSummary := template.HTML("0")

	g.Go(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var err error
		isContract, err = eth1data.IsContract(ctx, common.BytesToAddress(addressBytes))
		if err != nil {
			return fmt.Errorf("IsContract: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		txns, err = db.BigtableClient.GetAddressTransactionsTableData(addressBytes, "")
		if err != nil {
			return fmt.Errorf("GetAddressTransactionsTableData: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		internal, err = db.BigtableClient.GetAddressInternalTableData(addressBytes, "")
		if err != nil {
			return fmt.Errorf("GetAddressInternalTableData: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		zrc20, err = db.BigtableClient.GetAddressZrc20TableData(addressBytes, "")
		if err != nil {
			return fmt.Errorf("GetAddressZrc20TableData: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		zrc721, err = db.BigtableClient.GetAddressZrc721TableData(addressBytes, "")
		if err != nil {
			return fmt.Errorf("GetAddressZrc721TableData: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		zrc1155, err = db.BigtableClient.GetAddressZrc1155TableData(addressBytes, "")
		if err != nil {
			return fmt.Errorf("GetAddressZrc1155TableData: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		blocksMined, err = db.BigtableClient.GetAddressBlocksMinedTableData(address, "")
		if err != nil {
			return fmt.Errorf("GetAddressBlocksMinedTableData: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		withdrawals, err = db.GetAddressWithdrawalTableData(addressBytes, "")
		if err != nil {
			return fmt.Errorf("GetAddressWithdrawalTableData: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		sumWithdrawals, err := db.GetAddressWithdrawalsTotal(addressBytes)
		if err != nil {
			return fmt.Errorf("GetAddressWithdrawalsTotal: %w", err)
		}
		withdrawalSummary = utils.FormatClCurrency(sumWithdrawals, 6, true, false, false, true)
		return nil
	})

	if err = g.Wait(); err != nil {
		if handleTemplateError(w, r, "eth1Account.go", "Eth1Address", "g.Wait()", err) != nil {
			return // an error has occurred and was processed
		}
		return
	}

	pngStr, pngStrInverse, err := utils.GenerateQRCodeForAddress(addressBytes)
	if err != nil {
		logger.WithError(err).Errorf("error generating qr code for address %v", address)
	}

	tabs := []types.Eth1AddressPageTabs{}

	if internal != nil && len(internal.Data) != 0 {
		tabs = append(tabs, types.Eth1AddressPageTabs{
			Id:   "internalTxns",
			Href: "#internalTxns",
			Text: "Internal Txns",
			Data: internal,
		})
	}
	if zrc20 != nil && len(zrc20.Data) != 0 {
		tabs = append(tabs, types.Eth1AddressPageTabs{
			Id:   "zrc20Txns",
			Href: "#zrc20Txns",
			Text: "Zrc20 Token Txns",
			Data: zrc20,
		})
	}
	if zrc721 != nil && len(zrc721.Data) != 0 {
		tabs = append(tabs, types.Eth1AddressPageTabs{
			Id:   "zrc721Txns",
			Href: "#zrc721Txns",
			Text: "Zrc721 Token Txns",
			Data: zrc721,
		})
	}
	if blocksMined != nil && len(blocksMined.Data) != 0 {
		tabs = append(tabs, types.Eth1AddressPageTabs{
			Id:   "blocks",
			Href: "#blocks",
			Text: "Produced Blocks",
			Data: blocksMined,
		})
	}
	if zrc1155 != nil && len(zrc1155.Data) != 0 {
		tabs = append(tabs, types.Eth1AddressPageTabs{
			Id:   "zrc1155Txns",
			Href: "#zrc1155Txns",
			Text: "Zrc1155 Token Txns",
			Data: zrc1155,
		})
	}
	if withdrawals != nil && len(withdrawals.Data) != 0 {
		tabs = append(tabs, types.Eth1AddressPageTabs{
			Id:   "withdrawals",
			Href: "#withdrawals",
			Text: "Withdrawals",
			Data: withdrawals,
		})
	}

	data.Data = types.Eth1AddressPageData{
		Address: address,
		// TODO(now.youtrack.cloud/issue/TZB-1)
		// ZnsName:            znsData.Domain,
		IsContract:         isContract,
		QRCode:             pngStr,
		QRCodeInverse:      pngStrInverse,
		Metadata:           metadata,
		WithdrawalsSummary: withdrawalSummary,
		TransactionsTable:  txns,
		InternalTxnsTable:  internal,
		Zrc20Table:         zrc20,
		Zrc721Table:        zrc721,
		Zrc1155Table:       zrc1155,
		WithdrawalsTable:   withdrawals,
		BlocksMinedTable:   blocksMined,
		Tabs:               tabs,
	}

	if handleTemplateError(w, r, "eth1Account.go", "Eth1Address", "Done", eth1AddressTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}

func Eth1AddressTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	address, err := lowerAddressFromRequest(w, r)
	if err != nil {
		return
	}
	addressBytes := common.FromHex(address)

	errFields := map[string]interface{}{
		"route": r.URL.String()}

	pageToken := q.Get("pageToken")

	data, err := db.BigtableClient.GetAddressTransactionsTableData(addressBytes, pageToken)
	if err != nil {
		utils.LogError(err, "error getting eth1 tx table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func Eth1AddressBlocksMined(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	address, err := lowerAddressFromRequest(w, r)
	if err != nil {
		return
	}

	errFields := map[string]interface{}{
		"route": r.URL.String()}

	pageToken := q.Get("pageToken")

	data, err := db.BigtableClient.GetAddressBlocksMinedTableData(address, pageToken)
	if err != nil {
		utils.LogError(err, "error getting eth1 blocks mined table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func Eth1AddressWithdrawals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	address, err := lowerAddressFromRequest(w, r)
	if err != nil {
		return
	}

	errFields := map[string]interface{}{
		"route": r.URL.String()}

	addr, err := common.NewAddressFromString(address)
	if err != nil {
		return
	}
	data, err := db.GetAddressWithdrawalTableData(addr.Bytes(), q.Get("pageToken"))
	if err != nil {
		utils.LogError(err, "error getting address withdrawals data", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func Eth1AddressInternalTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	address, err := lowerAddressFromRequest(w, r)
	if err != nil {
		return
	}
	addressBytes := common.FromHex(address)

	errFields := map[string]interface{}{
		"route": r.URL.String()}

	pageToken := q.Get("pageToken")
	data, err := db.BigtableClient.GetAddressInternalTableData(addressBytes, pageToken)
	if err != nil {
		utils.LogError(err, "error getting eth1 internal tx table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func Eth1AddressZrc20Transactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	address, err := lowerAddressFromRequest(w, r)
	if err != nil {
		return
	}
	addressBytes := common.FromHex(address)

	errFields := map[string]interface{}{
		"route": r.URL.String()}

	pageToken := q.Get("pageToken")
	data, err := db.BigtableClient.GetAddressZrc20TableData(addressBytes, pageToken)
	if err != nil {
		utils.LogError(err, "error getting eth1 ZRC20 transactions table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func Eth1AddressZrc721Transactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	address, err := lowerAddressFromRequest(w, r)
	if err != nil {
		return
	}
	addressBytes := common.FromHex(address)

	errFields := map[string]interface{}{
		"route": r.URL.String()}

	pageToken := q.Get("pageToken")
	data, err := db.BigtableClient.GetAddressZrc721TableData(addressBytes, pageToken)
	if err != nil {
		utils.LogError(err, "error getting eth1 ZRC721 transactions table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func Eth1AddressZrc1155Transactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	address, err := lowerAddressFromRequest(w, r)
	if err != nil {
		return
	}
	addressBytes := common.FromHex(address)

	errFields := map[string]interface{}{
		"route": r.URL.String()}

	pageToken := q.Get("pageToken")
	data, err := db.BigtableClient.GetAddressZrc1155TableData(addressBytes, pageToken)
	if err != nil {
		utils.LogError(err, "error getting eth1 ZRC1155 transactions table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// takes the "address" parameter from the request and transforms it to lower case. The ZNS name can be used instead of the address
func lowerAddressFromRequest(w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	address := vars["address"]
	// if utils.IsValidZnsDomain(address) {
	// 	znsData, err := GetZnsDomain(address)
	// 	if err != nil {
	// 		handleNotFoundJson(address, w, r, err)
	// 		return "", err
	// 	}
	// 	if len(znsData.Address) > 0 {
	// 		address = znsData.Address
	// 	}
	// }
	return strings.ToLower(strings.Replace(address, "Z", "", -1)), nil
}

// TODO(now.youtrack.cloud/issue/TZB-1)
/*
func handleNotFoundJson(address string, w http.ResponseWriter, r *http.Request, err error) {
	logger.Errorf("error getting address for ZNS name [%v] not found for %v route: %v", address, r.URL.String(), err)
	http.Error(w, "Invalid ZNS name", http.StatusInternalServerError)
}
*/

func handleNotFoundHtml(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "execution/addressNotFound.html")
	data := InitPageData(w, r, "blockchain", "/address", "not found", templateFiles)

	if handleTemplateError(w, r, "eth1Account.go", "Eth1Address", "not valid", templates.GetTemplate(templateFiles...).ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}
