package utils

import (
	"github.com/labstack/echo/v4"
)

func ValidateRequest(c echo.Context) bool {
	return c.Request().ContentLength == 0
}
