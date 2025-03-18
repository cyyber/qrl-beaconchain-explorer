package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/theQRL/zond-beaconchain-explorer/price"
	"github.com/theQRL/zond-beaconchain-explorer/services"
	"github.com/theQRL/zond-beaconchain-explorer/types"
	"github.com/theQRL/zond-beaconchain-explorer/utils"
	"github.com/theQRL/zond-beaconchain-explorer/version"
)

var layoutTemplateFiles = []string{
	"layout.html",
	"layout/mainnavigation.html",
}

func InitPageData(w http.ResponseWriter, r *http.Request, active, path, title string, mainTemplates []string) *types.PageData {
	fullTitle := fmt.Sprintf("%v - %v - explorer.zond.theqrl.org - %v", title, utils.Config.Frontend.SiteName, time.Now().Year())

	if title == "" {
		fullTitle = fmt.Sprintf("%v - explorer.zond.theqrl.org - %v", utils.Config.Frontend.SiteName, time.Now().Year())
	}

	isMainnet := utils.Config.Chain.ClConfig.ConfigName == "mainnet"
	data := &types.PageData{
		Meta: &types.Meta{
			Title:       fullTitle,
			Description: "explorer.zond.theqrl.org makes Zond accessible to non-technical end users",
			Path:        path,
			GATag:       utils.Config.Frontend.GATag,
			NoTrack:     false,
			Templates:   strings.Join(mainTemplates, ","),
		},

		Active:                active,
		Data:                  &types.Empty{},
		Version:               version.Version,
		Year:                  time.Now().UTC().Year(),
		ChainSlotsPerEpoch:    utils.Config.Chain.ClConfig.SlotsPerEpoch,
		ChainSecondsPerSlot:   utils.Config.Chain.ClConfig.SecondsPerSlot,
		ChainGenesisTimestamp: utils.Config.Chain.GenesisTimestamp,
		CurrentEpoch:          services.LatestEpoch(),
		LatestFinalizedEpoch:  services.LatestFinalizedEpoch(),
		CurrentSlot:           services.LatestSlot(),
		FinalizationDelay:     services.FinalizationDelay(),
		Rates:                 &types.Rates{SelectedCurrency: "ZND"},
		// Rates:                 services.GetRates(GetCurrency(r)),
		Mainnet:             utils.Config.Chain.ClConfig.ConfigName == "mainnet",
		DepositContract:     utils.Config.Chain.ClConfig.DepositContractAddress,
		ChainConfig:         utils.Config.Chain.ClConfig,
		Lang:                "en-US",
		Debug:               utils.Config.Frontend.Debug,
		GasNow:              services.LatestGasNowData(),
		ShowSyncingMessage:  services.IsSyncing(),
		GlobalNotification:  services.GlobalNotificationMessage(),
		AvailableCurrencies: price.GetAvailableCurrencies(),
		MainMenuItems:       createMenuItems(active, isMainnet),
		TermsOfServiceUrl:   utils.Config.Frontend.Legal.TermsOfServiceUrl,
		PrivacyPolicyUrl:    utils.Config.Frontend.Legal.PrivacyPolicyUrl,
	}

	acceptedLangs := strings.Split(r.Header.Get("Accept-Language"), ",")
	if len(acceptedLangs) > 0 {
		if strings.Contains(acceptedLangs[0], "ru") || strings.Contains(acceptedLangs[0], "RU") {
			data.Lang = "ru-RU"
		}
	}

	for _, v := range r.Cookies() {
		if v.Name == "language" {
			data.Lang = v.Value
			break
		}
	}

	return data
}

func SetPageDataTitle(pageData *types.PageData, title string) {
	if title == "" {
		pageData.Meta.Title = fmt.Sprintf("%v - beaconcha.in - %v", utils.Config.Frontend.SiteName, time.Now().Year())
	} else {
		pageData.Meta.Title = fmt.Sprintf("%v - %v - beaconcha.in - %v", title, utils.Config.Frontend.SiteName, time.Now().Year())
	}
}

