package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/theQRL/qrl-beaconchain-explorer/metrics"
	"github.com/theQRL/qrl-beaconchain-explorer/sqrctf1"
	"github.com/theQRL/qrl-beaconchain-explorer/types"
	"github.com/theQRL/qrl-beaconchain-explorer/utils"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	qrl "github.com/theQRL/go-zond"
	"github.com/theQRL/go-zond/common"
	"github.com/theQRL/go-zond/common/hexutil"
	gzond_types "github.com/theQRL/go-zond/core/types"
	"github.com/theQRL/go-zond/qrlclient"
	gzond_rpc "github.com/theQRL/go-zond/rpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GzondClient struct {
	endpoint     string
	rpcClient    *gzond_rpc.Client
	qrlClient    *qrlclient.Client
	chainID      *big.Int
	multiChecker *Balance
}

type ParityTraceResult struct {
	Action struct {
		CallType      string `json:"callType"`
		From          string `json:"from"`
		Gas           string `json:"gas"`
		Input         string `json:"input"`
		To            string `json:"to"`
		Value         string `json:"value"`
		Init          string `json:"init"`
		Address       string `json:"address"`
		Balance       string `json:"balance"`
		RefundAddress string `json:"refundAddress"`
		Author        string `json:"author"`
		RewardType    string `json:"rewardType"`
	} `json:"action"`
	BlockHash   string `json:"blockHash"`
	BlockNumber int    `json:"blockNumber"`
	Error       string `json:"error"`
	Result      struct {
		GasUsed string `json:"gasUsed"`
		Code    string `json:"code"`
		Output  string `json:"output"`
		Address string `json:"address"`
	} `json:"result"`

	Subtraces           int     `json:"subtraces"`
	TraceAddress        []int64 `json:"traceAddress"`
	TransactionHash     string  `json:"transactionHash"`
	TransactionPosition int     `json:"transactionPosition"`
	Type                string  `json:"type"`
}

func (trace *ParityTraceResult) ConvertFields() ([]byte, []byte, []byte, string) {
	var from, to, value []byte
	tx_type := trace.Type

	switch trace.Type {
	case "create":
		from = common.FromHex(trace.Action.From)
		to = common.FromHex(trace.Result.Address)
		value = common.FromHex(trace.Action.Value)
	case "call":
		from = common.FromHex(trace.Action.From)
		to = common.FromHex(trace.Action.To)
		value = common.FromHex(trace.Action.Value)
		tx_type = trace.Action.CallType
	default:
		spew.Dump(trace)
		utils.LogFatal(nil, "unknown trace type", 0, map[string]interface{}{"trace type": trace.Type, "tx hash": trace.TransactionHash})
	}
	return from, to, value, tx_type
}

var CurrentGzondClient *GzondClient

func NewGzondClient(endpoint string) (*GzondClient, error) {
	logger.Infof("initializing gzond client at %v", endpoint)
	client := &GzondClient{
		endpoint: endpoint,
	}

	rpcClient, err := gzond_rpc.Dial(client.endpoint)
	if err != nil {
		return nil, fmt.Errorf("error dialing rpc node: %v", err)
	}

	client.rpcClient = rpcClient

	qrlClient, err := qrlclient.Dial(client.endpoint)
	if err != nil {
		return nil, fmt.Errorf("error dialing rpc node: %v", err)
	}
	client.qrlClient = qrlClient

	addr, _ := common.NewAddressFromString("Qb1F8e55c7f64D203C1400B9D8555d050F94aDF39")
	client.multiChecker, err = NewBalance(addr, client.qrlClient)
	if err != nil {
		return nil, fmt.Errorf("error initiation balance checker contract: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	chainID, err := client.qrlClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting chainid of rpcclient: %w", err)
	}
	client.chainID = chainID

	return client, nil
}

func (client *GzondClient) Close() {
	client.rpcClient.Close()
	client.qrlClient.Close()
}

func (client *GzondClient) GetNativeClient() *qrlclient.Client {
	return client.qrlClient
}

