package db

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/theQRL/zond-beaconchain-explorer/types"
	"github.com/theQRL/zond-beaconchain-explorer/utils"

	"github.com/jmoiron/sqlx"
)

// FrontendWriterDB is a pointer to the auth-database
var FrontendReaderDB *sqlx.DB
var FrontendWriterDB *sqlx.DB

func MustInitFrontendDB(writer *types.DatabaseConfig, reader *types.DatabaseConfig, driverName string, databaseBrand string) {
	FrontendWriterDB, FrontendReaderDB = mustInitDB(writer, reader, driverName, databaseBrand)
}

// GetUserEmailById returns the email of a user.
func GetUserEmailById(id uint64) (string, error) {
	var mail string = ""
	err := FrontendWriterDB.Get(&mail, "SELECT email FROM users WHERE id = $1", id)
	return mail, err
}

func MobileDeviceSettingsUpdate(userID, deviceID uint64, notifyEnabled, active string) (*sql.Rows, error) {
	var query = ""
	var args []interface{}

	args = append(args, userID)
	args = append(args, deviceID)

	if notifyEnabled != "" {
		args = append(args, notifyEnabled == "true")
		query = addParamToQuery(query, fmt.Sprintf("notify_enabled = $%d", len(args)))
	}

	if active != "" {
		args = append(args, active == "true")
		query = addParamToQuery(query, fmt.Sprintf("active = $%d", len(args)))
	}

	if query == "" {
		return nil, errors.New("no params for change provided")
	}

	rows, err := FrontendWriterDB.Query("UPDATE users_devices SET "+query+" WHERE user_id = $1 AND id = $2 RETURNING notify_enabled;",
		args...,
	)
	return rows, err
}

func MobileDeviceDelete(userID, deviceID uint64) error {
	_, err := FrontendWriterDB.Exec("DELETE FROM users_devices WHERE user_id = $1 AND id = $2 AND id != 2;", userID, deviceID)
	return err
}

func addParamToQuery(query, param string) string {
	var result = query
	if result != "" {
		result += ","
	}
	result += param
	return result
}

func MobileDeviceSettingsSelect(userID, deviceID uint64) (*sql.Rows, error) {
	rows, err := FrontendWriterDB.Query("SELECT notify_enabled FROM users_devices WHERE user_id = $1 AND id = $2;",
		userID, deviceID,
	)
	return rows, err
}

func NewTransaction() (*sql.Tx, error) {
	return FrontendWriterDB.Begin()
}

func getMachineStatsGap(resultCount uint64) int {
	if resultCount > 20160 { // more than 14 (31)
		return 8
	}
	if resultCount > 10080 { // more than 7 (14)
		return 7
	}
	if resultCount > 2880 { // more than 2 (7)
		return 5
	}
	if resultCount > 1440 { // more than 1 (2)
		return 4
	}
	if resultCount > 770 { // more than 12h
		return 2
	}
	return 1
}

func GetHistoricalPrice(chainId uint64, currency string, day uint64) (float64, error) {
	if chainId != 1 && chainId != 100 {
		// Don't show a historical price for testnets
		return 0.0, nil
	}
	if currency == utils.Config.Frontend.ClCurrency {
		currency = "USD"
	}
	currency = strings.ToLower(currency)

	if currency != "eur" && currency != "usd" && currency != "rub" && currency != "cny" && currency != "cad" && currency != "jpy" && currency != "gbp" && currency != "aud" {
		return 0.0, fmt.Errorf("currency %v not supported", currency)
	}

	// Convert day to ts
	genesisTime := time.Unix(int64(utils.Config.Chain.GenesisTimestamp), 0)
	dayStartGenesisTime := time.Date(genesisTime.Year(), genesisTime.Month(), genesisTime.Day(), 0, 0, 0, 0, time.UTC)
	ts := dayStartGenesisTime.Add(utils.Day * time.Duration(day))

	var value float64
	err := ReaderDb.Get(&value, fmt.Sprintf("SELECT %s FROM price WHERE ts = $1", currency), ts)
	if err != nil {
		return 0.0, err
	}
	return value, nil
}

