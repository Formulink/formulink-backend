package main

import (
	"fmt"
	"formulink-backend/internal"
	"formulink-backend/internal/config"
	"formulink-backend/pkg/db/postgres"
	"formulink-backend/pkg/db/redis"
	"formulink-backend/pkg/logger"
	"formulink-backend/pkg/mistral"
)

func main() {
	logger.Init()

	cfg, err := config.NewMainConfig()
	if err != nil {
		logger.Lg().Logf(0, "err: %v", err)
	}

	pgConn, err := postgres.New(cfg.POSTGRES)
	if err != nil {
		logger.Lg().Logf(0, "can't connect to db | err: %v", err)
	}
	redisClient := redis.NewRedisConn(cfg.REDIS)

	mistralClient := mistral.CreateMistralClient(cfg.MistralApiKey)
	fmt.Println(mistralClient)

	server := internal.NewServer(pgConn, redisClient, cfg)
	if err := server.Start(); err != nil {
		logger.Lg().Logf(0, "can' connect to db | err : %v", err)
	}
}
