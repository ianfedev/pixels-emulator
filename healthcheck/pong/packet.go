package pong

import "pixels-emulator/core/protocol"

// PacketCode is the unique identifier for the Pong packet.
// This constant is used to register and identify the Pong packet type in the packet registry.
const PacketCode = 2596

// Packet represents a Pong packet used in the ping-pong communication process.
// This type embeds the base protocol.Packet interface, indicating that it adheres
// to the standard packet structure of the Pixels Emulator application.
type Packet struct {
	protocol.Packet // Embeds the base protocol.Packet interface.
}

// Id returns the unique identifier of the Packet type.
//
// Returns:
//
//	uint16: The identifier associated with the Packet, defined by PacketCode.
func (p Packet) Id() uint16 {
	return PacketCode
}

// NewPacket creates a new instance of the Pong packet.
//
// Parameters:
//
//	_: The raw packet data (protocol.RawPacket) used to create the Pong packet.
//	   This parameter is unused as the Pong packet does not require additional data.
//
// Returns:
//
//	*Packet: A pointer to the newly created Pong packet.
//
// Example:
//
//	rawPacket := protocol.RawPacket{}
//	pongPacket := NewPacket(rawPacket)
//	fmt.Println("Created Pong packet with ID:", PacketCode)
func NewPacket(_ protocol.RawPacket) *Packet {
	return &Packet{}
}
