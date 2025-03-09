package encode

import "pixels-emulator/core/protocol"

// RoomData represents the essential data of a room.
type RoomData struct {
	// ID is the unique identifier of the room.
	ID int32

	// Name is the display name of the room.
	Name string

	// OwnerID is the unique identifier of the room owner.
	OwnerID int32

	// OwnerName is the display name of the room owner.
	OwnerName string

	// IsPublic indicates whether the room is public or private.
	IsPublic bool

	// State represents the room access state (0 = open, 1 = locked, 2 = password-protected).
	State int8

	// UserCount is the current number of users in the room.
	UserCount int16

	// UserMax is the maximum number of users allowed in the room.
	UserMax int16

	// Description is the textual description of the room.
	Description string

	// Score represents the popularity score of the room.
	Score int32

	// Category is the navigation category ID of the room.
	Category int16

	// Tags is a list of keywords associated with the room.
	Tags []string

	// GuildID is the unique identifier of the associated guild, or 0 if none.
	GuildID int32

	// GuildName is the name of the associated guild, if applicable.
	GuildName string

	// GuildBadge is the badge code of the associated guild, if applicable.
	GuildBadge string

	// PromotionTitle is the title of the active room promotion, if applicable.
	PromotionTitle string

	// PromotionDesc is the description of the active room promotion, if applicable.
	PromotionDesc string

	// PromotionTime is the remaining time (in minutes) for the active promotion.
	PromotionTime int32
}

// CalculateFlags computes the room's flag value based on its attributes.
func (r *RoomData) CalculateFlags() int8 {
	var flags int8

	if r.GuildID > 0 {
		flags |= 2 // Room is associated with a guild
	}
	if r.PromotionTitle != "" {
		flags |= 4 // Room has an active promotion
	}
	if !r.IsPublic {
		flags |= 8 // Room is not public
	}

	return flags
}

// Encode writes RoomData into a RawPacket.
func (r *RoomData) Encode(pck *protocol.RawPacket) {

	pck.AddInt(r.ID)
	pck.AddString(r.Name)
	pck.AddInt(r.OwnerID)
	pck.AddString(r.OwnerName)
	pck.AddShort(int16(r.State))
	pck.AddShort(r.UserCount)
	pck.AddShort(r.UserMax)
	pck.AddString(r.Description)
	pck.AddInt(r.Score)
	pck.AddShort(r.Category)

	// Serialize tags as an array
	pck.AddShort(int16(len(r.Tags)))
	for _, tag := range r.Tags {
		pck.AddString(tag)
	}

	// Calculate and add flags dynamically
	pck.AddShort(int16(r.CalculateFlags()))

	// Serialize Guild data if applicable
	if r.GuildID > 0 {
		pck.AddInt(r.GuildID)
		pck.AddString(r.GuildName)
		pck.AddString(r.GuildBadge)
	}

	// Serialize Promotion data if applicable
	if r.PromotionTitle != "" {
		pck.AddString(r.PromotionTitle)
		pck.AddString(r.PromotionDesc)
		pck.AddInt(r.PromotionTime)
	}

}
