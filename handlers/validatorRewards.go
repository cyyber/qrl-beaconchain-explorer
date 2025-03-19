package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/templates"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"github.com/gorilla/csrf"
)

// var supportedCurrencies = []string{"eur", "usd", "gbp", "cny", "cad", "jpy", "rub", "aud"}
const USER_SUBSCRIPTION_LIMIT = 8

type rewardsResp struct {
	Currencies       []string
	CsrfField        template.HTML
	MinDateTimestamp uint64
}

func ValidatorRewards(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "validatorRewards.html")
	var validatorRewardsServicesTemplate = templates.GetTemplate(templateFiles...)

	// var err error

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "services", "/rewards", "Zond Validator Rewards", templateFiles)

	// var supportedCurrencies []string
	// err = db.ReaderDb.Select(&supportedCurrencies,
	// 	`select column_name
	// 		from information_schema.columns
	// 		where table_name = 'price'`)
	// if err != nil {
	// 	logger.Errorf("error getting eth1-deposits-distribution for stake pools: %v", err)
	// }

	// var minTime time.Time
	// err = db.ReaderDb.Get(&minTime,
	// 	`select ts from price order by ts asc limit 1`)
	// if err != nil {
	// 	logger.Errorf("error getting min ts: %v", err)
	// }

	// data.Data = rewardsResp{Currencies: supportedCurrencies, CsrfField: csrf.TemplateField(r), MinDateTimestamp: uint64(minTime.Unix())}
	data.Data = rewardsResp{CsrfField: csrf.TemplateField(r)}

	if handleTemplateError(w, r, "validatorRewards.go", "ValidatorRewards", "", validatorRewardsServicesTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}

func RewardsHistoricalData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// q := r.URL.Query()

	validatorIndexArr, _, redirect, err := handleValidatorsQuery(w, r, true)
	if err != nil || redirect {
		return
	}

	// currency := q.Get("currency")

	// Set the default start and end time to the first day
	t := time.Unix(int64(utils.Config.Chain.GenesisTimestamp), 0)
	startGenesisDay := uint64(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).Unix())
	var start uint64 = startGenesisDay
	var end uint64 = startGenesisDay

	// TODO(rgeraldes24): reactivate
	// dateRange := strings.Split(q.Get("days"), "-")
	// if len(dateRange) == 2 {
	// 	start, err = strconv.ParseUint(dateRange[0], 10, 32) //Limit to uint32 for postgres
	// 	if err != nil || start < startGenesisDay {
	// 		logger.Warnf("error parsing days range: %v", err)
	// 		http.Error(w, "Error: Invalid parameter days.", http.StatusBadRequest)
	// 		return
	// 	}
	// 	end, err = strconv.ParseUint(dateRange[1], 10, 32) //Limit to uint32 for postgres
	// 	if err != nil || end < startGenesisDay {
	// 		logger.Warnf("error parsing days range: %v", err)
	// 		http.Error(w, "Error: Invalid parameter days.", http.StatusBadRequest)
	// 		return
	// 	}
	// }

	data := services.GetValidatorHist(validatorIndexArr /*currency,*/, start, end)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.WithError(err).WithField("route", r.URL.String()).Error("error encoding json response")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func DownloadRewardsHistoricalData(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	validatorIndexArr, _, redirect, err := handleValidatorsQuery(w, r, true)
	if err != nil || redirect {
		return
	}

	currency := q.Get("currency")

	// Set the default start and end time to the first day
	t := time.Unix(int64(utils.Config.Chain.GenesisTimestamp), 0)
	startGenesisDay := uint64(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).Unix())
	var start uint64 = startGenesisDay
	var end uint64 = startGenesisDay

	// TODO(rgeraldes24): reactivate
	// dateRange := strings.Split(q.Get("days"), "-")
	// if len(dateRange) == 2 {
	// 	start, err = strconv.ParseUint(dateRange[0], 10, 32) //Limit to uint32 for postgres
	// 	if err != nil || start < startGenesisDay {
	// 		logger.Warnf("error parsing days range: %v", err)
	// 		http.Error(w, "Error: Invalid parameter days.", http.StatusBadRequest)
	// 		return
	// 	}
	// 	end, err = strconv.ParseUint(dateRange[1], 10, 32) //Limit to uint32 for postgres
	// 	if err != nil || end < startGenesisDay {
	// 		logger.Warnf("error parsing days range: %v", err)
	// 		http.Error(w, "Error: Invalid parameter days.", http.StatusBadRequest)
	// 		return
	// 	}
	// }

	hist := services.GetValidatorHist(validatorIndexArr /*currency,*/, start, end)

	if len(hist.History) == 0 {
		w.Write([]byte("No data available"))
		return
	}

	s := time.Unix(int64(start), 0)
	e := time.Unix(int64(end), 0)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=income_history_%v_%v.pdf", s.Format("20060102"), e.Format("20060102")))
	w.Header().Set("Content-Type", "text/csv")

	_, err = w.Write(services.GeneratePdfReport(hist, currency))
	if err != nil {
		logger.WithError(err).WithField("route", r.URL.String()).Error("error writing response")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
