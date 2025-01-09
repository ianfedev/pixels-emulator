package protocol

import "pixels-emulator/core/util"

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

	// SendPacket pings a packet to a duplex collection.
	SendPacket(packet Packet)
}
