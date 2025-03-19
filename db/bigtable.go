package db

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/theQRL/zond-beaconchain-explorer/types"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	gcp_bigtable "cloud.google.com/go/bigtable"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	itypes "github.com/theQRL/zond-beaconchain-explorer/zond-rewards/types"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
)

var BigtableClient *Bigtable

const (
	DEFAULT_FAMILY                        = "f"
	VALIDATOR_BALANCES_FAMILY             = "vb"
	VALIDATOR_HIGHEST_ACTIVE_INDEX_FAMILY = "ha"
	ATTESTATIONS_FAMILY                   = "at"
	PROPOSALS_FAMILY                      = "pr"
	SYNC_COMMITTEES_FAMILY                = "sc"
	INCOME_DETAILS_COLUMN_FAMILY          = "id"
	STATS_COLUMN_FAMILY                   = "stats"
	SERIES_FAMILY                         = "series"

	SUM_COLUMN = "sum"

	MAX_CL_BLOCK_NUMBER = 1000000000 - 1
	MAX_EL_BLOCK_NUMBER = 1000000000
	MAX_EPOCH           = 1000000000 - 1

	max_block_number_v1 = 1000000000
	max_epoch_v1        = 1000000000

	MAX_BATCH_MUTATIONS   = 100000
	DEFAULT_BATCH_INSERTS = 10000

	REPORT_TIMEOUT = time.Second * 10
)

type Bigtable struct {
	client *gcp_bigtable.Client

	tableBeaconchain       *gcp_bigtable.Table
	tableValidators        *gcp_bigtable.Table
	tableValidatorsHistory *gcp_bigtable.Table

	tableData            *gcp_bigtable.Table
	tableBlocks          *gcp_bigtable.Table
	tableMetadataUpdates *gcp_bigtable.Table
	tableMetadata        *gcp_bigtable.Table

	redisCache *redis.Client

	LastAttestationCache    map[uint64]uint64
	LastAttestationCacheMux *sync.Mutex

	chainId string
}

func InitBigtable(project, instance, chainId, redisAddress string) (*Bigtable, error) {

	if utils.Config.Bigtable.Emulator {

		if utils.Config.Bigtable.EmulatorHost == "" {
			utils.Config.Bigtable.EmulatorHost = "127.0.0.1"
		}
		logger.Infof("using emulated local bigtable environment, setting BIGTABLE_EMULATOR_HOST env variable to %s:%d", utils.Config.Bigtable.EmulatorHost, utils.Config.Bigtable.EmulatorPort)
		err := os.Setenv("BIGTABLE_EMULATOR_HOST", fmt.Sprintf("%s:%d", utils.Config.Bigtable.EmulatorHost, utils.Config.Bigtable.EmulatorPort))

		if err != nil {
			logger.Fatalf("unable to set bigtable emulator environment variable: %v", err)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	poolSize := 50
	btClient, err := gcp_bigtable.NewClient(ctx, project, instance, option.WithGRPCConnectionPool(poolSize))
	// btClient, err := gcp_bigtable.NewClient(context.Background(), project, instance)

	if err != nil {
		return nil, err
	}

	rdc := redis.NewClient(&redis.Options{
		Addr:        redisAddress,
		ReadTimeout: time.Second * 20,
	})

	if err := rdc.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	bt := &Bigtable{
		client:                  btClient,
		tableData:               btClient.Open("data"),
		tableBlocks:             btClient.Open("blocks"),
		tableMetadataUpdates:    btClient.Open("metadata_updates"),
		tableMetadata:           btClient.Open("metadata"),
		tableBeaconchain:        btClient.Open("beaconchain"),
		tableValidators:         btClient.Open("beaconchain_validators"),
		tableValidatorsHistory:  btClient.Open("beaconchain_validators_history"),
		chainId:                 chainId,
		redisCache:              rdc,
		LastAttestationCacheMux: &sync.Mutex{},
	}

	BigtableClient = bt
	return bt, nil
}

func (bigtable *Bigtable) Close() {
	time.Sleep(time.Second * 5)
	bigtable.client.Close()
}

func (bigtable *Bigtable) GetClient() *gcp_bigtable.Client {
	return bigtable.client
}

func (bigtable *Bigtable) SaveValidatorBalances(epoch uint64, validators []*types.Validator) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	// start := time.Now()
	ts := gcp_bigtable.Timestamp(0)

	muts := types.NewBulkMutations(len(validators))

	highestActiveIndex := uint64(0)
	epochKey := bigtable.reversedPaddedEpoch(epoch)

	for _, validator := range validators {

		if validator.Balance > 0 && validator.Index > highestActiveIndex {
			highestActiveIndex = validator.Index
		}

		balanceEncoded := make([]byte, 8)
		binary.LittleEndian.PutUint64(balanceEncoded, validator.Balance)
		// TODO(rgeraldes24)
		effectiveBalanceEncoded := uint8(validator.EffectiveBalance / 1e9) // we can encode the effective balance in 1 byte as it is capped at 32ETH and only decrements in 1 ETH steps

		combined := append(balanceEncoded, effectiveBalanceEncoded)
		mut := &gcp_bigtable.Mutation{}
		mut.Set(VALIDATOR_BALANCES_FAMILY, "b", ts, combined)
		key := fmt.Sprintf("%s:%s:%s:%s", bigtable.chainId, bigtable.validatorIndexToKey(validator.Index), VALIDATOR_BALANCES_FAMILY, epochKey)

		muts.Add(key, mut)
	}

	err := bigtable.WriteBulk(muts, bigtable.tableValidatorsHistory, MAX_BATCH_MUTATIONS)
	if err != nil {
		return err
	}

	// store the highes active validator index for that epoch
	highestActiveIndexEncoded := make([]byte, 8)
	binary.LittleEndian.PutUint64(highestActiveIndexEncoded, highestActiveIndex)

	mut := &gcp_bigtable.Mutation{}
	mut.Set(VALIDATOR_HIGHEST_ACTIVE_INDEX_FAMILY, VALIDATOR_HIGHEST_ACTIVE_INDEX_FAMILY, ts, highestActiveIndexEncoded)
	key := fmt.Sprintf("%s:%s:%s", bigtable.chainId, VALIDATOR_HIGHEST_ACTIVE_INDEX_FAMILY, epochKey)
	err = bigtable.tableValidatorsHistory.Apply(ctx, key, mut)
	if err != nil {
		return err
	}
	return nil
}

func (bigtable *Bigtable) SaveProposalAssignments(epoch uint64, assignments map[uint64]uint64) error {

	start := time.Now()
	ts := gcp_bigtable.Timestamp(0)

	muts := types.NewBulkMutations(len(assignments))

	for slot, validator := range assignments {
		mut := gcp_bigtable.NewMutation()
		mut.Set(PROPOSALS_FAMILY, "p", ts, []byte{})

		key := fmt.Sprintf("%s:%s:%s:%s:%s", bigtable.chainId, bigtable.validatorIndexToKey(validator), PROPOSALS_FAMILY, bigtable.reversedPaddedEpoch(epoch), bigtable.reversedPaddedSlot(slot))

		muts.Add(key, mut)
	}

	err := bigtable.WriteBulk(muts, bigtable.tableValidatorsHistory, MAX_BATCH_MUTATIONS)

	if err != nil {
		return err
	}

	logger.Infof("exported proposal assignments to bigtable in %v", time.Since(start))
	return nil
}

func (bigtable *Bigtable) SaveAttestationDuties(duties map[types.Slot]map[types.ValidatorIndex][]types.Slot) error {

	// Initialize in memory last attestation cache lazily
	bigtable.LastAttestationCacheMux.Lock()
	if bigtable.LastAttestationCache == nil {
		t := time.Now()
		var err error
		bigtable.LastAttestationCache, err = bigtable.GetLastAttestationSlots([]uint64{})

		if err != nil {
			bigtable.LastAttestationCacheMux.Unlock()
			return err
		}
		logger.Infof("initialized in memory last attestation slot cache with %v validators in %v", len(bigtable.LastAttestationCache), time.Since(t))

	}
	bigtable.LastAttestationCacheMux.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	start := time.Now()

	mutsInclusionSlot := types.NewBulkMutations(MAX_BATCH_MUTATIONS)

	mutLastAttestationSlot := gcp_bigtable.NewMutation()
	mutLastAttestationSlotCount := 0

	for attestedSlot, validators := range duties {
		for validator, inclusions := range validators {

			epoch := utils.EpochOfSlot(uint64(attestedSlot))
			bigtable.LastAttestationCacheMux.Lock()
			if len(inclusions) == 0 { // for missed attestations we write the max block number which will yield a cell ts of 0
				inclusions = append(inclusions, MAX_CL_BLOCK_NUMBER)
			}
			for _, inclusionSlot := range inclusions {
				key := fmt.Sprintf("%s:%s:%s:%s", bigtable.chainId, bigtable.validatorIndexToKey(uint64(validator)), ATTESTATIONS_FAMILY, bigtable.reversedPaddedEpoch(epoch))

				mutInclusionSlot := gcp_bigtable.NewMutation()
				mutInclusionSlot.Set(ATTESTATIONS_FAMILY, fmt.Sprintf("%d", attestedSlot), gcp_bigtable.Timestamp((MAX_CL_BLOCK_NUMBER-inclusionSlot)*1000), []byte{})

				mutsInclusionSlot.Add(key, mutInclusionSlot)

				if inclusionSlot != MAX_CL_BLOCK_NUMBER && uint64(attestedSlot) > bigtable.LastAttestationCache[uint64(validator)] {
					mutLastAttestationSlot.Set(ATTESTATIONS_FAMILY, fmt.Sprintf("%d", validator), gcp_bigtable.Timestamp((attestedSlot)*1000), []byte{})
					bigtable.LastAttestationCache[uint64(validator)] = uint64(attestedSlot)
					mutLastAttestationSlotCount++

					if mutLastAttestationSlotCount == MAX_BATCH_MUTATIONS {
						mutStart := time.Now()
						err := bigtable.tableValidators.Apply(ctx, fmt.Sprintf("%s:lastAttestationSlot", bigtable.chainId), mutLastAttestationSlot)
						if err != nil {
							bigtable.LastAttestationCacheMux.Unlock()
							return fmt.Errorf("error applying last attestation slot mutations: %v", err)
						}
						mutLastAttestationSlot = gcp_bigtable.NewMutation()
						mutLastAttestationSlotCount = 0
						logger.Infof("applyied last attestation slot mutations in %v", time.Since(mutStart))
					}
				}

			}
			bigtable.LastAttestationCacheMux.Unlock()
		}
	}

	err := bigtable.WriteBulk(mutsInclusionSlot, bigtable.tableValidatorsHistory, MAX_BATCH_MUTATIONS)

	if err != nil {
		return fmt.Errorf("error writing attestation inclusion slot mutations: %v", err)
	}

	if mutLastAttestationSlotCount > 0 {
		err := bigtable.tableValidators.Apply(ctx, fmt.Sprintf("%s:lastAttestationSlot", bigtable.chainId), mutLastAttestationSlot)
		if err != nil {
			return fmt.Errorf("error applying last attestation slot mutations: %v", err)
		}
	}

	logger.Infof("exported %v attestations to bigtable in %v", mutsInclusionSlot.Len(), time.Since(start))
	return nil
}

// This method is only to be used for migrating the last attestation slot to bigtable and should not be used for any other purpose
func (bigtable *Bigtable) SetLastAttestationSlot(validator uint64, lastAttestationSlot uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	mutLastAttestationSlot := gcp_bigtable.NewMutation()
	mutLastAttestationSlot.Set(ATTESTATIONS_FAMILY, fmt.Sprintf("%d", validator), gcp_bigtable.Timestamp(lastAttestationSlot*1000), []byte{})
	err := bigtable.tableValidators.Apply(ctx, fmt.Sprintf("%s:lastAttestationSlot", bigtable.chainId), mutLastAttestationSlot)
	if err != nil {
		return err
	}

	return nil
}

func (bigtable *Bigtable) SaveProposal(block *types.Block) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	start := time.Now()

	if len(block.BlockRoot) != 32 { // skip dummy blocks
		return nil
	}
	mut := gcp_bigtable.NewMutation()
	mut.Set(PROPOSALS_FAMILY, "b", gcp_bigtable.Timestamp((MAX_CL_BLOCK_NUMBER-block.Slot)*1000), []byte{})
	key := fmt.Sprintf("%s:%s:%s:%s:%s", bigtable.chainId, bigtable.validatorIndexToKey(block.Proposer), PROPOSALS_FAMILY, bigtable.reversedPaddedEpoch(utils.EpochOfSlot(block.Slot)), bigtable.reversedPaddedSlot(block.Slot))

	err := bigtable.tableValidatorsHistory.Apply(ctx, key, mut)

	if err != nil {
		return err
	}

	logger.Infof("exported proposal to bigtable in %v", time.Since(start))
	return nil
}

