package ping

import "pixels-emulator/core/protocol"

// PacketCode is the unique identifier for the Ping packet.
// This constant is used to register and identify the Ping packet type in the packet registry.
const PacketCode = 3928

// Packet represents a Ping packet used to verify the status of a connection.
// This type embeds the base protocol.Packet interface, ensuring compatibility with
// the standard packet structure of the Pixels Emulator application.
type Packet struct {
	protocol.Packet // Embeds the base protocol.Packet interface.
}

// Id returns the unique identifier of the Packet type.
//
// Returns:
//
//	uint16: The identifier associated with the Packet, defined by PacketCode.
func (p *Packet) Id() uint16 {
	return PacketCode
}

// Serialize converts the Ping packet into a RawPacket that can be transmitted over the network.
//
// Returns:
//
//	protocol.RawPacket: A raw representation of the Ping packet, ready for transmission.
//
// Behavior:
//   - Creates a new RawPacket with the PacketCode for Ping.
//   - Adds a default boolean value (false) to the packet.
//
// Example:
//
//	pingPacket := NewPingPacket()
//	rawPacket := pingPacket.Serialize()
//	fmt.Println("Serialized Ping packet:", rawPacket)
func (p *Packet) Serialize() protocol.RawPacket {
	rawPack := protocol.NewPacket(PacketCode)
	rawPack.AddBoolean(false)
	return rawPack
}

// NewPingPacket creates a new instance of a Ping packet.
//
// Returns:
//
//	*Packet: A pointer to the newly created Ping packet.
//
// Example:
//
//	pingPacket := NewPingPacket()
//	fmt.Println("Created new Ping packet with ID:", PacketCode)
func NewPingPacket() *Packet {
	return &Packet{}
}
