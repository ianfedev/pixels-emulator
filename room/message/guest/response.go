package guest

import (
	"pixels-emulator/core/protocol"
	"pixels-emulator/room/encode"
)

// ResponseGuestRoomCode is the unique identifier for the packet
const ResponseGuestRoomCode = 687

// ResponseRoomPacket defines a packet structure where a no owned
// room is queried from a connection.
type ResponseRoomPacket struct {
	Enter         bool                     // Enter defines if the information obtaining involves user entering to it our just an update.
	Forward       bool                     // Forward defines if user is forwarding.
	Room          *encode.RoomData         // Room defines the room data part of the global room response
	StaffPick     bool                     // StaffPick defines is room is selected as featured
	GuildMember   bool                     // GuildMember if the connection querying the room is part of the guild associated to the server.
	GlobalMute    bool                     // GlobalMute if room is being global muted
	CanGlobalMute bool                     // CanGlobalMute if user has rights to create a global mute on the room
	Moderation    *encode.ModerationRights // Moderation contains deserialized rights of moderation.
	Settings      *encode.RoomChatSettings // Settings contains deserialized room chat settings.
	protocol.Packet
}

// Id returns the unique identifier of the Packet type.
func (p *ResponseRoomPacket) Id() uint16 {
	return ResponseGuestRoomCode
}

// Rate returns the rate limit for the Ping packet.
func (p *ResponseRoomPacket) Rate() (uint16, uint16) {
	return 0, 0
}

// Deadline provides the maximum time a packet can be processed in milliseconds.
func (p *ResponseRoomPacket) Deadline() uint {
	return 500
}

func (p *ResponseRoomPacket) Serialize() protocol.RawPacket {

	pck := protocol.NewPacket(ResponseGuestRoomCode)
	pck.AddBoolean(p.Enter)
	p.Room.Encode(&pck)
	pck.AddBoolean(p.Forward)
	pck.AddBoolean(p.StaffPick)
	pck.AddBoolean(p.GuildMember)
	pck.AddBoolean(p.GlobalMute)
	p.Moderation.Encode(&pck)
	pck.AddBoolean(p.CanGlobalMute)
	p.Settings.Encode(&pck)

	return pck

}
