package service

import (
	"database/sql"
	"formulink-backend/internal/service/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service struct {
	authHandler *handler.AuthHandler
}

func NewService(db *sql.DB) *Service {
	return &Service{
		authHandler: handler.NewAuthHandler(db),
	}
}

func (s *Service) Auth(c echo.Context) error {
	return s.authHandler.Auth(c)
}

func (s *Service) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "( ´ ꒳ ` )")
}
