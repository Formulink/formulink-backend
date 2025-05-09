package handler

import (
	"formulink-backend/internal/dto"
	"formulink-backend/pkg/logger"

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
		logger.Lg().Logf(0, "invalid request | err: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if req.TelegramId == 0 || req.Username == "" {
		logger.Lg().Logf(0, "missing tg_id or username")
		return echo.NewHTTPError(http.StatusBadRequest, "Missing telegram_id or username")
	}

	user, err := ah.userService.GetByTelegramID(req.TelegramId)
	if err != nil {
		logger.Lg().Logf(0, "ISE | err: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if user != nil {
		logger.Lg().Logf(0, "user succesfully logged in")
		return c.JSON(http.StatusOK, map[string]any{
			"user_id":         user.ID.String(),
			"need_onboarding": user.NeedOnboarding,
		})
	}

	newUser, err := ah.userService.CreateUser(req)
	if err != nil {
		logger.Lg().Logf(0, "ISE | err: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	logger.Lg().Logf(0, "ISE | err: %v", err)

	logger.Lg().Logf(0, "user succesfully created")
	return c.JSON(http.StatusCreated, map[string]any{
		"user_id":         newUser.ID.String(),
		"need_onboarding": newUser.NeedOnboarding,
	})
}

func (ah *AuthHandler) SetOnboardingFalse(c echo.Context) error {
	var req dto.OnboardingRequest
	if err := c.Bind(&req); err != nil {
		logger.Lg().Logf(0, "invalid request, err: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := ah.userService.SetNeedOnboardingFalse(req.TelegramId); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.NoContent(http.StatusOK)
}
