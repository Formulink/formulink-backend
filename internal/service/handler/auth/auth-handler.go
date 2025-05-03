package handler

import (
	"formulink-backend/internal/dto"

	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userService *UserService
}

func NewAuthHandler(userService *UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (ah *AuthHandler) Auth(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if req.TelegramId == 0 || req.Username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing telegram_id or username")
	}

	user, err := ah.userService.GetByTelegramID(req.TelegramId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if user != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"user_id": user.ID.String(),
		})
	}

	newUser, err := ah.userService.CreateUser(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"user_id": newUser.ID.String(),
	})
}
