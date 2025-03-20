package handler

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/server"
	"pixels-emulator/room/encode"
	"pixels-emulator/room/message/guest"
)

// GetGuestRoomHandler handles non-owned room join attempt.
type GetGuestRoomHandler struct {
	logger *zap.Logger // Logger for packet processing details.
}

// Handle processes the incoming navigation search packet.
func (h *GetGuestRoomHandler) Handle(raw protocol.Packet, conn protocol.Connection) {

	pck, ok := raw.(*guest.GetRoomPacket)
	if !ok {
		h.logger.Error("cannot cast navigator search packet, skipping processing")
		return
	}

	h.logger.Debug("Guest room event", zap.Int32("room", pck.RoomId), zap.Bool("enter", pck.Enter), zap.Bool("forward", pck.Forward))

	// This must not be handled by events due to being a mission-critical
	// packet, which can be queried many times not only for room entering,
	// but for information purposes. So, creating a cancellable event
	// may create conflicts...
	// Therefore, we decided to create cancellable events deeper in the
	// room logic (Like room entering).

	testPck := &guest.ResponseRoomPacket{
		Enter:   pck.Enter,
		Forward: pck.Forward,
		Room: &encode.RoomData{
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
		},
		StaffPick:     true,
		GuildMember:   true,
		GlobalMute:    true,
		CanGlobalMute: false,
		Moderation: &encode.ModerationRights{
			Mute: encode.None,
			Kick: encode.None,
			Ban:  encode.None,
		},
		Settings: &encode.RoomChatSettings{
			Mode:       encode.ChatModeFreeFlow,
			Weight:     encode.ChatBubbleWidthNormal,
			Speed:      encode.ChatScrollSpeedNormal,
			Distance:   10,
			Protection: encode.FloodFilterNormal,
		},
	}
	// TODO: Replace with real data.

	conn.SendPacket(testPck)

}

// NewNavigatorSearch creates a new handler instance.
func NewNavigatorSearch() *GetGuestRoomHandler {
	return &GetGuestRoomHandler{
		logger: server.GetServer().Logger(),
	}
}
