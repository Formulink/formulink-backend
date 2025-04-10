package handler

import (
	"database/sql"
	"formulink-backend/internal/dto"
	"formulink-backend/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(_db *sql.DB) *AuthHandler {
	return &AuthHandler{
		db: _db,
	}
}

func (ah *AuthHandler) Auth(c echo.Context) error {
	var user dto.CreateUserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE telegramid = $1)`
	err := ah.db.QueryRow(query, user.TelegramId).Scan(&exists)
	if err != nil {
		logger.Lg().Fatalf("err: %v", err)
		return err
	}

	if exists {
		return SendData(c)
	}

	uid := uuid.New()
	query = `INSERT INTO users values ($1, $2, $3, $4) ON CONFLICT DO NOTHING`

	_, err = ah.db.Exec(query, uid, user.TelegramId, user.Username, time.Now())
	if err != nil {
		logger.Lg().Fatalf("err: %v", err)
		return err
	}
	return c.JSON(http.StatusOK, dto.CreateUserResponse{
		Id:       uid,
		HaveData: false,
	})
}

func SendData(c echo.Context) error {
	return c.JSON(http.StatusOK, ":)")
}
