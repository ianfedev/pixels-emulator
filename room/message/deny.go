package message

import (
	"pixels-emulator/core/protocol"
)

// DenyRoomConnectionCode is the unique identifier for the packet
const DenyRoomConnectionCode = 899

// Type defines the connection message type to indicate Nitro which message to render.
type Type int

const (
	Default Type = iota
	Full         // Full defines if room is in max capacity.
	Closed       // Closed defines if room can not be opened.
	Queue        // Queue defines if the error has relation with queue.
	Banned       // Banned defines if user has a prohibition on the entry.
)

// DenyRoomConnectionPacket send to the user a connection rejecting message.
type DenyRoomConnectionPacket struct {
	Type        Type   // Type defines the default type of the connection rejection nature.
	QueryHolder string // QueryHolder sends a title placeholder interpreted by nitro when a custom message is given for queue errors.
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *DenyRoomConnectionPacket) Id() uint16 {
	return DenyRoomConnectionCode
}

// Rate returns the rate limit for the Ping packet.
func (p *DenyRoomConnectionPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *DenyRoomConnectionPacket) Deadline() uint {
	return 0
}

// Serialize transforms packet in byte.
func (p *DenyRoomConnectionPacket) Serialize() protocol.RawPacket {
	pck := protocol.NewPacket(CloseRoomConnectionCode)
	pck.AddInt(int32(p.Type))
	pck.AddString(p.QueryHolder)
	return pck
}
