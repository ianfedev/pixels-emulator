package event

import (
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
)

// RoomLoadRequestEventName is the identifiable name of the event for handler registry.
const RoomLoadRequestEventName = "room.load"

// RoomLoadRequestEvent must be fired when a user is validated internally or externally
// to be allowed to join a room lifecycle, and the room must be loaded.
type RoomLoadRequestEvent struct {
	*event.BaseEvent
	Conn protocol.Connection // Conn represents the connection which is joining the room.
	Room uint                // Room where the access is granted.
}

// NewRoomLoadRequestEvent creates a new instance.
func NewRoomLoadRequestEvent(conn protocol.Connection, id uint, owner uint16, metadata map[string]string) *RoomLoadRequestEvent {
	e := event.New(owner, metadata)
	return &RoomLoadRequestEvent{
		BaseEvent: e.(*event.BaseEvent),
		Conn:      conn,
		Room:      id,
	}
}
