package util

import (
	"sync"
	"time"
)

// AttemptLimiter controls attempt limits and temporary bans.
type AttemptLimiter struct {
	mu       sync.Mutex      // Ensures thread safety
	attempts map[string]int  // Tracks attempts per issuer-target pair
	frozen   map[string]bool // Tracks frozen status per issuer-target pair
}

// NewAttemptLimiter initializes an AttemptLimiter.
func NewAttemptLimiter() *AttemptLimiter {
	return &AttemptLimiter{
		attempts: make(map[string]int),
		frozen:   make(map[string]bool),
	}
}

// RegisterAttempt records an attempt for a given issuer and target.
// Returns false if further attempts are temporarily blocked.
func (al *AttemptLimiter) RegisterAttempt(issuer, target string) bool {
	key := issuer + ":" + target
	al.mu.Lock()
	defer al.mu.Unlock()

	if al.frozen[key] {
		return false
	}

	al.attempts[key]++
	if al.attempts[key] >= 3 {
		al.frozen[key] = true
		go al.Unfreeze(key)
	}
	return true
}

// IsFrozen checks if attempts are currently frozen for a given issuer and target.
func (al *AttemptLimiter) IsFrozen(issuer, target string) bool {
	key := issuer + ":" + target
	al.mu.Lock()
	defer al.mu.Unlock()
	return al.frozen[key]
}

// Unfreeze resets the attempt count and unfreezes attempts after a delay.
func (al *AttemptLimiter) Unfreeze(key string) {
	time.Sleep(30 * time.Second)
	al.mu.Lock()
	defer al.mu.Unlock()
	al.frozen[key] = false
	al.attempts[key] = 0
}
