package message

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

var fPck = &FloorHeightMapRequestPacket{
	WallHeight:  3,
	RelativeMap: "xxxx\nx22x\nx00x\nxxxx",
}

// parseFloor reads from raw packet to transform into floor heightmap request packet.
func parseFloor(pck protocol.RawPacket) (*FloorHeightMapRequestPacket, error) {

	h, err := pck.ReadInt()
	rm, err := pck.ReadString()

	return &FloorHeightMapRequestPacket{
		WallHeight:  h,
		RelativeMap: rm,
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
	assert.Equal(t, floorPck.RelativeMap, fPck.RelativeMap)
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
