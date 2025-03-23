package message

import "pixels-emulator/core/protocol"

// PingCode is the unique identifier for the Ping packet.
const PingCode = 3928

// PingPacket represents a Ping packet used to verify the status of a connection.
// This type embeds the base protocol.Packet interface.
type PingPacket struct {
	protocol.Packet // Embeds the base protocol.Packet interface.
}

// Id returns the unique identifier of the Packet type.
func (p *PingPacket) Id() uint16 {
	return PingCode
}

// Rate returns the rate limit for the Ping packet.
func (p *PingPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *PingPacket) Deadline() uint {
	return 10
}

// Serialize converts the Ping packet into a RawPacket that can be transmitted over the network.
func (p *PingPacket) Serialize() protocol.RawPacket {
	rawPack := protocol.NewPacket(PingCode)
	rawPack.AddBoolean(false)
	return rawPack
}

// ComposePing creates a new instance of a Ping packet.
func ComposePing() *PingPacket {
	return &PingPacket{}
}
