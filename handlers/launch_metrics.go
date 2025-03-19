package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/utils"
)

// SlotVizMetrics returns the metrics for the earliest epochs
func SlotVizMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", utils.Config.Chain.ClConfig.SecondsPerSlot)) // set local cache to the seconds per slot interval

	res := services.LatestSlotVizMetrics()

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Errorf("error enconding json response for %v route: %v", r.URL.String(), err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
