package message

import "pixels-emulator/core/protocol"

// GetGuestRoomCode is the unique identifier for the packet
const GetGuestRoomCode = 2230

// GetGuestRoomPacket represents a packet sent by client when
// a non-owned room is tried to be visited.
type GetGuestRoomPacket struct {
	RoomId  int32 // RoomId provides the room which the connection will try to enter.
	Enter   bool  // Enter defines if the information obtaining involves user entering to it our just an update.
	Forward bool  // Forward defines if user is forwarding.
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *GetGuestRoomPacket) Id() uint16 {
	return GetGuestRoomCode
}

// Rate returns the rate limit for the Ping packet.
func (p *GetGuestRoomPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// ComposeGuestRoomPacket composes a new instance of the packet.
func ComposeGuestRoomPacket(pck protocol.RawPacket) (*GetGuestRoomPacket, error) {

	rId, err := pck.ReadInt()
	enter, err := pck.ReadBoolean()
	forward, err := pck.ReadBoolean()

	return &GetGuestRoomPacket{
		RoomId:  rId,
		Enter:   enter,
		Forward: forward,
	}, err

}
