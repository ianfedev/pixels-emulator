package message

import "pixels-emulator/core/protocol"

// RoomEnterCode is the unique identifier for the packet
const RoomEnterCode = 2312

// RoomEnterPacket defines a connection opening request
// from client to access a room.
type RoomEnterPacket struct {
	RoomId   int32  // RoomId of the room to connect
	Password string // Password provided to access if pertinent.
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *RoomEnterPacket) Id() uint16 {
	return RoomEnterCode
}

// Rate returns the rate limit for the Ping packet.
func (p *RoomEnterPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *RoomEnterPacket) Deadline() uint {
	return 500
}

// ComposeRoomEnterPacket composes a new instance of the packet.
func ComposeRoomEnterPacket(pck protocol.RawPacket) (*RoomEnterPacket, error) {

	id, err := pck.ReadInt()
	pass, err := pck.ReadString()

	return &RoomEnterPacket{
		RoomId:   id,
		Password: pass,
	}, err

}
