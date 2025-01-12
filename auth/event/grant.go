package event

import (
	em "pixels-emulator/core/event"
)

const AuthGrantEventName = "auth.grant"

// AuthGrantEvent represents an event containing a user ID and metadata.
type AuthGrantEvent struct {
	*em.CancellableEvent     // CancellableEvent extends functionality.
	userID               int // userID defines the associated with the event.
}

// NewEvent creates a new AuthGrantEvent with the provided userID, owner, and metadata.
func NewEvent(userID int, owner uint16, metadata map[string]string) *AuthGrantEvent {
	ce := em.NewCancellable(owner, metadata)
	return &AuthGrantEvent{
		CancellableEvent: ce.(*em.CancellableEvent),
		userID:           userID,
	}
}

// UserID returns the ID of the user associated with this event.
func (e *AuthGrantEvent) UserID() int {
	return e.userID
}
