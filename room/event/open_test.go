package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewRoomOpenEvent verifies that a RoomOpenEvent is initialized properly.
func TestNewRoomOpenEvent(t *testing.T) {
	ev := NewRoomOpenEvent(123, 42, map[string]string{"source": "test"})

	assert.Equal(t, uint(123), ev.RoomId, "Room ID must match the one passed")
	assert.Equal(t, "test", ev.Metadata()["source"], "Metadata must be correctly passed")
	assert.Equal(t, uint16(42), ev.Owner(), "Owner ID must match")
}
