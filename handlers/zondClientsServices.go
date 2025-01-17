package handlers

import (
	"net/http"

	"github.com/theQRL/zond-beaconchain-explorer/db"
	"github.com/theQRL/zond-beaconchain-explorer/templates"
	"github.com/theQRL/zond-beaconchain-explorer/types"
	zondclients "github.com/theQRL/zond-beaconchain-explorer/zondClients"

	"github.com/gorilla/csrf"
)

func ZondClientsServices(w http.ResponseWriter, r *http.Request) {
	templateFiles := append(layoutTemplateFiles, "zondClientsServices.html")
	var zondClientsServicesTemplate = templates.GetTemplate(templateFiles...)

	var err error

	w.Header().Set("Content-Type", "text/html")

	data := InitPageData(w, r, "services", "/zondClientsServices", "Zond Clients Services Overview", templateFiles)

	pageData := zondclients.GetEthClientData()
	pageData.CsrfField = csrf.TemplateField(r)
	pageData.Banner = zondclients.GetBannerClients()
	if data.User.Authenticated {
		var dbData []string
		err = db.FrontendWriterDB.Select(&dbData,
			`select event_filter
			 from users_subscriptions 
			 where user_id = $1 AND event_name=$2
			`, data.User.UserID, string(types.EthClientUpdateEventName))
		if err != nil {
			logger.Errorf("error getting user subscriptions: %v route: %v", r.URL.String(), err)
		}

		for _, item := range dbData {
			switch item {
			case "gzond":
				pageData.Gzond.IsUserSubscribed = true
			case "qrysm":
				pageData.Qrysm.IsUserSubscribed = true
			default:
				continue
			}
		}

	}

	data.Data = pageData
	if handleTemplateError(w, r, "zondClientsServices.go", "ZondClientsServices", "", zondClientsServicesTemplate.ExecuteTemplate(w, "layout", data)) != nil {
		return // an error has occurred and was processed
	}
}
