package main

import "formulink-backend/internal/service"

func main() {
	svc := service.NewService()
	svc.StartServer()
}
