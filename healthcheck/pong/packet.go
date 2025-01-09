package pong

import "pixels-emulator/core/protocol"

// PacketCode represents the packet code for the Pong packet.
const PacketCode = 2596

// Packet is a struct that embeds the base protocol.Packet interface.
type Packet struct {
	protocol.Packet
}

// Id returns the packet code for the Pong packet.
func (p Packet) Id() uint16 {
	return PacketCode
}

// Rate returns the rate limit for the Pong packet.
func (p Packet) Rate() (uint16, uint16) {
	return 10, 5
}

// NewPacket creates a new instance of the Pong packet.
func NewPacket(_ protocol.RawPacket) *Packet {
	return &Packet{}
}
