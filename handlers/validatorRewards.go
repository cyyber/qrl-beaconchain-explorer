package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/templates"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"github.com/gorilla/csrf"
)

type rewardsResp struct {
	CsrfField template.HTML
}

func ValidatorRewards(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "validatorRewards.html")
	var validatorRewardsServicesTemplate = templates.GetTemplate(templateFiles...)

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "services", "/rewards", "Zond Validator Rewards", templateFiles)

	data.Data = rewardsResp{CsrfField: csrf.TemplateField(r)}

	if handleTemplateError(w, r, "validatorRewards.go", "ValidatorRewards", "", validatorRewardsServicesTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}

func RewardsHistoricalData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()

	validatorIndexArr, _, redirect, err := handleValidatorsQuery(w, r, true)
	if err != nil || redirect {
		return
	}

	// Set the default start and end time to the first day
	t := time.Unix(int64(utils.Config.Chain.GenesisTimestamp), 0)
	startGenesisDay := uint64(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).Unix())
	var start uint64 = startGenesisDay
	var end uint64 = startGenesisDay

	dateRange := strings.Split(q.Get("days"), "-")
	if len(dateRange) == 2 {
		start, err = strconv.ParseUint(dateRange[0], 10, 32) //Limit to uint32 for postgres
		if err != nil || start < startGenesisDay {
			logger.Warnf("error parsing days range: %v", err)
			http.Error(w, "Error: Invalid parameter days.", http.StatusBadRequest)
			return
		}
		end, err = strconv.ParseUint(dateRange[1], 10, 32) //Limit to uint32 for postgres
		if err != nil || end < startGenesisDay {
			logger.Warnf("error parsing days range: %v", err)
			http.Error(w, "Error: Invalid parameter days.", http.StatusBadRequest)
			return
		}
	}

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

	// Set the default start and end time to the first day
	t := time.Unix(int64(utils.Config.Chain.GenesisTimestamp), 0)
	startGenesisDay := uint64(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).Unix())
	var start uint64 = startGenesisDay
	var end uint64 = startGenesisDay

	dateRange := strings.Split(q.Get("days"), "-")
	if len(dateRange) == 2 {
		start, err = strconv.ParseUint(dateRange[0], 10, 32) //Limit to uint32 for postgres
		if err != nil || start < startGenesisDay {
			logger.Warnf("error parsing days range: %v", err)
			http.Error(w, "Error: Invalid parameter days.", http.StatusBadRequest)
			return
		}
		end, err = strconv.ParseUint(dateRange[1], 10, 32) //Limit to uint32 for postgres
		if err != nil || end < startGenesisDay {
			logger.Warnf("error parsing days range: %v", err)
			http.Error(w, "Error: Invalid parameter days.", http.StatusBadRequest)
			return
		}
	}

	hist := services.GetValidatorHist(validatorIndexArr, start, end)

	if len(hist.History) == 0 {
		w.Write([]byte("No data available"))
		return
	}

	s := time.Unix(int64(start), 0)
	e := time.Unix(int64(end), 0)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=income_history_%v_%v.pdf", s.Format("20060102"), e.Format("20060102")))
	w.Header().Set("Content-Type", "text/csv")

	_, err = w.Write(services.GeneratePdfReport(hist))
	if err != nil {
		logger.WithError(err).WithField("route", r.URL.String()).Error("error writing response")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