func (bigtable *Bigtable) SaveSyncComitteeDuties(duties map[types.Slot]map[types.ValidatorIndex]bool) error {
	start := time.Now()

	if len(duties) == 0 {
		logger.Infof("no sync duties to export")
		return nil
	}

	muts := types.NewBulkMutations(int(utils.Config.Chain.ClConfig.SlotsPerEpoch*utils.Config.Chain.ClConfig.SyncCommitteeSize + 1))

	for slot, validators := range duties {
		for validator, participated := range validators {
			mut := gcp_bigtable.NewMutation()
			if participated {
				mut.Set(SYNC_COMMITTEES_FAMILY, "s", gcp_bigtable.Timestamp((MAX_CL_BLOCK_NUMBER-slot)*1000), []byte{})
			} else {
				mut.Set(SYNC_COMMITTEES_FAMILY, "s", gcp_bigtable.Timestamp(0), []byte{})
			}
			key := fmt.Sprintf("%s:%s:%s:%s:%s", bigtable.chainId, bigtable.validatorIndexToKey(uint64(validator)), SYNC_COMMITTEES_FAMILY, bigtable.reversedPaddedEpoch(utils.EpochOfSlot(uint64(slot))), bigtable.reversedPaddedSlot(uint64(slot)))

			muts.Add(key, mut)
		}
	}

	err := bigtable.WriteBulk(muts, bigtable.tableValidatorsHistory, MAX_BATCH_MUTATIONS)
	if err != nil {
		return err
	}

	logger.Infof("exported %v sync committee duties to bigtable in %v", muts.Len(), time.Since(start))
	return nil
}

// GetMaxValidatorindexForEpoch returns the higest validatorindex with a balance at that epoch
func (bigtable *Bigtable) GetMaxValidatorindexForEpoch(epoch uint64) (uint64, error) {
	return bigtable.getMaxValidatorindexForEpochV2(epoch)
}

