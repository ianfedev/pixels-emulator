package message

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCloseRoomConnectionPacket_Serialize tests the correct serialization of the packet.
func TestCloseRoomConnectionPacket_Serialize(t *testing.T) {
	cPack := &CloseRoomConnectionPacket{}
	s := cPack.Serialize()
	assert.Equal(t, uint16(CloseRoomConnectionCode), s.GetHeader())
}

// TestCloseRoomConnectionPacket tests the correct attributes of the packet.
func TestCloseRoomConnectionPacket(t *testing.T) {
	cPack := &CloseRoomConnectionPacket{}
	assert.Equal(t, cPack.Id(), uint16(CloseRoomConnectionCode))
	assert.Equal(t, cPack.Deadline(), uint(0))
	mn, mx := cPack.Rate()
	if mn != 0 || mx != 0 {
		t.Errorf("AuthOkPacket.Rate() = (%d, %d), expected (0, 0)", mn, mx)
	}
}
