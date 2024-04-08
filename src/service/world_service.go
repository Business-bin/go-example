package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// World
// @Summary
// @Description
// @Accept json
// @Produce json
// @Success 200 {string} string "World!"
// @Failure 400
// @Failure 500
// @Router /worlds [get]
// @Tags World
func World(c echo.Context) error {
	// HTTP 응답
	return c.String(http.StatusOK, "World!")
}
