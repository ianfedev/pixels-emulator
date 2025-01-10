package protocol

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

// RateLimiterRegistry manages multiple instances of rate limiters associated with unique keys.
//
// This structure allows the creation and storage of rate limiters for different packet identifiers.
// Each packet can have its own rate limiting, which is used to control how many packets can be processed per second.
type RateLimiterRegistry struct {
	mu       sync.Mutex               // mu synchronizes access to the registry structure.
	limiters map[uint16]*rate.Limiter // limiters map that stores rate limiters associated with the packet Id.
}

// NewRateLimiterRegistry creates a new instance of RateLimiterRegistry.
//
// This function initializes a new RateLimiterRegistry, which is empty at the start.
// It returns a reference to the newly created RateLimiterRegistry.
func NewRateLimiterRegistry() *RateLimiterRegistry {
	return &RateLimiterRegistry{
		limiters: make(map[uint16]*rate.Limiter),
	}
}

// GetLimiter retrieves the rate limiter associated with a specific Id. If it doesn't exist, it creates one.
//
// The rate limiter is used to control the number of packets that can be processed per second for the given identifier.
// If the rate limiter does not already exist for the Id, the function will create it with the provided rate limit and return it.
func (r *RateLimiterRegistry) GetLimiter(id uint16, rateLimit uint16, duration uint16) *rate.Limiter {
	r.mu.Lock()
	defer r.mu.Unlock()

	// If the rate limiter for this Id doesn't exist, we create it
	if limiter, exists := r.limiters[id]; exists {
		return limiter
	}

	// Create a new rate limiter with the provided rate
	t := rate.Every(time.Duration(rateLimit) * time.Second)
	limiter := rate.NewLimiter(t, int(duration))
	r.limiters[id] = limiter
	return limiter
}
