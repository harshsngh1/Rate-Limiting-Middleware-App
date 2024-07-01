package routes

import (
	"rate_limiting_middleware/handlers"

	"github.com/labstack/echo/v4"
)

func SetupServerRoutes(e *echo.Echo) {
	e.GET("/endpoint1", handlers.Endpoint1)
	e.GET("/endpoint2", handlers.Endpoint2)
	SetupRateLimitRoutes(e)
}
