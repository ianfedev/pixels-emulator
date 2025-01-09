package hello

import "pixels-emulator/core/protocol"

// PacketCode is the unique identifier for the Packet.
// This constant is used to register and identify the Packet type in the packet registry.
const PacketCode = 4000

// Packet represents a packet sent by a client to initiate a connection.
// It includes version information about the client.
type Packet struct {
	protocol.Packet        // Embeds the base protocol.Packet interface.
	Version         string // Version of the client software.
}

// Id returns the unique identifier of the Packet type.
//
// Returns:
//
//	uint16: The identifier associated with the Packet, defined by PacketCode.
func (p Packet) Id() uint16 {
	return PacketCode
}

// NewPacket creates a new instance of Packet from the provided raw packet data.
//
// Parameters:
//
//	packet: The raw packet containing the data to initialize a Packet.
//
// Returns:
//
//	*Packet: A pointer to the newly created Packet containing parsed data.
//	error: An error if the raw packet cannot be parsed correctly, such as if the version string
//	       cannot be read.
//
// Example:
//
//	rawPacket := protocol.RawPacket{}
//	clientHello, err := NewPacket(rawPacket)
//	if err != nil {
//	    fmt.Println("Error creating Packet:", err)
//	} else {
//	    fmt.Println("Client version:", clientHello.Version)
//	}
func NewPacket(packet protocol.RawPacket) (*Packet, error) {
	// Read the version string from the raw packet
	ver, err := packet.ReadString()
	if err != nil {
		return nil, err
	}

	// Return a new Packet with the parsed version
	return &Packet{
		Version: ver,
	}, nil
}
