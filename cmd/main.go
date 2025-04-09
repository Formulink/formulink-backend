package main

import (
	"formulink-backend/internal/service"
	"formulink-backend/pkg/logger"
)

func main() {
	logger.Init()

	svc := service.NewService()
	svc.StartServer()
}