func (bigtable *Bigtable) getMaxValidatorindexForEpochV2(epoch uint64) (uint64, error) {

	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"epoch": epoch,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*5))
	defer cancel()

	key := fmt.Sprintf("%s:%s:%s", bigtable.chainId, VALIDATOR_HIGHEST_ACTIVE_INDEX_FAMILY, bigtable.reversedPaddedEpoch(epoch))

	row, err := bigtable.tableValidatorsHistory.ReadRow(ctx, key)
	if err != nil {
		return 0, err
	}

	for _, ri := range row[VALIDATOR_HIGHEST_ACTIVE_INDEX_FAMILY] {
		return binary.LittleEndian.Uint64(ri.Value), nil
	}

	return 0, nil
}

func (bigtable *Bigtable) GetValidatorBalanceHistory(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64][]*types.ValidatorBalance, error) {
	return bigtable.getValidatorBalanceHistoryV2(validators, startEpoch, endEpoch)
}

func (bigtable *Bigtable) getValidatorBalanceHistoryV2(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64][]*types.ValidatorBalance, error) {

	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"validators_count": len(validators),
			"startEpoch":       startEpoch,
			"endEpoch":         endEpoch,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	if len(validators) == 0 {
		return nil, fmt.Errorf("passing empty validator array is unsupported")
	}

	batchSize := 1000
	concurrency := 10

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*5))
	defer cancel()

	res := make(map[uint64][]*types.ValidatorBalance, len(validators))
	resMux := &sync.Mutex{}

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)

	for i := 0; i < len(validators); i += batchSize {

		upperBound := i + batchSize
		if len(validators) < upperBound {
			upperBound = len(validators)
		}
		vals := validators[i:upperBound]

		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
			}
			ranges := bigtable.getValidatorsEpochRanges(vals, VALIDATOR_BALANCES_FAMILY, startEpoch, endEpoch)
			ro := gcp_bigtable.LimitRows(int64(endEpoch-startEpoch+1) * int64(len(vals)))

			handleRow := func(r gcp_bigtable.Row) bool {
				// logger.Info(r.Key())
				keySplit := strings.Split(r.Key(), ":")

				epoch, err := strconv.ParseUint(keySplit[3], 10, 64)
				if err != nil {
					logger.Errorf("error parsing epoch from row key %v: %v", r.Key(), err)
					return false
				}

				validator, err := bigtable.validatorKeyToIndex(keySplit[1])
				if err != nil {
					logger.Errorf("error parsing validator index from row key %v: %v", r.Key(), err)
					return false
				}
				resMux.Lock()
				if res[validator] == nil {
					res[validator] = make([]*types.ValidatorBalance, 0)
				}
				resMux.Unlock()

				for _, ri := range r[VALIDATOR_BALANCES_FAMILY] {
					balances := ri.Value

					balanceBytes := balances[0:8]
					balance := binary.LittleEndian.Uint64(balanceBytes)
					var effectiveBalance uint64
					if len(balances) == 9 { // in new schema the effective balance is encoded in 1 byte
						effectiveBalance = uint64(balances[8]) * 1e9
					} else {
						effectiveBalanceBytes := balances[8:16]
						effectiveBalance = binary.LittleEndian.Uint64(effectiveBalanceBytes)
					}

					resMux.Lock()
					res[validator] = append(res[validator], &types.ValidatorBalance{
						Epoch:            MAX_EPOCH - epoch,
						Balance:          balance,
						EffectiveBalance: effectiveBalance,
						Index:            validator,
						PublicKey:        []byte{},
					})
					resMux.Unlock()
				}
				return true
			}

			err := bigtable.tableValidatorsHistory.ReadRows(gCtx, ranges, handleRow, ro)
			if err != nil {
				return err
			}

			// logrus.Infof("retrieved data for validators %v - %v", vals[0], vals[len(vals)-1])
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}

func (bigtable *Bigtable) GetValidatorAttestationHistory(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64][]*types.ValidatorAttestation, error) {
	return bigtable.getValidatorAttestationHistoryV2(validators, startEpoch, endEpoch)
}

func (bigtable *Bigtable) getValidatorAttestationHistoryV2(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64][]*types.ValidatorAttestation, error) {
	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"validatorsCount": len(validators),
			"startEpoch":      startEpoch,
			"endEpoch":        endEpoch,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	if len(validators) == 0 {
		return nil, fmt.Errorf("passing empty validator array is unsupported")
	}

	batchSize := 1000
	concurrency := 10

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*5))
	defer cancel()

	res := make(map[uint64][]*types.ValidatorAttestation, len(validators))
	resMux := &sync.Mutex{}

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)

	attestationsMap := make(map[types.ValidatorIndex]map[types.Slot][]*types.ValidatorAttestation)

	for i := 0; i < len(validators); i += batchSize {

		upperBound := i + batchSize
		if len(validators) < upperBound {
			upperBound = len(validators)
		}
		vals := validators[i:upperBound]

		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
			}
			ranges := bigtable.getValidatorsEpochRanges(vals, ATTESTATIONS_FAMILY, startEpoch, endEpoch)
			filter := gcp_bigtable.LimitRows(int64(endEpoch-startEpoch+1) * int64(len(vals))) // max is one row per epoch
			err := bigtable.tableValidatorsHistory.ReadRows(ctx, ranges, func(r gcp_bigtable.Row) bool {
				keySplit := strings.Split(r.Key(), ":")

				validator, err := bigtable.validatorKeyToIndex(keySplit[1])
				if err != nil {
					logger.Errorf("error parsing validator from row key %v: %v", r.Key(), err)
					return false
				}

				for _, ri := range r[ATTESTATIONS_FAMILY] {
					attesterSlotString := strings.Replace(ri.Column, ATTESTATIONS_FAMILY+":", "", 1)
					attesterSlot, err := strconv.ParseUint(attesterSlotString, 10, 64)
					if err != nil {
						logger.Errorf("error parsing slot from row key %v: %v", r.Key(), err)
						return false
					}
					inclusionSlot := MAX_CL_BLOCK_NUMBER - uint64(ri.Timestamp)/1000

					status := uint64(1)
					if inclusionSlot == MAX_CL_BLOCK_NUMBER {
						inclusionSlot = 0
						status = 0
					}

					resMux.Lock()
					if attestationsMap[types.ValidatorIndex(validator)] == nil {
						attestationsMap[types.ValidatorIndex(validator)] = make(map[types.Slot][]*types.ValidatorAttestation)
					}

					if attestationsMap[types.ValidatorIndex(validator)][types.Slot(attesterSlot)] == nil {
						attestationsMap[types.ValidatorIndex(validator)][types.Slot(attesterSlot)] = make([]*types.ValidatorAttestation, 0)
					}

					attestationsMap[types.ValidatorIndex(validator)][types.Slot(attesterSlot)] = append(attestationsMap[types.ValidatorIndex(validator)][types.Slot(attesterSlot)], &types.ValidatorAttestation{
						InclusionSlot: inclusionSlot,
						Status:        status,
					})
					resMux.Unlock()

				}
				return true
			}, filter)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Find all missed and orphaned slots
	slots := []uint64{}
	maxSlot := ((endEpoch + 1) * utils.Config.Chain.ClConfig.SlotsPerEpoch) - 1
	for slot := startEpoch * utils.Config.Chain.ClConfig.SlotsPerEpoch; slot <= maxSlot; slot++ {
		slots = append(slots, slot)
	}

	var missedSlotsMap map[uint64]bool
	var orphanedSlotsMap map[uint64]bool

	g = new(errgroup.Group)

	g.Go(func() error {
		var err error
		missedSlotsMap, err = GetMissedSlotsMap(slots)
		return err
	})

	g.Go(func() error {
		var err error
		orphanedSlotsMap, err = GetOrphanedSlotsMap(slots)
		return err
	})
	err := g.Wait()
	if err != nil {
		return nil, err
	}

	// Convert the attestationsMap info to the return format
	// Set the delay of the inclusionSlot
	for validator, attestations := range attestationsMap {
		if res[uint64(validator)] == nil {
			res[uint64(validator)] = make([]*types.ValidatorAttestation, 0)
		}
		for attesterSlot, att := range attestations {
			currentAttInfo := att[0]
			for _, attInfo := range att {
				if orphanedSlotsMap[attInfo.InclusionSlot] {
					attInfo.Status = 0
				}

				if currentAttInfo.Status != 1 && attInfo.Status == 1 {
					currentAttInfo.Status = attInfo.Status
					currentAttInfo.InclusionSlot = attInfo.InclusionSlot
				}
			}

			missedSlotsCount := uint64(0)
			for slot := uint64(attesterSlot) + 1; slot < currentAttInfo.InclusionSlot; slot++ {
				if missedSlotsMap[slot] || orphanedSlotsMap[slot] {
					missedSlotsCount++
				}
			}
			currentAttInfo.Index = uint64(validator)
			currentAttInfo.Epoch = uint64(attesterSlot) / utils.Config.Chain.ClConfig.SlotsPerEpoch
			currentAttInfo.CommitteeIndex = 0
			currentAttInfo.AttesterSlot = uint64(attesterSlot)
			currentAttInfo.Delay = int64(currentAttInfo.InclusionSlot - uint64(attesterSlot) - missedSlotsCount - 1)

			res[uint64(validator)] = append(res[uint64(validator)], currentAttInfo)
		}
	}

	// Sort the result by attesterSlot desc
	for validator, att := range res {
		sort.Slice(att, func(i, j int) bool {
			return att[i].AttesterSlot > att[j].AttesterSlot
		})
		res[validator] = att
	}

	return res, nil
}

