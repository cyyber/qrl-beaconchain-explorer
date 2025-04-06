package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/theQRL/zond-beaconchain-explorer/db"
	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/types"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
)

// ApiHealthz godoc
// @Summary Health of the explorer
// @Tags Misc
// @Description Health endpoint for monitoring if the explorer is in sync
// @Produce  text/plain
// @Success 200 {object} types.ApiResponse
// @Router /api/healthz [get]
func ApiHealthz(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")

	modules := []string{
		"monitoring_app",
		"monitoring_el_data",
		"monitoring_services",
		"monitoring_cl_data",
		"monitoring_api",
		"monitoring_redis",
	}

	res := []struct {
		Name   string
		Status string
	}{}
	err := db.WriterDb.Select(&res, "SELECT name, status FROM service_status WHERE name = ANY($1) AND last_update > NOW() - INTERVAL '5 MINUTES' ORDER BY last_update DESC", pq.Array(modules))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No monitoring data available", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	modulesMap := make(map[string]string)
	for _, module := range modules {
		modulesMap[module] = ""
	}

	hasError := false
	response := strings.Builder{}
	for _, status := range res {

		if modulesMap[status.Name] == "" {
			modulesMap[status.Name] = status.Status

			if status.Status != "OK" {
				hasError = true
			}

			response.WriteString(fmt.Sprintf("module %s: %s\n", status.Name, status.Status))
		}
	}

	for _, module := range modules {
		if modulesMap[module] == "" {
			hasError = true
			response.WriteString(fmt.Sprintf("module %s: %s\n", module, "No monitoring data available"))
		}
	}

	if !hasError {
		_, err = fmt.Fprint(w, response.String())

		if err != nil {
			logger.Debugf("error writing status: %v", err)
		}
	} else {
		http.Error(w, response.String(), http.StatusInternalServerError)
		return
	}
}

// ApiHealthzLoadbalancer godoc
// @Summary Health of the explorer-api regarding having a healthy connection to the database
// @Tags Misc
// @Description Health endpoint for monitoring if the explorer-api
// @Produce  text/plain
// @Success 200 {object} types.ApiResponse
// @Router /api/healthz-loadbalancer [get]
func ApiHealthzLoadbalancer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	lastEpoch, err := db.GetLatestEpoch()

	if err != nil {
		http.Error(w, "Internal server error: could not retrieve latest epoch from the db", http.StatusInternalServerError)
		return
	}

	if utils.Config.Chain.GenesisTimestamp == 18446744073709551615 {
		fmt.Fprint(w, "OK. No GENESIS_TIMESTAMP defined yet")
		return
	}

	genesisTime := time.Unix(int64(utils.Config.Chain.GenesisTimestamp), 0)
	if genesisTime.After(time.Now()) {
		fmt.Fprintf(w, "OK. Genesis in %v (%v)", time.Until(genesisTime), genesisTime)
		return
	}

	fmt.Fprintf(w, "OK. Last epoch is from %v ago", time.Since(utils.EpochToTime(lastEpoch)))
}

