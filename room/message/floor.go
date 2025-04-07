package message

import (
	"pixels-emulator/core/protocol"
)

// FloorHeightMapPacketCode is the unique identifier for the packet
const FloorHeightMapPacketCode = 1301

type FloorHeightMapRequestPacket struct {
	WallHeight  int32  // WallHeight provides the configured wall height for a room.
	RelativeMap string // RelativeMap provides the heightmap for a wall.
}

// Id returns the unique identifier of the Packet type.
func (p *FloorHeightMapRequestPacket) Id() uint16 {
	return FloorHeightMapPacketCode
}

// Rate returns the rate limit for the Ping packet.
func (p *FloorHeightMapRequestPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *FloorHeightMapRequestPacket) Deadline() uint {
	return 0
}

// Serialize transforms the packet into protocol RawPacket.
func (p *FloorHeightMapRequestPacket) Serialize() protocol.RawPacket {
	pck := protocol.NewPacket(FloorHeightMapPacketCode)
	pck.AddInt(p.WallHeight)
	pck.AddString(p.RelativeMap)
	return pck
}