func (bigtable *Bigtable) GetLastAttestationSlots(validators []uint64) (map[uint64]uint64, error) {

	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"validatorsCount": len(validators),
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	valLen := len(validators)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*5))
	defer cancel()

	res := make(map[uint64]uint64, len(validators))

	columnFilters := []gcp_bigtable.Filter{}
	if valLen < 1000 {
		columnFilters = make([]gcp_bigtable.Filter, 0, len(validators))
		for _, validator := range validators {
			columnFilters = append(columnFilters, gcp_bigtable.ColumnFilter(fmt.Sprintf("%d", validator)))
		}
	}

	filter := gcp_bigtable.ChainFilters(
		gcp_bigtable.FamilyFilter(ATTESTATIONS_FAMILY),
		gcp_bigtable.InterleaveFilters(columnFilters...),
		gcp_bigtable.LatestNFilter(1),
	)

	if len(columnFilters) == 1 { // special case to retrieve data for one validators
		filter = gcp_bigtable.ChainFilters(
			gcp_bigtable.FamilyFilter(ATTESTATIONS_FAMILY),
			columnFilters[0],
			gcp_bigtable.LatestNFilter(1),
		)
	} else if len(columnFilters) == 0 { // special case to retrieve data for all validators
		filter = gcp_bigtable.ChainFilters(
			gcp_bigtable.FamilyFilter(ATTESTATIONS_FAMILY),
			gcp_bigtable.LatestNFilter(1),
		)
	}

	key := fmt.Sprintf("%s:lastAttestationSlot", bigtable.chainId)

	row, err := bigtable.tableValidators.ReadRow(ctx, key, gcp_bigtable.RowFilter(filter))
	if err != nil {
		return nil, err
	}

	for _, ri := range row[ATTESTATIONS_FAMILY] {
		attestedSlot := uint64(ri.Timestamp) / 1000

		validator, err := strconv.ParseUint(strings.TrimPrefix(ri.Column, ATTESTATIONS_FAMILY+":"), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing validator from column key %v: %v", ri.Column, err)
		}

		res[validator] = attestedSlot
	}

	return res, nil
}

func (bigtable *Bigtable) GetValidatorMissedAttestationHistory(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64]map[uint64]bool, error) {
	return bigtable.getValidatorMissedAttestationHistoryV2(validators, startEpoch, endEpoch)
}

func (bigtable *Bigtable) getValidatorMissedAttestationHistoryV2(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64]map[uint64]bool, error) {

	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"validatorsCount": len(validators),
			"startEpoch":      startEpoch,
			"endEpoch":        endEpoch,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	if len(validators) == 0 {
		return nil, fmt.Errorf("passing empty validator array is unsupported")
	}

	batchSize := 1000
	concurrency := 10

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*20))
	defer cancel()

	slots := []uint64{}

	for slot := startEpoch * utils.Config.Chain.ClConfig.SlotsPerEpoch; slot < (endEpoch+1)*utils.Config.Chain.ClConfig.SlotsPerEpoch; slot++ {
		slots = append(slots, slot)
	}
	orphanedSlotsMap, err := GetOrphanedSlotsMap(slots)
	if err != nil {
		return nil, err
	}

	res := make(map[uint64]map[uint64]bool)
	foundValid := make(map[uint64]map[uint64]bool)

	resMux := &sync.Mutex{}

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)

	for i := 0; i < len(validators); i += batchSize {

		upperBound := i + batchSize
		if len(validators) < upperBound {
			upperBound = len(validators)
		}
		vals := validators[i:upperBound]

		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
			}
			ranges := bigtable.getValidatorsEpochRanges(vals, ATTESTATIONS_FAMILY, startEpoch, endEpoch)

			filter := gcp_bigtable.LimitRows(int64(endEpoch-startEpoch+1) * int64(len(vals))) // max is one row per epoch

			err = bigtable.tableValidatorsHistory.ReadRows(ctx, ranges, func(r gcp_bigtable.Row) bool {
				keySplit := strings.Split(r.Key(), ":")

				validator, err := bigtable.validatorKeyToIndex(keySplit[1])
				if err != nil {
					logger.Errorf("error parsing validator from row key %v: %v", r.Key(), err)
					return false
				}

				for _, ri := range r[ATTESTATIONS_FAMILY] {
					attesterSlotString := strings.Replace(ri.Column, ATTESTATIONS_FAMILY+":", "", 1)
					attesterSlot, err := strconv.ParseUint(attesterSlotString, 10, 64)
					if err != nil {
						logger.Errorf("error parsing slot from row key %v: %v", r.Key(), err)
						return false
					}

					inclusionSlot := MAX_CL_BLOCK_NUMBER - uint64(ri.Timestamp)/1000

					status := uint64(1)
					if inclusionSlot == MAX_CL_BLOCK_NUMBER {
						status = 0
					}

					resMux.Lock()
					// only if the attestation was not included in another slot we count it as missed
					if (status == 0 || orphanedSlotsMap[inclusionSlot]) && (foundValid[validator] == nil || !foundValid[validator][attesterSlot]) {
						if res[validator] == nil {
							res[validator] = make(map[uint64]bool, 0)
						}
						res[validator][attesterSlot] = true
					} else {
						if res[validator] != nil && res[validator][attesterSlot] {
							delete(res[validator], attesterSlot)
						}
						if foundValid[validator] == nil {
							foundValid[validator] = make(map[uint64]bool, 0)
						}
						foundValid[validator][attesterSlot] = true
					}
					resMux.Unlock()
				}
				return true
			}, filter)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}

// GetValidatorSyncDutiesHistory returns the sync participation status for the given validators ranging from startSlot to endSlot (both inclusive)
//
// The returned map uses the following keys: [validatorIndex][slot]
func (bigtable *Bigtable) GetValidatorSyncDutiesHistory(validators []uint64, startSlot uint64, endSlot uint64) (map[uint64]map[uint64]*types.ValidatorSyncParticipation, error) {
	return bigtable.getValidatorSyncDutiesHistoryV2(validators, startSlot, endSlot)

}

