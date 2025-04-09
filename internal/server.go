package internal

import (
	"formulink-backend/internal/service"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e       *echo.Echo
	service *service.Service
}

func NewServer() *Server {
	server := &Server{
		service: service.NewService(),
	}
	configureServer(server)
	return server
}

func configureServer(s *Server) {
	e := echo.New()
	e.GET("/", s.service.Hello)
	s.e = e
}

func (s *Server) Start() {
	s.e.Logger.Fatal(s.e.Start(":8081"))
}
