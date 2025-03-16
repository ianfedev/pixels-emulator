package guest

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"pixels-emulator/room"
	"pixels-emulator/room/encode"
	"testing"
)

var mockPacket = &ResponseRoomPacket{
	Enter:   true,
	Forward: true,
	Room: &encode.RoomData{
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

// decodeGuestResponsePacket acts as helper function to decode the packet.
func decodeGuestResponsePacket(pck *protocol.RawPacket) (*ResponseRoomPacket, error) {

	enter, err := pck.ReadBoolean()

	roomData := &encode.RoomData{}
	err = roomData.Decode(pck)

	forward, err := pck.ReadBoolean()
	staffPick, err := pck.ReadBoolean()
	guildMember, err := pck.ReadBoolean()
	globalMute, err := pck.ReadBoolean()

	moderation := &encode.ModerationRights{}
	err = moderation.Decode(pck)

	canMute, err := pck.ReadBoolean()

	settings := &encode.RoomChatSettings{}
	err = settings.Decode(pck)

	return &ResponseRoomPacket{
		Enter:         enter,
		Forward:       forward,
		Room:          roomData,
		StaffPick:     staffPick,
		GuildMember:   guildMember,
		GlobalMute:    globalMute,
		CanGlobalMute: canMute,
		Moderation:    moderation,
		Settings:      settings,
	}, err

}

// TestResponseRoomPacket_Serialize checks if packet serialization and deserialization is working correctly.
func TestResponseRoomPacket_Serialize(t *testing.T) {
	pck := mockPacket.Serialize()
	bytes := pck.ToBytes()
	raw, err := protocol.FromBytes(bytes)
	assert.NoError(t, err, "Must not have error on raw decoding")
	res, err := decodeGuestResponsePacket(raw)
	assert.Equal(t, mockPacket, res)
}
