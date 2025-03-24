package db

import (
	"github.com/0xPolygonHermez/zkevm-bridge-service/db/pgstorage"
	"github.com/0xPolygonHermez/zkevm-bridge-service/db/sqlitestorage"
)

// Config struct
type Config struct {
	// Database type
	Database string `mapstructure:"Database"`

	PgStorage pgstorage.Config

	SqliteStorage sqlitestorage.Config
}
