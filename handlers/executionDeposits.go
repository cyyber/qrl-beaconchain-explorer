package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/theQRL/qrl-beaconchain-explorer/db"
	"github.com/theQRL/qrl-beaconchain-explorer/services"
	"github.com/theQRL/qrl-beaconchain-explorer/templates"
	"github.com/theQRL/qrl-beaconchain-explorer/types"
	"github.com/theQRL/qrl-beaconchain-explorer/utils"
)

// Deposits will return information about deposits using a go template
func Deposits(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "deposits.html", "index/depositChart.html")
	var DepositsTemplate = templates.GetTemplate(templateFiles...)

	w.Header().Set("Content-Type", "text/html")

	pageData := &types.DepositsPageData{}

	latestChartsPageData := services.LatestChartsPageData()
	if len(latestChartsPageData) != 0 {
		for _, c := range latestChartsPageData {
			if c.Path == "deposits" {
				pageData.DepositChart = c
				break
			}
		}
	}

	pageData.Stats = services.GetLatestStats()
	pageData.DepositContract = utils.Config.Chain.ClConfig.DepositContractAddress

	data := InitPageData(w, r, "blockchain", "/deposits", "Deposits", templateFiles)
	data.Data = pageData

	if handleTemplateError(w, r, "executionDeposits.go", "Deposits", "", DepositsTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}

func ExecutionDeposits(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/validators/deposits", http.StatusMovedPermanently)
}

// ExecutionDepositsData will return execution-deposits as json
func ExecutionDepositsData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()

	// TODO(now.youtrack.cloud/issue/TZB-1)
	// search := ReplaceQrnsNameWithAddress(q.Get("search[value]"))
	search := q.Get("search[value]")
	search = strings.Replace(search, "0x", "", -1)

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
	length, err := strconv.ParseUint(q.Get("length"), 10, 64)
	if err != nil {
		logger.Warnf("error converting datatables length parameter from string to int: %v", err)
		http.Error(w, "Error: Missing or invalid parameter length", http.StatusBadRequest)
		return
	}
	if length > 100 {
		length = 100
	}

	orderColumn := q.Get("order[0][column]")
	orderByMap := map[string]string{
		"0": "from_address",
		"1": "publickey",
		"2": "withdrawal_credential",
		"3": "amount",
		"4": "tx_hash",
		"5": "block_ts",
		"6": "block_number",
		"7": "state",
		"8": "valid_signature",
	}
	orderBy, exists := orderByMap[orderColumn]
	if !exists {
		orderBy = "block_ts"
	}

	orderDir := q.Get("order[0][dir]")

	latestEpoch := services.LatestEpoch()
	validatorOnlineThresholdSlot := GetValidatorOnlineThresholdSlot()

	deposits, depositCount, err := db.GetExecutionDepositsJoinConsensusDeposits(search, length, start, orderBy, orderDir, latestEpoch, validatorOnlineThresholdSlot)
	if err != nil {
		logger.Errorf("GetExecutionDeposits error retrieving execution_deposit data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tableData := make([][]interface{}, len(deposits))
	for i, d := range deposits {
		valid := "❌"
		if d.ValidSignature {
			valid = "✅"
		}
		tableData[i] = []interface{}{
			utils.FormatExecutionAddress(d.FromAddress),
			utils.FormatPublicKey(d.PublicKey),
			utils.FormatWithdawalCredentials(d.WithdrawalCredentials, true),
			utils.FormatDepositAmount(d.Amount, "Quanta"),
			utils.FormatExecutionTxHash(d.TxHash),
			utils.FormatTimestamp(d.BlockTs.Unix()),
			utils.FormatExecutionBlock(d.BlockNumber),
			utils.FormatValidatorStatus(d.State),
			valid,
		}
	}

	data := &types.DataTableResponse{
		Draw:            draw,
		RecordsTotal:    depositCount,
		RecordsFiltered: depositCount,
		Data:            tableData,
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Errorf("error enconding json response for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
