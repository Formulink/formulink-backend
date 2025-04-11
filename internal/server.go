package internal

import (
	"database/sql"
	"formulink-backend/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	e       *echo.Echo
	service *service.Service
}

func NewServer(db *sql.DB, redis *redis.Client) *Server {
	server := &Server{
		service: service.NewService(db, redis),
	}
	configureServer(server)
	return server
}

func configureServer(s *Server) {
	e := echo.New()
	e.GET("/", s.service.Hello)

	//auth
	e.GET("/auth", s.service.Auth)

	//sections
	e.GET("/sections", s.service.GetSections)
	e.GET("/:subject/sections", s.service.GetSectionsBySubjectId)
	e.GET("/subjects", s.service.GetSubjects)

	//formulas
	e.GET("/:id/formulas", s.service.GetFormulaByFormulaId)
	e.GET("/formulas/:id", s.service.GetFormulaById)
	e.GET("/formulas/fday", s.service.GetFormulaOfTheDay)
	e.GET("/formulas/all", s.service.GetAllFormulas)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://localhost:5173",
			"http://localhost:5174",
			"https://localhost:5174",
		},
		AllowHeaders: []string{"*"},
	}))

	s.e = e
}

func (s *Server) Start() {
	s.e.Logger.Fatal(s.e.Start(":8082"))
}
