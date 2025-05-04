package service

import (
	"database/sql"
	"formulink-backend/internal/config"
	"formulink-backend/internal/service/handler"
	handler2 "formulink-backend/internal/service/handler/auth"
	"formulink-backend/internal/service/handler/conversations"
	"formulink-backend/internal/service/handler/formulas"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type Service struct {
	db                  *sql.DB
	redis               *redis.Client
	authHandler         *handler2.AuthHandler
	formulaHandler      *formulas.FormulaHandler
	formulaLikeHandler  *formulas.FormulaLikesHandler
	sectionHandler      *handler.SectionHandler
	taskHandler         *handler.TaskHandler
	mistralHandler      *handler.MistralHandler
	conversationHandler *conversations.ConversationHandler
}

func NewService(db *sql.DB, redis *redis.Client, cfg *config.MainConfig) *Service {
	return &Service{
		db:                  db,
		redis:               redis,
		authHandler:         handler2.NewAuthHandler(handler2.NewUserService(db)),
		formulaHandler:      formulas.NewFormulaHandler(db, redis),
		formulaLikeHandler:  formulas.NewFormulaLikesHandler(db),
		sectionHandler:      handler.NewSectionHandler(db, redis),
		taskHandler:         handler.NewTaskHandler(db, redis),
		mistralHandler:      handler.NewMistralHandler(db, redis, cfg.MistralApiKey),
		conversationHandler: conversations.NewConversationHandler(db),
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

// likes functions
func (s *Service) HandleLike(c echo.Context) error {
	return s.formulaLikeHandler.HandleLike(c)
}

func (s *Service) GetStatus(c echo.Context) error {
	return s.formulaLikeHandler.GetStatus(c)
}

func (s *Service) GetAllLikedFormulas(c echo.Context) error {
	return s.formulaLikeHandler.GetAllLikedFormulas(c)
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

// messages

func (s *Service) CreateNewConversation(c echo.Context) error {
	return s.conversationHandler.CreateNewConversation(c)
}

func (s *Service) GetConversation(c echo.Context) error {
	return s.conversationHandler.GetConversation(c)
}

func (s *Service) GetAllConversations(c echo.Context) error {
	return s.conversationHandler.GetAllConversations(c)
}

func (s *Service) AddMessage(c echo.Context) error {
	return s.conversationHandler.AddMessage(c)
}

func (s *Service) DeleteConversation(c echo.Context) error {
	return s.conversationHandler.DeleteConversation(c)
}
