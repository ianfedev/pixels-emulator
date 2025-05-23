package event

import (
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
)

const RoomJoinEventName = "room.join"

// RoomJoinEvent represents the action.go when a connection
// queries to log the room.
type RoomJoinEvent struct {
	*event.CancellableEvent                     // Extends functionality for cancellation.
	Id                      int32               // Id represents the id to connect
	Conn                    protocol.Connection // Conn represents the connection which is joining the room.
	Password                string              // Password represents the hashed password which enters to the room.
	OverrideCheck           bool                // OverrideCheck overrides the common checks
}

// NewRoomJoinEvent creates a new instance.
func NewRoomJoinEvent(conn protocol.Connection, id int32, password string, owner uint16, metadata map[string]string) *RoomJoinEvent {
	ce := event.NewCancellable(owner, metadata)
	return &RoomJoinEvent{
		CancellableEvent: ce.(*event.CancellableEvent),
		Id:               id,
		Conn:             conn,
		Password:         password,
	}
}
