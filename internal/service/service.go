package service

import (
	"formulink-backend/internal/service/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service struct {
	authHandler *handler.AuthHandler
}

func NewService() *Service {
	return &Service{
		authHandler: handler.NewAuthHandler(),
	}
}

func (s *Service) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "( ´ ꒳ ` )")
}
