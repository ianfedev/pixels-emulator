package misc

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

var visSettingsPck = &RoomVisualizationSettingsPacket{
	HideWall:  true,
	WallSize:  1,
	FloorSize: 2,
}

// parseVisSettings reconstructs the packet for testing.
func parseVisSettings(pck protocol.RawPacket) (*RoomVisualizationSettingsPacket, error) {
	hideWall, err := pck.ReadBoolean()
	if err != nil {
		return nil, err
	}
	wallSize, err := pck.ReadInt()
	if err != nil {
		return nil, err
	}
	floorSize, err := pck.ReadInt()
	if err != nil {
		return nil, err
	}
	return &RoomVisualizationSettingsPacket{
		HideWall:  hideWall,
		WallSize:  wallSize,
		FloorSize: floorSize,
	}, nil
}

// TestRoomVisualizationSettingsPacket_Serialize checks if serialization works correctly.
func TestRoomVisualizationSettingsPacket_Serialize(t *testing.T) {
	raw := visSettingsPck.Serialize()
	bytes := raw.ToBytes()
	pck, err := protocol.FromBytes(bytes)
	assert.NoError(t, err)
	parsePck, err := parseVisSettings(*pck)
	assert.NoError(t, err)
	assert.Equal(t, visSettingsPck.HideWall, parsePck.HideWall)
	assert.Equal(t, visSettingsPck.WallSize, parsePck.WallSize)
	assert.Equal(t, visSettingsPck.FloorSize, parsePck.FloorSize)
}

// TestRoomVisualizationSettingsPacket checks static values.
func TestRoomVisualizationSettingsPacket(t *testing.T) {
	assert.Equal(t, visSettingsPck.Id(), uint16(RoomVisualizationSettingsPacketCode))
	assert.Equal(t, visSettingsPck.Deadline(), uint(0))
	mn, mx := visSettingsPck.Rate()
	if mn != 0 || mx != 0 {
		t.Errorf("RoomVisualizationSettingsPacket.Rate() = (%d, %d), expected (0, 0)", mn, mx)
	}
}
