package handler

import (
	"database/sql"
	"formulink-backend/internal/model"
	"formulink-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strconv"
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

func (sh *SectionHandler) GetSectionsBySubjectId(c echo.Context) error {
	var sections []model.Section

	id, err := strconv.Atoi(c.Param("subject"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid subject id")
	}

	query := `SELECT * from sections where subjectid = $1`
	rows, err := sh.db.Query(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "can't send queries to db")
	}

	for rows.Next() {
		var section model.Section
		err = rows.Scan(&section.SubjectId, &section.Name, &section.Description, &section.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't parse data from db")
		}

		sections = append(sections, section)
	}

	return c.JSON(http.StatusOK, sections)
}

func (sh *SectionHandler) GetSubjects(c echo.Context) error {
	var subjects []model.Subject

	query := `SELECT * from subject`
	rows, err := sh.db.Query(query)
	if err != nil {
		logger.Lg().Logf(zapcore.InfoLevel, "err: %v", err)
	}

	for rows.Next() {
		var subject model.Subject
		err = rows.Scan(&subject.Id, &subject.Name)
		if err != nil {
			logger.Lg().Logf(zapcore.InfoLevel, "err: %v", err)
			return err
		}
		subjects = append(subjects, subject)
	}

	return c.JSON(http.StatusOK, subjects)
}
