package ok

import "pixels-emulator/core/protocol"

// PacketCode is the unique identifier for the Auth OK message.
const PacketCode = 2491

// Packet represents a Ping packet used to verify the status of a connection.
// This type embeds the base protocol.Packet interface.
type Packet struct {
	protocol.Packet // Embeds the base protocol.Packet interface.
}

// Id returns the unique identifier of the Packet type.
func (p *Packet) Id() uint16 {
	return PacketCode
}

// Rate returns the rate limit for the Ping packet.
func (p *Packet) Rate() (uint16, uint16) {
	return 0, 0
}

// Serialize converts the Auth OK packet into a RawPacket that can be transmitted over the network.
func (p *Packet) Serialize() protocol.RawPacket {
	return protocol.NewPacket(PacketCode)
}

// NewAuthOkPacket creates a new instance of Auth OK packet.
func NewAuthOkPacket() *Packet {
	return &Packet{}
}
