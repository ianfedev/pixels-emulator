package message

import "pixels-emulator/core/protocol"

// HelloCode is the unique identifier for the Hello packet.
const HelloCode = 4000

// HelloPacket represents a Hello packet that includes the client software version.
// This type embeds the base protocol.Packet interface.
type HelloPacket struct {
	protocol.Packet        // Embeds the base protocol.Packet interface.
	Version         string // Version of the client software.
}

// Id returns the unique identifier of the Packet type.
func (p *HelloPacket) Id() uint16 {
	return HelloCode
}

// Rate returns the rate limit for the Pong packet.
func (p *HelloPacket) Rate() (uint16, uint16) {
	return 300, 1
}

// ComposeHello creates a new instance of hello packet.
func ComposeHello(packet protocol.RawPacket) (*HelloPacket, error) {
	ver, err := packet.ReadString()
	if err != nil {
		return nil, err
	}

	return &HelloPacket{
		Version: ver,
	}, nil
}
