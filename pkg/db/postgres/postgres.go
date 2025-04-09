package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PgConfig struct {
	Host         string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" env-default:"localhost"`
	Port         uint16 `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" env-default:"5432"`
	Username     string `yaml:"POSTGRES_USER" env:"POSTGRES_USER" env-default:"root"`
	Password     string `yaml:"POSTGRES_PASSWORD" env:"POSTGRES_PASSWORD" env-default:"1234"`
	Database     string `yaml:"POSTGRES_DB" env:"POSTGRES_DB" env-default:"postgres"`
	MinConns     int32  `yaml:"POSTGRES_MIN_CONNS" env:"POSTGRES_MIN_CONNS" env-default:"5"`
	MaxConns     int32  `yaml:"POSTGRES_MAX_CONNS" env:"POSTGRES_MAX_CONNS" env-default:"10"`
	SearchSchema string `yaml:"POSTGRES_MAIN_SCHEMA" env:"POSTGRES_MAIN_SCHEMA" env-default:"public"`
}

func NewPostgres(c PgConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("can't connect to database | err:  %v", err)
	}

	// TODO: realise migrate
	//m, err := migrate.New(
	//	"file://db/migrations",
	//	connStr,
	//)
	//if err != nil {
	//	return nil, fmt.Errorf("can't create migration | err: %v", err)
	//}
	//
	//if err := m.Up(); err != nil {
	//	if err == migrate.ErrNoChange {
	//		log.Println("no new migrations")
	//	} else {
	//		return nil, fmt.Errorf("can't migrate database | err: %v", err)
	//	}
	//}

	return conn, nil

}
