package message

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

var connPck = &OpenRoomConnectionPacket{}

// TestOpenRoomConnectionPacket_Serialize checks if serialization is made correctly.
func TestOpenRoomConnectionPacket_Serialize(t *testing.T) {
	raw := connPck.Serialize()
	bytes := raw.ToBytes()
	_, err := protocol.FromBytes(bytes)
	assert.NoError(t, err)
}

// TestOpenRoomConnectionPacket checks packet integrity.
func TestOpenRoomConnectionPacket(t *testing.T) {
	assert.Equal(t, connPck.Id(), uint16(OpenRoomConnectionPacketCode))
	assert.Equal(t, connPck.Deadline(), uint(0))
	mn, mx := connPck.Rate()
	if mn != 0 || mx != 0 {
		t.Errorf("RoomReadyPacketCode.Rate() = (%d, %d), expected (0, 0)", mn, mx)
	}
}
