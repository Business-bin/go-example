package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Hello
// @Summary
// @Description
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello,"
// @Failure 400
// @Failure 500
// @Router /hellos [get]
// @Tags Hello
func Hello(c echo.Context) error {
	// HTTP 응답
	return c.String(http.StatusOK, "Hello,")
}
