package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/theQRL/qrl-beaconchain-explorer/db"
	"github.com/theQRL/qrl-beaconchain-explorer/metrics"
	"github.com/theQRL/qrl-beaconchain-explorer/rpc"
	"github.com/theQRL/qrl-beaconchain-explorer/services"
	"github.com/theQRL/qrl-beaconchain-explorer/types"
	"github.com/theQRL/qrl-beaconchain-explorer/utils"
	"github.com/theQRL/qrl-beaconchain-explorer/version"

	"github.com/coocood/freecache"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	_ "net/http/pprof"
)

func main() {
	gzondEndpoint := flag.String("gzond", "", "Gzond archive node enpoint")
	block := flag.Int64("block", 0, "Index a specific block")

	reorgDepth := flag.Int("reorg.depth", 20, "Lookback to check and handle chain reorgs")

	concurrencyBlocks := flag.Int64("blocks.concurrency", 30, "Concurrency to use when indexing blocks from gzond")
	startBlocks := flag.Int64("blocks.start", 0, "Block to start indexing")
	endBlocks := flag.Int64("blocks.end", 0, "Block to finish indexing")
	bulkBlocks := flag.Int64("blocks.bulk", 8000, "Maximum number of blocks to be processed before saving")
	offsetBlocks := flag.Int64("blocks.offset", 100, "Blocks offset")
	checkBlocksGaps := flag.Bool("blocks.gaps", false, "Check for gaps in the blocks table")
	checkBlocksGapsLookback := flag.Int("blocks.gaps.lookback", 1000000, "Lookback for gaps check of the blocks table")
	// TODO(now.youtrack.cloud/issue/TZB-7)
	// traceMode := flag.String("blocks.tracemode", "parity/gzond", "Trace mode to use, can be either 'parity', 'gzond' or 'parity/gzond' for both")

	concurrencyData := flag.Int64("data.concurrency", 30, "Concurrency to use when indexing data from bigtable")
	startData := flag.Int64("data.start", 0, "Block to start indexing")
	endData := flag.Int64("data.end", 0, "Block to finish indexing")
	bulkData := flag.Int64("data.bulk", 8000, "Maximum number of blocks to be processed before saving")
	offsetData := flag.Int64("data.offset", 1000, "Data offset")
	checkDataGaps := flag.Bool("data.gaps", false, "Check for gaps in the data table")
	checkDataGapsLookback := flag.Int("data.gaps.lookback", 1000000, "Lookback for gaps check of the blocks table")

	enableBalanceUpdater := flag.Bool("balances.enabled", false, "Enable balance update process")
	enableFullBalanceUpdater := flag.Bool("balances.full.enabled", false, "Enable full balance update process")
	balanceUpdaterBatchSize := flag.Int("balances.batch", 1000, "Batch size for balance updates")

	versionFlag := flag.Bool("version", false, "Print version and exit")

	configPath := flag.String("config", "", "Path to the config file, if empty string defaults will be used")

	// TODO(now.youtrack.cloud/issue/TZB-1)
	// enableZnsUpdater := flag.Bool("zns.enabled", false, "Enable zns update process")
	// znsBatchSize := flag.Int64("zns.batch", 200, "Batch size for zns updates")

	flag.Parse()

	if *versionFlag {
		fmt.Println(version.Version)
		fmt.Println(version.GoVersion)
		return
	}

	cfg := &types.Config{}
	err := utils.ReadConfig(cfg, *configPath)
	if err != nil {
		logrus.Fatalf("error reading config file: %v", err)
	}
	utils.Config = cfg
	logrus.WithField("config", *configPath).WithField("version", version.Version).WithField("chainName", utils.Config.Chain.ClConfig.ConfigName).Printf("starting")

	if utils.Config.Metrics.Enabled {
		go func(addr string) {
			logrus.Infof("serving metrics on %v", addr)
			if err := metrics.Serve(addr); err != nil {
				logrus.WithError(err).Fatal("Error serving metrics")
			}
		}(utils.Config.Metrics.Address)
	}

	// enable pprof endpoint if requested
	if utils.Config.Pprof.Enabled {
		go func() {
			logrus.Infof("starting pprof http server on port %s", utils.Config.Pprof.Port)
			logrus.Info(http.ListenAndServe(fmt.Sprintf("localhost:%s", utils.Config.Pprof.Port), nil))
		}()
	}

	db.MustInitDB(&types.DatabaseConfig{
		Username:     cfg.WriterDatabase.Username,
		Password:     cfg.WriterDatabase.Password,
		Name:         cfg.WriterDatabase.Name,
		Host:         cfg.WriterDatabase.Host,
		Port:         cfg.WriterDatabase.Port,
		MaxOpenConns: cfg.WriterDatabase.MaxOpenConns,
		MaxIdleConns: cfg.WriterDatabase.MaxIdleConns,
		SSL:          cfg.WriterDatabase.SSL,
	}, &types.DatabaseConfig{
		Username:     cfg.ReaderDatabase.Username,
		Password:     cfg.ReaderDatabase.Password,
		Name:         cfg.ReaderDatabase.Name,
		Host:         cfg.ReaderDatabase.Host,
		Port:         cfg.ReaderDatabase.Port,
		MaxOpenConns: cfg.ReaderDatabase.MaxOpenConns,
		MaxIdleConns: cfg.ReaderDatabase.MaxIdleConns,
		SSL:          cfg.ReaderDatabase.SSL,
	}, "pgx", "postgres")
	defer db.ReaderDb.Close()
	defer db.WriterDb.Close()

	if gzondEndpoint == nil || *gzondEndpoint == "" {
		if utils.Config.ELNodeEndpoint == "" {
			utils.LogFatal(nil, "no EL node url provided", 0)
		} else {
			logrus.Info("applying EL endpoint from config")
			*gzondEndpoint = utils.Config.ELNodeEndpoint
		}
	}

	logrus.Infof("using gzond node at %v", *gzondEndpoint)
	client, err := rpc.NewGzondClient(*gzondEndpoint)
	if err != nil {
		utils.LogFatal(err, "gzond client creation error", 0)
	}

	chainId := strconv.FormatUint(utils.Config.Chain.ClConfig.DepositChainID, 10)

	balanceUpdaterPrefix := chainId + ":B:"

	nodeChainId, err := client.GetNativeClient().ChainID(context.Background())
	if err != nil {
		utils.LogFatal(err, "node chain id error", 0)
	}

	if nodeChainId.String() != chainId {
		logrus.Fatalf("node chain id mismatch, wanted %v got %v", chainId, nodeChainId.String())
	}

	bt, err := db.InitBigtable(utils.Config.Bigtable.Project, utils.Config.Bigtable.Instance, chainId, utils.Config.RedisCacheEndpoint)
	if err != nil {
		logrus.Fatalf("error connecting to bigtable: %v", err)
	}
	defer bt.Close()

	if *enableFullBalanceUpdater {
		ProcessMetadataUpdates(bt, client, balanceUpdaterPrefix, *balanceUpdaterBatchSize, -1)
		return
	}

	transforms := make([]func(blk *types.Eth1Block, cache *freecache.Cache) (*types.BulkMutations, *types.BulkMutations, error), 0)
	transforms = append(transforms,
		bt.TransformBlock,
		bt.TransformTx,
		bt.TransformItx,
		bt.TransformZRC20,
		bt.TransformZRC721,
		bt.TransformZRC1155,
		bt.TransformWithdrawals,
		// TODO(now.youtrack.cloud/issue/TZB-1)
		// bt.TransformZnsNameRegistered,
		bt.TransformContract)

	cache := freecache.NewCache(100 * 1024 * 1024) // 100 MB limit

	if *block != 0 {
		// TODO(now.youtrack.cloud/issue/TZB-7)
		err = IndexFromNode(bt, client, *block, *block, *concurrencyBlocks /*, *traceMode*/)
		if err != nil {
			logrus.WithError(err).Fatalf("error indexing from node, start: %v end: %v concurrency: %v", *block, *block, *concurrencyBlocks)
		}
		err = bt.IndexEventsWithTransformers(*block, *block, transforms, *concurrencyData, cache)
		if err != nil {
			logrus.WithError(err).Fatalf("error indexing from bigtable")
		}
		cache.Clear()

		logrus.Infof("indexing of block %v completed", *block)
		return
	}

	if *checkBlocksGaps {
		bt.CheckForGapsInBlocksTable(*checkBlocksGapsLookback)
		return
	}

	if *checkDataGaps {
		bt.CheckForGapsInDataTable(*checkDataGapsLookback)
		return
	}

	if *endBlocks != 0 && *startBlocks < *endBlocks {
		// TODO(now.youtrack.cloud/issue/TZB-7)
		err = IndexFromNode(bt, client, *startBlocks, *endBlocks, *concurrencyBlocks /*, *traceMode*/)
		if err != nil {
			logrus.WithError(err).Fatalf("error indexing from node, start: %v end: %v concurrency: %v", *startBlocks, *endBlocks, *concurrencyBlocks)
		}
		return
	}

	if *endData != 0 && *startData < *endData {
		err = bt.IndexEventsWithTransformers(int64(*startData), int64(*endData), transforms, *concurrencyData, cache)
		if err != nil {
			logrus.WithError(err).Fatalf("error indexing from bigtable")
		}
		cache.Clear()
		return
	}

	var lastBlockFromNodeOld uint64
	var lastBlockFromNodeSameCount uint64
	lastSuccessulBlockIndexingTs := time.Now()
	for ; ; time.Sleep(time.Second * 14) {
		err := HandleChainReorgs(bt, client, *reorgDepth)
		if err != nil {
			logrus.Errorf("error handling chain reorgs: %v", err)
			continue
		}

		lastBlockFromNode, err := client.GetLatestEth1BlockNumber()
		if err != nil {
			lastBlockFromNodeSameCount++
			if lastBlockFromNodeSameCount > 20 { // nearly 5 minutes no new block
				utils.LogFatal(err, "no new block in 20 tries", 0, map[string]interface{}{
					"lastBlockFromNode": lastBlockFromNodeOld,
				})
			}
			logrus.Errorf("error retrieving latest eth block number: %v", err)
			continue
		}
		if lastBlockFromNode != lastBlockFromNodeOld {
			lastBlockFromNodeOld = lastBlockFromNode
			lastBlockFromNodeSameCount = 0
		} else {
			lastBlockFromNodeSameCount++
			if lastBlockFromNodeSameCount > 20 { // nearly 5 minutes no new block
				utils.LogFatal(nil, "no new block in 20 tries", 0, map[string]interface{}{
					"lastBlockFromNode": lastBlockFromNodeOld,
				})
			}
		}

		lastBlockFromBlocksTable, err := bt.GetLastBlockInBlocksTable()
		if err != nil {
			logrus.Errorf("error retrieving last blocks from blocks table: %v", err)
			continue
		}

		lastBlockFromDataTable, err := bt.GetLastBlockInDataTable()
		if err != nil {
			logrus.Errorf("error retrieving last blocks from data table: %v", err)
			continue
		}

		logrus.WithFields(
			logrus.Fields{
				"node":   lastBlockFromNode,
				"blocks": lastBlockFromBlocksTable,
				"data":   lastBlockFromDataTable,
			},
		).Infof("last blocks")

		continueAfterError := false
		if lastBlockFromNode > 0 {
			if lastBlockFromBlocksTable < int(lastBlockFromNode) {
				logrus.Infof("missing blocks %v to %v in blocks table, indexing ...", lastBlockFromBlocksTable+1, lastBlockFromNode)

				startBlock := int64(lastBlockFromBlocksTable+1) - *offsetBlocks
				if startBlock < 0 {
					startBlock = 0
				}

				if *bulkBlocks <= 0 || *bulkBlocks > int64(lastBlockFromNode)-startBlock+1 {
					*bulkBlocks = int64(lastBlockFromNode) - startBlock + 1
				}

				for startBlock <= int64(lastBlockFromNode) && !continueAfterError {
					endBlock := startBlock + *bulkBlocks - 1
					if endBlock > int64(lastBlockFromNode) {
						endBlock = int64(lastBlockFromNode)
					}

					// TODO(now.youtrack.cloud/issue/TZB-7)
					err = IndexFromNode(bt, client, startBlock, endBlock, *concurrencyBlocks /*, *traceMode*/)
					if err != nil {
						errMsg := "error indexing from node"
						errFields := map[string]interface{}{
							"start":       startBlock,
							"end":         endBlock,
							"concurrency": *concurrencyBlocks}
						if time.Since(lastSuccessulBlockIndexingTs) > time.Minute*30 {
							utils.LogFatal(err, errMsg, 0, errFields)
						} else {
							utils.LogError(err, errMsg, 0, errFields)
						}
						continueAfterError = true
						continue
					} else {
						lastSuccessulBlockIndexingTs = time.Now()
					}

					startBlock = endBlock + 1
				}
				if continueAfterError {
					continue
				}
			}

			if lastBlockFromDataTable < int(lastBlockFromNode) {
				logrus.Infof("missing blocks %v to %v in data table, indexing ...", lastBlockFromDataTable+1, lastBlockFromNode)

				startBlock := int64(lastBlockFromDataTable+1) - *offsetData
				if startBlock < 0 {
					startBlock = 0
				}

				if *bulkData <= 0 || *bulkData > int64(lastBlockFromNode)-startBlock+1 {
					*bulkData = int64(lastBlockFromNode) - startBlock + 1
				}

				for startBlock <= int64(lastBlockFromNode) && !continueAfterError {
					endBlock := startBlock + *bulkData - 1
					if endBlock > int64(lastBlockFromNode) {
						endBlock = int64(lastBlockFromNode)
					}

					err = bt.IndexEventsWithTransformers(startBlock, endBlock, transforms, *concurrencyData, cache)
					if err != nil {
						utils.LogError(err, "error indexing from bigtable", 0, map[string]interface{}{"start": startBlock, "end": endBlock, "concurrency": *concurrencyData})
						cache.Clear()
						continueAfterError = true
						continue
					}
					cache.Clear()

					startBlock = endBlock + 1
				}
				if continueAfterError {
					continue
				}
			}
		}

		if *enableBalanceUpdater {
			ProcessMetadataUpdates(bt, client, balanceUpdaterPrefix, *balanceUpdaterBatchSize, 10)
		}

		// TODO(now.youtrack.cloud/issue/TZB-1)
		// if *enableZnsUpdater {
		// 	err := bt.ImportZnsUpdates(client.GetNativeClient(), *znsBatchSize)
		// 	if err != nil {
		// 		utils.LogError(err, "error importing zns updates", 0, nil)
		// 		continue
		// 	}
		// }

		logrus.Infof("index run completed")
		services.ReportStatus("elIndexer", "Running", nil)
	}

	// utils.WaitForCtrlC()
}

