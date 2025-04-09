package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
