package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"sort"
	"syscall"
	"time"

	"github.com/lib/pq"
	"github.com/theQRL/zond-beaconchain-explorer/db"
	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/types"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"github.com/gorilla/mux"
	utilMath "github.com/protolambda/zrnt/eth2/util/math"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func GetValidatorOnlineThresholdSlot() uint64 {
	latestProposedSlot := services.LatestProposedSlot()
	threshold := utils.Config.Chain.ClConfig.SlotsPerEpoch * 2

	var validatorOnlineThresholdSlot uint64
	if latestProposedSlot < 1 || latestProposedSlot < threshold {
		validatorOnlineThresholdSlot = 0
	} else {
		validatorOnlineThresholdSlot = latestProposedSlot - threshold
	}

	return validatorOnlineThresholdSlot
}

// GetValidatorEarnings will return the earnings (last day, week, month and total) of selected validators, including proposal and statisic information - infused with data from the current day. all values are
func GetValidatorEarnings(validators []uint64) (*types.ValidatorEarnings, map[uint64]*types.Validator, error) {
	if len(validators) == 0 {
		return nil, nil, errors.New("no validators provided")
	}
	latestFinalizedEpoch := services.LatestFinalizedEpoch()

	firstSlot := uint64(0)
	lastStatsDay, lastExportedStatsErr := services.LatestExportedStatisticDay()
	if lastExportedStatsErr == nil {
		firstSlot = utils.GetLastBalanceInfoSlotForDay(lastStatsDay) + 1
	}

	lastSlot := latestFinalizedEpoch * utils.Config.Chain.ClConfig.SlotsPerEpoch

	balancesMap := make(map[uint64]*types.Validator, 0)
	totalBalance := uint64(0)

	g := errgroup.Group{}
	g.Go(func() error {
		latestBalances, err := db.BigtableClient.GetValidatorBalanceHistory(validators, latestFinalizedEpoch, latestFinalizedEpoch)
		if err != nil {
			logger.Errorf("error getting validator balance data in GetValidatorEarnings: %v", err)
			return err
		}

		for balanceIndex, balance := range latestBalances {
			if len(balance) == 0 {
				continue
			}

			if balancesMap[balanceIndex] == nil {
				balancesMap[balanceIndex] = &types.Validator{}
			}
			balancesMap[balanceIndex].Balance = balance[0].Balance
			balancesMap[balanceIndex].EffectiveBalance = balance[0].EffectiveBalance

			totalBalance += balance[0].Balance
		}
		return nil
	})

	income := types.ValidatorIncomePerformance{}
	g.Go(func() error {
		return db.GetValidatorIncomePerformance(validators, &income)
	})

	var totalDeposits uint64
	g.Go(func() error {
		return db.GetTotalValidatorDeposits(validators, &totalDeposits)
	})

	var firstActivationEpoch uint64
	g.Go(func() error {
		return db.GetFirstActivationEpoch(validators, &firstActivationEpoch)
	})

	var lastDeposits uint64
	var lastWithdrawals uint64
	var lastBalance uint64
	g.Go(func() error {
		if lastExportedStatsErr == db.ErrNoStats {
			err := db.GetValidatorActivationBalance(validators, &lastBalance)
			if err != nil {
				return err
			}
		} else {
			err := db.GetValidatorBalanceForDay(validators, lastStatsDay, &lastBalance)
			if err != nil {
				return err
			}
		}
		err := db.GetValidatorDepositsForSlots(validators, firstSlot, lastSlot, &lastDeposits)
		if err != nil {
			return err
		}
		return db.GetValidatorWithdrawalsForSlots(validators, firstSlot, lastSlot, &lastWithdrawals)
	})

	proposals := []types.ValidatorProposalInfo{}
	g.Go(func() error {
		return db.GetValidatorPropsosals(validators, &proposals)
	})

	err := g.Wait()
	if err != nil {
		return nil, nil, err
	}

	if totalDeposits == 0 {
		totalDeposits = utils.Config.Chain.ClConfig.MaxEffectiveBalance * uint64(len(validators))
	}

	clApr7d := income.ClIncomePlanck7d.DivRound(decimal.NewFromInt(1e9), 18).DivRound(decimal.NewFromInt(int64(totalDeposits)), 18).Mul(decimal.NewFromInt(365)).Div(decimal.NewFromInt(7)).InexactFloat64()
	if clApr7d < float64(-1) {
		clApr7d = float64(-1)
	}
	if math.IsNaN(clApr7d) {
		clApr7d = float64(0)
	}

	elApr7d := income.ElIncomePlanck7d.DivRound(decimal.NewFromInt(1e9), 18).DivRound(decimal.NewFromInt(int64(totalDeposits)), 18).Mul(decimal.NewFromInt(365)).Div(decimal.NewFromInt(7)).InexactFloat64()
	if elApr7d < float64(-1) {
		elApr7d = float64(-1)
	}
	if math.IsNaN(elApr7d) {
		elApr7d = float64(0)
	}

	clApr31d := income.ClIncomePlanck31d.DivRound(decimal.NewFromInt(1e9), 18).DivRound(decimal.NewFromInt(int64(totalDeposits)), 18).Mul(decimal.NewFromInt(365)).Div(decimal.NewFromInt(31)).InexactFloat64()
	if clApr31d < float64(-1) {
		clApr31d = float64(-1)
	}
	if math.IsNaN(clApr31d) {
		clApr31d = float64(0)
	}

	elApr31d := income.ElIncomePlanck31d.DivRound(decimal.NewFromInt(1e9), 18).DivRound(decimal.NewFromInt(int64(totalDeposits)), 18).Mul(decimal.NewFromInt(365)).Div(decimal.NewFromInt(31)).InexactFloat64()
	if elApr31d < float64(-1) {
		elApr31d = float64(-1)
	}
	if math.IsNaN(elApr31d) {
		elApr31d = float64(0)
	}

	clApr365d := income.ClIncomePlanck365d.DivRound(decimal.NewFromInt(1e9), 18).DivRound(decimal.NewFromInt(int64(totalDeposits)), 18).InexactFloat64()
	if clApr365d < float64(-1) {
		clApr365d = float64(-1)
	}
	if math.IsNaN(clApr365d) {
		clApr365d = float64(0)
	}

	elApr365d := income.ElIncomePlanck365d.DivRound(decimal.NewFromInt(1e9), 18).DivRound(decimal.NewFromInt(int64(totalDeposits)), 18).InexactFloat64()
	if elApr365d < float64(-1) {
		elApr365d = float64(-1)
	}
	if math.IsNaN(elApr365d) {
		elApr365d = float64(0)
	}

	proposedToday := []uint64{}
	todayStartEpoch := uint64(0)
	if lastExportedStatsErr == nil {
		todayStartEpoch = uint64(lastStatsDay+1) * utils.EpochsPerDay()
	}
	validatorProposalData := types.ValidatorProposalData{}
	validatorProposalData.Proposals = make([][]uint64, len(proposals))
	for i, b := range proposals {
		validatorProposalData.Proposals[i] = []uint64{
			uint64(utils.SlotToTime(b.Slot).Unix()),
			b.Status,
		}
		if b.Status == 0 {
			validatorProposalData.LastScheduledSlot = utilMath.MaxU64(validatorProposalData.LastScheduledSlot, b.Slot)
			validatorProposalData.ScheduledBlocksCount++
		} else if b.Status == 1 {
			validatorProposalData.ProposedBlocksCount++
			// add to list of blocks proposed today if epoch hasn't been exported into stats yet
			if utils.EpochOfSlot(b.Slot) >= todayStartEpoch && b.ExecBlockNumber.Int64 > 0 {
				proposedToday = append(proposedToday, uint64(b.ExecBlockNumber.Int64))
			}
		} else if b.Status == 2 {
			validatorProposalData.MissedBlocksCount++
		} else if b.Status == 3 {
			validatorProposalData.OrphanedBlocksCount++
		}
	}

	validatorProposalData.BlocksCount = uint64(len(proposals))
	if validatorProposalData.BlocksCount > 0 {
		validatorProposalData.UnmissedBlocksPercentage = float64(validatorProposalData.BlocksCount-validatorProposalData.MissedBlocksCount-validatorProposalData.OrphanedBlocksCount) / float64(validatorProposalData.BlocksCount)
	} else {
		validatorProposalData.UnmissedBlocksPercentage = 1.0
	}

	var slots []uint64
	for _, p := range proposals {
		if p.ExecBlockNumber.Int64 > 0 {
			slots = append(slots, p.Slot)
		}
	}

	validatorProposalData.ProposalLuck, _ = getProposalLuck(slots, len(validators), firstActivationEpoch)
	avgSlotInterval := uint64(getAvgSlotInterval(len(validators)))
	avgSlotIntervalAsDuration := time.Duration(utils.Config.Chain.ClConfig.SecondsPerSlot*avgSlotInterval) * time.Second
	validatorProposalData.AvgSlotInterval = &avgSlotIntervalAsDuration
	if len(slots) > 0 {
		nextSlotEstimate := utils.SlotToTime(slots[len(slots)-1] + avgSlotInterval)
		validatorProposalData.ProposalEstimate = &nextSlotEstimate
	}

	currentDayClIncome := decimal.NewFromInt(int64(totalBalance - lastBalance - lastDeposits + lastWithdrawals)).Mul(decimal.NewFromInt(1e9))
	incomeToday := types.ClEl{
		El:    decimal.NewFromInt(0),
		Cl:    currentDayClIncome,
		Total: currentDayClIncome,
	}
	if len(proposedToday) > 0 {
		// get el data
		execBlocks, err := db.BigtableClient.GetBlocksIndexedMultiple(proposedToday, 10000)
		if err != nil {
			return nil, nil, fmt.Errorf("error retrieving execution blocks data from bigtable: %v", err)
		}

		incomeTodayEl := new(big.Int)
		for _, execBlock := range execBlocks {

			blockEpoch := utils.TimeToEpoch(execBlock.Time.AsTime())
			if blockEpoch > int64(latestFinalizedEpoch) {
				continue
			}
			incomeTodayEl = new(big.Int).Add(incomeTodayEl, new(big.Int).SetBytes(execBlock.GetTxReward()))
		}
		incomeToday.El = decimal.NewFromBigInt(incomeTodayEl, 0)
		incomeToday.Total = incomeToday.Total.Add(incomeToday.El)
	}

	earnings := &types.ValidatorEarnings{
		IncomeToday: incomeToday,
		Income1d: types.ClEl{
			El:    income.ElIncomePlanck1d,
			Cl:    income.ClIncomePlanck1d,
			Total: income.ElIncomePlanck1d.Add(income.ClIncomePlanck1d),
		},
		Income7d: types.ClEl{
			El:    income.ElIncomePlanck7d,
			Cl:    income.ClIncomePlanck7d,
			Total: income.ElIncomePlanck7d.Add(income.ClIncomePlanck7d),
		},
		Income31d: types.ClEl{
			El:    income.ElIncomePlanck31d,
			Cl:    income.ClIncomePlanck31d,
			Total: income.ElIncomePlanck31d.Add(income.ClIncomePlanck31d),
		},
		IncomeTotal: types.ClEl{
			El:    income.ElIncomePlanckTotal.Add(incomeToday.El),
			Cl:    income.ClIncomePlanckTotal.Add(incomeToday.Cl),
			Total: income.ElIncomePlanckTotal.Add(incomeToday.El).Add(income.ClIncomePlanckTotal.Add(incomeToday.Cl)),
		},
		Apr7d: types.ClElFloat64{
			El:    elApr7d,
			Cl:    clApr7d,
			Total: clApr7d + elApr7d,
		},
		Apr31d: types.ClElFloat64{
			El:    elApr31d,
			Cl:    clApr31d,
			Total: clApr31d + elApr31d,
		},
		Apr365d: types.ClElFloat64{
			El:    elApr365d,
			Cl:    clApr365d,
			Total: clApr365d + elApr365d,
		},
		TotalDeposits: int64(totalDeposits),
		ProposalData:  validatorProposalData,
	}
	earnings.LastDayFormatted = utils.FormatIncomeClEl(earnings.Income1d, "ZND")
	earnings.LastWeekFormatted = utils.FormatIncomeClEl(earnings.Income7d, "ZND")
	earnings.LastMonthFormatted = utils.FormatIncomeClEl(earnings.Income31d, "ZND")
	earnings.TotalFormatted = utils.FormatIncomeClEl(earnings.IncomeTotal, "ZND")
	earnings.TotalBalance = "<b>" + utils.FormatClCurrency(totalBalance, "ZND", 5, true, false, false, false) + "</b>"
	return earnings, balancesMap, nil
}