func HandleChainReorgs(bt *db.Bigtable, client *rpc.GzondClient, depth int) error {
	ctx := context.Background()
	// get latest block from the node
	latestNodeBlock, err := client.GetNativeClient().BlockByNumber(ctx, nil)
	if err != nil {
		return err
	}
	latestNodeBlockNumber := latestNodeBlock.NumberU64()

	// for each block check if block node hash and block db hash match
	if depth > int(latestNodeBlockNumber) {
		depth = int(latestNodeBlockNumber)
	}
	for i := latestNodeBlockNumber - uint64(depth); i <= latestNodeBlockNumber; i++ {
		nodeBlock, err := client.GetNativeClient().HeaderByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			return err
		}

		dbBlock, err := bt.GetBlockFromBlocksTable(i)
		if err != nil {
			if err == db.ErrBlockNotFound { // exit if we hit a block that is not yet in the db
				return nil
			}
			return err
		}

		if !bytes.Equal(nodeBlock.Hash().Bytes(), dbBlock.Hash) {
			logrus.Warnf("found incosistency at height %v, node block hash: %x, db block hash: %x", i, nodeBlock.Hash().Bytes(), dbBlock.Hash)

			// first we set the cached marker of the last block in the blocks/data table to the block prior to the forked one
			if i > 0 {
				previousBlock := i - 1
				err := bt.SetLastBlockInBlocksTable(int64(previousBlock))
				if err != nil {
					return fmt.Errorf("error setting last block [%v] in blocks table: %w", previousBlock, err)
				}
				err = bt.SetLastBlockInDataTable(int64(previousBlock))
				if err != nil {
					return fmt.Errorf("error setting last block [%v] in data table: %w", previousBlock, err)
				}
				// now we can proceed to delete all blocks including and after the forked block
			}
			// delete all blocks starting from the fork block up to the latest block in the db
			for j := i; j <= latestNodeBlockNumber; j++ {
				dbBlock, err := bt.GetBlockFromBlocksTable(j)
				if err != nil {
					if err == db.ErrBlockNotFound { // exit if we hit a block that is not yet in the db
						return nil
					}
					return err
				}
				logrus.Infof("deleting block at height %v with hash %x", dbBlock.Number, dbBlock.Hash)

				err = bt.DeleteBlock(dbBlock.Number, dbBlock.Hash)
				if err != nil {
					return err
				}
			}
		} else {
			logrus.Infof("height %v, node block hash: %x, db block hash: %x", i, nodeBlock.Hash().Bytes(), dbBlock.Hash)
		}
	}

	return nil
}

