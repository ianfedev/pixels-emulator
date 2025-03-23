package message

import "pixels-emulator/core/protocol"

// NavigatorSearchCode is the unique identifier for the packet
const NavigatorSearchCode = 249

// NavigatorSearchPacket represents a packet sent by navigator to
// query requested rooms.
type NavigatorSearchPacket struct {
	protocol.Packet

	View string // View represents the context of the view where the query was made.

	Query string // Query represents the tags containing the query conditioning.

}

// Id returns the unique identifier of the Packet type.
func (p *NavigatorSearchPacket) Id() uint16 {
	return NavigatorSearchCode
}

// Rate returns the rate limit for the packet.
func (p *NavigatorSearchPacket) Rate() (uint16, uint16) {
	return 5, 5
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *NavigatorSearchPacket) Deadline() uint {
	return 10
}

// ComposeNavigatorSearch composes a new instance of the packet.
func ComposeNavigatorSearch(pck protocol.RawPacket) (*NavigatorSearchPacket, error) {

	view, err := pck.ReadString()
	if err != nil {
		return nil, err
	}

	query, err := pck.ReadString()
	if err != nil {
		return nil, err
	}

	return &NavigatorSearchPacket{View: view, Query: query}, nil

}