// Timeframe constants
const fiveDays = utils.Day * 5
const oneWeek = utils.Week
const oneMonth = utils.Month
const sixWeeks = utils.Day * 45
const twoMonths = utils.Month * 2
const threeMonths = utils.Month * 3
const fourMonths = utils.Month * 4
const fiveMonths = utils.Month * 5
const sixMonths = utils.Month * 6
const year = utils.Year

// getProposalLuck calculates the luck of a given set of proposed blocks for a certain number of validators
// given the blocks proposed by the validators and the number of validators
//
// precondition: slots is sorted by ascending block number
func getProposalLuck(slots []uint64, validatorsCount int, fromEpoch uint64) (float64, time.Duration) {
	// Return 0 if there are no proposed blocks or no validators
	if len(slots) == 0 || validatorsCount == 0 {
		return 0, 0
	}

	activeValidatorsCount := *services.GetLatestStats().ActiveValidatorCount
	// Calculate the expected number of slot proposals for 30 days
	expectedSlotProposals := calcExpectedSlotProposals(oneMonth, validatorsCount, activeValidatorsCount)

	// Get the timeframe for which we should consider qualified proposals
	var proposalTimeFrame time.Duration
	// Time since the first epoch of the related validators
	timeSinceFirstEpoch := time.Since(utils.EpochToTime(fromEpoch))

	targetBlocks := 8.0

	// Determine the appropriate timeframe based on the time since the first block and the expected slot proposals
	switch {
	case timeSinceFirstEpoch < fiveDays:
		proposalTimeFrame = fiveDays
	case timeSinceFirstEpoch < oneWeek:
		proposalTimeFrame = oneWeek
	case timeSinceFirstEpoch < oneMonth:
		proposalTimeFrame = oneMonth
	case timeSinceFirstEpoch > year && expectedSlotProposals <= targetBlocks/12:
		proposalTimeFrame = year
	case timeSinceFirstEpoch > sixMonths && expectedSlotProposals <= targetBlocks/6:
		proposalTimeFrame = sixMonths
	case timeSinceFirstEpoch > fiveMonths && expectedSlotProposals <= targetBlocks/5:
		proposalTimeFrame = fiveMonths
	case timeSinceFirstEpoch > fourMonths && expectedSlotProposals <= targetBlocks/4:
		proposalTimeFrame = fourMonths
	case timeSinceFirstEpoch > threeMonths && expectedSlotProposals <= targetBlocks/3:
		proposalTimeFrame = threeMonths
	case timeSinceFirstEpoch > twoMonths && expectedSlotProposals <= targetBlocks/2:
		proposalTimeFrame = twoMonths
	case timeSinceFirstEpoch > sixWeeks && expectedSlotProposals <= targetBlocks/1.5:
		proposalTimeFrame = sixWeeks
	default:
		proposalTimeFrame = oneMonth
	}

	// Recalculate expected slot proposals for the new timeframe
	expectedSlotProposals = calcExpectedSlotProposals(proposalTimeFrame, validatorsCount, activeValidatorsCount)
	if expectedSlotProposals == 0 {
		return 0, 0
	}
	// Cutoff time for proposals to be considered qualified
	blockProposalCutoffTime := time.Now().Add(-proposalTimeFrame)

	// Count the number of qualified proposals
	qualifiedProposalCount := 0
	for _, slot := range slots {
		if utils.SlotToTime(slot).After(blockProposalCutoffTime) {
			qualifiedProposalCount++
		}
	}
	// Return the luck as the ratio of qualified proposals to expected slot proposals
	return float64(qualifiedProposalCount) / expectedSlotProposals, proposalTimeFrame
}

