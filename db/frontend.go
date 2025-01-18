package db

import (
	"context"
	"database/sql"
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

func addParamToQuery(query, param string) string {
	var result = query
	if result != "" {
		result += ","
	}
	result += param
	return result
}

func NewTransaction() (*sql.Tx, error) {
	return FrontendWriterDB.Begin()
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
