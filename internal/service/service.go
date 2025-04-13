package service

import (
	"database/sql"
	"formulink-backend/internal/config"
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
	taskHandler    *handler.TaskHandler
	mistralHandler *handler.MistralHandler
}

func NewService(db *sql.DB, redis *redis.Client, cfg *config.MainConfig) *Service {
	return &Service{
		db:             db,
		redis:          redis,
		authHandler:    handler.NewAuthHandler(db),
		formulaHandler: handler.NewFormulaHandler(db, redis),
		sectionHandler: handler.NewSectionHandler(db, redis),
		taskHandler:    handler.NewTaskHandler(db, redis),
		mistralHandler: handler.NewMistralHandler(db, redis, cfg.MistralApiKey),
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
func (s *Service) GetAllFormulas(c echo.Context) error {
	return s.formulaHandler.GetAllFormulas(c)
}

func (s *Service) GetFormulaByFormulaId(c echo.Context) error {
	return s.formulaHandler.GetFormulasBySectionId(c)
}

func (s *Service) GetFormulaById(c echo.Context) error {
	return s.formulaHandler.GetFormulaById(c)
}

func (s *Service) GetFormulaOfTheDay(c echo.Context) error {
	return s.formulaHandler.GetFormulaOfTheDay(c)
}

// tasks
func (s *Service) GetTasksByFormulaId(c echo.Context) error {
	return s.taskHandler.GetTasksByFormulaId(c)
}
func (s *Service) GetTaskById(c echo.Context) error {
	return s.taskHandler.GetTaskById(c)
}

// neuro
func (s *Service) MistralChat(c echo.Context) error {
	return s.mistralHandler.Chat(c)
}

// other
func (s *Service) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "( ´ ꒳ ` )")
}