func (client *GzondClient) GetBlockNumberByHash(hash string) (uint64, error) {
	startTime := time.Now()
	defer func() {
		metrics.TaskDuration.WithLabelValues("rpc_el_get_block_number_by_hash").Observe(time.Since(startTime).Seconds())
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	block, err := client.qrlClient.BlockByHash(ctx, common.HexToHash(hash))
	if err != nil {
		return 0, err
	}
	return block.NumberU64(), nil
}

func (client *GzondClient) GetBlock(number int64) (*types.ExecutionBlock, *types.GetBlockTimings, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	start := time.Now()
	timings := &types.GetBlockTimings{}

	block, err := client.qrlClient.BlockByNumber(ctx, big.NewInt(int64(number)))
	if err != nil {
		return nil, nil, err
	}

	timings.Headers = time.Since(start)
	start = time.Now()

	c := &types.ExecutionBlock{
		Hash:         block.Hash().Bytes(),
		ParentHash:   block.ParentHash().Bytes(),
		Coinbase:     block.Coinbase().Bytes(),
		Root:         block.Root().Bytes(),
		TxHash:       block.TxHash().Bytes(),
		ReceiptHash:  block.ReceiptHash().Bytes(),
		Number:       block.NumberU64(),
		GasLimit:     block.GasLimit(),
		GasUsed:      block.GasUsed(),
		Time:         timestamppb.New(time.Unix(int64(block.Time()), 0)),
		Extra:        block.Extra(),
		Random:       block.Random().Bytes(),
		Bloom:        block.Bloom().Bytes(),
		Transactions: []*types.ExecutionTransaction{},
	}

	if block.BaseFee() != nil {
		c.BaseFee = block.BaseFee().Bytes()
	}

	receipts := make([]*gzond_types.Receipt, len(block.Transactions()))
	reqs := make([]gzond_rpc.BatchElem, len(block.Transactions()))

	txs := block.Transactions()

	for _, tx := range txs {

		var from []byte
		sender, err := gzond_types.Sender(gzond_types.NewShanghaiSigner(tx.ChainId()), tx)
		if err != nil {
			from, _ = hex.DecodeString("abababababababababababababababababababab")
			logrus.Errorf("error converting tx %v to msg: %v", tx.Hash(), err)
		} else {
			from = sender.Bytes()
		}

		pbTx := &types.ExecutionTransaction{
			Type:                 uint32(tx.Type()),
			Nonce:                tx.Nonce(),
			MaxPriorityFeePerGas: tx.GasTipCap().Bytes(),
			MaxFeePerGas:         tx.GasFeeCap().Bytes(),
			Gas:                  tx.Gas(),
			Value:                tx.Value().Bytes(),
			Data:                 tx.Data(),
			From:                 from,
			ChainId:              tx.ChainId().Bytes(),
			AccessList:           []*types.AccessList{},
			Hash:                 tx.Hash().Bytes(),
			Itx:                  []*types.ExecutionInternalTransaction{},
		}

		if tx.To() != nil {
			pbTx.To = tx.To().Bytes()
		}
		c.Transactions = append(c.Transactions, pbTx)

	}

	for i := range reqs {
		reqs[i] = gzond_rpc.BatchElem{
			Method: "qrl_getTransactionReceipt",
			Args:   []interface{}{txs[i].Hash().String()},
			Result: &receipts[i],
		}
	}

	if len(reqs) > 0 {
		if err := client.rpcClient.BatchCallContext(ctx, reqs); err != nil {
			return nil, nil, fmt.Errorf("error retrieving receipts for block %v: %v", block.Number(), err)
		}
	}
	timings.Receipts = time.Since(start)

	for i := range reqs {
		if reqs[i].Error != nil {
			return nil, nil, fmt.Errorf("error retrieving receipt %v for block %v: %v", i, block.Number(), reqs[i].Error)
		}
		if receipts[i] == nil {
			return nil, nil, fmt.Errorf("got null value for receipt %d of block %v", i, block.Number())
		}

		r := receipts[i]
		c.Transactions[i].ContractAddress = r.ContractAddress[:]
		c.Transactions[i].CommulativeGasUsed = r.CumulativeGasUsed
		c.Transactions[i].GasUsed = r.GasUsed
		c.Transactions[i].LogsBloom = r.Bloom[:]
		c.Transactions[i].Logs = make([]*types.ExecutionLog, 0, len(r.Logs))

		for _, l := range r.Logs {
			pbLog := &types.ExecutionLog{
				Address: l.Address.Bytes(),
				Data:    l.Data,
				Removed: l.Removed,
				Topics:  make([][]byte, 0, len(l.Topics)),
			}

			for _, t := range l.Topics {
				pbLog.Topics = append(pbLog.Topics, t.Bytes())
			}
			c.Transactions[i].Logs = append(c.Transactions[i].Logs, pbLog)
		}
	}

	return c, timings, nil
}

func (client *GzondClient) GetLatestExecutionBlockNumber() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	latestBlock, err := client.qrlClient.BlockByNumber(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error getting latest block: %v", err)
	}

	return latestBlock.NumberU64(), nil
}

