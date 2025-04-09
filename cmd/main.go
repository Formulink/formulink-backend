package main

import (
	"formulink-backend/internal"
	"formulink-backend/pkg/logger"
)

func main() {
	logger.Init()

	server := internal.NewServer()
	server.Start()
}
