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

	latestBurn := services.LatestBurnData()
	data.Data = latestBurn
	if handleTemplateError(w, r, "burn.go", "Burn", "", burnTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}

func BurnPageData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	latestBurn := services.LatestBurnData()

	err := json.NewEncoder(w).Encode(latestBurn)
	if err != nil {
		logger.Errorf("error sending latest burn page data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
