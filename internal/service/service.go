package service

import (
	"formulink-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service struct {
	e *echo.Echo
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ConfigureEcho() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	s.e = e
}

func (s *Service) StartServer() {
	if s.e == nil {
		logger.Lg().Fatal("can't start server because it didn't configured")
		return
	}
	s.e.Logger.Fatal(s.e.Start(":8081"))
}
