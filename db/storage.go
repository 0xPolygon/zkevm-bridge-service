package db

import (
	"context"
	"fmt"

	"github.com/0xPolygon/zkevm-bridge-service/db/pgstorage"
	"github.com/0xPolygon/zkevm-bridge-service/utils/gerror"
)

// Storage interface
type Storage interface{}

// NewStorage creates a new Storage
func NewStorage(ctx context.Context, cfg Config) (Storage, error) {
	if cfg.Database == "postgres" {
		pg, err := pgstorage.NewPostgresStorage(ctx, pgstorage.Config{
			Name:     cfg.PgStorage.Name,
			User:     cfg.PgStorage.User,
			Password: cfg.PgStorage.Password,
			Host:     cfg.PgStorage.Host,
			Port:     cfg.PgStorage.Port,
			MaxConns: cfg.PgStorage.MaxConns,
		})
		return pg, err
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
	}
	return fmt.Errorf("database type not supported")
}
