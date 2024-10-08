package main

import (
	"log"
	"rate_limiting_middleware/config"
	"rate_limiting_middleware/middleware"
	"rate_limiting_middleware/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	e := echo.New()
	e.Use(middleware.RateLimitMiddleware)
	routes.SetupServerRoutes(e)
	log.Printf("Server starting on %s", cfg.Port)
	e.Logger.Fatal(e.Start(cfg.Port))
}