func ProcessMetadataUpdates(bt *db.Bigtable, client *rpc.GzondClient, prefix string, batchSize int, iterations int) {
	lastKey := prefix

	its := 0
	for {
		start := time.Now()
		keys, pairs, err := bt.GetMetadataUpdates(prefix, lastKey, batchSize)
		if err != nil {
			logrus.Errorf("error retrieving metadata updates from bigtable: %v", err)
			return
		}

		if len(keys) == 0 {
			return
		}

		balances := make([]*types.Eth1AddressBalance, 0, len(pairs))
		for b := 0; b < len(pairs); b += batchSize {
			start := b
			end := b + batchSize
			if len(pairs) < end {
				end = len(pairs)
			}

			logrus.Infof("processing batch %v with start %v and end %v", b, start, end)

			b, err := client.GetBalances(pairs[start:end])

			if err != nil {
				logrus.Errorf("error retrieving balances from node: %v", err)
				return
			}
			balances = append(balances, b...)
		}

		err = bt.SaveBalances(balances, keys)
		if err != nil {
			logrus.Errorf("error saving balances to bigtable: %v", err)
			return
		}

		lastKey = keys[len(keys)-1]
		logrus.Infof("retrieved %v balances in %v, currently at %v", len(balances), time.Since(start), lastKey)

		its++

		if iterations != -1 && its > iterations {
			return
		}
	}
}

