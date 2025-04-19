package room

import (
	"context"
	"pixels-emulator/core/model"
	"pixels-emulator/room/encode"
	"pixels-emulator/room/unit"
	"pixels-emulator/user"
	"strconv"
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

func EncodePlayer(ctx context.Context, p *user.Player, r *Room) (*encode.UnitDetail, *encode.PlayerDetail, error) {

	id, err := strconv.ParseInt(p.Id, 10, 32)
	if err != nil {
		return nil, nil, err
	}

	uData := <-p.Record(ctx)
	if uData.Error != nil {
		return nil, nil, uData.Error
	}
	u := uData.Data

	uDetail := &encode.UnitDetail{
		Id:        int32(id),
		Username:  u.Username,
		Custom:    u.Motto,
		Figure:    u.Look,
		RoomIndex: int32(r.Id),
		UnitX:     int32(p.Unit().Current.X()),
		UnitY:     int32(p.Unit().Current.Y()),
		UnitZ:     int32(p.Unit().Current.Z()),
		Rot:       int32(p.Unit().Current.Dir()),
		Type:      encode.User,
	}

	pDetail := &encode.PlayerDetail{
		Gender:         u.Gender,
		GroupId:        0,
		GroupName:      "",   // TODO: Groups
		SwimFigure:     "",   // INVESTIGATION
		ActivityPoints: 0,    // TODO: Achievements
		Moderator:      true, // TODO: Permissions
	}

	return uDetail, pDetail, nil
}
