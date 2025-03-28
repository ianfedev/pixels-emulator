package message

import (
	"pixels-emulator/core/protocol"
)

// CloseRoomConnectionCode is the unique identifier for the packet
const CloseRoomConnectionCode = 122

// CloseRoomConnectionPacket closes the current room connection of the user.
type CloseRoomConnectionPacket struct {
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *CloseRoomConnectionPacket) Id() uint16 {
	return CloseRoomConnectionCode
}

// Rate returns the rate limit for the Ping packet.
func (p *CloseRoomConnectionPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *CloseRoomConnectionPacket) Deadline() uint {
	return 0
}

func (p *CloseRoomConnectionPacket) Serialize() protocol.RawPacket {
	pck := protocol.NewPacket(CloseRoomConnectionCode)
	return pck
}
