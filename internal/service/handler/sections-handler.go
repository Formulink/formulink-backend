package handler

import (
	"database/sql"
	"formulink-backend/internal/model"
	"formulink-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap/zapcore"
	"net/http"
)

type SectionHandler struct {
	db    *sql.DB
	redis *redis.Client
}

func NewSectionHandler(_db *sql.DB, _redis *redis.Client) *SectionHandler {
	return &SectionHandler{
		db:    _db,
		redis: _redis,
	}
}

func (sh *SectionHandler) GetSections(c echo.Context) error {
	var sections []model.Section

	rows, err := sh.db.Query("select * from sections")
	if err != nil {
		logger.Lg().Log(zapcore.InfoLevel, "can't parse data from db")
		return c.JSON(http.StatusInternalServerError, "idk what is wrong (maybe db is dead)`")
	}
	for rows.Next() {
		var section model.Section
		if err := rows.Scan(
			&section.SubjectId,
			&section.Name,
			&section.Description,
			&section.Id,
		); err != nil {
			logger.Lg().Log(zapcore.InfoLevel, "can't parse data from db")
			return c.JSON(http.StatusInternalServerError, "can't parse data from db")
		}
		sections = append(sections, section)
	}

	return c.JSON(http.StatusOK, sections)
}
