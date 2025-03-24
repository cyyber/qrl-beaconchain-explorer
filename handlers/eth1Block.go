package handlers

import (
	"database/sql"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/theQRL/zond-beaconchain-explorer/db"
	"github.com/theQRL/zond-beaconchain-explorer/rpc"
	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/templates"
	"github.com/theQRL/zond-beaconchain-explorer/types"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"github.com/gorilla/mux"
)

func Eth1Block(w http.ResponseWriter, r *http.Request) {
	blockTemplateFiles := append(layoutTemplateFiles,
		"slot/slot.html",
		"slot/attestations.html",
		"slot/deposits.html",
		"slot/votes.html",
		"slot/attesterSlashing.html",
		"slot/proposerSlashing.html",
		"slot/exits.html",
		"components/timestamp.html",
		"slot/overview.html",
		"slot/execTransactions.html",
		"slot/withdrawals.html")
	var blockTemplate = templates.GetTemplate(
		blockTemplateFiles...,
	)
	notFountTemplateFiles := append(layoutTemplateFiles, "slotnotfound.html")
	var blockNotFoundTemplate = templates.GetTemplate(notFountTemplateFiles...)

	w.Header().Set("Content-Type", "text/html")
	vars := mux.Vars(r)

	// parse block number from url
	numberString := strings.Replace(vars["block"], "0x", "", -1)
	var number uint64
	var err error
	if len(numberString) == 64 {
		number, err = rpc.CurrentGzondClient.GetBlockNumberByHash(numberString)
	} else {
		number, err = strconv.ParseUint(numberString, 10, 64)
	}

	if err != nil {
		data := InitPageData(w, r, "blockchain", "/block", fmt.Sprintf("Block %d", 0), notFountTemplateFiles)
		data.Data = "block"

		if handleTemplateError(w, r, "eth1Block.go", "Eth1Block", "number", blockNotFoundTemplate.ExecuteTemplate(w, "layout", data)) != nil {
			return // an error has occurred and was processed
		}
		return
	}

	eth1BlockPageData, err := GetExecutionBlockPageData(number, 10)
	if err != nil {
		data := InitPageData(w, r, "blockchain", "/block", fmt.Sprintf("Block %d", 0), notFountTemplateFiles)
		data.Data = "block"
		if handleTemplateError(w, r, "eth1Block.go", "Eth1Block", "GetExecutionBlockPageData", blockNotFoundTemplate.ExecuteTemplate(w, "layout", data)) != nil {
			return // an error has occurred and was processed
		}
		return
	}

	data := InitPageData(w, r, "blockchain", "/block", fmt.Sprintf("Block %d", number), blockTemplateFiles)

	// blockSlot := uint64(0)
	blockSlot := utils.TimeToSlot(uint64(eth1BlockPageData.Ts.Unix()))

	// retrieve consensus data
	blockPageData, err := GetSlotPageData(blockSlot)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("error retrieving slot page data: %v", err)
		}

		data.Data = "block"
		if handleTemplateError(w, r, "eth1Block.go", "Eth1Block", "GetSlotPageData", blockNotFoundTemplate.ExecuteTemplate(w, "layout", data)) != nil {
			return // an error has occurred and was processed
		}
		return
	}
	blockPageData.ExecutionData = eth1BlockPageData

	data.Data = blockPageData

	if handleTemplateError(w, r, "eth1Block.go", "Eth1Block", "Done (Post Merge)", blockTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}

}