func (bigtable *Bigtable) getValidatorSyncDutiesHistoryV2(validators []uint64, startSlot uint64, endSlot uint64) (map[uint64]map[uint64]*types.ValidatorSyncParticipation, error) {

	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"validatorsCount": len(validators),
			"startSlot":       startSlot,
			"endSlot":         endSlot,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	if len(validators) == 0 {
		return nil, fmt.Errorf("passing empty validator array is unsupported")
	}

	batchSize := 1000
	concurrency := 10

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*20))
	defer cancel()

	res := make(map[uint64]map[uint64]*types.ValidatorSyncParticipation, len(validators))
	resMux := &sync.Mutex{}

	filter := gcp_bigtable.LatestNFilter(1)

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)

	for i := 0; i < len(validators); i += batchSize {

		i := i
		upperBound := i + batchSize
		if len(validators) < upperBound {
			upperBound = len(validators)
		}
		vals := validators[i:upperBound]

		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
			}
			ranges := bigtable.getValidatorSlotRanges(vals, SYNC_COMMITTEES_FAMILY, startSlot, endSlot)

			err := bigtable.tableValidatorsHistory.ReadRows(ctx, ranges, func(r gcp_bigtable.Row) bool {
				keySplit := strings.Split(r.Key(), ":")

				validator, err := bigtable.validatorKeyToIndex(keySplit[1])
				if err != nil {
					logger.Errorf("error parsing validator from row key %v: %v", r.Key(), err)
					return false
				}
				slot, err := strconv.ParseUint(keySplit[4], 10, 64)
				if err != nil {
					logger.Errorf("error parsing slot from row key %v: %v", r.Key(), err)
					return false
				}
				slot = MAX_CL_BLOCK_NUMBER - slot

				for _, ri := range r[SYNC_COMMITTEES_FAMILY] {

					inclusionSlot := MAX_CL_BLOCK_NUMBER - uint64(ri.Timestamp)/1000

					status := uint64(1) // 1: participated
					if inclusionSlot == MAX_CL_BLOCK_NUMBER {
						inclusionSlot = 0
						status = 0 // 0: missed
					}

					resMux.Lock()
					if res[validator] == nil {
						res[validator] = make(map[uint64]*types.ValidatorSyncParticipation, 0)
					}

					if len(res[validator]) > 0 && res[validator][slot] != nil {
						res[validator][slot].Status = status
					} else {
						res[validator][slot] = &types.ValidatorSyncParticipation{
							Slot:   slot,
							Status: status,
						}
					}
					resMux.Unlock()

				}
				return true
			}, gcp_bigtable.RowFilter(filter))

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}

func (bigtable *Bigtable) GetValidatorMissedAttestationsCount(validators []uint64, firstEpoch uint64, lastEpoch uint64) (map[uint64]*types.ValidatorMissedAttestationsStatistic, error) {

	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"validatorsCount": len(validators),
			"startEpoch":      firstEpoch,
			"endEpoch":        lastEpoch,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	if firstEpoch > lastEpoch {
		return nil, fmt.Errorf("GetValidatorMissedAttestationsCount received an invalid firstEpoch (%d) and lastEpoch (%d) combination", firstEpoch, lastEpoch)
	}

	res := make(map[uint64]*types.ValidatorMissedAttestationsStatistic)

	data, err := bigtable.GetValidatorMissedAttestationHistory(validators, firstEpoch, lastEpoch)

	if err != nil {
		return nil, err
	}

	// logger.Infof("retrieved missed attestation history for epochs %v - %v", firstEpoch, lastEpoch)

	for validator, attestations := range data {
		if len(attestations) == 0 {
			continue
		}
		res[validator] = &types.ValidatorMissedAttestationsStatistic{
			Index:              validator,
			MissedAttestations: uint64(len(attestations)),
		}
	}

	return res, nil
}

func (bigtable *Bigtable) GetValidatorSyncDutiesStatistics(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64]*types.ValidatorSyncDutiesStatistic, error) {

	data, err := bigtable.GetValidatorSyncDutiesHistory(validators, startEpoch*utils.Config.Chain.ClConfig.SlotsPerEpoch, ((endEpoch+1)*utils.Config.Chain.ClConfig.SlotsPerEpoch)-1)

	if err != nil {
		return nil, err
	}

	slotsMap := make(map[uint64]bool)
	for _, duties := range data {
		for _, duty := range duties {
			slotsMap[duty.Slot] = true
		}
	}
	slots := []uint64{}
	for slot := range slotsMap {
		slots = append(slots, slot)
	}

	orphanedSlots, err := GetOrphanedSlots(slots)
	if err != nil {
		return nil, err
	}

	orphanedSlotsMap := make(map[uint64]bool)
	for _, slot := range orphanedSlots {
		orphanedSlotsMap[slot] = true
	}

	res := make(map[uint64]*types.ValidatorSyncDutiesStatistic)

	for validator, duties := range data {
		if res[validator] == nil && len(duties) > 0 {
			res[validator] = &types.ValidatorSyncDutiesStatistic{
				Index: validator,
			}
		}

		for _, duty := range duties {
			if orphanedSlotsMap[duty.Slot] {
				res[validator].OrphanedSync++
			} else if duty.Status == 0 {
				res[validator].MissedSync++
			} else {
				res[validator].ParticipatedSync++
			}
		}
	}

	return res, nil
}

// returns the validator attestation effectiveness in %
func (bigtable *Bigtable) GetValidatorEffectiveness(validators []uint64, epoch uint64) ([]*types.ValidatorEffectiveness, error) {
	end := epoch
	start := uint64(0)
	lookback := uint64(99)
	if end > lookback {
		start = end - lookback
	}
	data, err := bigtable.GetValidatorAttestationHistory(validators, start, end)

	if err != nil {
		return nil, err
	}

	res := make([]*types.ValidatorEffectiveness, 0, len(validators))
	type readings struct {
		Count uint64
		Sum   float64
	}

	aggEffectiveness := make(map[uint64]*readings)

	for validator, history := range data {
		for _, attestation := range history {
			if aggEffectiveness[validator] == nil {
				aggEffectiveness[validator] = &readings{}
			}
			if attestation.InclusionSlot > 0 {
				// logger.Infof("adding %v for epoch %v %.2f%%", attestation.InclusionSlot, attestation.AttesterSlot, 1.0/float64(attestation.InclusionSlot-attestation.AttesterSlot)*100)
				aggEffectiveness[validator].Sum += 1.0 / float64(attestation.InclusionSlot-attestation.AttesterSlot)
				aggEffectiveness[validator].Count++
			} else {
				aggEffectiveness[validator].Sum += 0 // missed attestations get a penalty of 32 slots // TODO(rgeraldes24)
				aggEffectiveness[validator].Count++
			}
		}
	}
	for validator, reading := range aggEffectiveness {
		res = append(res, &types.ValidatorEffectiveness{
			Validatorindex:        validator,
			AttestationEfficiency: float64(reading.Sum) / float64(reading.Count) * 100,
		})
	}

	return res, nil
}

