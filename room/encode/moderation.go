package encode

import "pixels-emulator/core/protocol"

// Level defines as enum the moderation type
type Level int32

const (
	// None defines no moderation permissions for user
	None = iota
	// Rights defines shared permissions by the owner.
	Rights = 1
	// Administrator defines full permissions on the room.
	Administrator = 2
)

// ModerationRights define an encodable wrapper which can be
// provided as part of room querying.
type ModerationRights struct {
	protocol.Encodable
	Mute Level // Mute check permission level for room muting.
	Kick Level // Kick check permission level for room kicking.
	Ban  Level // Ban check permission level for room banning.
}

// Encode adds current data to a packet.
func (e *ModerationRights) Encode(pck *protocol.RawPacket) {
	pck.AddInt(int32(e.Mute))
	pck.AddInt(int32(e.Kick))
	pck.AddInt(int32(e.Ban))
}

func (e *ModerationRights) Decode(pck *protocol.RawPacket) error {
	mute, err := pck.ReadInt()
	e.Mute = Level(mute)
	kick, err := pck.ReadInt()
	e.Kick = Level(kick)
	ban, err := pck.ReadInt()
	e.Ban = Level(ban)
	return err
}
