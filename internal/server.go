package internal

import (
	"database/sql"
	"formulink-backend/internal/config"
	"formulink-backend/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type Server struct {
	e       *echo.Echo
	service *service.Service
}

type Res struct {
	CounterMain                     int
	CounterStudentRules             int
	CounterLyceumRegulation         int
	CounterResponsibleInfo          int
	CounterLicenseRegistry          int
	CounterAccreditation            int
	CounterEducationLicense         int
	CounterCharterChanges           int
	CounterEGRUL                    int
	CounterEducationActivityLicense int
	CounterAccreditationAppendix    int
}

var counters = Res{}

func NewServer(db *sql.DB, redis *redis.Client, cfg *config.MainConfig) *Server {
	server := &Server{
		service: service.NewService(db, redis, cfg),
	}
	configureServer(server)
	return server
}

func configureServer(s *Server) {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", s.service.Hello)

	//auth
	e.POST("/auth", s.service.Auth)
	e.POST("/onboarding-requirement", s.service.UpdateOnboarding)

	//user-stasts
	e.GET("/user/stats/:user_id", s.service.GetUserRecords)
	e.POST("/user/stats/new", s.service.CreateNewRecord)

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

	e.GET("/lyceum-6/main", func(c echo.Context) error {
		counters.CounterMain++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru")
	})
	//here
	e.GET("/lyceum-6/student-rules", func(c echo.Context) error {
		counters.CounterStudentRules++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/wp-content/uploads/2023/08/%D0%9F%D1%80%D0%B0%D0%B2%D0%B8%D0%BB%D0%B0-%D0%B2%D0%BD%D1%83%D1%82%D1%80%D0%B5%D0%BD%D0%BD%D0%B5%D0%B3%D0%BE-%D1%80%D0%B0%D1%81%D0%BF%D0%BE%D1%80%D1%8F%D0%B4%D0%BA%D0%B0-%D0%BE%D0%B1%D1%83%D1%87%D0%B0%D1%8E%D1%89%D0%B8%D1%85%D1%81%D1%8F-2024.pdf")
	})

	//here
	e.GET("/lyceum-6/lyceum-regulation", func(c echo.Context) error {
		counters.CounterLyceumRegulation++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/wp-content/uploads/2023/08/%D0%A3%D1%81%D1%82%D0%B0%D0%B2-%D0%9B%D0%B8%D1%86%D0%B5%D1%8F-6-15.09.2021.pdf")
	})

	e.GET("/lyceum-6/responsible-info", func(c echo.Context) error {
		counters.CounterResponsibleInfo++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/2024/09/16/%d0%b8%d0%bd%d1%84%d0%be%d1%80%d0%bc%d0%b0%d1%86%d0%b8%d1%8f-%d0%be%d0%b1-%d0%be%d1%82%d0%b2%d0%b5%d1%82%d1%81%d1%82%d0%b2%d0%b5%d0%bd%d0%bd%d1%8b%d1%85-%d0%b7%d0%b0-%d0%be%d1%80%d0%b3%d0%b0%d0%bd/")
	})

	e.GET("/lyceum-6/license-registry", func(c echo.Context) error {
		counters.CounterLicenseRegistry++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/2024/09/02/%d0%b2%d1%8b%d0%bf%d0%b8%d1%81%d0%ba%d0%b0-%d0%b8%d0%b7-%d1%80%d0%b5%d0%b5%d1%81%d1%82%d1%80%d0%b0-%d0%bb%d0%b8%d1%86%d0%b5%d0%bd%d0%b7%d0%b8%d0%b9-%e2%84%96-%d0%bb035-01255-50-00215348-2/")
	})

	e.GET("/lyceum-6/accreditation", func(c echo.Context) error {
		counters.CounterAccreditation++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/2024/09/02/%d1%81%d0%b2%d0%b8%d0%b4%d0%b5%d1%82%d0%b5%d0%bb%d1%8c%d1%81%d1%82%d0%b2%d0%be-%d0%be-%d0%b3%d0%be%d1%81%d1%83%d0%b4%d0%b0%d1%80%d1%81%d1%82%d0%b2%d0%b5%d0%bd%d0%bd%d0%be%d0%b9-%d0%b0%d0%ba%d0%ba-3/")
	})

	e.GET("/lyceum-6/education-license", func(c echo.Context) error {
		counters.CounterEducationLicense++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/2024/09/03/%d0%bb%d0%b8%d1%86%d0%b5%d0%bd%d0%b7%d0%b8%d1%8f-%d0%bd%d0%b0-%d0%be%d1%81%d1%83%d1%89%d0%b5%d1%81%d1%82%d0%b2%d0%bb%d0%b5%d0%bd%d0%b8%d0%b5-%d0%be%d0%b1%d1%80%d0%b0%d0%b7%d0%be%d0%b2%d0%b0%d1%82/")
	})

	e.GET("/lyceum-6/charter-changes", func(c echo.Context) error {
		counters.CounterCharterChanges++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/2024/11/15/%d0%b8%d0%b7%d0%bc%d0%b5%d0%bd%d0%b5%d0%bd%d0%b8%d1%8f-%d0%b2-%d1%83%d1%81%d1%82%d0%b0%d0%b2-%d0%bc%d0%b0%d0%be%d1%83-%d0%be%d0%b4%d0%b8%d0%bd%d1%86%d0%be%d0%b2%d1%81%d0%ba%d0%be%d0%b3%d0%be-%d0%bb/")
	})

	e.GET("/lyceum-6/egrul", func(c echo.Context) error {
		counters.CounterEGRUL++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/2023/08/15/%d0%bb%d0%b8%d1%81%d1%82-%d0%b7%d0%b0%d0%bf%d0%b8%d1%81%d0%b8-%d0%b5%d0%b3%d1%80%d1%8e%d0%bb-%d0%be-%d0%b2%d0%bd%d0%b5%d1%81%d0%b5%d0%bd%D0%b8%D0%b8-%d0%b7%d0%b0%d0%bf%d0%b8%d1%81%D0%b8-%d0%be%d0%b1/")
	})

	e.GET("/lyceum-6/education-activity-license", func(c echo.Context) error {
		counters.CounterEducationActivityLicense++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/2024/09/02/%d0%bb%d0%b8%d1%86%d0%b5%d0%bd%d0%b7%d0%b8%d1%8f-%d0%bd%d0%b0-%d0%bf%d1%80%d0%b0%d0%b2%d0%be-%d0%b2%d0%b5%d0%b4%d0%b5%d0%bd%d0%b8%d1%8f-%d0%be%d0%b1%d1%80%d0%b0%d0%b7%d0%be%d0%b2%d0%b0%d1%82%d0%b5/")
	})

	e.GET("/lyceum-6/accreditation-appendix", func(c echo.Context) error {
		counters.CounterAccreditationAppendix++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/2024/09/02/%d1%81%d0%b2%d0%b8%d0%b4%d0%b5%d1%82%d0%b5%d0%bb%d1%8c%d1%81%d1%82%d0%b2%d0%be-%d0%be-%d0%b3%d0%be%d1%81%d1%83%d0%b4%d0%b0%d1%80%d1%81%d1%82%d0%b2%d0%b5%d0%bd%d0%bd%d0%be%d0%b9-%d0%b0%d0%ba%d0%ba-4/")
	})

	e.GET("/counters", func(c echo.Context) error {
		return c.JSON(http.StatusOK, counters)
	})

	//CORS

	s.e = e
}

func (s *Server) Start() error {
	err := s.e.Start(":8082")
	return err
}
