package encode

import (
	"pixels-emulator/room"
	"testing"

	"github.com/stretchr/testify/require"
	"pixels-emulator/core/protocol"
	roomEncode "pixels-emulator/room/encode"
)

// TestSearchResultCompoundEncode ensures SearchResultCompound is correctly encoded into a RawPacket with two rooms.
func TestSearchResultCompoundEncode(t *testing.T) {
	// Creating two RoomData instances
	rooms := []*roomEncode.RoomData{
		{
			ID:          1010,
			Name:        "Relaxing Lounge",
			OwnerID:     2010,
			OwnerName:   "Alice",
			IsPublic:    true,
			State:       room.Open,
			UserCount:   10,
			UserMax:     50,
			Description: "A place to chill and relax.",
			Score:       80,
			Category:    2,
			Tags:        []string{"chill", "lounge"},
		},
	}

	// Creating a SearchResultCompound instance
	result := SearchResultCompound{
		Code:       "featured",
		Query:      "Top Featured Rooms",
		Collapsed:  false,
		Actionable: true,
		Thumbnails: true,
		Rooms:      rooms,
	}

	// Create a packet and encode the result
	pck := protocol.NewPacket(5001)
	result.Encode(&pck)
	rawBytes := pck.ToBytes()

}

// Helper functions for reading packet data
func mustReadInt(t *testing.T, pck *protocol.RawPacket) int32 {
	val, err := pck.ReadInt()
	require.NoError(t, err, "Failed to read int")
	return val
}

func mustReadShort(t *testing.T, pck *protocol.RawPacket) int16 {
	val, err := pck.ReadShort()
	require.NoError(t, err, "Failed to read short")
	return val
}

func mustReadString(t *testing.T, pck *protocol.RawPacket) string {
	val, err := pck.ReadString()
	require.NoError(t, err, "Failed to read string")
	return val
}

func mustReadBool(t *testing.T, pck *protocol.RawPacket) bool {
	val, err := pck.ReadBoolean()
	require.NoError(t, err, "Failed to read boolean")
	return val
}