func (bigtable *Bigtable) GetValidatorBalanceStatistics(validators []uint64, startEpoch, endEpoch uint64) (map[uint64]*types.ValidatorBalanceStatistic, error) {

	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"validatorsCount": len(validators),
			"startEpoch":      startEpoch,
			"endEpoch":        endEpoch,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	type ResultContainer struct {
		mu  sync.Mutex
		res map[uint64]*types.ValidatorBalanceStatistic
	}
	resultContainer := ResultContainer{}
	resultContainer.res = make(map[uint64]*types.ValidatorBalanceStatistic)

	// g, gCtx := errgroup.WithContext(ctx)
	batchSize := 10000
	// g.SetLimit(1)
	for i := 0; i < len(validators); i += batchSize {

		upperBound := i + batchSize
		if len(validators) < upperBound {
			upperBound = len(validators)
		}
		vals := validators[i:upperBound]

		// logrus.Infof("retrieving validator balance stats for validators %v - %v", vals[0], vals[len(vals)-1])

		res, err := bigtable.GetValidatorBalanceHistory(vals, startEpoch, endEpoch)
		if err != nil {
			return nil, err
		}
		resultContainer.mu.Lock()
		for validator, balances := range res {
			for _, balance := range balances {
				if resultContainer.res[validator] == nil {
					resultContainer.res[validator] = &types.ValidatorBalanceStatistic{
						Index:                 validator,
						MinEffectiveBalance:   balance.EffectiveBalance,
						MaxEffectiveBalance:   0,
						MinBalance:            balance.Balance,
						MaxBalance:            0,
						StartEffectiveBalance: 0,
						EndEffectiveBalance:   0,
						StartBalance:          0,
						EndBalance:            0,
					}
				}

				if balance.Epoch == startEpoch {
					resultContainer.res[validator].StartBalance = balance.Balance
					resultContainer.res[validator].StartEffectiveBalance = balance.EffectiveBalance
				}

				if balance.Epoch == endEpoch {
					resultContainer.res[validator].EndBalance = balance.Balance
					resultContainer.res[validator].EndEffectiveBalance = balance.EffectiveBalance
				}

				if balance.Balance > resultContainer.res[validator].MaxBalance {
					resultContainer.res[validator].MaxBalance = balance.Balance
				}
				if balance.Balance < resultContainer.res[validator].MinBalance {
					resultContainer.res[validator].MinBalance = balance.Balance
				}

				if balance.EffectiveBalance > resultContainer.res[validator].MaxEffectiveBalance {
					resultContainer.res[validator].MaxEffectiveBalance = balance.EffectiveBalance
				}
				if balance.EffectiveBalance < resultContainer.res[validator].MinEffectiveBalance {
					resultContainer.res[validator].MinEffectiveBalance = balance.EffectiveBalance
				}
			}
		}

		resultContainer.mu.Unlock()

	}

	return resultContainer.res, nil
}

func (bigtable *Bigtable) GetValidatorProposalHistory(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64][]*types.ValidatorProposal, error) {
	return bigtable.getValidatorProposalHistoryV2(validators, startEpoch, endEpoch)
}

func (bigtable *Bigtable) getValidatorProposalHistoryV2(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64][]*types.ValidatorProposal, error) {
	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"validatorsCount": len(validators),
			"startEpoch":      startEpoch,
			"endEpoch":        endEpoch,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	if len(validators) == 0 {
		return nil, fmt.Errorf("passing empty validator array is unsupported")
	}

	batchSize := 1000
	concurrency := 10

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
	defer cancel()

	res := make(map[uint64][]*types.ValidatorProposal, len(validators))
	resMux := &sync.Mutex{}

	filter := gcp_bigtable.LatestNFilter(1)

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)

	for i := 0; i < len(validators); i += batchSize {

		upperBound := i + batchSize
		if len(validators) < upperBound {
			upperBound = len(validators)
		}
		vals := validators[i:upperBound]

		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
			}
			ranges := bigtable.getValidatorsEpochSlotRanges(vals, PROPOSALS_FAMILY, startEpoch, endEpoch)
			err := bigtable.tableValidatorsHistory.ReadRows(ctx, ranges, func(r gcp_bigtable.Row) bool {
				for _, ri := range r[PROPOSALS_FAMILY] {
					keySplit := strings.Split(r.Key(), ":")

					proposalSlot, err := strconv.ParseUint(keySplit[4], 10, 64)
					if err != nil {
						logger.Errorf("error parsing slot from row key %v: %v", r.Key(), err)
						return false
					}
					proposalSlot = MAX_CL_BLOCK_NUMBER - proposalSlot
					inclusionSlot := MAX_CL_BLOCK_NUMBER - uint64(r[PROPOSALS_FAMILY][0].Timestamp)/1000

					status := uint64(1)
					if inclusionSlot == MAX_CL_BLOCK_NUMBER {
						inclusionSlot = 0
						status = 2
					}

					validator, err := bigtable.validatorKeyToIndex(keySplit[1])
					if err != nil {
						logger.Errorf("error parsing validator from column key %v: %v", ri.Column, err)
						return false
					}

					resMux.Lock()
					if res[validator] == nil {
						res[validator] = make([]*types.ValidatorProposal, 0)
					}

					if len(res[validator]) > 0 && res[validator][len(res[validator])-1].Slot == proposalSlot {
						res[validator][len(res[validator])-1].Slot = proposalSlot
						res[validator][len(res[validator])-1].Status = status
					} else {
						res[validator] = append(res[validator], &types.ValidatorProposal{
							Index:  validator,
							Status: status,
							Slot:   proposalSlot,
						})
					}
					resMux.Unlock()

				}
				return true
			}, gcp_bigtable.RowFilter(filter))

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}

func (bigtable *Bigtable) SaveValidatorIncomeDetails(epoch uint64, rewards map[uint64]*itypes.ValidatorEpochIncome) error {

	start := time.Now()
	ts := gcp_bigtable.Timestamp(utils.EpochToTime(epoch).UnixMicro())

	total := &itypes.ValidatorEpochIncome{}

	muts := types.NewBulkMutations(len(rewards))

	for i, rewardDetails := range rewards {

		data, err := proto.Marshal(rewardDetails)

		if err != nil {
			return err
		}

		mut := &gcp_bigtable.Mutation{}
		mut.Set(INCOME_DETAILS_COLUMN_FAMILY, "i", ts, data)
		key := fmt.Sprintf("%s:%s:%s:%s", bigtable.chainId, bigtable.validatorIndexToKey(i), INCOME_DETAILS_COLUMN_FAMILY, bigtable.reversedPaddedEpoch(epoch))

		muts.Add(key, mut)

		total.AttestationHeadReward += rewardDetails.AttestationHeadReward
		total.AttestationSourceReward += rewardDetails.AttestationSourceReward
		total.AttestationSourcePenalty += rewardDetails.AttestationSourcePenalty
		total.AttestationTargetReward += rewardDetails.AttestationTargetReward
		total.AttestationTargetPenalty += rewardDetails.AttestationTargetPenalty
		total.FinalityDelayPenalty += rewardDetails.FinalityDelayPenalty
		total.ProposerSlashingInclusionReward += rewardDetails.ProposerSlashingInclusionReward
		total.ProposerAttestationInclusionReward += rewardDetails.ProposerAttestationInclusionReward
		total.ProposerSyncInclusionReward += rewardDetails.ProposerSyncInclusionReward
		total.SyncCommitteeReward += rewardDetails.SyncCommitteeReward
		total.SyncCommitteePenalty += rewardDetails.SyncCommitteePenalty
		total.SlashingReward += rewardDetails.SlashingReward
		total.SlashingPenalty += rewardDetails.SlashingPenalty
		total.TxFeeRewardPlanck = utils.AddBigInts(total.TxFeeRewardPlanck, rewardDetails.TxFeeRewardPlanck)
	}

	sum, err := proto.Marshal(total)
	if err != nil {
		return err
	}

	mut := &gcp_bigtable.Mutation{}
	mut.Set(STATS_COLUMN_FAMILY, SUM_COLUMN, ts, sum)

	muts.Add(fmt.Sprintf("%s:%s:%s", bigtable.chainId, SUM_COLUMN, bigtable.reversedPaddedEpoch(epoch)), mut)

	err = bigtable.WriteBulk(muts, bigtable.tableValidatorsHistory, MAX_BATCH_MUTATIONS)

	if err != nil {
		return err
	}

	logger.Infof("exported validator income details for epoch %v to bigtable in %v", epoch, time.Since(start))
	return nil
}

