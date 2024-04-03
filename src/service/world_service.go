package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func World(c echo.Context) error {
	// HTTP 응답
	return c.String(http.StatusOK, "World!")
}
