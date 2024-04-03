package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(c echo.Context) error {
	// HTTP 응답
	return c.String(http.StatusOK, "Hello,")
}
