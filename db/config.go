package db

import (
	"github.com/0xPolygon/zkevm-bridge-service/db/pgstorage"
)

// Config struct
type Config struct {
	// Database type
	Database string `mapstructure:"Database"`

	PgStorage pgstorage.Config
}
