package message

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"pixels-emulator/navigator/encode"
	roomEncode "pixels-emulator/room/encode"
	"testing"
)

// Parse reads result packet for testing.
func Parse(pck *protocol.RawPacket) (*NavigatorSearchResultPacket, error) {

	v, err := pck.ReadString()
	q, err := pck.ReadString()
	l, err := pck.ReadInt()

	res := make([]*encode.SearchResultCompound, l)
	for i := 0; i < int(l); i++ {
		compound := &encode.SearchResultCompound{}
		er := compound.Decode(pck)
		if er != nil {
			err = er
		} else {
			res[i] = compound
		}
	}

	return ComposeNavigatorSearchResult(v, q, res), err

}

func Test_Serialize(t *testing.T) {

	r := []*encode.SearchResultCompound{
		{
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
		},
		{
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
		},
	}

	res := ComposeNavigatorSearchResult("test_code", "test_query", r)
	pck := res.Serialize()
	bytes := pck.ToBytes()

	raw, err := protocol.FromBytes(bytes)
	assert.NoError(t, err, "Encoding should not portray error")

	enc, err := Parse(raw)
	assert.NoError(t, err, "Parsing should not portray error")
	assert.Equal(t, res.SearchCode, enc.SearchCode, "Search code should be matched upon serialization")
	assert.Equal(t, res.SearchQuery, enc.SearchQuery, "Search query should be matched upon serialization")
	assert.Equal(t, len(res.Results), len(enc.Results), "Results length should match")

}
