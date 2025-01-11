package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewRateLimiterRegistry tests the creation of a new RateLimiterRegistry.
func TestNewRateLimiterRegistry(t *testing.T) {
	registry := NewRateLimiterRegistry()
	assert.NotNil(t, registry)
	assert.Len(t, registry.limiters, 0)
}

// TestGetLimiter tests the retrieval or creation of a rate limiter.
func TestGetLimiter(t *testing.T) {
	registry := NewRateLimiterRegistry()
	limiter := registry.GetLimiter(1, 10, 5)
	assert.NotNil(t, limiter)
	assert.Len(t, registry.limiters, 1)
	assert.Contains(t, registry.limiters, uint16(1))
}

// TestGetLimiter_CreatesNewLimiter tests that GetLimiter creates a new limiter if not exists.
func TestGetLimiter_CreatesNewLimiter(t *testing.T) {
	registry := NewRateLimiterRegistry()
	limiter1 := registry.GetLimiter(1, 10, 5)
	assert.NotNil(t, limiter1)
	assert.Len(t, registry.limiters, 1)
	limiter2 := registry.GetLimiter(1, 10, 5)
	assert.Equal(t, limiter1, limiter2)
}

// TestGetLimiter_DifferentId tests that GetLimiter creates different limiters for different IDs.
func TestGetLimiter_DifferentId(t *testing.T) {
	registry := NewRateLimiterRegistry()
	limiter1 := registry.GetLimiter(1, 10, 5)
	assert.NotNil(t, limiter1)
	limiter2 := registry.GetLimiter(2, 15, 10)
	assert.NotNil(t, limiter2)
	assert.NotEqual(t, limiter1, limiter2)
	assert.Len(t, registry.limiters, 2)
}

// TestRateLimiterBehaviour tests the behaviour of the rate limiter in terms of packet limiting.
func TestRateLimiterBehaviour(t *testing.T) {
	registry := NewRateLimiterRegistry()
	limiter := registry.GetLimiter(1, 5, 5)
	assert.NotNil(t, limiter)
	for i := 0; i < 5; i++ {
		assert.True(t, limiter.Allow())
	}
	assert.False(t, limiter.Allow())
}

// TestLimiterConcurrency tests concurrent access and creation of rate limiters.
func TestLimiterConcurrency(t *testing.T) {
	registry := NewRateLimiterRegistry()
	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go func(id uint16) {
			limiter := registry.GetLimiter(id, 10, 5)
			assert.NotNil(t, limiter)
			done <- true
		}(uint16(i))
	}
	for i := 0; i < 100; i++ {
		<-done
	}
	assert.Len(t, registry.limiters, 100)
}
