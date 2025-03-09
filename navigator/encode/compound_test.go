package encode

import (
	"pixels-emulator/room"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pixels-emulator/core/protocol"
	roomEncode "pixels-emulator/room/encode"
)

// TestSearchResultCompoundEncode ensures SearchResultCompound is correctly encoded into a RawPacket.
func TestSearchResultCompoundEncode(t *testing.T) {
	rooms := []*roomEncode.RoomData{
		{
			ID:             1023,
			Name:           "My Awesome Room",
			OwnerID:        5021,
			OwnerName:      "Juanito",
			IsPublic:       false,
			State:          room.Locked,
			UserCount:      5,
			UserMax:        20,
			Description:    "Welcome to my room!",
			Score:          40,
			Category:       1,
			Tags:           []string{"fun", "party"},
			GuildID:        3001,
			GuildName:      "The Warriors",
			GuildBadge:     "badge123",
			PromotionTitle: "Amazing Party",
			PromotionDesc:  "Join the best party of the night",
			PromotionTime:  120,
		},
	}

	result := SearchResultCompound{
		Code:       "popular",
		Query:      "Popular Rooms",
		Collapsed:  true,
		Actionable: true,
		Thumbnails: true,
		Rooms:      rooms,
	}

	// Create a packet and encode the result
	pck := protocol.NewPacket(5001)
	result.Encode(&pck)
	rawBytes := pck.ToBytes()

	// Convert bytes back to a RawPacket
	receivedPck, err := protocol.FromBytes(rawBytes)
	require.NoError(t, err, "Failed to parse RawPacket")

	// Validate serialized data
	assert.Equal(t, result.Code, mustReadString(t, receivedPck), "Code mismatch")
	assert.Equal(t, result.Query, mustReadString(t, receivedPck), "Query mismatch")

	action := mustReadInt(t, receivedPck)
	assert.Equal(t, int32(1), action, "Actionable mismatch")

	assert.Equal(t, result.Collapsed, mustReadBool(t, receivedPck), "Collapsed mismatch")

	thumbnails := mustReadInt(t, receivedPck)
	assert.Equal(t, int32(1), thumbnails, "Thumbnails mismatch")

	// Validate room count
	roomCount := mustReadInt(t, receivedPck)
	assert.Equal(t, int32(len(result.Rooms)), roomCount, "Room count mismatch")

	// Validate RoomData encoding
	for _, expectedRoom := range result.Rooms {
		assert.Equal(t, expectedRoom.ID, mustReadInt(t, receivedPck), "Room ID mismatch")
		assert.Equal(t, expectedRoom.Name, mustReadString(t, receivedPck), "Room Name mismatch")
		assert.Equal(t, expectedRoom.OwnerID, mustReadInt(t, receivedPck), "Owner ID mismatch")
		assert.Equal(t, expectedRoom.OwnerName, mustReadString(t, receivedPck), "Owner Name mismatch")
		assert.Equal(t, int16(expectedRoom.State), mustReadShort(t, receivedPck), "State mismatch")
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

		// Validate flags
		assert.Equal(t, int16(expectedRoom.CalculateFlags()), mustReadShort(t, receivedPck), "Flags mismatch")

		// Validate Guild data
		if expectedRoom.GuildID > 0 {
			assert.Equal(t, expectedRoom.GuildID, mustReadInt(t, receivedPck), "GuildID mismatch")
			assert.Equal(t, expectedRoom.GuildName, mustReadString(t, receivedPck), "GuildName mismatch")
			assert.Equal(t, expectedRoom.GuildBadge, mustReadString(t, receivedPck), "GuildBadge mismatch")
		}

		// Validate Promotion data
		if expectedRoom.PromotionTitle != "" {
			assert.Equal(t, expectedRoom.PromotionTitle, mustReadString(t, receivedPck), "PromotionTitle mismatch")
			assert.Equal(t, expectedRoom.PromotionDesc, mustReadString(t, receivedPck), "PromotionDesc mismatch")
			assert.Equal(t, expectedRoom.PromotionTime, mustReadInt(t, receivedPck), "PromotionTime mismatch")
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
