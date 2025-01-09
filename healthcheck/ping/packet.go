package ping

import "pixels-emulator/core/protocol"

// PacketCode is the unique identifier for the Ping packet.
const PacketCode = 3928

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

// Serialize converts the Ping packet into a RawPacket that can be transmitted over the network.
func (p *Packet) Serialize() protocol.RawPacket {
	rawPack := protocol.NewPacket(PacketCode)
	rawPack.AddBoolean(false)
	return rawPack
}

// NewPingPacket creates a new instance of a Ping packet.
func NewPingPacket() *Packet {
	return &Packet{}
}
