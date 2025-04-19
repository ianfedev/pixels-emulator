package unit

import (
	"pixels-emulator/core/protocol"
	"pixels-emulator/room/encode"
)

// UpdateStatusCode is the unique identifier for the packet
const UpdateStatusCode = 1640

// UpdateStatusPacket defines the update for multiple room units.
type UpdateStatusPacket struct {
	Units []encode.UnitMessage
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *UpdateStatusPacket) Id() uint16 {
	return UpdateStatusCode
}

// Rate returns the rate limit for the packet.
func (p *UpdateStatusPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *UpdateStatusPacket) Deadline() uint {
	return 0
}

// Serialize transforms packet in byte.
func (p *UpdateStatusPacket) Serialize() protocol.RawPacket {
	pck := protocol.NewPacket(UpdateStatusCode)
	for _, unit := range p.Units {
		unit.Encode(&pck)
	}
	return pck
}
