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
	"github.com/theQRL/qrl-beaconchain-explorer/executiondata"
	"github.com/theQRL/qrl-beaconchain-explorer/templates"
	"github.com/theQRL/qrl-beaconchain-explorer/types"
	"github.com/theQRL/qrl-beaconchain-explorer/utils"

	"github.com/gorilla/mux"
	"github.com/theQRL/go-zond/common"
	"golang.org/x/sync/errgroup"
)

func ExecutionAddress(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "execution/address.html")
	var executionAddressTemplate = templates.GetTemplate(templateFiles...)

	w.Header().Set("Content-Type", "text/html")
	vars := mux.Vars(r)
	address := template.HTMLEscapeString(vars["address"])
	// TODO(now.youtrack.cloud/issue/TZB-1)
	// qrnsData, err := GetQrnsDomain(address)
	// if err != nil && utils.IsValidQrnsDomain(address) {
	// 	handleNotFoundHtml(w, r)
	// 	return
	// }
	// if len(qrnsData.Address) > 0 {
	// 	address = qrnsData.Address
	// }

	isValid := utils.IsAddress(address)
	if !isValid {
		handleNotFoundHtml(w, r)
		return
	}

	addressHex := strings.Replace(address, "Q", "", -1)
	addressHex = strings.ToLower(addressHex)

	addressBytes := common.FromHex(addressHex)
	data := InitPageData(w, r, "blockchain", "/address", fmt.Sprintf("Address Q%x", addressBytes), templateFiles)

	metadata, err := db.BigtableClient.GetMetadataForAddress(addressBytes, 0, db.SQRCTF1TokensPerAddressLimit)
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
	sqrctf1 := &types.DataTableResponse{}
	sqrctn1 := &types.DataTableResponse{}
	sqrctb1 := &types.DataTableResponse{}
	blocksMined := &types.DataTableResponse{}
	withdrawals := &types.DataTableResponse{}
	withdrawalSummary := template.HTML("0")

	g.Go(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var err error
		isContract, err = executiondata.IsContract(ctx, common.BytesToAddress(addressBytes))
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
		sqrctf1, err = db.BigtableClient.GetAddressSqrcTf1TableData(addressBytes, "")
		if err != nil {
			return fmt.Errorf("GetAddressSqrcTf1TableData: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		sqrctn1, err = db.BigtableClient.GetAddressSqrcTn1TableData(addressBytes, "")
		if err != nil {
			return fmt.Errorf("GetAddressSqrcTn1TableData: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		sqrctb1, err = db.BigtableClient.GetAddressSqrcTb1TableData(addressBytes, "")
		if err != nil {
			return fmt.Errorf("GetAddressSqrcTb1TableData: %w", err)
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
		if handleTemplateError(w, r, "executionAccount.go", "ExecutionAddress", "g.Wait()", err) != nil {
			return // an error has occurred and was processed
		}
		return
	}

	pngStr, pngStrInverse, err := utils.GenerateQRCodeForAddress(addressBytes)
	if err != nil {
		logger.WithError(err).Errorf("error generating qr code for address %v", address)
	}

	tabs := []types.ExecutionAddressPageTabs{}

	if internal != nil && len(internal.Data) != 0 {
		tabs = append(tabs, types.ExecutionAddressPageTabs{
			Id:   "internalTxns",
			Href: "#internalTxns",
			Text: "Internal Txns",
			Data: internal,
		})
	}
	if sqrctf1 != nil && len(sqrctf1.Data) != 0 {
		tabs = append(tabs, types.ExecutionAddressPageTabs{
			Id:   "sqrctf1Txns",
			Href: "#sqrctf1Txns",
			Text: "SqrcTf1 Token Txns",
			Data: sqrctf1,
		})
	}
	if sqrctn1 != nil && len(sqrctn1.Data) != 0 {
		tabs = append(tabs, types.ExecutionAddressPageTabs{
			Id:   "sqrctn1Txns",
			Href: "#sqrctn1Txns",
			Text: "SqrcTn1 Token Txns",
			Data: sqrctn1,
		})
	}
	if blocksMined != nil && len(blocksMined.Data) != 0 {
		tabs = append(tabs, types.ExecutionAddressPageTabs{
			Id:   "blocks",
			Href: "#blocks",
			Text: "Produced Blocks",
			Data: blocksMined,
		})
	}
	if sqrctb1 != nil && len(sqrctb1.Data) != 0 {
		tabs = append(tabs, types.ExecutionAddressPageTabs{
			Id:   "sqrctb1Txns",
			Href: "#sqrctb1Txns",
			Text: "SqrcTb1 Token Txns",
			Data: sqrctb1,
		})
	}
	if withdrawals != nil && len(withdrawals.Data) != 0 {
		tabs = append(tabs, types.ExecutionAddressPageTabs{
			Id:   "withdrawals",
			Href: "#withdrawals",
			Text: "Withdrawals",
			Data: withdrawals,
		})
	}

	data.Data = types.ExecutionAddressPageData{
		Address: address,
		// TODO(now.youtrack.cloud/issue/TZB-1)
		// QrnsName:           qrnsData.Domain,
		IsContract:         isContract,
		QRCode:             pngStr,
		QRCodeInverse:      pngStrInverse,
		Metadata:           metadata,
		WithdrawalsSummary: withdrawalSummary,
		TransactionsTable:  txns,
		InternalTxnsTable:  internal,
		SqrcTf1Table:       sqrctf1,
		SqrcTn1Table:       sqrctn1,
		SqrcTb1Table:       sqrctb1,
		WithdrawalsTable:   withdrawals,
		BlocksMinedTable:   blocksMined,
		Tabs:               tabs,
	}

	if handleTemplateError(w, r, "executionAccount.go", "ExecutionAddress", "Done", executionAddressTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}

func ExecutionAddressTransactions(w http.ResponseWriter, r *http.Request) {
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
		utils.LogError(err, "error getting execution tx table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func ExecutionAddressBlocksMined(w http.ResponseWriter, r *http.Request) {
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
		utils.LogError(err, "error getting execution blocks mined table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func ExecutionAddressWithdrawals(w http.ResponseWriter, r *http.Request) {
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

func ExecutionAddressInternalTransactions(w http.ResponseWriter, r *http.Request) {
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
		utils.LogError(err, "error getting execution internal tx table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func ExecutionAddressSqrcTf1Transactions(w http.ResponseWriter, r *http.Request) {
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
	data, err := db.BigtableClient.GetAddressSqrcTf1TableData(addressBytes, pageToken)
	if err != nil {
		utils.LogError(err, "error getting execution SQRCTF1 transactions table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func ExecutionAddressSqrcTn1Transactions(w http.ResponseWriter, r *http.Request) {
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
	data, err := db.BigtableClient.GetAddressSqrcTn1TableData(addressBytes, pageToken)
	if err != nil {
		utils.LogError(err, "error getting execution SQRCTN1 transactions table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func ExecutionAddressSqrcTb1Transactions(w http.ResponseWriter, r *http.Request) {
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
	data, err := db.BigtableClient.GetAddressSqrcTb1TableData(addressBytes, pageToken)
	if err != nil {
		utils.LogError(err, "error getting execution SQRCTB1 transactions table data", 0, errFields)
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		utils.LogError(err, "error enconding json response", 0, errFields)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// takes the "address" parameter from the request and transforms it to lower case. The QRNS name can be used instead of the address
func lowerAddressFromRequest(w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	address := vars["address"]
	// if utils.IsValidQrnsDomain(address) {
	// 	qrnsData, err := GetQrnsDomain(address)
	// 	if err != nil {
	// 		handleNotFoundJson(address, w, r, err)
	// 		return "", err
	// 	}
	// 	if len(qrnsData.Address) > 0 {
	// 		address = qrnsData.Address
	// 	}
	// }
	return strings.ToLower(strings.Replace(address, "Q", "", -1)), nil
}

// TODO(now.youtrack.cloud/issue/TZB-1)
/*
func handleNotFoundJson(address string, w http.ResponseWriter, r *http.Request, err error) {
	logger.Errorf("error getting address for QRNS name [%v] not found for %v route: %v", address, r.URL.String(), err)
	http.Error(w, "Invalid QRNS name", http.StatusInternalServerError)
}
*/

func handleNotFoundHtml(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "execution/addressNotFound.html")
	data := InitPageData(w, r, "blockchain", "/address", "not found", templateFiles)

	if handleTemplateError(w, r, "executionAccount.go", "ExecutionAddress", "not valid", templates.GetTemplate(templateFiles...).ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}
