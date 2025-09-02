package handlers

// TODO(now.youtrack.cloud/issue/TZB-1)
/*
import (
	"net/http"

	"github.com/theQRL/qrl-beaconchain-explorer/templates"
	"github.com/theQRL/qrl-beaconchain-explorer/types"

	"github.com/gorilla/mux"
)

// Will return the QrnsPage
func QrnsSearch(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "qrnsSearch.html")
	var qrnsSearchTemplate = templates.GetTemplate(templateFiles...)

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "qrnssearch", "/qrns", "Qrns search", templateFiles)

	vars := mux.Vars(r)
	search := vars["search"]

	result, err := GetQrnsDomain(search)

	var pageData types.QrnsSearchPageData
	if err != nil {
		pageData.Error = "No matching QRNS registration found"
	} else {
		pageData.Result = result
	}
	pageData.Search = search

	data.Data = pageData

	if handleTemplateError(w, r, "qrnsSearch.go", "QrnsSearch", "", qrnsSearchTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}
*/