func getExpectedSyncCommitteeSlots(validators []uint64, epoch uint64) (expectedSlots uint64, err error) {
	lastFinalizedEpoch := services.LatestFinalizedEpoch()
	if epoch > lastFinalizedEpoch {
		epoch = lastFinalizedEpoch
	}

	// retrieve activation and exit epochs from database per validator
	type ValidatorInfo struct {
		Id                         int64  `db:"validatorindex"`
		ActivationEpoch            uint64 `db:"activationepoch"`
		ExitEpoch                  uint64 `db:"exitepoch"`
		FirstPossibleSyncCommittee uint64 // calculated
		LastPossibleSyncCommittee  uint64 // calculated
	}

	var validatorsInfoFromDb = []ValidatorInfo{}
	query, args, err := sqlx.In(`SELECT validatorindex, activationepoch, exitepoch FROM validators WHERE validatorindex IN (?) ORDER BY validatorindex ASC`, validators)
	if err != nil {
		return 0, err
	}

	err = db.ReaderDb.Select(&validatorsInfoFromDb, db.ReaderDb.Rebind(query), args...)
	if err != nil {
		return 0, err
	}

	// only check validators that are/have been active and that did not exit before altair
	const noEpoch = uint64(9223372036854775807)
	var validatorsInfo = make([]ValidatorInfo, 0, len(validatorsInfoFromDb))
	for _, v := range validatorsInfoFromDb {
		if v.ActivationEpoch != noEpoch && v.ActivationEpoch < epoch {
			validatorsInfo = append(validatorsInfo, v)
		}
	}

	if len(validatorsInfo) == 0 {
		// no validators relevant for sync duties
		return 0, nil
	}

	// we need all related and unique timeframes (activation and exit sync period) for all validators
	uniquePeriods := make(map[uint64]bool)
	for i := range validatorsInfo {
		// first epoch (activation epoch or Altair if Altair was later as there were no sync committees pre Altair)
		firstSyncEpoch := validatorsInfo[i].ActivationEpoch
		validatorsInfo[i].FirstPossibleSyncCommittee = utils.SyncPeriodOfEpoch(firstSyncEpoch)
		uniquePeriods[validatorsInfo[i].FirstPossibleSyncCommittee] = true

		// last epoch (exit epoch or current epoch if not exited yet)
		lastSyncEpoch := epoch
		if validatorsInfo[i].ExitEpoch != noEpoch && validatorsInfo[i].ExitEpoch <= epoch {
			lastSyncEpoch = validatorsInfo[i].ExitEpoch
		}
		validatorsInfo[i].LastPossibleSyncCommittee = utils.SyncPeriodOfEpoch(lastSyncEpoch)
		uniquePeriods[validatorsInfo[i].LastPossibleSyncCommittee] = true
	}

	// transform map to slice; this will be used to query sync_committees_count_per_validator
	periodSlice := make([]uint64, 0, len(uniquePeriods))
	for period := range uniquePeriods {
		periodSlice = append(periodSlice, period)
	}

	// get aggregated count for all relevant committees from sync_committees_count_per_validator
	var countStatistics []struct {
		Period     uint64  `db:"period"`
		CountSoFar float64 `db:"count_so_far"`
	}

	query, args, errs := sqlx.In(`SELECT period, count_so_far FROM sync_committees_count_per_validator WHERE period IN (?) ORDER BY period ASC`, periodSlice)
	if errs != nil {
		return 0, errs
	}
	err = db.ReaderDb.Select(&countStatistics, db.ReaderDb.Rebind(query), args...)
	if err != nil {
		return 0, err
	}
	if len(countStatistics) != len(periodSlice) {
		return 0, fmt.Errorf("unable to retrieve all sync committee count statistics, required %v entries but got %v entries (epoch: %v)", len(periodSlice), len(countStatistics), epoch)
	}

	// transform query result to map for easy access
	periodInfoMap := make(map[uint64]float64)
	for _, pl := range countStatistics {
		periodInfoMap[pl.Period] = pl.CountSoFar
	}

	// calculate expected committies for every single validator and aggregate them
	expectedCommitties := 0.0
	for _, vi := range validatorsInfo {
		expectedCommitties += periodInfoMap[vi.LastPossibleSyncCommittee] - periodInfoMap[vi.FirstPossibleSyncCommittee]
	}

	// transform committees to slots
	expectedSlots = uint64(expectedCommitties * float64(utils.SlotsPerSyncCommittee()))

	return expectedSlots, nil
}

type Cached struct {
	Data interface{}
	Ts   int64
}

func DecodeMapStructure(input interface{}, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   output,
		TagName:  "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func SendBadRequestResponse(w http.ResponseWriter, route, message string) {
	sendErrorWithCodeResponse(w, route, message, http.StatusBadRequest)
}

func sendErrorWithCodeResponse(w http.ResponseWriter, route, message string, errorcode int) {
	w.WriteHeader(errorcode)
	j := json.NewEncoder(w)
	response := &types.ApiResponse{}
	response.Status = "ERROR: " + message
	err := j.Encode(response)

	if err != nil {
		logger.Errorf("error serializing json error for API %v route: %v", route, err)
	}
}

func SendOKResponse(j *json.Encoder, route string, data []interface{}) {
	response := &types.ApiResponse{}
	response.Status = "OK"

	if len(data) == 1 {
		response.Data = data[0]
	} else {
		response.Data = data
	}
	err := j.Encode(response)

	if err != nil {
		logger.Errorf("error serializing json data for API %v route: %v", route, err)
	}
}
