package event

import (
	em "pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
	"strconv"
)

const UserDisconnectEventName = "user.disconnect"

// DisconnectReason defines the possible reasons for a user disconnecting
type DisconnectReason int

const (
	// SECURITY indicates a security-related disconnection
	SECURITY DisconnectReason = iota
	// INTERNAL indicates an internal system issue
	INTERNAL
	// OPERATIONAL indicates an operational issue
	OPERATIONAL
)

// UserDisconnectEvent represents an event triggered when a user disconnects
type UserDisconnectEvent struct {
	*em.BaseEvent
	ID     int              // ID of the user
	Reason DisconnectReason // Reason for the disconnection
}

// NewEvent creates a new UserDisconnectEvent with the given connection and reason
func NewEvent(conn protocol.Connection, reason DisconnectReason) *UserDisconnectEvent {

	id := 0
	// Convert connection identifier to int
	convId, err := strconv.Atoi(conn.Identifier())
	if err == nil {
		id = convId
	}

	// Return new UserDisconnectEvent with the ID and reason
	return &UserDisconnectEvent{
		ID:     id,
		Reason: reason,
	}
}
