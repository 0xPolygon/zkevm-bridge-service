package db

import (
	"github.com/0xPolygonHermez/zkevm-bridge-service/db/pgstorage"
	"github.com/0xPolygonHermez/zkevm-bridge-service/db/sqlitestorage"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils/gerror"
)

// Storage interface
type Storage interface{}

// NewStorage creates a new Storage
func NewStorage(cfg Config) (Storage, error) {
	if cfg.Database == "postgres" {
		pg, err := pgstorage.NewPostgresStorage(pgstorage.Config{
			Name:     cfg.PgStorage.Name,
			User:     cfg.PgStorage.User,
			Password: cfg.PgStorage.Password,
			Host:     cfg.PgStorage.Host,
			Port:     cfg.PgStorage.Port,
			MaxConns: cfg.PgStorage.MaxConns,
		})
		return pg, err
	} else if cfg.Database == "sqlite" {
		return sqlitestorage.NewSQLiteStorage(sqlitestorage.Config{
			DBFile: cfg.SqliteStorage.DBFile,})
	}
	return nil, gerror.ErrStorageNotRegister
}

// RunMigrations will execute pending migrations if needed to keep
// the database updated with the latest changes
func RunMigrations(cfg Config) error {
	if cfg.Database == "postgres" {
		config := pgstorage.Config{
			Name:     cfg.PgStorage.Name,
			User:     cfg.PgStorage.User,
			Password: cfg.PgStorage.Password,
			Host:     cfg.PgStorage.Host,
			Port:     cfg.PgStorage.Port,
		}
		return pgstorage.RunMigrationsUp(config)
	} else if cfg.Database == "sqlite" {
		config := sqlitestorage.Config{
			DBFile: cfg.SqliteStorage.DBFile,
		}
		return sqlitestorage.RunMigrationsUp(config)
	}
	return nil
}
