package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

type Config struct {
	Host         string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" env-default:"formulink_postgres-container"`
	Port         uint16 `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" env-default:"5432"`
	Username     string `yaml:"POSTGRES_USER" env:"POSTGRES_USER" env-default:"root"`
	Password     string `yaml:"POSTGRES_PASSWORD" env:"POSTGRES_PASSWORD" env-default:"password"`
	Database     string `yaml:"POSTGRES_DB" env:"POSTGRES_DB" env-default:"postgres"`
	MinConns     int32  `yaml:"POSTGRES_MIN_CONNS" env:"POSTGRES_MIN_CONNS" env-default:"5"`
	MaxConns     int32  `yaml:"POSTGRES_MAX_CONNS" env:"POSTGRES_MAX_CONNS" env-default:"10"`
	SearchSchema string `yaml:"POSTGRES_MAIN_SCHEMA" env:"POSTGRES_MAIN_SCHEMA" env-default:"base_schema"`
}

func New(config Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, err := sql.Open("postgres", connStr)

	connStr2 := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.SearchSchema,
	)

	m, err := migrate.New("file://db/migrations", connStr2)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("can't create migration instance: %w", err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("can't apply migrations: %w", err)
	}

	return db, nil
}

