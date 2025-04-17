package event

import (
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
)

const RoomCloseConnectionEventName = "room.connection.close"

// RoomCloseConnectionEvent is triggered when a specific connection
// is about to be removed from a room.
type RoomCloseConnectionEvent struct {
	*event.BaseEvent                     // Allows the closure to be cancelled.
	Connection       protocol.Connection // Connection to be removed from the room.
}

// NewRoomCloseConnectionEvent creates a new RoomCloseConnectionEvent.
//
// Parameters:
//   - conn: the connection that is being removed from the room.
//   - owner: the ID of the event owner, used for tracing.
//   - metadata: optional event metadata.
//
// Returns:
//   - A pointer to the initialized RoomCloseConnectionEvent.
func NewRoomCloseConnectionEvent(conn protocol.Connection, owner uint16, metadata map[string]string) *RoomCloseConnectionEvent {
	ce := event.New(owner, metadata)
	return &RoomCloseConnectionEvent{
		BaseEvent:  ce.(*event.BaseEvent),
		Connection: conn,
	}
}
