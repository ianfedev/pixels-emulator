package message

import "pixels-emulator/core/protocol"

// PongCode represents the packet code for the Pong packet.
const PongCode = 2596

// PongPacket is a struct that embeds the base protocol.Packet interface.
type PongPacket struct {
	protocol.Packet
}

// Id returns the packet code for the Pong packet.
func (p *PongPacket) Id() uint16 {
	return PongCode
}

// Rate returns the rate limit for the Pong packet.
func (p *PongPacket) Rate() (uint16, uint16) {
	return 10, 5
}

// ComposePong composes a new instance of the Pong packet.
func ComposePong(_ protocol.RawPacket) *PongPacket {
	return &PongPacket{}
}
