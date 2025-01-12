package sso

import (
	"pixels-emulator/core/protocol"
)

// PacketCode represents the packet code for the SSO Auth.
const PacketCode = 2419

type Packet struct {
	protocol.Packet // Embeds the base protocol.Packet interface.

	// Ticket is the SSO auth ticket generated.
	Ticket string

	// Time containing tick of packet issuing.
	Time int
}

// Id returns the packet code for the Pong packet.
func (p Packet) Id() uint16 {
	return PacketCode
}

// Rate returns the rate limit for the Pong packet.
func (p Packet) Rate() (uint16, uint16) {
	return 10, 5
}

// NewPacket creates a new instance of the Pong packet.
func NewPacket(pack protocol.RawPacket) (*Packet, error) {

	ticket, err := pack.ReadString()
	if err != nil {
		return nil, err
	}

	time, err := pack.ReadInt()
	if err != nil {
		return nil, err
	}

	return &Packet{
		Ticket: ticket,
		Time:   int(time),
	}, err

}
