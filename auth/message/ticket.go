package message

import (
	"pixels-emulator/core/protocol"
)

// AuthTicketCode represents the packet code for the SSO Auth.
const AuthTicketCode = 2419

type AuthTicketPacket struct {
	protocol.Packet // Embeds the base protocol.Packet interface.

	// Ticket is the SSO auth ticket generated.
	Ticket string

	// Time containing tick of packet issuing.
	Time int
}

// Id returns the packet code for the Pong packet.
func (p *AuthTicketPacket) Id() uint16 {
	return AuthTicketCode
}

// Rate returns the rate limit for the Pong packet.
func (p *AuthTicketPacket) Rate() (uint16, uint16) {
	return 10, 5
}

// ComposeTicket creates a new instance of the ticket packet.
func ComposeTicket(pack protocol.RawPacket) (*AuthTicketPacket, error) {

	ticket, err := pack.ReadString()
	if err != nil {
		return nil, err
	}

	time, err := pack.ReadInt()
	if err != nil {
		return nil, err
	}

	return &AuthTicketPacket{
		Ticket: ticket,
		Time:   int(time),
	}, err

}
