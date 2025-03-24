package handlers

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/theQRL/go-zond"
	"github.com/theQRL/zond-beaconchain-explorer/eth1data"
	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/templates"
	"github.com/theQRL/zond-beaconchain-explorer/types"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"github.com/gorilla/mux"
	"github.com/theQRL/go-zond/common"
)

// Tx will show the tx using a go template
func Eth1TransactionTx(w http.ResponseWriter, r *http.Request) {
	txNotFoundTemplateFiles := append(layoutTemplateFiles, "eth1txnotfound.html")
	txTemplateFiles := append(layoutTemplateFiles, "eth1tx.html")
	mempoolTxTemplateFiles := append(layoutTemplateFiles, "mempoolTx.html")
	var txNotFoundTemplate = templates.GetTemplate(txNotFoundTemplateFiles...)
	var txTemplate = templates.GetTemplate(txTemplateFiles...)
	var mempoolTxTemplate = templates.GetTemplate(mempoolTxTemplateFiles...)

	w.Header().Set("Content-Type", "text/html")
	vars := mux.Vars(r)
	txHashString := vars["hash"]
	var data *types.PageData
	title := fmt.Sprintf("Transaction %v", txHashString)
	path := fmt.Sprintf("/tx/%v", txHashString)

	errFields := map[string]interface{}{
		"route":        r.URL.String(),
		"txHashString": txHashString,
	}

	txHash, err := hex.DecodeString(strings.ReplaceAll(txHashString, "0x", ""))
	if err != nil {
		logger.Warnf("error parsing tx hash %v: %v", txHashString, err)
		data = InitPageData(w, r, "blockchain", path, title, txNotFoundTemplateFiles)
		txTemplate = txNotFoundTemplate
	} else {
		txData, err := eth1data.GetEth1Transaction(common.BytesToHash(txHash), "ZND")
		if err != nil {
			mempool := services.LatestMempoolTransactions()
			mempoolTx := mempool.FindTxByHash(txHashString)
			if mempoolTx != nil {
				data = InitPageData(w, r, "blockchain", path, title, mempoolTxTemplateFiles)
				mempoolPageData := &types.MempoolTxPageData{RawMempoolTransaction: *mempoolTx}
				txTemplate = mempoolTxTemplate
				if mempoolTx.To == nil {
					mempoolPageData.IsContractCreation = true
				}
				if mempoolTx.Input != nil {
					mempoolPageData.TargetIsContract = true
				}

				data.Data = mempoolPageData
			} else {
				if !errors.Is(err, zond.NotFound) && !errors.Is(err, eth1data.ErrTxIsPending) {
					utils.LogError(err, "error getting eth1 transaction data", 0, errFields)
				}
				data = InitPageData(w, r, "blockchain", path, title, txNotFoundTemplateFiles)
				txTemplate = txNotFoundTemplate
			}
		} else {
			data = InitPageData(w, r, "blockchain", path, title, txTemplateFiles)
			data.Data = txData
		}
	}

	if utils.IsApiRequest(r) {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data.Data)
	} else {
		err = txTemplate.ExecuteTemplate(w, "layout", data)
	}

	if handleTemplateError(w, r, "eth1tx.go", "Eth1TransactionTx", "Done", err) != nil {
		return // an error has occurred and was processed
	}
}

func Eth1TransactionTxData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	txHashString := vars["hash"]

	err := json.NewEncoder(w).Encode(getEth1TransactionTxData(txHashString, "ZND"))
	if err != nil {
		logger.Errorf("error enconding json response for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func getEth1TransactionTxData(txhash, currency string) *types.DataTableResponse {
	tableData := make([][]interface{}, 0, minimumTransactionsPerUpdate)
	txHash, err := hex.DecodeString(strings.ReplaceAll(txhash, "0x", ""))
	if err != nil {
		logger.Warnf("error parsing tx hash %v: %v", txhash, err)
	} else {
		txData, err := eth1data.GetEth1Transaction(common.BytesToHash(txHash), currency)
		its := txData.InternalTxns
		if err != nil {
			utils.LogError(err, "error getting transaction data", 0, map[string]interface{}{"txhash": txHash})
		} else {
			for _, i := range its {
				tableData = append(tableData, []interface{}{
					i.TracePath,
					i.From,
					i.To,
					i.Amount,
					i.Gas.Limit,
				})
			}
		}
	}

	return &types.DataTableResponse{
		Data: tableData,
	}
}
