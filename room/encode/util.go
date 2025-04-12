package encode

import (
	"pixels-emulator/core/model"
	"pixels-emulator/room"
)

func RoomToEncodable(r *model.Room, t *room.Room) *RoomData {

	var s Door
	switch r.State {
	case "closed":
		s = Locked
		break
	case "password_Protected":
		s = PasswordProtected
		break
	default:
		s = Open
		break
	} // TODO: Other two cases?

	enc := &RoomData{
		ID:                int32(r.ID),
		Name:              r.Name,
		OwnerID:           int32(r.OwnerID),
		OwnerName:         r.Owner.Username,
		IsPublic:          r.IsPublic,
		DoorMode:          s,
		UserCount:         int32(len(t.Players)),
		UserMax:           int32(r.UsersMax),
		Description:       r.Description,
		Score:             0, // TODO: Get this
		Category:          0,
		Tags:              make([]string, 0),
		GuildID:           0,
		GuildName:         "",
		GuildBadge:        "",
		PromotionTitle:    "",
		PromotionDesc:     "",
		PromotionTime:     120,
		Thumbnail:         "",
		FeaturedPromotion: false,
		AllowPets:         r.Configuration.AllowPets, // End of get this
	}

	return enc

}

func SettingsToEncodable(c *model.RoomConfiguration) *RoomChatSettings {
	return &RoomChatSettings{
		Mode:       ChatMode(c.ChatMode),
		Weight:     ChatWidth(c.ChatWeight),
		Speed:      ChatSpeed(c.ChatSpeed),
		Distance:   int32(c.ChatHearingDistance),
		Protection: ChatFilter(c.ChatProtection),
	}
}
