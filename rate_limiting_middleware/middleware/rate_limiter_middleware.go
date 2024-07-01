package middleware

import (
	"fmt"
	"net/http"
	"rate_limiting_middleware/config"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

var rateLimiters = struct {
	sync.Mutex
	clients map[string]*RateLimiter
}{
	clients: make(map[string]*RateLimiter),
}

type RateLimiter struct {
	mu           sync.Mutex
	lastReset    time.Time
	requestCount int
	requestLimit int
}

func NewRateLimiter(limit int) *RateLimiter {
	return &RateLimiter{
		requestLimit: limit,
	}
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	// Reset the counter if a new window has started
	if now.Sub(r.lastReset) > time.Minute {
		r.lastReset = now
		r.requestCount = 0
	}

	if r.requestCount < r.requestLimit {
		r.requestCount++
		return true
	}

	return false
}

func RateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		endpoint := c.Path()
		ip := c.RealIP()

		logMessage := fmt.Sprintf("Request from IP: %s to endpoint: %s", ip, endpoint)
		fmt.Println(logMessage)

		limit, exists := config.RateLimiterConfig.GetRateLimit(endpoint, ip)
		if !exists {
			limit, _ = config.RateLimiterConfig.GetRateLimit(endpoint, "default")
		}

		if limit == 0 {
			return next(c)
		}

		rateLimiters.Lock()
		limiter, exists := rateLimiters.clients[ip+endpoint]
		if !exists {
			limiter = NewRateLimiter(limit)
			rateLimiters.clients[ip+endpoint] = limiter
		}
		rateLimiters.Unlock()

		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded for:", ip)
			return c.JSON(http.StatusTooManyRequests, echo.Map{"error": "rate limit exceeded"})
		}

		return next(c)
	}
}
