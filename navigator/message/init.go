package message

import "pixels-emulator/core/protocol"

// NavigatorInitCode is the unique identifier for the packet
const NavigatorInitCode = 2110

// NavigatorInitPacket represents a packet sent by client when
// the user opens the navigator.
type NavigatorInitPacket struct {
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *NavigatorInitPacket) Id() uint16 {
	return NavigatorInitCode
}

// Rate returns the rate limit for the Ping packet.
func (p *NavigatorInitPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *NavigatorInitPacket) Deadline() uint {
	return 10
}

// ComposeNavigatorInit composes a new instance of the packet.
func ComposeNavigatorInit(_ protocol.RawPacket) *NavigatorInitPacket {
	return &NavigatorInitPacket{}
}
