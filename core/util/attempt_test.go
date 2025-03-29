package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestAttemptLimiter tests if limiter is working.
func TestAttemptLimiter(t *testing.T) {
	limiter := NewAttemptLimiter()
	issuer := "user1"
	target := "room123"

	// First attempt should be allowed
	assert.True(t, limiter.RegisterAttempt(issuer, target))
	// Second attempt should be allowed
	assert.True(t, limiter.RegisterAttempt(issuer, target))
	// Third attempt should be allowed
	assert.True(t, limiter.RegisterAttempt(issuer, target))
	// Fourth attempt should be blocked
	assert.False(t, limiter.RegisterAttempt(issuer, target))
	assert.True(t, limiter.IsFrozen(issuer, target))

	// Wait for unfreeze (reducing delay for test efficiency)
	go func() {
		time.Sleep(31 * time.Second) // Ensures it surpasses the unfreeze time
		assert.True(t, limiter.RegisterAttempt(issuer, target))
	}()
}
