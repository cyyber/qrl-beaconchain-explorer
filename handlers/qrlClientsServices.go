package handlers

import (
	"net/http"

	qrlclients "github.com/theQRL/qrl-beaconchain-explorer/qrlClients"
	"github.com/theQRL/qrl-beaconchain-explorer/templates"

	"github.com/gorilla/csrf"
)

func QRLClientsServices(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "qrlClientsServices.html")
	var qrlClientsServicesTemplate = templates.GetTemplate(templateFiles...)

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "services", "/qrlClientsServices", "QRL Clients Services Overview", templateFiles)

	pageData := qrlclients.GetQRLClientData()
	pageData.CsrfField = csrf.TemplateField(r)

	data.Data = pageData
	if handleTemplateError(w, r, "qrlClientsServices.go", "QRLClientsServices", "", qrlClientsServicesTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}
