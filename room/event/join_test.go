package event

import (
	"github.com/stretchr/testify/assert"
	mockproto "pixels-emulator/core/protocol/mock"
	"testing"
)

func TestNewRoomJoinEvent(t *testing.T) {
	con := &mockproto.MockConnection{}
	ev := NewRoomJoinEvent(con, 1, "pass", 0, make(map[string]string))

	assert.NotNil(t, ev.Conn, "Connection must be passed")
	assert.Equal(t, ev.Id, int32(1), "Room id must match")
	assert.Equal(t, ev.Password, "pass", "Password must match") // TODO: Hash this
}
