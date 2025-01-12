package grant

import (
	em "pixels-emulator/core/event"
)

const AuthEventName = "auth.grant"

// AuthEvent represents an event containing a user ID and metadata.
type AuthEvent struct {
	*em.CancellableEvent     // CancellableEvent extends functionality.
	userID               int // userID defines the associated with the event.
}

// NewEvent creates a new AuthEvent with the provided userID, owner, and metadata.
func NewEvent(userID int, owner uint16, metadata map[string]string) *AuthEvent {
	ce := em.NewCancellable(owner, metadata)
	return &AuthEvent{
		CancellableEvent: ce.(*em.CancellableEvent),
		userID:           userID,
	}
}

// UserID returns the ID of the user associated with this event.
func (e *AuthEvent) UserID() int {
	return e.userID
}
