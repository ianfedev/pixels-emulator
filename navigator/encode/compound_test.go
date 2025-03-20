package encode_test

import (
	"pixels-emulator/navigator/encode"
	"testing"

	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	roomEncode "pixels-emulator/room/encode"
)

func TestSearchResultCompoundEncodeDecode(t *testing.T) {
	r := encode.SearchResultCompound{
		Code:       "test_code",
		Query:      "test_query",
		Collapsed:  true,
		Actionable: true,
		Thumbnails: true,
		Rooms: []*roomEncode.RoomData{
			{
				ID:                1,
				Name:              "Test Room",
				OwnerID:           100,
				OwnerName:         "Owner",
				IsPublic:          false,
				DoorMode:          roomEncode.Door(1),
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
			},
			{
				ID:                2,
				Name:              "The promax room",
				OwnerID:           10,
				OwnerName:         "Owner",
				IsPublic:          true,
				DoorMode:          roomEncode.Door(0),
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
			},
		},
	}

	pck := protocol.NewPacket(100)
	r.Encode(&pck)

	decodedResult := encode.SearchResultCompound{}
	err := decodedResult.Decode(&pck)
	assert.NoError(t, err, "Decode should not return an error")
	assert.Equal(t, r.Code, decodedResult.Code, "Code should match")
	assert.Equal(t, r.Query, decodedResult.Query, "Query should match")
	assert.Equal(t, r.Collapsed, decodedResult.Collapsed, "Collapsed should match")
	assert.Equal(t, r.Actionable, decodedResult.Actionable, "Actionable should match")
	assert.Equal(t, r.Thumbnails, decodedResult.Thumbnails, "Thumbnails should match")
	assert.Equal(t, len(r.Rooms), len(decodedResult.Rooms), "Number of rooms should match")
}