// calcExpectedSlotProposals calculates the expected number of slot proposals for a certain time frame and validator count
func calcExpectedSlotProposals(timeframe time.Duration, validatorCount int, activeValidatorsCount uint64) float64 {
	if validatorCount == 0 || activeValidatorsCount == 0 {
		return 0
	}
	slotsInTimeframe := timeframe.Seconds() / float64(utils.Config.Chain.ClConfig.SecondsPerSlot)
	return (slotsInTimeframe / float64(activeValidatorsCount)) * float64(validatorCount)
}

// getAvgSlotInterval will return the average block interval for a certain number of validators
//
// result of the function should be interpreted as "1 in every X slots will be proposed by this amount of validators on avg."
func getAvgSlotInterval(validatorsCount int) float64 {
	// don't estimate if there are no proposed blocks or no validators
	activeValidatorsCount := *services.GetLatestStats().ActiveValidatorCount
	if activeValidatorsCount == 0 {
		return 0
	}

	probability := float64(validatorsCount) / float64(activeValidatorsCount)
	// in a geometric distribution, the expected value of the number of trials needed until first success is 1/p
	// you can think of this as the average interval of blocks until you get a proposal
	return 1 / probability
}

// getAvgSyncCommitteeInterval will return the average sync committee interval for a certain number of validators
//
// result of the function should be interpreted as "there will be one validator included in every X committees, on average"
func getAvgSyncCommitteeInterval(validatorsCount int) float64 {
	activeValidatorsCount := *services.GetLatestStats().ActiveValidatorCount
	if activeValidatorsCount == 0 {
		return 0
	}

	probability := (float64(utils.Config.Chain.ClConfig.SyncCommitteeSize) / float64(activeValidatorsCount)) * float64(validatorsCount)
	// in a geometric distribution, the expected value of the number of trials needed until first success is 1/p
	// you can think of this as the average interval of sync committees until you expect to have been part of one
	return 1 / probability
}

