package handlers

import (
	"net/http"
	"rate_limiting_middleware/config"

	"github.com/labstack/echo/v4"
)

type RateLimitRequest struct {
	Endpoint string `json:"endpoint,omitempty"`
	IP       string `json:"ip,omitempty"`
	Limit    int    `json:"limit"`
}

func SetRateLimit(c echo.Context) error {
	req := new(RateLimitRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if req.Endpoint == "" && req.IP == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "either endpoint or IP must be specified"})
	}

	if req.IP == "" {
		req.IP = "default"
	}
	config.RateLimiterConfig.SetRateLimit(req.Endpoint, req.IP, req.Limit)

	return c.JSON(http.StatusOK, echo.Map{"message": "rate limit updated"})
}

func GetRateLimits(c echo.Context) error {
	return c.JSON(http.StatusOK, config.RateLimiterConfig.Limits)
}
