package misc

import (
	"pixels-emulator/core/protocol"
)

// RoomVisualizationSettingsPacketCode is the unique identifier for the packet
const RoomVisualizationSettingsPacketCode = 3547

// RoomVisualizationSettingsPacket defines visualization settings for a room.
type RoomVisualizationSettingsPacket struct {
	HideWall  bool
	WallSize  int32
	FloorSize int32
}

// Id returns the unique identifier of the Packet type.
func (p *RoomVisualizationSettingsPacket) Id() uint16 {
	return RoomVisualizationSettingsPacketCode
}

// Rate returns the rate limit for the packet.
func (p *RoomVisualizationSettingsPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *RoomVisualizationSettingsPacket) Deadline() uint {
	return 0
}

// Serialize transforms the packet into protocol RawPacket.
func (p *RoomVisualizationSettingsPacket) Serialize() protocol.RawPacket {
	pck := protocol.NewPacket(RoomVisualizationSettingsPacketCode)
	pck.AddBoolean(p.HideWall)
	pck.AddInt(p.WallSize)
	pck.AddInt(p.FloorSize)
	return pck
}
