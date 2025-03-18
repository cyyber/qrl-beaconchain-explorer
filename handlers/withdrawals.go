package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/theQRL/zond-beaconchain-explorer/db"
	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/templates"
	"github.com/theQRL/zond-beaconchain-explorer/types"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"golang.org/x/sync/errgroup"
)

// Withdrawals will return information about recent withdrawals
func Withdrawals(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "withdrawals.html", "validator/withdrawalOverviewRow.html", "components/charts.html")
	var withdrawalsTemplate = templates.GetTemplate(templateFiles...)

	w.Header().Set("Content-Type", "text/html")

	pageData := &types.WithdrawalsPageData{}
	pageData.Stats = services.GetLatestStats()

	data := InitPageData(w, r, "validators", "/withdrawals", "Validator Withdrawals", templateFiles)

	latestChartsPageData := services.LatestChartsPageData()
	if len(latestChartsPageData) != 0 {
		for _, c := range latestChartsPageData {
			if c.Path == "withdrawals" {
				pageData.WithdrawalChart = c
				break
			}
		}
	}

	data.Data = pageData

	err := withdrawalsTemplate.ExecuteTemplate(w, "layout", data)
	if handleTemplateError(w, r, "withdrawals.go", "withdrawals", "", err) != nil {
		return // an error has occurred and was processed
	}
}