func GetUserAPIKeyStatistics(apikey *string) (*types.ApiStatistics, error) {
	stats := &types.ApiStatistics{}

	query := `
	SELECT (
		SELECT 
			COALESCE(SUM(count), 0) as daily 
		FROM 
			api_statistics 
		WHERE 
			ts >= DATE_TRUNC('day', NOW()) AND apikey = $1 AND bucket = 'default'
	), (
		SELECT 
			COALESCE(SUM(count),0) as monthly 
		FROM 
			api_statistics 
		WHERE 
			ts >= DATE_TRUNC('month', NOW()) AND apikey = $1 AND bucket = 'default'
	)`

	err := FrontendWriterDB.Get(stats, query, apikey)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func GetSubsForEventFilter(eventName types.EventName) ([][]byte, map[string][]types.Subscription, error) {
	var subs []types.Subscription
	subQuery := `
		SELECT id, user_id, event_filter, last_sent_epoch, created_epoch, event_threshold, ENCODE(unsubscribe_hash, 'hex') as unsubscribe_hash, internal_state from users_subscriptions where event_name = $1
		`

	subMap := make(map[string][]types.Subscription, 0)
	err := FrontendWriterDB.Select(&subs, subQuery, utils.GetNetwork()+":"+string(eventName))
	if err != nil {
		return nil, nil, err
	}

	filtersEncode := make([][]byte, 0, len(subs))
	for _, sub := range subs {
		if _, ok := subMap[sub.EventFilter]; !ok {
			subMap[sub.EventFilter] = make([]types.Subscription, 0)
		}
		subMap[sub.EventFilter] = append(subMap[sub.EventFilter], types.Subscription{
			UserID:         sub.UserID,
			ID:             sub.ID,
			LastEpoch:      sub.LastEpoch,
			EventFilter:    sub.EventFilter,
			CreatedEpoch:   sub.CreatedEpoch,
			EventThreshold: sub.EventThreshold,
			State:          sub.State,
		})

		b, _ := hex.DecodeString(sub.EventFilter)
		filtersEncode = append(filtersEncode, b)
	}
	return filtersEncode, subMap, nil
}

// SaveDataTableState saves the state of the current datatable state update
func SaveDataTableState(user uint64, key string, state types.DataTableSaveState) error {
	ctx, done := context.WithTimeout(context.Background(), time.Second*30)
	defer done()

	// check how many table states are stored
	count := 0
	err := FrontendReaderDB.GetContext(ctx, &count, `
		SELECT count(*)
		FROM users_datatable
		WHERE user_id = $1
	`, user)
	if err != nil {
		return err
	}

	// only store the most recent 100 table states across all networks
	if count > 100 {
		_, err := FrontendWriterDB.ExecContext(ctx, `
			DELETE FROM users_datatable 
			WHERE user_id = $1 
			ORDER by updated_at asc 
			LIMIT 10
		`)
		if err != nil {
			return err
		}
	}
	// append network prefix
	key = utils.GetNetwork() + ":" + key

	_, err = FrontendWriterDB.ExecContext(ctx, `
		INSERT INTO 
			users_datatable (user_id, key, state) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (user_id, key) DO UPDATE SET state = $3, updated_at = now()
	`, user, key, state)

	return err
}

// GetDataTablesState retrieves the state for a given user and table
func GetDataTablesState(user uint64, key string) (*types.DataTableSaveState, error) {
	var state types.DataTableSaveState

	// append network prefix
	key = utils.GetNetwork() + ":" + key

	ctx, done := context.WithTimeout(context.Background(), time.Second*30)
	defer done()

	err := FrontendReaderDB.GetContext(ctx, &state, `
		SELECT state 
		FROM users_datatable
		WHERE user_id = $1 and key = $2
	`, user, key)

	return &state, err
}