// TODO(now.youtrack.cloud/issue/TZB-7)
func IndexFromNode(bt *db.Bigtable, client *rpc.GzondClient, start, end, concurrency int64 /*, traceMode string*/) error {
	ctx := context.Background()
	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(int(concurrency))

	startTs := time.Now()
	lastTickTs := time.Now()

	processedBlocks := int64(0)

	for i := start; i <= end; i++ {
		i := i
		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
			}

			startTime := time.Now()
			defer func() {
				metrics.TaskDuration.WithLabelValues("bt_index_from_node").Observe(time.Since(startTime).Seconds())
			}()

			blockStartTs := time.Now()
			// TODO(now.youtrack.cloud/issue/TZB-7)
			bc, timings, err := client.GetBlock(i /*, traceMode*/)
			if err != nil {
				return fmt.Errorf("error getting block: %v from zond node err: %w", i, err)
			}

			metrics.TaskDuration.WithLabelValues("rpc_el_get_block_headers").Observe(timings.Headers.Seconds())
			metrics.TaskDuration.WithLabelValues("rpc_el_get_block_receipts").Observe(timings.Receipts.Seconds())
			metrics.TaskDuration.WithLabelValues("rpc_el_get_block_traces").Observe(timings.Traces.Seconds())

			dbStart := time.Now()
			err = bt.SaveBlock(bc)
			if err != nil {
				return fmt.Errorf("error saving block: %v to bigtable: %w", i, err)
			}
			current := atomic.AddInt64(&processedBlocks, 1)
			if current%100 == 0 {
				r := end - start
				if r == 0 {
					r = 1
				}
				perc := float64(i-start) * 100 / float64(r)

				logrus.Infof("retrieved & saved block %v (0x%x) in %v (header: %v, receipts: %v, traces: %v, db: %v)", bc.Number, bc.Hash, time.Since(blockStartTs), timings.Headers, timings.Receipts, timings.Traces, time.Since(dbStart))
				logrus.Infof("processed %v blocks in %v (%.1f blocks / sec); sync is %.1f%% complete", current, time.Since(startTs), float64((current))/time.Since(lastTickTs).Seconds(), perc)

				lastTickTs = time.Now()
				atomic.StoreInt64(&processedBlocks, 0)
			}
			return nil
		})

	}

	err := g.Wait()

	if err != nil {
		return err
	}

	lastBlockInCache, err := bt.GetLastBlockInBlocksTable()
	if err != nil {
		return err
	}

	if end > int64(lastBlockInCache) {
		err := bt.SetLastBlockInBlocksTable(end)

		if err != nil {
			return err
		}
	}
	return nil
}
