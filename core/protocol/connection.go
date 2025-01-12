package protocol

import (
	"pixels-emulator/core/util"
)

// Connection acts a generic interface which describes
// a protocol-able connection.
//
// This is thought in order to future-proof any further
// implementation of upcoming protocols like what happened
// with Nitro client and old SWF-styled packets.
type Connection interface {
	util.Disposable

	// Identifier provides an unique identifier of this connection.
	Identifier() string

	// GrantIdentifier provides a new identifier for connection.
	GrantIdentifier(identifier string)

	// SendPacket pings an outgoing packet.
	SendPacket(packet Packet)

	// SendRaw pings a raw packet with custom restriction.
	SendRaw(packet RawPacket, period uint16, rate uint16)

	// RateRegistry limit rates outgoing packets.
	RateRegistry() *RateLimiterRegistry
}