// LatestState will return common information that about the current state of the eth2 chain
func LatestState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", utils.Config.Chain.ClConfig.SecondsPerSlot)) // set local cache to the seconds per slot interval

	data := services.LatestState()
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Errorf("error sending latest index page data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func GetDataTableStateChanges(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	errMsgPrefix := "error loading data table state"

	response := &types.ApiResponse{}
	response.Status = errMsgPrefix
	response.Data = ""

	defer json.NewEncoder(w).Encode(response)

	response.Status = "OK"
}

func SetDataTableStateChanges(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	tableKey := vars["tableId"]

	errMsgPrefix := "error saving data table state"

	response := &types.ApiResponse{}
	response.Status = errMsgPrefix
	response.Data = ""

	defer json.NewEncoder(w).Encode(response)

	settings := types.DataTableSaveState{}
	err := json.NewDecoder(r.Body).Decode(&settings)
	if err != nil {
		logger.Warnf(errMsgPrefix+", could not parse body for tableKey %v: %v", tableKey, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	settings.Key = tableKey

	// never store the page number
	settings.Start = 0

	response.Data = settings

	response.Status = "OK"
}

// used to handle errors constructed by Template.ExecuteTemplate correctly
func handleTemplateError(w http.ResponseWriter, r *http.Request, fileIdentifier string, functionIdentifier string, infoIdentifier string, err error) error {
	// ignore network related errors
	if err != nil && !errors.Is(err, syscall.EPIPE) && !errors.Is(err, syscall.ETIMEDOUT) {
		logger.WithFields(logrus.Fields{
			"file":       fileIdentifier,
			"function":   functionIdentifier,
			"info":       infoIdentifier,
			"error type": fmt.Sprintf("%T", err),
			"route":      r.URL.String(),
		}).WithError(err).Error("error executing template")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	return err
}

func GetWithdrawableCountFromCursor(epoch uint64, validatorindex uint64, cursor uint64) (uint64, error) {
	// the validators' balance will not be checked here as this is only a rough estimation
	// checking the balance for hundreds of thousands of validators is too expensive

	var maxValidatorIndex uint64
	err := db.WriterDb.Get(&maxValidatorIndex, "SELECT COALESCE(MAX(validatorindex), 0) FROM validators")
	if err != nil {
		return 0, fmt.Errorf("error getting withdrawable validator count from cursor: %w", err)
	}

	if maxValidatorIndex == 0 {
		return 0, nil
	}

	activeValidators := services.LatestIndexPageData().ActiveValidators
	if activeValidators == 0 {
		activeValidators = maxValidatorIndex
	}

	if validatorindex > cursor {
		// if the validatorindex is after the cursor, simply return the number of validators between the cursor and the validatorindex
		// the returned data is then scaled using the number of currently active validators in order to account for exited / entering validators
		return (validatorindex - cursor) * activeValidators / maxValidatorIndex, nil
	} else if validatorindex < cursor {
		// if the validatorindex is before the cursor (wraparound case) return the number of validators between the cursor and the most recent validator plus the amount of validators from the validator 0 to the validatorindex
		// the returned data is then scaled using the number of currently active validators in order to account for exited / entering validators
		return (maxValidatorIndex - cursor + validatorindex) * activeValidators / maxValidatorIndex, nil
	} else {
		return 0, nil
	}
}

func getExecutionChartData(indices []uint64, lowerBoundDay uint64) ([]*types.ChartDataPoint, error) {
	var limit uint64 = 300
	blockList, consMap, err := findExecBlockNumbersByProposerIndex(indices, 0, limit, false, true, lowerBoundDay)
	if err != nil {
		return nil, err
	}

	blocks, err := db.BigtableClient.GetBlocksIndexedMultiple(blockList, limit)
	if err != nil {
		return nil, err
	}

	var chartData = []*types.ChartDataPoint{}
	epochsPerDay := utils.EpochsPerDay()
	color := "#90ed7d"

	// Map to keep track of the cumulative reward for each day
	dayRewardMap := make(map[int64]float64)

	for _, block := range blocks {
		consData := consMap[block.Number]
		day := int64(consData.Epoch / epochsPerDay)

		totalReward := utils.PlanckToZND(utils.Eth1TotalReward(block)).InexactFloat64()

		// Add the reward to the existing reward for the day or set it if not previously set
		dayRewardMap[day] += totalReward
	}

	// Now populate the chartData array using the dayRewardMap
	for day, reward := range dayRewardMap {
		ts := float64(utils.DayToTime(day).Unix() * 1000)
		chartData = append(chartData, &types.ChartDataPoint{
			X:     ts,
			Y:     reward,
			Color: color,
		})
	}

	// If needed, sort chartData based on X values
	sort.Slice(chartData, func(i, j int) bool {
		return chartData[i].X < chartData[j].X
	})

	return chartData, nil
}

func findExecBlockNumbersByProposerIndex(indices []uint64, offset, limit uint64, isSortAsc bool, onlyFinalized bool, lowerBoundDay uint64) ([]uint64, map[uint64]types.ExecBlockProposer, error) {
	var blockListSub []types.ExecBlockProposer

	lowerBoundEpoch := lowerBoundDay * utils.EpochsPerDay()

	order := "DESC"
	if isSortAsc {
		order = "ASC"
	}

	status := "status != '3'"
	if onlyFinalized {
		status = `status = '1'`
	}

	query := fmt.Sprintf(`
		SELECT 
			exec_block_number,
			proposer,
			slot,
			epoch  
		FROM blocks 
		WHERE proposer = ANY($1)
			AND exec_block_number IS NOT NULL AND exec_block_number > 0
			AND epoch >= $4
			AND %s
		ORDER BY exec_block_number %s
		OFFSET $2 LIMIT $3`, status, order)

	err := db.ReaderDb.Select(&blockListSub,
		query,
		pq.Array(indices),
		offset,
		limit,
		lowerBoundEpoch,
	)
	if err != nil {
		return nil, nil, err
	}
	blockList, blockProposerMap := getBlockNumbersAndMapProposer(blockListSub)
	return blockList, blockProposerMap, nil
}

func getBlockNumbersAndMapProposer(data []types.ExecBlockProposer) ([]uint64, map[uint64]types.ExecBlockProposer) {
	blockList := []uint64{}
	blockToProposerMap := make(map[uint64]types.ExecBlockProposer)
	for _, execBlock := range data {
		blockList = append(blockList, execBlock.ExecBlock)
		blockToProposerMap[execBlock.ExecBlock] = execBlock
	}
	return blockList, blockToProposerMap
}
