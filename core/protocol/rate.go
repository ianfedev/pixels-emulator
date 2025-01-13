package protocol

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

// RateLimiter defines the behavior of a rate limiter.
type RateLimiter interface {
	// GetLimiter retrieves the rate limiter for a given ID and rate parameters.
	GetLimiter(id uint16, rateLimit uint16, duration uint16) *rate.Limiter
}

// RateLimiterRegistry manages multiple rate limiters associated with unique keys.
type RateLimiterRegistry struct {
	mu       sync.Mutex               // Synchronizes access to the registry.
	limiters map[uint16]*rate.Limiter // Stores rate limiters for each ID.
}

// NewRateLimiter creates a new RateLimiterRegistry instance.
func NewRateLimiter() RateLimiter {
	return &RateLimiterRegistry{
		limiters: make(map[uint16]*rate.Limiter),
	}
}

// GetLimiter retrieves or creates a rate limiter for the specified ID.
func (r *RateLimiterRegistry) GetLimiter(id uint16, rateLimit uint16, duration uint16) *rate.Limiter {
	r.mu.Lock()
	defer r.mu.Unlock()

	if limiter, exists := r.limiters[id]; exists {
		return limiter
	}

	t := rate.Every(time.Duration(rateLimit/duration) * time.Second)
	limiter := rate.NewLimiter(t, int(duration))
	r.limiters[id] = limiter
	return limiter
}
