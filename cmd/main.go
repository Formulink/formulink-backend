package main

import (
	"formulink-backend/internal"
	"formulink-backend/internal/config"
	"formulink-backend/pkg/db/postgres"
	"formulink-backend/pkg/db/redis"
	"formulink-backend/pkg/logger"
)

func main() {
	logger.Init()

	cfg, err := config.NewMainConfig()
	if err != nil {
		logger.Lg().Fatalf("err: %v", err)
	}

	pgConn, err := postgres.NewPostgres(cfg.POSTGRES)
	redisClient := redis.NewRedisConn(cfg.REDIS)

	server := internal.NewServer(pgConn, redisClient)
	server.Start()
}
