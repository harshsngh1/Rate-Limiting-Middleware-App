package tests

import (
	"net/http"
	"net/http/httptest"
	"rate_limiting_middleware/config"
	"rate_limiting_middleware/middleware"
	"rate_limiting_middleware/routes"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RateLimitMiddleware)
	routes.SetupServerRoutes(e)
	return e
}

func TestRateLimiterMiddleware(t *testing.T) {
	e := setupEcho()

	// Test valid requests within limit
	req := httptest.NewRequest(http.MethodGet, "/endpoint1", nil)
	req.Header.Set("X-Real-IP", "192.168.0.1")
	rec := httptest.NewRecorder()

	// First request should pass
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Now we will send multiple request to check the rate limit exceed
	for i := 0; i < 10; i++ {
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)
	}

	// 11th request should be rate limited
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusTooManyRequests, rec.Code)
}

func TestDynamicRateLimitChange(t *testing.T) {
	e := setupEcho()

	// Set a new rate limit for a specific IP with limit value
	config.RateLimiterConfig.SetRateLimit("/endpoint1", "192.168.0.1", 5)

	req := httptest.NewRequest(http.MethodGet, "/endpoint1", nil)
	req.Header.Set("X-Real-IP", "192.168.0.1")
	rec := httptest.NewRecorder()

	// First 5 requests should pass
	for i := 0; i < 5; i++ {
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	// 6th request should be rate limited
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusTooManyRequests, rec.Code)
}

func TestDifferentIPAddresses(t *testing.T) {
	e := setupEcho()

	req1 := httptest.NewRequest(http.MethodGet, "/endpoint1", nil)
	req1.Header.Set("X-Real-IP", "192.168.0.1")
	rec1 := httptest.NewRecorder()

	req2 := httptest.NewRequest(http.MethodGet, "/endpoint1", nil)
	req2.Header.Set("X-Real-IP", "192.168.0.2")
	rec2 := httptest.NewRecorder()

	// First request from each IP should pass
	e.ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusOK, rec1.Code)

	e.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)

	// Exceed rate limit for first IP
	for i := 0; i < 10; i++ {
		e.ServeHTTP(rec1, req1)
		rec1 = httptest.NewRecorder() // reset recorder for the next iteration
	}

	// 11th request from first IP should be rate limited
	e.ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusTooManyRequests, rec1.Code)

	// Second IP should still be allowed
	e.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)
}
