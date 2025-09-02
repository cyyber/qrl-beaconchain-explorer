package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/theQRL/qrl-beaconchain-explorer/cache"
	"github.com/theQRL/qrl-beaconchain-explorer/db"
	"github.com/theQRL/qrl-beaconchain-explorer/utils"
)

const latestBlockNumberCacheKey = "latestExecutionBlockNumber"
const latestBlockHashRootCacheKey = "latestExecutionBlockRootHash"

// latestBlockUpdater updates the most recent execution block number variable
func latestBlockUpdater(wg *sync.WaitGroup) {
	firstRun := true

	for {
		recent, err := db.BigtableClient.GetMostRecentBlockFromDataTable()
		if err != nil {
			utils.LogError(err, "error getting most recent execution block", 0)
		}
		cacheKey := fmt.Sprintf("%d:frontend:%s", utils.Config.Chain.ClConfig.DepositChainID, latestBlockNumberCacheKey)
		err = cache.TieredCache.SetUint64(cacheKey, recent.GetNumber(), utils.Day)
		if err != nil {
			utils.LogError(err, fmt.Sprintf("error caching latest block number with cache key %s", latestBlockNumberCacheKey), 0)
		}

		if firstRun {
			logger.Info("initialized execution block updater")
			wg.Done()
			firstRun = false
		}
		ReportStatus("latestBlockUpdater", "Running", nil)
		time.Sleep(time.Second * 10)
	}
}

// LatestExecutionBlockNumber will return most recent execution block number
func LatestExecutionBlockNumber() uint64 {
	cacheKey := fmt.Sprintf("%d:frontend:%s", utils.Config.Chain.ClConfig.DepositChainID, latestBlockNumberCacheKey)

	if wanted, err := cache.TieredCache.GetUint64WithLocalTimeout(cacheKey, time.Second*5); err == nil {
		return wanted
	} else {
		utils.LogError(err, fmt.Sprintf("error retrieving latest block number from cache with key %s", latestBlockNumberCacheKey), 0)
	}
	return 0
}

// headBlockRootHashUpdater updates the hash of the current chain head block
func headBlockRootHashUpdater(wg *sync.WaitGroup) {
	firstRun := true

	for {
		blockRootHash := []byte{}
		err := db.ReaderDb.Get(&blockRootHash, `
		SELECT blockroot
		FROM blocks
		WHERE status = '1'
		ORDER BY slot DESC
		LIMIT 1`)

		if err != nil {
			utils.LogError(err, "error getting blockroot hash for chain head", 0)
		}
		cacheKey := fmt.Sprintf("%d:frontend:%s", utils.Config.Chain.ClConfig.DepositChainID, latestBlockHashRootCacheKey)
		err = cache.TieredCache.SetString(cacheKey, string(blockRootHash), utils.Day)
		if err != nil {
			utils.LogError(err, fmt.Sprintf("error caching latest blockroot hash with cache key %s", latestBlockHashRootCacheKey), 0)
		}

		if firstRun {
			logger.Info("initialized execution head block root hash updater")
			wg.Done()
			firstRun = false
		}
		ReportStatus("headBlockRootHashUpdater", "Running", nil)
		time.Sleep(time.Second * 10)
	}
}
