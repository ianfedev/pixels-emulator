package encode

import (
	"pixels-emulator/core/protocol"
	roomEncode "pixels-emulator/room/encode"
)

// SearchResultCompound defines a single navigator group of rooms which
// can be encoded and sent to the Nitro client.
type SearchResultCompound struct {

	// Encodable asserts the result encoding ability.
	protocol.Encodable

	// Code is the same id used by the query and filter used in the granted room selection.
	Code string

	// Query defines additional tags or parameters used by the client to get the search result.
	Query string

	// Collapsed if desired result should be collapsed or expanded based on user preferences.
	Collapsed bool

	// Actionable if it should be any further action for the user (Handled by Nitro client. e.g: "More")
	Actionable bool

	// Thumbnails showing the room image or default listing.
	Thumbnails bool

	// Rooms to be provided in the result list.
	Rooms []*roomEncode.RoomData
}

func (r *SearchResultCompound) Encode(pck *protocol.RawPacket) {
	pck.AddString(r.Code)
	pck.AddString(r.Query)

	var action int32 = 0
	if r.Actionable {
		action = 1
	}
	pck.AddInt(action)

	pck.AddBoolean(r.Collapsed)

	var thumbnails int32 = 0
	if r.Thumbnails {
		thumbnails = 1
	}
	pck.AddInt(thumbnails)

	pck.AddInt(int32(len(r.Rooms)))
	for _, room := range r.Rooms {
		room.Encode(pck)
	}
}
