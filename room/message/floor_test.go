package message

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

var fPck = &FloorHeightMapRequestPacket{
	WallHeight: 3,
	Scale:      false,
	Layout:     "xxxx\r\nx22x\r\nx00x\r\nxxxx",
}

// parseFloor reads from raw packet to transform into floor heightmap request packet.
func parseFloor(pck protocol.RawPacket) (*FloorHeightMapRequestPacket, error) {

	s, err := pck.ReadBoolean()
	h, err := pck.ReadInt()
	l, err := pck.ReadString()

	return &FloorHeightMapRequestPacket{
		Scale:      s,
		WallHeight: h,
		Layout:     l,
	}, err

}

// TestFloorHeightMapRequestPacket_Serialize checks if serialization is made correctly.
func TestFloorHeightMapRequestPacket_Serialize(t *testing.T) {
	raw := fPck.Serialize()
	bytes := raw.ToBytes()
	pck, err := protocol.FromBytes(bytes)
	assert.NoError(t, err)
	floorPck, err := parseFloor(*pck)
	assert.NoError(t, err)
	assert.Equal(t, floorPck.WallHeight, fPck.WallHeight)
	assert.Equal(t, floorPck.Scale, fPck.Scale)
	assert.Equal(t, floorPck.Layout, fPck.Layout)
}

// TestFloorHeightMapRequestPacket check packet integrity.
func TestFloorHeightMapRequestPacket(t *testing.T) {
	assert.Equal(t, fPck.Id(), uint16(FloorHeightMapPacketCode))
	assert.Equal(t, fPck.Deadline(), uint(0))
	mn, mx := fPck.Rate()
	if mn != 0 || mx != 0 {
		t.Errorf("Packet.Rate() = (%d, %d), expected (0, 0)", mn, mx)
	}
}
