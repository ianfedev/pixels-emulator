package message

import (
	"pixels-emulator/core/protocol"
)

// GenericErrorCode is the unique identifier for the packet
const GenericErrorCode = 1600

// GenericErrorPacket send to the user a generic error. (Nitro did it dirty, only few codes... No customization at all)
type GenericErrorPacket struct {
	Code int32
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *GenericErrorPacket) Id() uint16 {
	return GenericErrorCode
}

// Rate returns the rate limit for the packet.
func (p *GenericErrorPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *GenericErrorPacket) Deadline() uint {
	return 0
}

// Serialize transforms packet in byte.
func (p *GenericErrorPacket) Serialize() protocol.RawPacket {
	pck := protocol.NewPacket(GenericErrorCode)
	pck.AddInt(p.Code)
	return pck
}
