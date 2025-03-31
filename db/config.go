package db

import (
	"github.com/0xPolygonHermez/zkevm-bridge-service/db/pgstorage"
)

// Config struct
type Config struct {
	// Database type
	Database string `mapstructure:"Database"`

	PgStorage pgstorage.Config
}
