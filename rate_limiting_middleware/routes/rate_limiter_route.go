package routes

import (
	"rate_limiting_middleware/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRateLimitRoutes(e *echo.Echo) {
	e.POST("/set-rate-limit", handlers.SetRateLimit)
	e.GET("/get-rate-limits", handlers.GetRateLimits)
}
