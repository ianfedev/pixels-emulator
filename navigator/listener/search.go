package listener

import (
	"errors"
	"go.uber.org/zap"
	"pixels-emulator/core/event"
	"pixels-emulator/core/server"
	"pixels-emulator/navigator/encode"
	eventNav "pixels-emulator/navigator/event"
	"pixels-emulator/navigator/message"
	"pixels-emulator/room"
	roomEncode "pixels-emulator/room/encode"
)

func ProvideSearch() func(event event.Event) {
	return func(event event.Event) {
		OnNavigatorSearch(event)
	}
}

func OnNavigatorSearch(ev event.Event) {

	var err error
	defer func() {
		if err != nil {
			server.GetServer().Logger().Error("error during authentication handle", zap.Error(err))
		}
	}()

	queryEv, valid := ev.(*eventNav.NavigatorQueryEvent)
	if !valid {
		err = errors.New("event proportioned was not authentication")
		return
	}

	r := []*encode.SearchResultCompound{
		{
			Code:       queryEv.Realm(),
			Query:      queryEv.RawQuery(),
			Collapsed:  false,
			Actionable: false,
			Thumbnails: true,
			Rooms: []*roomEncode.RoomData{
				{
					ID:                1,
					Name:              "Test Room",
					OwnerID:           100,
					OwnerName:         "Owner",
					IsPublic:          false,
					DoorMode:          room.Door(1),
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
					DoorMode:          room.Door(0),
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
			Code:       queryEv.Realm(),
			Query:      queryEv.RawQuery(),
			Collapsed:  true,
			Actionable: true,
			Thumbnails: false,
			Rooms: []*roomEncode.RoomData{
				{
					ID:                1,
					Name:              "Test Room",
					OwnerID:           100,
					OwnerName:         "Owner",
					IsPublic:          false,
					DoorMode:          room.Door(1),
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
					DoorMode:          room.Door(0),
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

	cr := message.ComposeNavigatorSearchResult(queryEv.Realm(), queryEv.RawQuery(), r)
	queryEv.Conn().SendPacket(cr)

}
