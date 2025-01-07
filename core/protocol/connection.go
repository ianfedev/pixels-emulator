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

	// GetIdentifier provides an unique identifier of this connection.
	GetIdentifier() string
}