type GzondTraceCallResultWrapper struct {
	Result *GzondTraceCallResult
}

type GzondTraceCallResult struct {
	TransactionPosition int
	Time                string
	GasUsed             string
	From                common.Address
	To                  common.Address
	Value               string
	Gas                 string
	Input               string
	Output              string
	Error               string
	Type                string
}

func toCallArg(msg qrl.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}

	return arg
}

func (client *GzondClient) GetBalances(pairs []*types.ExecutionAddressBalance) ([]*types.ExecutionAddressBalance, error) {
	batchElements := make([]gzond_rpc.BatchElem, 0, len(pairs))

	ret := make([]*types.ExecutionAddressBalance, len(pairs))

	for i, pair := range pairs {
		result := ""

		ret[i] = &types.ExecutionAddressBalance{
			Address: pair.Address,
			Token:   pair.Token,
		}

		if len(pair.Token) < 20 {
			addr := common.BytesToAddress(pair.Address)
			batchElements = append(batchElements, gzond_rpc.BatchElem{
				Method: "qrl_getBalance",
				Args:   []interface{}{addr, "latest"},
				Result: &result,
			})
		} else {
			to := common.BytesToAddress(pair.Token)
			msg := qrl.CallMsg{
				To:   &to,
				Gas:  1000000,
				Data: common.Hex2Bytes(fmt.Sprintf("70a08231000000000000000000000000%x", pair.Address)),
			}

			batchElements = append(batchElements, gzond_rpc.BatchElem{
				Method: "qrl_call",
				Args:   []interface{}{toCallArg(msg), "latest"},
				Result: &result,
			})
		}
	}

	err := client.rpcClient.BatchCall(batchElements)
	if err != nil {
		return nil, fmt.Errorf("error during batch request: %v", err)
	}

	for i, el := range batchElements {
		if el.Error != nil {
			logrus.Warnf("error in batch call: %v", el.Error) // PPR: are smart contracts that pretend to implement the sqrctf1 standard but are somehow buggy
		}

		res := strings.TrimPrefix(*el.Result.(*string), "0x")
		ret[i].Balance = new(big.Int).SetBytes(common.Hex2Bytes(res)).Bytes()
	}

	return ret, nil
}

func (client *GzondClient) GetSQRCTF1TokenMetadata(token []byte) (*types.SQRCTF1Metadata, error) {
	logger.Infof("retrieving metadata for token %x", token)

	contract, err := sqrctf1.NewSqrcTf1(common.BytesToAddress(token), client.qrlClient)
	if err != nil {
		return nil, fmt.Errorf("error getting token-contract: sqrctf1.NewSqrcTf1: %w", err)
	}

	g := new(errgroup.Group)

	ret := &types.SQRCTF1Metadata{}

	g.Go(func() error {
		symbol, err := contract.Symbol(nil)
		if err != nil {
			if strings.Contains(err.Error(), "abi") {
				ret.Symbol = "UNKNOWN"
				return nil
			}

			return fmt.Errorf("error retrieving token symbol: %w", err)
		}

		ret.Symbol = symbol
		return nil
	})

	g.Go(func() error {
		totalSupply, err := contract.TotalSupply(nil)
		if err != nil {
			return fmt.Errorf("error retrieving token total supply: %w", err)
		}
		ret.TotalSupply = totalSupply.Bytes()
		return nil
	})

	g.Go(func() error {
		decimals, err := contract.Decimals(nil)
		if err != nil {
			return fmt.Errorf("error retrieving token decimals: %w", err)
		}
		ret.Decimals = big.NewInt(int64(decimals)).Bytes()
		return nil
	})

	err = g.Wait()
	if err != nil {
		return ret, err
	}

	if len(ret.Decimals) == 0 && ret.Symbol == "" && len(ret.TotalSupply) == 0 {
		// it's possible that a token contract implements the SQRCTF1 interfaces but does not return any values; we use a backup in this case
		ret = &types.SQRCTF1Metadata{
			Decimals:    []byte{0x0},
			Symbol:      "UNKNOWN",
			TotalSupply: []byte{0x0}}
	}

	return ret, err
}
