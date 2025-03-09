package encode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pixels-emulator/core/protocol"
)

// TestRoomDataEncode ensures RoomData is correctly encoded into a RawPacket.
func TestRoomDataEncode(t *testing.T) {
	room := RoomData{
		ID:             1023,
		Name:           "Coolest room",
		OwnerID:        5021,
		OwnerName:      "Juanito",
		IsPublic:       false,
		State:          1,
		UserCount:      5,
		UserMax:        20,
		Description:    "Welecome to our room",
		Score:          40,
		Category:       1,
		Tags:           []string{"fun", "party"},
		GuildID:        3001,
		GuildName:      "Warriors",
		GuildBadge:     "badge123",
		PromotionTitle: "Awesome Party",
		PromotionDesc:  "Wleocme to the best pixel party",
		PromotionTime:  120,
	}

	pck := protocol.NewPacket(4001)
	pck.AddShort(99) // Extra data before RoomData
	room.Encode(&pck)
	pck.AddInt(777) // Extra data after RoomData

	rawBytes := pck.ToBytes()

	// Convert bytes back to a RawPacket
	receivedPck, err := protocol.FromBytes(rawBytes)
	require.NoError(t, err, "Failed to parse RawPacket")

	// Read extra data before RoomData
	beforeRoomData, err := receivedPck.ReadShort()
	require.NoError(t, err, "Failed to read beforeRoomData")
	assert.Equal(t, int16(99), beforeRoomData, "Incorrect beforeRoomData value")

	// Validate encoded data
	assert.Equal(t, room.ID, mustReadInt(t, receivedPck), "ID mismatch")
	assert.Equal(t, room.Name, mustReadString(t, receivedPck), "Name mismatch")
	assert.Equal(t, room.OwnerID, mustReadInt(t, receivedPck), "OwnerID mismatch")
	assert.Equal(t, room.OwnerName, mustReadString(t, receivedPck), "OwnerName mismatch")
	assert.Equal(t, int16(room.State), mustReadShort(t, receivedPck), "State mismatch")
	assert.Equal(t, room.UserCount, mustReadShort(t, receivedPck), "UserCount mismatch")
	assert.Equal(t, room.UserMax, mustReadShort(t, receivedPck), "UserMax mismatch")
	assert.Equal(t, room.Description, mustReadString(t, receivedPck), "Description mismatch")
	assert.Equal(t, room.Score, mustReadInt(t, receivedPck), "Score mismatch")
	assert.Equal(t, room.Category, mustReadShort(t, receivedPck), "Category mismatch")

	// Validate tags
	tagCount := mustReadShort(t, receivedPck)
	assert.Equal(t, int16(len(room.Tags)), tagCount, "Tag count mismatch")
	for i := 0; i < int(tagCount); i++ {
		assert.Equal(t, room.Tags[i], mustReadString(t, receivedPck), "Tag mismatch")
	}

	// Validate flags
	assert.Equal(t, int16(room.CalculateFlags()), mustReadShort(t, receivedPck), "Flags mismatch")

	// Validate Guild data
	if room.GuildID > 0 {
		assert.Equal(t, room.GuildID, mustReadInt(t, receivedPck), "GuildID mismatch")
		assert.Equal(t, room.GuildName, mustReadString(t, receivedPck), "GuildName mismatch")
		assert.Equal(t, room.GuildBadge, mustReadString(t, receivedPck), "GuildBadge mismatch")
	}

	// Validate Promotion data
	if room.PromotionTitle != "" {
		assert.Equal(t, room.PromotionTitle, mustReadString(t, receivedPck), "PromotionTitle mismatch")
		assert.Equal(t, room.PromotionDesc, mustReadString(t, receivedPck), "PromotionDesc mismatch")
		assert.Equal(t, room.PromotionTime, mustReadInt(t, receivedPck), "PromotionTime mismatch")
	}

	// Read extra data after RoomData
	afterRoomData, err := receivedPck.ReadInt()
	require.NoError(t, err, "Failed to read afterRoomData")
	assert.Equal(t, int32(777), afterRoomData, "Incorrect afterRoomData value")
}

// TestRoomDataFlags ensures CalculateFlags correctly computes room flags.
func TestRoomDataFlags(t *testing.T) {
	tests := []struct {
		room     RoomData
		expected int8
	}{
		{
			room:     RoomData{IsPublic: true, GuildID: 0, PromotionTitle: ""},
			expected: 0,
		},
		{
			room:     RoomData{IsPublic: false, GuildID: 0, PromotionTitle: ""},
			expected: 8, // Not public
		},
		{
			room:     RoomData{IsPublic: true, GuildID: 1234, PromotionTitle: ""},
			expected: 2, // Guild only
		},
		{
			room:     RoomData{IsPublic: true, GuildID: 0, PromotionTitle: "Promo"},
			expected: 4, // Promotion only
		},
		{
			room:     RoomData{IsPublic: false, GuildID: 5678, PromotionTitle: "Promo"},
			expected: 8 | 2 | 4, // Not public + Guild + Promotion
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.room.CalculateFlags(), "Flags mismatch for test case %+v", test.room)
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
