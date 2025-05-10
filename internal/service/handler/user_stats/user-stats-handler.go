package user_stats

import (
	"database/sql"
	"formulink-backend/internal/dto"
	"formulink-backend/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserStatsHandler struct {
	db   *sql.DB
	repo *UserStatsRepository
}

func NewUserStatsHandler(db *sql.DB) *UserStatsHandler {
	return &UserStatsHandler{db: db, repo: NewUserStatsRepository(db)}
}

func (ush *UserStatsHandler) CreateNewRecord(c echo.Context) error {
	var req dto.NewRecordRequest
	if err := c.Bind(&req); err != nil {
		logger.Lg().Logf(0, "invalid request | err: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	resp, err := ush.repo.createNewRecord(req)
	if err != nil {
		logger.Lg().Logf(0, "can't insert stats into bd | err: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (ush *UserStatsHandler) GetUserRecords(c echo.Context) error {
	id, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		logger.Lg().Logf(0, "invalid request | err: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	stats := ush.repo.getUserStats(id)
	if stats == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, stats)
}