// WithdrawalsData will return eth1-deposits as json
func WithdrawalsData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	currency := GetCurrency(r)
	q := r.URL.Query()

	// TODO(rgeraldes24)
	// search := ReplaceZnsNameWithAddress(q.Get("search[value]"))
	search := q.Get("search[value]")

	draw, err := strconv.ParseUint(q.Get("draw"), 10, 64)
	if err != nil {
		logger.Warnf("error converting datatables draw parameter from string to int: %v", err)
		http.Error(w, "Error: Missing or invalid parameter draw", http.StatusBadRequest)
		return
	}

	start, err := strconv.ParseUint(q.Get("start"), 10, 64)
	if err != nil {
		logger.Warnf("error converting datatables start parameter from string to int: %v", err)
		http.Error(w, "Error: Missing or invalid parameter start", http.StatusBadRequest)
		return
	}
	if start > db.WithdrawalsQueryLimit {
		start = db.WithdrawalsQueryLimit
	}
	length, err := strconv.ParseUint(q.Get("length"), 10, 64)
	if err != nil {
		logger.Warnf("error converting datatables length parameter from string to int: %v", err)
		http.Error(w, "Error: Missing or invalid parameter length", http.StatusBadRequest)
		return
	}
	if length > 100 {
		length = 100
	}

	orderBy := q.Get("order[0][column]")
	orderDir := q.Get("order[0][dir]")

	data, err := WithdrawalsTableData(draw, search, length, start, orderBy, orderDir, currency)
	if err != nil {
		logger.Errorf("error getting withdrawal table data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Errorf("error enconding json response for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func WithdrawalsTableData(draw uint64, search string, length, start uint64, orderBy, orderDir string, currency string) (*types.DataTableResponse, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*10))
	defer cancel()

	orderByMap := map[string]string{
		"0": "epoch",
		"1": "slot",
		"2": "index",
		"3": "validator",
		"4": "address",
		"5": "amount",
	}
	orderColumn, exists := orderByMap[orderBy]
	if !exists {
		orderBy = "index"
	}

	if orderDir != "asc" {
		orderDir = "desc"
	}

	g, gCtx := errgroup.WithContext(ctx)
	withdrawalCount := uint64(0)
	g.Go(func() error {
		select {
		case <-gCtx.Done():
			return nil
		default:
		}
		var err error
		withdrawalCount, err = db.GetTotalWithdrawals()
		if err != nil {
			return fmt.Errorf("error getting total withdrawals: %w", err)
		}
		return nil
	})

	filteredCount := uint64(0)
	trimmedSearch := strings.ToLower(strings.TrimPrefix(search, "0x"))
	if trimmedSearch != "" {
		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return nil
			default:
			}
			var err error
			filteredCount, err = db.GetWithdrawalsCountForQuery(search)
			if err != nil {
				return fmt.Errorf("error getting withdrwal count for filter [%v]: %w", search, err)
			}
			return nil
		})
	}

	withdrawals := []*types.Withdrawals{}
	g.Go(func() error {
		select {
		case <-gCtx.Done():
			return nil
		default:
		}
		var err error
		withdrawals, err = db.GetWithdrawals(search, length, start, orderColumn, orderDir)
		if err != nil {
			return fmt.Errorf("error getting withdrawals: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	if trimmedSearch == "" {
		filteredCount = withdrawalCount
	}

	formatCurrency := currency
	// TODO(rgeraldes24)
	if currency == "ZND" {
		formatCurrency = "ZND"
	}

	var err error
	names := make(map[string]string)
	for _, v := range withdrawals {
		names[string(v.Address)] = ""
	}
	names, _, err = db.BigtableClient.GetAddressesNamesArMetadata(&names, nil)
	if err != nil {
		return nil, err
	}

	tableData := make([][]interface{}, len(withdrawals))
	for i, w := range withdrawals {
		tableData[i] = []interface{}{
			utils.FormatEpoch(utils.EpochOfSlot(w.Slot)),
			utils.FormatBlockSlot(w.Slot),
			template.HTML(fmt.Sprintf("%v", w.Index)),
			utils.FormatValidator(w.ValidatorIndex),
			utils.FormatTimestamp(utils.SlotToTime(w.Slot).Unix()),
			utils.FormatAddressWithLimits(w.Address, names[string(w.Address)], false, "address", visibleDigitsForHash+5, 18, true),
			utils.FormatAmount(new(big.Int).Mul(new(big.Int).SetUint64(w.Amount), big.NewInt(1e9)), formatCurrency, 6),
		}
	}

	if filteredCount > db.WithdrawalsQueryLimit {
		filteredCount = db.WithdrawalsQueryLimit
	}
	if withdrawalCount > db.WithdrawalsQueryLimit {
		withdrawalCount = db.WithdrawalsQueryLimit
	}

	data := &types.DataTableResponse{
		Draw:            draw,
		RecordsTotal:    withdrawalCount,
		RecordsFiltered: filteredCount,
		Data:            tableData,
		PageLength:      length,
		DisplayStart:    start,
	}
	return data, nil
}

// Eth1DepositsData will return eth1-deposits as json
func DilithiumChangeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query()

	// TODO(rgeraldes24)
	// search := ReplaceZnsNameWithAddress(q.Get("search[value]"))
	search := q.Get("search[value]")

	draw, err := strconv.ParseUint(q.Get("draw"), 10, 64)
	if err != nil {
		logger.Warnf("error converting datatables draw parameter from string to int: %v", err)
		http.Error(w, "Error: Missing or invalid parameter draw", http.StatusBadRequest)
		return
	}
	start, err := strconv.ParseUint(q.Get("start"), 10, 64)
	if err != nil {
		logger.Warnf("error converting datatables start parameter from string to int: %v", err)
		http.Error(w, "Error: Missing or invalid parameter start", http.StatusBadRequest)
		return
	}
	if start > db.DilithiumChangeQueryLimit {
		// limit offset to 10000, otherwise the query will be too slow
		start = db.DilithiumChangeQueryLimit
	}
	length, err := strconv.ParseUint(q.Get("length"), 10, 64)
	if err != nil {
		logger.Warnf("error converting datatables length parameter from string to int: %v", err)
		http.Error(w, "Error: Missing or invalid parameter length", http.StatusBadRequest)
		return
	}
	if length > 100 {
		length = 100
	}

	orderBy := q.Get("order[0][column]")
	orderDir := q.Get("order[0][dir]")

	data, err := DilithiumTableData(draw, search, length, start, orderBy, orderDir)
	if err != nil {
		logger.Errorf("Error getting dilithium changes: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Errorf("error enconding json response for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func DilithiumTableData(draw uint64, search string, length, start uint64, orderBy, orderDir string) (*types.DataTableResponse, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*10))
	defer cancel()
	orderByMap := map[string]string{
		"0": "block_slot",
		"1": "block_slot",
		"2": "validatorindex",
	}
	orderVar, exists := orderByMap[orderBy]
	if !exists {
		orderBy = "block_slot"
	}

	if orderDir != "asc" {
		orderDir = "desc"
	}

	g, gCtx := errgroup.WithContext(ctx)
	totalCount := uint64(0)
	g.Go(func() error {
		select {
		case <-gCtx.Done():
			return nil
		default:
		}
		var err error
		totalCount, err = db.GetTotalDilithiumChanges()
		if err != nil {
			return fmt.Errorf("error getting total dilithium changes: %w", err)
		}
		return nil
	})

	filteredCount := uint64(0)
	trimmedSearch := strings.ToLower(strings.TrimPrefix(search, "0x"))
	if trimmedSearch != "" {
		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return nil
			default:
			}
			var err error
			filteredCount, err = db.GetDilithiumChangesCountForQuery(search)
			if err != nil {
				return fmt.Errorf("error getting dilithium changes count for filter [%v]: %w", search, err)
			}
			return nil
		})
	}

	dilithiumChange := []*types.DilithiumChange{}
	g.Go(func() error {
		select {
		case <-gCtx.Done():
			return nil
		default:
		}
		var err error
		dilithiumChange, err = db.GetDilithiumChanges(search, length, start, orderVar, orderDir)
		if err != nil {
			return fmt.Errorf("error getting dilithium changes: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	if trimmedSearch == "" {
		filteredCount = totalCount
	}

	var err error
	names := make(map[string]string)
	for _, v := range dilithiumChange {
		names[string(v.Address)] = ""
	}
	names, _, err = db.BigtableClient.GetAddressesNamesArMetadata(&names, nil)
	if err != nil {
		return nil, err
	}

	tableData := make([][]interface{}, len(dilithiumChange))
	for i, dilithium := range dilithiumChange {
		tableData[i] = []interface{}{
			utils.FormatEpoch(utils.EpochOfSlot(dilithium.Slot)),
			utils.FormatBlockSlot(dilithium.Slot),
			utils.FormatValidator(dilithium.Validatorindex),
			utils.FormatHashWithCopy(dilithium.Signature),
			utils.FormatHashWithCopy(dilithium.DilithiumPubkey),
			utils.FormatAddressWithLimits(dilithium.Address, names[string(dilithium.Address)], false, "address", visibleDigitsForHash+5, 18, true),
		}
	}

	if totalCount > db.DilithiumChangeQueryLimit {
		totalCount = db.DilithiumChangeQueryLimit
	}
	if filteredCount > db.DilithiumChangeQueryLimit {
		filteredCount = db.DilithiumChangeQueryLimit
	}

	data := &types.DataTableResponse{
		Draw:            draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: filteredCount,
		Data:            tableData,
		PageLength:      length,
		DisplayStart:    start,
	}
	return data, nil
}
