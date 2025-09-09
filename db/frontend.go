package db

import (
	"github.com/theQRL/qrl-beaconchain-explorer/types"

	"github.com/jmoiron/sqlx"
)

// FrontendWriterDB is a pointer to the auth-database
var FrontendReaderDB *sqlx.DB
var FrontendWriterDB *sqlx.DB

func MustInitFrontendDB(writer *types.DatabaseConfig, reader *types.DatabaseConfig, driverName string, databaseBrand string) {
	FrontendWriterDB, FrontendReaderDB = mustInitDB(writer, reader, driverName, databaseBrand)
}
