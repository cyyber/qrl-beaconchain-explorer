package db

import (
	"context"
	"database/sql"
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
