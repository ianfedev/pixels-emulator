package room

import (
	"pixels-emulator/core/model"
	"pixels-emulator/room/encode"
	"pixels-emulator/room/unit"
)

// EncodeUnit codifies a room unit into a binary unit wrapper.
func EncodeUnit(unit *unit.Unit) (*encode.UnitMessage, error) {

	c := unit.Current
	h, b := unit.Rotation()

	return &encode.UnitMessage{
		Id:     unit.Id,
		X:      int32(c.X()),
		Y:      int32(c.Y()),
		Z:      int32(c.Z()),
		Head:   int32(h),
		Body:   int32(b),
		Status: unit.Status,
	}, nil

}

// EncodeSettings creates a protocol version encoded settings.
func EncodeSettings(c *model.RoomConfiguration) *encode.RoomChatSettings {
	return &encode.RoomChatSettings{
		Mode:       encode.ChatMode(c.ChatMode),
		Weight:     encode.ChatWidth(c.ChatWeight),
		Speed:      encode.ChatSpeed(c.ChatSpeed),
		Distance:   int32(c.ChatHearingDistance),
		Protection: encode.ChatFilter(c.ChatProtection),
	}
}

func EncodeRoom(r *model.Room, t *Room) *encode.RoomData {

	var s encode.Door
	switch r.State {
	case "closed":
		s = encode.Locked
		break
	case "password_Protected":
		s = encode.PasswordProtected
		break
	default:
		s = encode.Open
		break
	} // TODO: Other two cases?

	enc := &encode.RoomData{
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
