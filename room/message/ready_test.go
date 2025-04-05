package message

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

var readyPck = &RoomReadyPacket{
	Room:   1,
	Layout: "xxxx\nx22x\nx00x\nxxxx",
}

// parseReady reads from raw packet to transform into denying packet.
func parseReady(pck protocol.RawPacket) (*RoomReadyPacket, error) {
	l, err := pck.ReadString()
	id, err := pck.ReadInt()
	return &RoomReadyPacket{
		Room:   id,
		Layout: l,
	}, err
}

// TestRoomReadyPacket_Serialize checks if serialization is made correctly.
func TestRoomReadyPacket_Serialize(t *testing.T) {
	raw := readyPck.Serialize()
	bytes := raw.ToBytes()
	pck, err := protocol.FromBytes(bytes)
	assert.NoError(t, err)
	parsePck, err := parseReady(*pck)
	assert.NoError(t, err)
	assert.Equal(t, parsePck.Room, readyPck.Room)
	assert.Equal(t, parsePck.Layout, readyPck.Layout)
}

// TestRoomReadyPacket check packet integrity.
func TestRoomReadyPacket(t *testing.T) {
	assert.Equal(t, readyPck.Id(), uint16(RoomReadyPacketCode))
	assert.Equal(t, readyPck.Deadline(), uint(0))
	mn, mx := readyPck.Rate()
	if mn != 0 || mx != 0 {
		t.Errorf("AuthOkPacket.Rate() = (%d, %d), expected (0, 0)", mn, mx)
	}
}
