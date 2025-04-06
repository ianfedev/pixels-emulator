package message

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
)

// HeightMapRequestPacketCode is the unique identifier for the packet
const HeightMapRequestPacketCode = 2753

type HeightMapRequestPacket struct {
	Width   int32   // Width of the heightmap.
	Total   int32   // Total tiles of the heightmap.
	Heights []int16 // Heights defines the map height for every tile (Must be parsed Y,X instead of X,Y)
}

// Id returns the unique identifier of the Packet type.
func (p *HeightMapRequestPacket) Id() uint16 {
	return HeightMapRequestPacketCode
}

// Rate returns the rate limit for the Ping packet.
func (p *HeightMapRequestPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *HeightMapRequestPacket) Deadline() uint {
	return 0
}

// Serialize transforms the packet into protocol RawPacket.
func (p *HeightMapRequestPacket) Serialize() protocol.RawPacket {

	if int(p.Total) != len(p.Heights) {
		zap.L().Warn("Mismatch on height map request packet")
	}

	pck := protocol.NewPacket(HeightMapRequestPacketCode)
	pck.AddInt(p.Width)
	pck.AddInt(p.Total)

	for _, h := range p.Heights {
		pck.AddShort(h)
	}

	return pck
}
