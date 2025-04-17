package event

import (
	"pixels-emulator/core/event"
)

const RoomOpenEventName = "room.open"

// RoomOpenEvent is triggered when a room is about to be opened.
type RoomOpenEvent struct {
	*event.CancellableEvent      // CancellableEvent provides the ability to cancel the operation.
	RoomId                  uint // RoomId holds the identifier of the room being opened.
}

// NewRoomOpenEvent creates a new RoomOpenEvent.
//
// Parameters:
//   - id: the ID of the room that is being opened.
//   - owner: the ID of the event owner, used for tracing.
//   - metadata: additional metadata associated with the event.
//
// Returns:
//   - A pointer to the initialized RoomOpenEvent.
func NewRoomOpenEvent(id uint, owner uint16, metadata map[string]string) *RoomOpenEvent {
	ce := event.NewCancellable(owner, metadata)
	return &RoomOpenEvent{
		CancellableEvent: ce.(*event.CancellableEvent),
		RoomId:           id,
	}
}
