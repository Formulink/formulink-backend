package internal

import (
	"database/sql"
	"formulink-backend/internal/config"
	"formulink-backend/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	e       *echo.Echo
	service *service.Service
}

func NewServer(db *sql.DB, redis *redis.Client, cfg *config.MainConfig) *Server {
	server := &Server{
		service: service.NewService(db, redis, cfg),
	}
	configureServer(server)
	return server
}

func configureServer(s *Server) {
	e := echo.New()
	e.GET("/", s.service.Hello)

	//auth
	e.POST("/auth", s.service.Auth)

	//sections
	e.GET("/sections", s.service.GetSections)
	e.GET("/:subject/sections", s.service.GetSectionsBySubjectId)
	e.GET("/subjects", s.service.GetSubjects)

	//formulas
	e.GET("/:id/formulas", s.service.GetFormulaByFormulaId)
	e.GET("/formulas/:id", s.service.GetFormulaById)
	e.GET("/formulas/fday", s.service.GetFormulaOfTheDay)
	e.GET("/formulas/all", s.service.GetAllFormulas)

	//likes
	e.POST("/like", s.service.HandleLike)
	e.POST("/like-status", s.service.GetStatus)
	e.GET("/liked-formulas/:user_id", s.service.GetAllLikedFormulas)

	//task
	e.GET("/tasks/:id", s.service.GetTasksByFormulaId)
	e.GET("task/:id", s.service.GetTaskById)

	//neuro
	e.POST("/ai", s.service.MistralChat)

	//messages
	e.POST("/conversation/new", s.service.CreateNewConversation)
	e.GET("/conversation/:id", s.service.GetConversation)
	e.GET("/conversations/:user_id", s.service.GetAllConversations)
	e.POST("/conversation/message", s.service.AddMessage)
	e.DELETE("/conversation/:id", s.service.DeleteConversation)

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://localhost:5173",
			"http://localhost:5174",
			"https://localhost:5174",
			"0886-5-104-75-74.ngrok-free.app",
			"https://73cb-5-104-75-74.ngrok-free.app",
		},
		AllowHeaders: []string{"*"},
	}))

	s.e = e
}

func (s *Server) Start() error {
	err := s.e.Start(":8082")
	return err
}
