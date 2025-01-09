package hello

import "pixels-emulator/core/protocol"

// PacketCode is the unique identifier for the Hello packet.
const PacketCode = 4000

// Packet represents a Hello packet that includes the client software version.
// This type embeds the base protocol.Packet interface.
type Packet struct {
	protocol.Packet        // Embeds the base protocol.Packet interface.
	Version         string // Version of the client software.
}

// Id returns the unique identifier of the Packet type.
func (p Packet) Id() uint16 {
	return PacketCode
}

// Rate returns the rate limit for the Pong packet.
func (p Packet) Rate() (uint16, uint16) {
	return 300, 1
}

// NewPacket creates a new Hello packet from a RawPacket, extracting the client software version.
//
// Returns:
//
//	*Packet: A pointer to the newly created Hello packet with the extracted version.
//	error: Any error that occurs during the process of reading the version string.
func NewPacket(packet protocol.RawPacket) (*Packet, error) {
	ver, err := packet.ReadString()
	if err != nil {
		return nil, err
	}

	return &Packet{
		Version: ver,
	}, nil
}
