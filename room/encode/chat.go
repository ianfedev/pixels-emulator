package encode

import "pixels-emulator/core/protocol"

// ChatMode defines the different chat modes
const (
	ChatModeFreeFlow   = iota // Free flow chat mode
	ChatModeLineByLine = 1    // Line by line chat mode
)

// ChatBubbleWidth defines the different chat bubble widths
const (
	ChatBubbleWidthWide   = iota // Wide chat bubble
	ChatBubbleWidthNormal = 1    // Normal chat bubble
	ChatBubbleWidthThin   = 2    // Thin chat bubble
)

// ChatScrollSpeed defines the different chat scroll speeds
const (
	ChatScrollSpeedFast   = iota // Fast scrolling
	ChatScrollSpeedNormal = 1    // Normal scrolling
	ChatScrollSpeedSlow   = 2    // Slow scrolling
)

// FloodFilter defines the different flood protection levels
const (
	FloodFilterStrict = iota // Strict flood protection
	FloodFilterNormal = 1    // Normal flood protection
	FloodFilterLoose  = 2    // Loose flood protection
)

// RoomChatSettings represents chat settings for a room
type RoomChatSettings struct {
	protocol.Encodable
	Mode       int32 // Mode defines the chat mode
	Weight     int32 // Weight defines the hat bubble weight
	Speed      int32 // Speed defines the chat scroll speed
	Distance   int32 // Distance defines the chat distance limit
	Protection int32 // Protection defines the flood protection level
}

// Encode adds the RoomChatSettings data to a packet
func (e *RoomChatSettings) Encode(pck *protocol.RawPacket) {
	pck.AddInt(e.Mode)
	pck.AddInt(e.Weight)
	pck.AddInt(e.Speed)
	pck.AddInt(e.Distance)
	pck.AddInt(e.Protection)
}

// Decode the RoomChatSettings from a packet.
func (e *RoomChatSettings) Decode(pck *protocol.RawPacket) error {
	mode, err := pck.ReadInt()
	e.Mode = mode
	weight, err := pck.ReadInt()
	e.Weight = weight
	speed, err := pck.ReadInt()
	e.Speed = speed
	distance, err := pck.ReadInt()
	e.Distance = distance
	protection, err := pck.ReadInt()
	e.Protection = protection
	return err
}
