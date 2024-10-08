package config

import "sync"

type ServerConfig struct {
	Port string
}

type RateLimitConfig struct {
	Limits map[string]map[string]int
	sync.RWMutex
}

var RateLimiterConfig = &RateLimitConfig{
	Limits: map[string]map[string]int{
		"/endpoint1": {
			"default": 10,
		},
		"/endpoint2": {
			"default": 20,
		},
	},
}

func (r *RateLimitConfig) SetRateLimit(endpoint string, ip string, limit int) {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.Limits[endpoint]; !exists {
		r.Limits[endpoint] = make(map[string]int)
	}
	r.Limits[endpoint][ip] = limit
}

func (r *RateLimitConfig) GetRateLimit(endpoint string, ip string) (int, bool) {
	r.RLock()
	defer r.RUnlock()

	if limits, exists := r.Limits[endpoint]; exists {
		if limit, exists := limits[ip]; exists {
			return limit, true
		}
		if limit, exists := limits["default"]; exists {
			return limit, true
		}
	}
	return 0, false
}

func LoadConfig() ServerConfig {
	return ServerConfig{
		Port: ":8080",
	}
}
