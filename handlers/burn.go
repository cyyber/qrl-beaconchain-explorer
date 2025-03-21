package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/templates"
)

func Burn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	templateFiles := append(layoutTemplateFiles, "burn.html")
	data := InitPageData(w, r, "burn", "/burn", "ZND Burned", templateFiles)

	var burnTemplate = templates.GetTemplate(templateFiles...)

	// data.Meta.Tdata1 = utils.FormatAmount((data.Data.(*types.BurnPageData).TotalBurned / 1e18) * data.Data.(*types.BurnPageData).Price)
	// data.Meta.Tdata2 = utils.FormatAmount(data.Data.(*types.BurnPageData).BurnRate24h/1e18) + " ZND/min"
	// data.Meta.Description = "The current zond burn rate is " + data.Meta.Tdata2 + ". A total of " + utils.FormatUSD(data.Data.(*types.BurnPageData).TotalBurned/1e18) + "ZND with a market value of $" + data.Meta.Tdata1 + " has been burned. " + data.Meta.Description

	latestBurn := services.LatestBurnData()

	// if currency == utils.MainCurrency {
	// 	currency = "USD"
	// }

	// TODO(rgeraldes24)
	// latestBurn.Price = price.GetPrice(utils.MainCurrency, currency)
	// latestBurn.Price = 1.0
	// latestBurn.Currency = currency

	data.Data = latestBurn
	if handleTemplateError(w, r, "burn.go", "Burn", "", burnTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}

func BurnPageData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	latestBurn := services.LatestBurnData()

	// currency := GetCurrency(r)

	// if currency == "ZND" {
	// 	currency = "USD"
	// }

	// latestBurn.Price = price.GetPrice(utils.MainCurrency, currency)
	// latestBurn.Price = 1.0
	// latestBurn.Currency = currency

	err := json.NewEncoder(w).Encode(latestBurn)
	if err != nil {
		logger.Errorf("error sending latest burn page data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
