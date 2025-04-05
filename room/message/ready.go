package message

import (
	"golang.org/x/crypto/openpgp/packet"
	"pixels-emulator/core/protocol"
)

// RoomReadyPacketCode is the unique identifier for the packet
const RoomReadyPacketCode = 2031

// RoomReadyPacket defines when the room is loaded and ready
// accept connections, sending the related layout and room identifiers
type RoomReadyPacket struct {
	Layout string // Layout defines the heightmap id to be queried further.
	Room   int32  // Room defines the room id which is ready.
	packet.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *RoomReadyPacket) Id() uint16 {
	return RoomReadyPacketCode
}

// Rate returns the rate limit for the Ping packet.
func (p *RoomReadyPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *RoomReadyPacket) Deadline() uint {
	return 0
}

// Serialize transforms the packet into protocol RawPacket.
func (p *RoomReadyPacket) Serialize() protocol.RawPacket {
	pck := protocol.NewPacket(RoomReadyPacketCode)
	pck.AddString(p.Layout)
	pck.AddInt(p.Room)
	return pck
}
