package handlers

// TODO(now.youtrack.cloud/issue/TZB-1)
/*
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

	result, err := GetZnsDomain(search)

	var pageData types.ZnsSearchPageData
	if err != nil {
		pageData.Error = "No matching ZNS registration found"
	} else {
		pageData.Result = result
	}
	pageData.Search = search

	data.Data = pageData

	if handleTemplateError(w, r, "znsSearch.go", "ZnsSearch", "", znsSearchTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}
*/
