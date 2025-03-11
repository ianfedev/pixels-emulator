package message

import (
	"pixels-emulator/room"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pixels-emulator/core/protocol"
	"pixels-emulator/navigator/encode"
	roomEncode "pixels-emulator/room/encode"
)

// TestNavigatorSearchResultPacketEncode ensures NavigatorSearchResultPacket is correctly encoded into a RawPacket.
func TestNavigatorSearchResultPacketEncode(t *testing.T) {
	rooms := []*roomEncode.RoomData{
		{
			ID:          1023,
			Name:        "My Awesome Room",
			OwnerID:     5021,
			OwnerName:   "Juanito",
			IsPublic:    true,
			DoorMode:    room.Open,
			UserCount:   10,
			UserMax:     30,
			Description: "Welcome to my room!",
			Score:       50,
			Category:    1,
			Tags:        []string{"fun", "party"},
		},
	}

	results := []*encode.SearchResultCompound{
		{
			Code:       "popular",
			Query:      "Popular Rooms",
			Collapsed:  false,
			Actionable: true,
			Thumbnails: true,
			Rooms:      rooms,
		},
	}

	packet := ComposeNavigatorSearchResult("popular", "Popular Rooms", results)

	// Create a RawPacket and serialize the packet
	pck := protocol.NewPacket(NavigatorSearchResultCode)
	packet.Serialize(&pck)
	rawBytes := pck.ToBytes()

	// Convert bytes back to a RawPacket
	receivedPck, err := protocol.FromBytes(rawBytes)
	require.NoError(t, err, "Failed to parse RawPacket")

	// Validate serialized data
	assert.Equal(t, packet.SearchCode, mustReadString(t, receivedPck), "SearchCode mismatch")
	assert.Equal(t, packet.SearchQuery, mustReadString(t, receivedPck), "SearchQuery mismatch")

	// Validate the number of result blocks
	resultCount := mustReadInt(t, receivedPck)
	assert.Equal(t, int32(len(packet.Results)), resultCount, "Result count mismatch")

	// Validate SearchResultCompound encoding
	for _, expectedResult := range packet.Results {
		assert.Equal(t, expectedResult.Code, mustReadString(t, receivedPck), "Result Code mismatch")
		assert.Equal(t, expectedResult.Query, mustReadString(t, receivedPck), "Result Query mismatch")
		assert.Equal(t, expectedResult.Collapsed, mustReadBool(t, receivedPck), "Collapsed mismatch")

		actionable := mustReadInt(t, receivedPck)
		assert.Equal(t, int32(1), actionable, "Actionable mismatch")

		thumbnails := mustReadInt(t, receivedPck)
		assert.Equal(t, int32(1), thumbnails, "Thumbnails mismatch")

		// Validate the number of rooms
		roomCount := mustReadInt(t, receivedPck)
		assert.Equal(t, int32(len(expectedResult.Rooms)), roomCount, "Room count mismatch")

		// Validate RoomData encoding
		for _, expectedRoom := range expectedResult.Rooms {
			assert.Equal(t, expectedRoom.ID, mustReadInt(t, receivedPck), "Room ID mismatch")
			assert.Equal(t, expectedRoom.Name, mustReadString(t, receivedPck), "Room Name mismatch")
			assert.Equal(t, expectedRoom.OwnerID, mustReadInt(t, receivedPck), "Owner ID mismatch")
			assert.Equal(t, expectedRoom.OwnerName, mustReadString(t, receivedPck), "Owner Name mismatch")
			assert.Equal(t, int16(expectedRoom.DoorMode), mustReadShort(t, receivedPck), "State mismatch")
			assert.Equal(t, expectedRoom.UserCount, mustReadShort(t, receivedPck), "UserCount mismatch")
			assert.Equal(t, expectedRoom.UserMax, mustReadShort(t, receivedPck), "UserMax mismatch")
			assert.Equal(t, expectedRoom.Description, mustReadString(t, receivedPck), "Description mismatch")
			assert.Equal(t, expectedRoom.Score, mustReadInt(t, receivedPck), "Score mismatch")
			assert.Equal(t, expectedRoom.Category, mustReadShort(t, receivedPck), "Category mismatch")

			// Validate tags
			tagCount := mustReadShort(t, receivedPck)
			assert.Equal(t, int16(len(expectedRoom.Tags)), tagCount, "Tag count mismatch")
			for i := 0; i < int(tagCount); i++ {
				assert.Equal(t, expectedRoom.Tags[i], mustReadString(t, receivedPck), "Tag mismatch")
			}
		}
	}
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