// GetValidatorIncomeDetailsHistory returns the validator income details
// startEpoch & endEpoch are inclusive
func (bigtable *Bigtable) GetValidatorIncomeDetailsHistory(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64]map[uint64]*itypes.ValidatorEpochIncome, error) {
	return bigtable.getValidatorIncomeDetailsHistoryV2(validators, startEpoch, endEpoch)
}

func (bigtable *Bigtable) getValidatorIncomeDetailsHistoryV2(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64]map[uint64]*itypes.ValidatorEpochIncome, error) {

	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"validatorsCount": len(validators),
			"startEpoch":      startEpoch,
			"endEpoch":        endEpoch,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	if len(validators) == 0 {
		return nil, fmt.Errorf("passing empty validator array is unsupported")
	}

	batchSize := 1000
	concurrency := 10

	if startEpoch > endEpoch {
		startEpoch = 0
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	res := make(map[uint64]map[uint64]*itypes.ValidatorEpochIncome, len(validators))
	resMux := &sync.Mutex{}

	filter := gcp_bigtable.LatestNFilter(1)

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)

	for i := 0; i < len(validators); i += batchSize {

		upperBound := i + batchSize
		if len(validators) < upperBound {
			upperBound = len(validators)
		}
		vals := validators[i:upperBound]

		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
			}
			ranges := bigtable.getValidatorsEpochRanges(vals, INCOME_DETAILS_COLUMN_FAMILY, startEpoch, endEpoch)
			err := bigtable.tableValidatorsHistory.ReadRows(ctx, ranges, func(r gcp_bigtable.Row) bool {
				keySplit := strings.Split(r.Key(), ":")

				validator, err := bigtable.validatorKeyToIndex(keySplit[1])
				if err != nil {
					logger.Errorf("error parsing validator from row key %v: %v", r.Key(), err)
					return false
				}

				epoch, err := strconv.ParseUint(keySplit[3], 10, 64)
				if err != nil {
					logger.Errorf("error parsing epoch from row key %v: %v", r.Key(), err)
					return false
				}

				for _, ri := range r[INCOME_DETAILS_COLUMN_FAMILY] {
					incomeDetails := &itypes.ValidatorEpochIncome{}
					err = proto.Unmarshal(ri.Value, incomeDetails)
					if err != nil {
						logger.Errorf("error decoding validator income data for row %v: %v", r.Key(), err)
						return false
					}

					resMux.Lock()
					if res[validator] == nil {
						res[validator] = make(map[uint64]*itypes.ValidatorEpochIncome)
					}

					res[validator][MAX_EPOCH-epoch] = incomeDetails
					resMux.Unlock()
				}
				return true
			}, gcp_bigtable.RowFilter(filter))

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}

// GetAggregatedValidatorIncomeDetailsHistory returns aggregated validator income details
// startEpoch & endEpoch are inclusive
func (bigtable *Bigtable) GetAggregatedValidatorIncomeDetailsHistory(validators []uint64, startEpoch uint64, endEpoch uint64) (map[uint64]*itypes.ValidatorEpochIncome, error) {
	if startEpoch > endEpoch {
		startEpoch = 0
	}

	type ResultContainer struct {
		mu  sync.Mutex
		res map[uint64]*itypes.ValidatorEpochIncome
	}
	resultContainer := ResultContainer{}
	resultContainer.res = make(map[uint64]*itypes.ValidatorEpochIncome, len(validators))

	batchSize := 10000
	for i := 0; i < len(validators); i += batchSize {

		upperBound := i + batchSize
		if len(validators) < upperBound {
			upperBound = len(validators)
		}
		vals := validators[i:upperBound]

		logrus.Infof("retrieving validator income stats for validators %v - %v", vals[0], vals[len(vals)-1])

		res, err := bigtable.GetValidatorIncomeDetailsHistory(vals, startEpoch, endEpoch)

		if err != nil {
			return nil, err
		}
		resultContainer.mu.Lock()
		for validator, epochs := range res {
			for _, rewardDetails := range epochs {

				if resultContainer.res[validator] == nil {
					resultContainer.res[validator] = &itypes.ValidatorEpochIncome{}
				}

				resultContainer.res[validator].AttestationHeadReward += rewardDetails.AttestationHeadReward
				resultContainer.res[validator].AttestationSourceReward += rewardDetails.AttestationSourceReward
				resultContainer.res[validator].AttestationSourcePenalty += rewardDetails.AttestationSourcePenalty
				resultContainer.res[validator].AttestationTargetReward += rewardDetails.AttestationTargetReward
				resultContainer.res[validator].AttestationTargetPenalty += rewardDetails.AttestationTargetPenalty
				resultContainer.res[validator].FinalityDelayPenalty += rewardDetails.FinalityDelayPenalty
				resultContainer.res[validator].ProposerSlashingInclusionReward += rewardDetails.ProposerSlashingInclusionReward
				resultContainer.res[validator].ProposerAttestationInclusionReward += rewardDetails.ProposerAttestationInclusionReward
				resultContainer.res[validator].ProposerSyncInclusionReward += rewardDetails.ProposerSyncInclusionReward
				resultContainer.res[validator].SyncCommitteeReward += rewardDetails.SyncCommitteeReward
				resultContainer.res[validator].SyncCommitteePenalty += rewardDetails.SyncCommitteePenalty
				resultContainer.res[validator].SlashingReward += rewardDetails.SlashingReward
				resultContainer.res[validator].SlashingPenalty += rewardDetails.SlashingPenalty
				resultContainer.res[validator].TxFeeRewardPlanck = utils.AddBigInts(resultContainer.res[validator].TxFeeRewardPlanck, rewardDetails.TxFeeRewardPlanck)
			}
		}
		resultContainer.mu.Unlock()
	}

	return resultContainer.res, nil
}

