package postgres

import (
    "database/sql"
    "errors"
    "fmt"
    "github.com/golang-migrate/migrate"
    _ "github.com/golang-migrate/migrate/database/postgres"
    _ "github.com/golang-migrate/migrate/source/file"
    _ "github.com/lib/pq"
    "net/url"
)

type Config struct {
    Host         string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" env-default:"formulink_postgres-container"`
    Port         uint16 `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" env-default:"5432"`
    Username     string `yaml:"POSTGRES_USER" env:"POSTGRES_USER" env-default:"root"`
    Password     string `yaml:"POSTGRES_PASSWORD" env:"POSTGRES_PASSWORD" env-default:"password"`
    Database     string `yaml:"POSTGRES_DB" env:"POSTGRES_DB" env-default:"formulink"`
    MinConns     int32  `yaml:"POSTGRES_MIN_CONNS" env:"POSTGRES_MIN_CONNS" env-default:"5"`
    MaxConns     int32  `yaml:"POSTGRES_MAX_CONNS" env:"POSTGRES_MAX_CONNS" env-default:"10"`
    SearchSchema string `yaml:"POSTGRES_MAIN_SCHEMA" env:"POSTGRES_MAIN_SCHEMA" env-default:"base_schema"`
}

func New(config Config) (*sql.DB, error) {
    // Формируем строку подключения
    connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
        config.Username,
        config.Password,
        config.Host,
        config.Port,
        config.Database,
    )

    // Открываем соединение с базой
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Пингуем базу данных
    if err := db.Ping(); err != nil {
        return nil, err
    }

    // Создаем строку подключения с дополнительным параметром для схемы поиска
    user := url.QueryEscape(config.Username)
    pass := url.QueryEscape(config.Password)
    schema := url.QueryEscape(config.SearchSchema)

    connStr2 := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s",
        user,
        pass,
        config.Host,
        config.Port,
        config.Database,
        schema,
    )

    // Создаем объект миграции
    m, err := migrate.New("file:///app/db/migrations", connStr2)
    if err != nil {
        return nil, err
    }

    // Принудительно устанавливаем версию 2 (чтобы исправить грязную версию)
    err = m.Force(2)
    if err != nil {
        return nil, err
    }

    // Выполняем миграции (если нет изменений, они не будут применены)
    if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
        return nil, err
    }

    // Возвращаем объект базы данных
    return db, nil
}

