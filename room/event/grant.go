package event

import (
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
	"pixels-emulator/room"
)

// RoomAccessGrantEventName is the identifiable name of the event for handler registry.
const RoomAccessGrantEventName = "room_access"

// RoomAccessGrantEvent must be fired when a user is validated internally or externally
// to be allowed to join a room lifecycle.
type RoomAccessGrantEvent struct {
	*event.BaseEvent
	Conn         protocol.Connection // Conn represents the connection which is joining the room.
	Room         uint                // Room where the access is granted.
	Relationship room.Relationship   // Relationship defines the permission case the access will handle.
}

// NewRoomAccessGrantEvent creates a new instance.
func NewRoomAccessGrantEvent(conn protocol.Connection, id uint, rel room.Relationship, owner uint16, metadata map[string]string) *RoomAccessGrantEvent {
	e := event.New(owner, metadata)
	return &RoomAccessGrantEvent{
		BaseEvent:    e.(*event.BaseEvent),
		Conn:         conn,
		Room:         id,
		Relationship: rel,
	}
}
