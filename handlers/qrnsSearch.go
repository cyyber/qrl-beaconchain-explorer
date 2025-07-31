package handlers

import (
	"net/http"

	"github.com/theQRL/zond-beaconchain-explorer/templates"
	"github.com/theQRL/zond-beaconchain-explorer/types"

	"github.com/gorilla/mux"
)

// Will return the ZnsPage
func ZnsSearch(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "znsSearch.html")
	var znsSearchTemplate = templates.GetTemplate(templateFiles...)

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "znssearch", "/zns", "Zns search", templateFiles)

	vars := mux.Vars(r)
	search := vars["search"]

	result, err := GetQrnsDomain(search)

	var pageData types.ZnsSearchPageData
	if err != nil {
		pageData.Error = "No matching QRNS registration found"
	} else {
		pageData.Result = result
	}
	pageData.Search = search

	data.Data = pageData

	if handleTemplateError(w, r, "znsSearch.go", "ZnsSearch", "", znsSearchTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}
