package handlers

import (
	"net/http"

	"github.com/theQRL/zond-beaconchain-explorer/templates"
	zondclients "github.com/theQRL/zond-beaconchain-explorer/zondClients"

	"github.com/gorilla/csrf"
)

func ZondClientsServices(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "zondClientsServices.html")
	var zondClientsServicesTemplate = templates.GetTemplate(templateFiles...)

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "services", "/zondClientsServices", "Zond Clients Services Overview", templateFiles)

	pageData := zondclients.GetZondClientData()
	pageData.CsrfField = csrf.TemplateField(r)

	data.Data = pageData
	if handleTemplateError(w, r, "zondClientsServices.go", "ZondClientsServices", "", zondClientsServicesTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}
