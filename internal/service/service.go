package service

import (
	"database/sql"
	"formulink-backend/internal/service/handler"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type Service struct {
	db             *sql.DB
	redis          *redis.Client
	authHandler    *handler.AuthHandler
	formulaHandler *handler.FormulaHandler
	sectionHandler *handler.SectionHandler
}

func NewService(db *sql.DB, redis *redis.Client) *Service {
	return &Service{
		db:             db,
		redis:          redis,
		authHandler:    handler.NewAuthHandler(db),
		formulaHandler: handler.NewFormulaHandler(db, redis),
		sectionHandler: handler.NewSectionHandler(db, redis),
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

func (s *Service) GetSectionsBySubjectId(c echo.Context) error {
	return s.sectionHandler.GetSectionsBySubjectId(c)
}

func (s *Service) GetSubjects(c echo.Context) error {
	return s.sectionHandler.GetSubjects(c)
}

// formula functions
func (s *Service) GetFormulaByFormulaId(c echo.Context) error {
	return s.formulaHandler.GetFormulasBySectionId(c)
}

func (s *Service) GetFormulaById(c echo.Context) error {
	return s.formulaHandler.GetFormulaById(c)
}

func (s *Service) GetFormulaOfTheDay(c echo.Context) error {
	return s.formulaHandler.GetFormulaOfTheDay(c)
}

// other
func (s *Service) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "( ´ ꒳ ` )")
}