// GetTotalValidatorIncomeDetailsHistory returns the total validator income for a given range of epochs
// It is considerably faster than fetching the individual income for each validator and aggregating it
// startEpoch & endEpoch are inclusive
func (bigtable *Bigtable) GetTotalValidatorIncomeDetailsHistory(startEpoch uint64, endEpoch uint64) (map[uint64]*itypes.ValidatorEpochIncome, error) {
	tmr := time.AfterFunc(REPORT_TIMEOUT, func() {
		logger.WithFields(logrus.Fields{
			"startEpoch": startEpoch,
			"endEpoch":   endEpoch,
		}).Warnf("%s call took longer than %v", utils.GetCurrentFuncName(), REPORT_TIMEOUT)
	})
	defer tmr.Stop()

	if startEpoch > endEpoch {
		startEpoch = 0
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	res := make(map[uint64]*itypes.ValidatorEpochIncome, endEpoch-startEpoch+1)

	filter := gcp_bigtable.LimitRows(int64(endEpoch - startEpoch + 1))

	rowRange := bigtable.getTotalIncomeEpochRanges(startEpoch, endEpoch)
	err := bigtable.tableValidatorsHistory.ReadRows(ctx, rowRange, func(r gcp_bigtable.Row) bool {
		keySplit := strings.Split(r.Key(), ":")

		epoch, err := strconv.ParseUint(keySplit[2], 10, 64)
		if err != nil {
			logger.Errorf("error parsing epoch from row key %v: %v", r.Key(), err)
			return false
		}

		for _, ri := range r[STATS_COLUMN_FAMILY] {
			incomeDetails := &itypes.ValidatorEpochIncome{}
			err = proto.Unmarshal(ri.Value, incomeDetails)
			if err != nil {
				logger.Errorf("error decoding validator income data for row %v: %v", r.Key(), err)
				return false
			}

			res[MAX_EPOCH-epoch] = incomeDetails
		}
		return true
	}, filter)

	if err != nil {
		return nil, err
	}
	return res, nil
}

// Deletes all block data from bigtable
func (bigtable *Bigtable) DeleteEpoch(epoch uint64) error {
	// TOTO: Implement
	return fmt.Errorf("NOT IMPLEMENTED")
}

func (bigtable *Bigtable) getValidatorsEpochRanges(validatorIndices []uint64, prefix string, startEpoch uint64, endEpoch uint64) gcp_bigtable.RowRangeList {
	if endEpoch > math.MaxInt64 {
		endEpoch = 0
	}
	if endEpoch < startEpoch { // handle overflows
		startEpoch = 0
	}

	ranges := make(gcp_bigtable.RowRangeList, 0, int((endEpoch-startEpoch+1))*len(validatorIndices))

	for _, validatorIndex := range validatorIndices {
		validatorKey := bigtable.validatorIndexToKey(validatorIndex)

		// epochs are sorted descending, so start with the largest epoch and end with the smallest
		// add \x00 to make the range inclusive
		rangeEnd := fmt.Sprintf("%s:%s:%s:%s%s", bigtable.chainId, validatorKey, prefix, bigtable.reversedPaddedEpoch(startEpoch), "\x00")
		rangeStart := fmt.Sprintf("%s:%s:%s:%s", bigtable.chainId, validatorKey, prefix, bigtable.reversedPaddedEpoch(endEpoch))
		ranges = append(ranges, gcp_bigtable.NewRange(rangeStart, rangeEnd))
	}
	return ranges
}

func (bigtable *Bigtable) getTotalIncomeEpochRanges(startEpoch uint64, endEpoch uint64) gcp_bigtable.RowRange {
	if endEpoch > math.MaxInt64 {
		endEpoch = 0
	}
	if endEpoch < startEpoch { // handle overflows
		startEpoch = 0
	}

	rangeEnd := fmt.Sprintf("%s:%s:%s%s", bigtable.chainId, SUM_COLUMN, bigtable.reversedPaddedEpoch(startEpoch), "\x00")
	rangeStart := fmt.Sprintf("%s:%s:%s", bigtable.chainId, SUM_COLUMN, bigtable.reversedPaddedEpoch(endEpoch))

	return gcp_bigtable.NewRange(rangeStart, rangeEnd)
}

func (bigtable *Bigtable) getValidatorsEpochSlotRanges(validatorIndices []uint64, prefix string, startEpoch uint64, endEpoch uint64) gcp_bigtable.RowRangeList {

	if endEpoch > math.MaxInt64 {
		endEpoch = 0
	}
	if endEpoch < startEpoch { // handle overflows
		startEpoch = 0
	}

	ranges := make(gcp_bigtable.RowRangeList, 0, int((endEpoch-startEpoch+1))*len(validatorIndices))

	for _, validatorIndex := range validatorIndices {
		validatorKey := bigtable.validatorIndexToKey(validatorIndex)

		rangeEnd := fmt.Sprintf("%s:%s:%s:%s:%s%s", bigtable.chainId, validatorKey, prefix, bigtable.reversedPaddedEpoch(startEpoch), bigtable.reversedPaddedSlot(startEpoch*utils.Config.Chain.ClConfig.SlotsPerEpoch), "\x00")
		rangeStart := fmt.Sprintf("%s:%s:%s:%s:%s", bigtable.chainId, validatorKey, prefix, bigtable.reversedPaddedEpoch(endEpoch), bigtable.reversedPaddedSlot(endEpoch*utils.Config.Chain.ClConfig.SlotsPerEpoch+utils.Config.Chain.ClConfig.SlotsPerEpoch-1))
		ranges = append(ranges, gcp_bigtable.NewRange(rangeStart, rangeEnd))

	}
	return ranges
}

func (bigtable *Bigtable) getValidatorSlotRanges(validatorIndices []uint64, prefix string, startSlot uint64, endSlot uint64) gcp_bigtable.RowRangeList {
	if endSlot > math.MaxInt64 {
		endSlot = 0
	}
	if endSlot < startSlot { // handle overflows
		startSlot = 0
	}

	startEpoch := utils.EpochOfSlot(startSlot)
	endEpoch := utils.EpochOfSlot(endSlot)

	ranges := make(gcp_bigtable.RowRangeList, 0, len(validatorIndices))

	for _, validatorIndex := range validatorIndices {
		validatorKey := bigtable.validatorIndexToKey(validatorIndex)

		rangeEnd := fmt.Sprintf("%s:%s:%s:%s:%s%s", bigtable.chainId, validatorKey, prefix, bigtable.reversedPaddedEpoch(startEpoch), bigtable.reversedPaddedSlot(startSlot), "\x00")
		rangeStart := fmt.Sprintf("%s:%s:%s:%s:%s", bigtable.chainId, validatorKey, prefix, bigtable.reversedPaddedEpoch(endEpoch), bigtable.reversedPaddedSlot(endSlot))
		ranges = append(ranges, gcp_bigtable.NewRange(rangeStart, rangeEnd))

	}
	return ranges
}

func (bigtable *Bigtable) validatorIndexToKey(index uint64) string {
	return utils.ReverseString(fmt.Sprintf("%d", index))
}

func (bigtable *Bigtable) validatorKeyToIndex(key string) (uint64, error) {
	key = utils.ReverseString(key)
	indexKey, err := strconv.ParseUint(key, 10, 64)

	if err != nil {
		return 0, err
	}
	return indexKey, nil
}

func GetCurrentDayClIncome(validator_indices []uint64) (map[uint64]int64, error) {
	dayIncome := make(map[uint64]int64)
	lastDay, err := GetLastExportedStatisticDay()
	if err != nil {
		if err == ErrNoStats {
			return dayIncome, nil
		}
		return dayIncome, err
	}

	currentDay := uint64(lastDay + 1)
	startEpoch := currentDay * utils.EpochsPerDay()
	endEpoch := startEpoch + utils.EpochsPerDay() - 1
	income, err := BigtableClient.GetValidatorIncomeDetailsHistory(validator_indices, startEpoch, endEpoch)
	if err != nil {
		return dayIncome, err
	}

	// agregate all epoch income data to total day income for each validator
	for validatorIndex, validatorIncome := range income {
		if len(validatorIncome) == 0 {
			continue
		}
		for _, validatorEpochIncome := range validatorIncome {
			dayIncome[validatorIndex] += validatorEpochIncome.TotalClRewards()
		}
	}

	return dayIncome, nil
}

func (bigtable *Bigtable) reversePaddedUserID(userID uint64) string {
	return fmt.Sprintf("%09d", ^uint64(0)-userID)
}

func (bigtable *Bigtable) reversedPaddedEpoch(epoch uint64) string {
	return fmt.Sprintf("%09d", MAX_EPOCH-epoch)
}

func (bigtable *Bigtable) reversedPaddedSlot(slot uint64) string {
	return fmt.Sprintf("%09d", MAX_CL_BLOCK_NUMBER-slot)
}
