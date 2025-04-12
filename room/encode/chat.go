package encode

import "pixels-emulator/core/protocol"

type ChatMode int32

// ChatMode defines the different chat modes
const (
	ChatModeFreeFlow   = iota // Free flow chat mode
	ChatModeLineByLine = 1    // Line by line chat mode
)

type ChatWidth int32

// ChatBubbleWidth defines the different chat bubble widths
const (
	ChatBubbleWidthWide   = iota // Wide chat bubble
	ChatBubbleWidthNormal = 1    // Normal chat bubble
	ChatBubbleWidthThin   = 2    // Thin chat bubble
)

type ChatSpeed int32

// ChatScrollSpeed defines the different chat scroll speeds
const (
	ChatScrollSpeedFast   = iota // Fast scrolling
	ChatScrollSpeedNormal = 1    // Normal scrolling
	ChatScrollSpeedSlow   = 2    // Slow scrolling
)

type ChatFilter int32

// FloodFilter defines the different flood protection levels
const (
	FloodFilterStrict = iota // Strict flood protection
	FloodFilterNormal = 1    // Normal flood protection
	FloodFilterLoose  = 2    // Loose flood protection
)

// RoomChatSettings represents chat settings for a room
type RoomChatSettings struct {
	protocol.Encodable
	Mode       ChatMode   // Mode defines the chat mode
	Weight     ChatWidth  // Weight defines the hat bubble weight
	Speed      ChatSpeed  // Speed defines the chat scroll speed
	Distance   int32      // Distance defines the chat distance limit
	Protection ChatFilter // Protection defines the flood protection level
}

// Encode adds the RoomChatSettings data to a packet
func (e *RoomChatSettings) Encode(pck *protocol.RawPacket) {
	pck.AddInt(int32(e.Mode))
	pck.AddInt(int32(e.Weight))
	pck.AddInt(int32(e.Speed))
	pck.AddInt(e.Distance)
	pck.AddInt(int32(e.Protection))
}

// Decode the RoomChatSettings from a packet.
func (e *RoomChatSettings) Decode(pck *protocol.RawPacket) error {
	mode, err := pck.ReadInt()
	e.Mode = ChatMode(mode)
	weight, err := pck.ReadInt()
	e.Weight = ChatWidth(weight)
	speed, err := pck.ReadInt()
	e.Speed = ChatSpeed(speed)
	distance, err := pck.ReadInt()
	e.Distance = distance
	protection, err := pck.ReadInt()
	e.Protection = ChatFilter(protection)
	return err
}
