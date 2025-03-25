package sqlitestorage

import (
	"os"
	"database/sql"
	"fmt"

	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	packr "github.com/gobuffalo/packr/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	migrate "github.com/rubenv/sql-migrate"
	_ "github.com/mattn/go-sqlite3"
)

// RunMigrationsUp migrate up.
func RunMigrationsUp(cfg Config) error {
	return runMigrations(cfg, migrate.Up)
}

// RunMigrationsDown migrate down.
func RunMigrationsDown(cfg Config) error {
	return runMigrations(cfg, migrate.Down)
}

// ResetDB.
func ResetDB(cfg Config) error {
	file := fmt.Sprintf("file:%s?_txlock=exclusive&_foreign_keys=on&_journal_mode=WAL", cfg.DBFile)
	log.Debug("Resetting database: ", file)
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS token_wrapped;")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS remove_exit_root;")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS exit_root;")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS claim;")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS deposit;")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS block;")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS root;")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS rollup_exit;")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS rht;")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS gorp_migrations;")
	if err != nil {
		return err
	}
	return nil
}

// runMigrations will execute pending migrations if needed to keep
// the database updated with the latest changes in either direction of up or down.
func runMigrations(cfg Config, direction migrate.MigrationDirection) error {
	file := fmt.Sprintf("file:%s?_txlock=exclusive&_foreign_keys=on&_journal_mode=WAL", cfg.DBFile)
	log.Debug("Running migrations: ", file)
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return err
	}

	var migrations = &migrate.PackrMigrationSource{Box: packr.New("bridge-service-db-migrations", "./migrations")}
	nMigrations, err := migrate.Exec(db, "sqlite3", migrations, direction)
	if err != nil {
		return err
	}

	log.Info("successfully ran ", nMigrations, " migrations")
	return nil
}

// InitOrReset will initializes the db running the migrations or
// will reset all the known data and rerun the migrations
func InitOrReset(cfg Config) error {
	// connect to database
	_, err := NewSQLiteStorage(cfg)
	if err != nil {
		return err
	}

	// run migrations
	if err := ResetDB(cfg); err != nil {
		return err
	}
	return RunMigrationsUp(cfg)
}

// NewConfigFromEnv creates config from standard postgres environment variables,
func NewConfigFromEnv() Config {
	return Config{
		DBFile: getEnv("ZKEVM_BRIDGE_DATABASE_SQLITESTORAGE_DBFILE", "/tmp/database.sqlite"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

// nolint
var (
	String, _ = abi.NewType("string", "", nil)
	Uint8, _  = abi.NewType("uint8", "", nil)
)

func getDecodedToken(metadata []byte) (*etherman.TokenMetadata, error) {
	//nolint
	args := abi.Arguments{
		{"name", String, false},
		{"symbol", String, false},
		{"decimals", Uint8, false},
	}
	token := make(map[string]interface{})
	err := args.UnpackIntoMap(token, metadata)
	if err != nil {
		return nil, err
	}

	return &etherman.TokenMetadata{
		Name:     token["name"].(string),
		Symbol:   token["symbol"].(string),
		Decimals: token["decimals"].(uint8),
	}, nil
}
