package guest

import "pixels-emulator/core/protocol"

// GetGuestRoomCode is the unique identifier for the packet
const GetGuestRoomCode = 2230

// GetRoomPacket represents a packet sent by client when a non-owned room is tried to be visited.
//
// This packet is called in many different ways at client, but, most of them, used when forwarded
// from navigator, guild rooms, user following, etc...
type GetRoomPacket struct {
	RoomId  int32 // RoomId provides the room which the connection will try to enter.
	Enter   bool  // Enter defines if the user is in enter processing (Already accessed and is being processed)
	Forward bool  // Forward defines if user is being forwarded to the room (First attempt to enter)
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *GetRoomPacket) Id() uint16 {
	return GetGuestRoomCode
}

// Rate returns the rate limit for the Ping packet.
func (p *GetRoomPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// ComposeGuestRoomPacket composes a new instance of the packet.
func ComposeGuestRoomPacket(pck protocol.RawPacket) (*GetRoomPacket, error) {

	rId, err := pck.ReadInt()
	enterRaw, err := pck.ReadInt()
	enter := enterRaw == 1
	forwardRaw, err := pck.ReadInt()
	forward := forwardRaw == 1

	return &GetRoomPacket{
		RoomId:  rId,
		Enter:   enter,
		Forward: forward,
	}, err

}
