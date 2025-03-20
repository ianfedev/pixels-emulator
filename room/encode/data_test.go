package encode_test

import (
	"pixels-emulator/room/encode"
	"testing"

	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
)

func TestGenerateBitmask(t *testing.T) {
	r := encode.RoomData{
		Thumbnail:         "example.png",
		GuildID:           1,
		PromotionTitle:    "Promo Title",
		IsPublic:          true,
		AllowPets:         true,
		FeaturedPromotion: true,
	}

	bitmask := r.GenerateBitmask()

	assert.True(t, encode.Has(bitmask, encode.Thumbnail), "Expected bitmask to include Thumbnail")
	assert.True(t, encode.Has(bitmask, encode.Guild), "Expected bitmask to include Guild")
	assert.True(t, encode.Has(bitmask, encode.Promotion), "Expected bitmask to include Promotion")
	assert.True(t, encode.Has(bitmask, encode.Owner), "Expected bitmask to include Owner")
	assert.True(t, encode.Has(bitmask, encode.Pets), "Expected bitmask to include Pets")
	assert.True(t, encode.Has(bitmask, encode.FeaturedPromotion), "Expected bitmask to include FeaturedPromotion")
}

func TestEncodeDecode(t *testing.T) {
	r := encode.RoomData{
		ID:                1,
		Name:              "Test Room",
		OwnerID:           100,
		OwnerName:         "Owner",
		IsPublic:          false,
		DoorMode:          encode.Door(1),
		UserCount:         10,
		UserMax:           50,
		Description:       "A test room",
		Score:             200,
		Category:          3,
		Tags:              []string{"fun", "game"},
		GuildID:           2,
		GuildName:         "Guild Name",
		GuildBadge:        "Badge123",
		PromotionTitle:    "Special Promo",
		PromotionDesc:     "Limited time event",
		PromotionTime:     120,
		Thumbnail:         "thumbnail.png",
		AllowPets:         true,
		FeaturedPromotion: true,
	}

	pck := protocol.NewPacket(100)
	r.Encode(&pck)

	decodedRoom := encode.RoomData{}
	pck.ResetOffset()
	err := decodedRoom.Decode(&pck)
	assert.NoError(t, err, "Decode should not return an error")
	assert.Equal(t, r, decodedRoom, "Decoded RoomData should match original")
}