func createMenuItems(active string, isMain bool) []types.MainMenuItem {
	hiddenFor := []string{"confirmation", "login", "register"}

	if utils.SliceContains(hiddenFor, active) {
		return []types.MainMenuItem{}
	}
	return []types.MainMenuItem{
		{
			Label:    "Blockchain",
			IsActive: active == "blockchain",
			Groups: []types.NavigationGroup{
				{
					Links: []types.NavigationLink{
						{
							Label: "Epochs",
							Path:  "/epochs",
							Icon:  "fa-history",
						},
						{
							Label: "Slots",
							Path:  "/slots",
							Icon:  "fa-cube",
						},
					},
				}, {
					Links: []types.NavigationLink{
						{
							Label: "Blocks",
							Path:  "/blocks",
							Icon:  "fa-cubes",
						},
						{
							Label: "Txs",
							Path:  "/transactions",
							Icon:  "fa-credit-card",
						},
						{
							Label: "Mempool",
							Path:  "/mempool",
							Icon:  "fa-upload",
						},
					},
				},
			},
		},
		{
			Label:    "Validators",
			IsActive: active == "validators",
			Groups: []types.NavigationGroup{
				{
					Links: []types.NavigationLink{
						{
							Label: "Overview",
							Path:  "/validators",
							Icon:  "fa-table",
						},
						{
							Label: "Slashings",
							Path:  "/validators/slashings",
							Icon:  "fa-user-slash",
						},
					},
				}, {
					Links: []types.NavigationLink{
						{
							Label: "Validator Leaderboard",
							Path:  "/validators/leaderboard",
							Icon:  "fa-medal",
						},
						{
							Label: "Deposit Leaderboard",
							Path:  "/validators/deposit-leaderboard",
							Icon:  "fa-file-import",
						},
					},
				}, {
					Links: []types.NavigationLink{
						{
							Label: "Deposits",
							Path:  "/validators/deposits",
							Icon:  "fa-file-signature",
						},
						{
							Label: "Withdrawals",
							Path:  "/validators/withdrawals",
							Icon:  "fa-money-bill",
						},
					},
				},
			},
		},
		{
			Label:    "Dashboard",
			IsActive: active == "dashboard",
			Path:     "/dashboard",
		},
		{
			Label:        "More",
			IsActive:     active == "more",
			HasBigGroups: true,
			Groups: []types.NavigationGroup{
				{
					Label: "Stats",
					Links: []types.NavigationLink{
						{
							Label: "Charts",
							Path:  "/charts",
							Icon:  "fa-chart-bar",
						},
						{
							Label: "Income History",
							Path:  "/rewards",
							Icon:  "fa-money-bill-alt",
						},
						{
							Label: "Profit Calculator",
							Path:  "/calculator",
							Icon:  "fa-calculator",
						},
						{
							Label: "Block Viz",
							Path:  "/vis",
							Icon:  "fa-project-diagram",
						},
						{
							Label: "EIP-1559 Burn",
							Path:  "/burn",
							Icon:  "fa-burn",
						},
						{
							Label:    "Correlations",
							Path:     "/correlations",
							Icon:     "fa-chart-line",
							IsHidden: !isMain,
						},
					},
				}, {
					Label: "Tools",
					Links: []types.NavigationLink{
						{
							Label: "Unit Converter",
							Path:  "/tools/unitConverter",
							Icon:  "fa-sync",
						},
						{
							Label: "GasNow",
							Path:  "/gasnow",
							Icon:  "fa-gas-pump",
						},
						{
							Label: "Broadcast Signed Messages",
							Path:  "/tools/broadcast",
							Icon:  "fa-bullhorn",
						},
					},
				}, {
					Label: "Services",
					Links: []types.NavigationLink{
						{
							Label: "Zond Clients",
							Path:  "/zondClients",
							Icon:  "fa-desktop",
						},
						{
							Label: "Slot Finder",
							Path:  "/slots/finder",
							Icon:  "fa-cube",
						},
						// TODO(rgeraldes24)
						/*
							{
								Label: "Report a scam",
								Path:  "https://www.chainabuse.com/report?source=bitfly",
								Icon:  "fa-flag",
							},
						*/
					},
				},
			},
		},
	}
}
