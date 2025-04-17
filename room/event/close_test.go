package event

import (
	mockproto "pixels-emulator/core/protocol/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewRoomCloseConnectionEvent ensures the event is created properly.
func TestNewRoomCloseConnectionEvent(t *testing.T) {
	mockConn := &mockproto.MockConnection{}
	ev := NewRoomCloseConnectionEvent(mockConn, 99, map[string]string{"reason": "timeout"})

	assert.Equal(t, mockConn, ev.Connection, "Connection should match the one passed")
	assert.Equal(t, "timeout", ev.Metadata()["reason"], "Metadata must be correctly set")
	assert.Equal(t, uint16(99), ev.Owner(), "Owner must match")
}
