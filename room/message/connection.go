package message

import (
	"golang.org/x/crypto/openpgp/packet"
	"pixels-emulator/core/protocol"
)

// OpenRoomConnectionPacketCode is the unique identifier for the packet
const OpenRoomConnectionPacketCode = 758

type OpenRoomConnectionPacket struct {
	packet.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *OpenRoomConnectionPacket) Id() uint16 {
	return OpenRoomConnectionPacketCode
}

// Rate returns the rate limit for the Ping packet.
func (p *OpenRoomConnectionPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *OpenRoomConnectionPacket) Deadline() uint {
	return 0
}

// Serialize transforms the packet into protocol RawPacket.
func (p *OpenRoomConnectionPacket) Serialize() protocol.RawPacket {
	pck := protocol.NewPacket(OpenRoomConnectionPacketCode)
	return pck
}
