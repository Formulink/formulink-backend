package service

import (
	"database/sql"
	"formulink-backend/internal/service/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service struct {
	authHandler    *handler.AuthHandler
	formulaHandler *handler.FormulaHandler
	sectionHandler *handler.SectionHandler
}

func NewService(db *sql.DB) *Service {
	return &Service{
		authHandler:    handler.NewAuthHandler(db),
		formulaHandler: handler.NewFormulaHandler(db),
		sectionHandler: handler.NewSectionHandler(db),
	}
}

// auth functions
func (s *Service) Auth(c echo.Context) error {
	return s.authHandler.Auth(c)
}

// section functions
func (s *Service) GetSections(c echo.Context) error {
	return s.sectionHandler.GetSections(c)
}

// formula functions
func (s *Service) GetFormulaByFormulaId(c echo.Context) error {
	return s.formulaHandler.GetFormulasBySectionId(c)
}

func (s *Service) GetFormulaById(c echo.Context) error {
	return s.formulaHandler.GetFormulaById(c)
}

// other
func (s *Service) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "( ´ ꒳ ` )")
}
