package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Endpoint1(c echo.Context) error {
	return c.String(http.StatusOK, "This is a rate-limited endpoint.")
}

func Endpoint2(c echo.Context) error {
	return c.String(http.StatusOK, "This is another rate-limited endpoint.")
}