func GetExecutionBlockPageData(number uint64, limit int) (*types.Eth1BlockPageData, error) {
	block, err := db.BigtableClient.GetBlockFromBlocksTable(number)
	if diffToHead := int64(services.LatestEth1BlockNumber()) - int64(number); err != nil && diffToHead < 0 && diffToHead >= -5 {
		block, _, err = rpc.CurrentGzondClient.GetBlock(int64(number) /*, "parity/gzond"*/)
	}
	if err != nil {
		return nil, err
	}

	// retrieve address names from bigtable
	names := make(map[string]string)
	names[string(block.Coinbase)] = ""
	for _, tx := range block.Transactions {
		names[string(tx.From)] = ""
		names[string(tx.To)] = ""
	}
	names, _, err = db.BigtableClient.GetAddressesNamesArMetadata(&names, nil)
	if err != nil {
		return nil, err
	}

	// calculate total block reward and set lowest gas price
	txs := []types.Eth1BlockPageTransaction{}
	txFees := new(big.Int)
	lowestGasPrice := big.NewInt(1 << 62)

	contractInteractionTypes, err := db.BigtableClient.GetAddressContractInteractionsAtBlock(block)
	if err != nil {
		utils.LogError(err, "error getting contract states", 0)
	}

	for i, tx := range block.Transactions {

		// sum txFees
		txFee := db.CalculateTxFeeFromTransaction(tx, new(big.Int).SetBytes(block.BaseFee))
		txFees.Add(txFees, txFee)

		effectiveGasPrice := big.NewInt(0)
		if gasUsed := new(big.Int).SetUint64(tx.GasUsed); gasUsed.Cmp(big.NewInt(0)) != 0 {
			// calculate effective gas price
			effectiveGasPrice = new(big.Int).Div(txFee, gasUsed)
			if effectiveGasPrice.Cmp(lowestGasPrice) < 0 {
				lowestGasPrice = effectiveGasPrice
			}
		}

		contractCreation := tx.GetTo() == nil
		// set tx to if tx is contract creation
		if contractCreation {
			tx.To = tx.ContractAddress
		}

		var contractInteraction types.ContractInteractionType
		if len(contractInteractionTypes) > i {
			contractInteraction = contractInteractionTypes[i]
		}

		txs = append(txs, types.Eth1BlockPageTransaction{
			Hash:          fmt.Sprintf("%#x", tx.Hash),
			HashFormatted: utils.FormatTransactionHash(tx.Hash, tx.ErrorMsg == ""),
			From:          fmt.Sprintf("%#x", tx.From),
			FromFormatted: utils.FormatAddressWithLimits(tx.From, names[string(tx.From)], false, "address", 15, 20, true),
			To:            fmt.Sprintf("%#x", tx.To),
			ToFormatted:   utils.FormatAddressWithLimits(tx.To, db.BigtableClient.GetAddressLabel(names[string(tx.To)], contractInteraction), contractInteraction != types.CONTRACT_NONE, "address", 15, 20, true),
			Value:         new(big.Int).SetBytes(tx.Value),
			Fee:           txFee,
			GasPrice:      effectiveGasPrice,
			Method:        db.BigtableClient.GetMethodLabel(tx.GetData(), contractInteraction),
		})
	}

	if limit > 0 {
		if len(txs) > limit {
			txs = txs[:limit]
		} else {
			txs = txs[:0]
		}
	}

	burnedTxFees := new(big.Int).Mul(new(big.Int).SetBytes(block.BaseFee), big.NewInt(int64(block.GasUsed)))
	reward := new(big.Int).Sub(txFees, burnedTxFees)

	nextBlock := number + 1
	if nextBlock > services.LatestEth1BlockNumber() {
		nextBlock = 0
	}
	eth1BlockPageData := types.Eth1BlockPageData{
		Number:         number,
		PreviousBlock:  number - 1,
		NextBlock:      nextBlock,
		TxCount:        uint64(len(block.Transactions)),
		Hash:           fmt.Sprintf("%#x", block.Hash),
		ParentHash:     fmt.Sprintf("%#x", block.ParentHash),
		MinerAddress:   fmt.Sprintf("%#x", block.Coinbase),
		MinerFormatted: utils.FormatAddressWithLimits(block.Coinbase, names[string(block.Coinbase)], false, "address", 42, 42, true),
		Reward:         reward,
		TxFees:         txFees,
		GasUsage:       utils.FormatBlockUsage(block.GasUsed, block.GasLimit),
		GasLimit:       block.GasLimit,
		LowestGasPrice: lowestGasPrice,
		Ts:             block.GetTime().AsTime(),
		BaseFeePerGas:  new(big.Int).SetBytes(block.BaseFee),
		BurnedFees:     burnedTxFees,
		Extra:          fmt.Sprintf("%#x", block.Extra),
		Txs:            txs,
	}

	return &eth1BlockPageData, nil
}
