package encode

import (
	"pixels-emulator/core/protocol"
	"pixels-emulator/room"
)

// RoomData represents the essential data of a room.
type RoomData struct {
	protocol.Encodable
	ID                int32     // ID is the unique identifier of the room.
	Name              string    // Name is the display name of the room.
	OwnerID           int32     // OwnerID is the unique identifier of the room owner.
	OwnerName         string    // OwnerName is the display name of the room owner.
	IsPublic          bool      // IsPublic indicates whether the room is public or private.
	DoorMode          room.Door // DoorMode represents the room access state (0 = open, 1 = locked, 2 = password-protected, 3 = invisible).
	UserCount         int32     // UserCount is the current number of users in the room.
	UserMax           int32     // UserMax is the maximum number of users allowed in the room.
	Description       string    // Description is the textual description of the room.
	Score             int32     // Score represents the popularity score of the room.
	Category          int32     // Category is the navigation category ID of the room.
	Tags              []string  // Tags is a list of keywords associated with the room.
	GuildID           int32     // GuildID is the unique identifier of the associated guild, or 0 if none.
	GuildName         string    // GuildName is the name of the associated guild, if applicable.
	GuildBadge        string    // GuildBadge is the badge code of the associated guild, if applicable.
	PromotionTitle    string    // PromotionTitle is the title of the active room promotion, if applicable.
	PromotionDesc     string    // PromotionDesc is the description of the active room promotion, if applicable.
	PromotionTime     int32     // PromotionTime is the remaining time (in minutes) for the active promotion.
	Thumbnail         string    // Thumbnail is room has staff pick banner.
	AllowPets         bool      // AllowPets indicates whether pets are allowed in the room.
	FeaturedPromotion bool      // FeaturedPromotion indicates whether the room has a featured promotion.
}

// GenerateBitmask generates the bitmask for the room based on its attributes.
func (r *RoomData) GenerateBitmask() int {

	bitmask := 0

	if r.Thumbnail != "" {
		bitmask = Enable(bitmask, Thumbnail)
	}

	if r.GuildID > 0 {
		bitmask = Enable(bitmask, Guild)
	}

	if r.PromotionTitle != "" {
		bitmask = Enable(bitmask, Promotion)
	}

	if r.IsPublic {
		bitmask = Enable(bitmask, Owner)
	}

	if r.AllowPets {
		bitmask = Enable(bitmask, Pets)
	}

	if r.FeaturedPromotion {
		bitmask = Enable(bitmask, FeaturedPromotion)
	}

	return bitmask
}

// Encode writes RoomData into a RawPacket.
func (r *RoomData) Encode(pck *protocol.RawPacket) {

	// Basic room data
	pck.AddInt(r.ID)
	pck.AddString(r.Name)
	pck.AddInt(r.OwnerID)
	pck.AddString(r.OwnerName)

	// Numbers
	pck.AddInt(int32(r.DoorMode))
	pck.AddInt(r.UserCount)
	pck.AddInt(r.UserMax)

	pck.AddString(r.Description)
	pck.AddInt(0) // TradeMode is empty (INVESTIGATION: Use)
	pck.AddInt(r.Score)
	pck.AddInt(0) // Ranking is empty (NITRO: Useless)
	pck.AddInt(r.Category)

	// Tags
	pck.AddInt(int32(len(r.Tags)))
	for _, tag := range r.Tags {
		pck.AddString(tag)
	}

	bitmask := r.GenerateBitmask()
	pck.AddInt(int32(bitmask))

	// Add data for Thumbnail (if bitmask includes it)
	if Has(bitmask, Thumbnail) {
		pck.AddString(r.Thumbnail)
	}

	// Add data for Group data (if bitmask includes it)
	if Has(bitmask, Guild) {
		pck.AddInt(r.GuildID)
		pck.AddString(r.GuildName)
		pck.AddString(r.GuildBadge)
	}

	// Add data for Room Advertisement (if bitmask includes it)
	if Has(bitmask, Promotion) {
		pck.AddString(r.PromotionTitle)
		pck.AddString(r.PromotionDesc)
		pck.AddInt(r.PromotionTime)
	}

}

// Decode enforces type assertion.
func (r *RoomData) Decode(pck *protocol.RawPacket) error {

	// Basic room data
	id, err := pck.ReadInt()
	r.ID = id
	name, err := pck.ReadString()
	r.Name = name
	owner, err := pck.ReadInt()
	r.OwnerID = owner
	oName, err := pck.ReadString()
	r.OwnerName = oName

	mode, err := pck.ReadInt()
	r.DoorMode = room.Door(mode)
	count, err := pck.ReadInt()
	r.UserCount = count
	uMax, err := pck.ReadInt()
	r.UserMax = uMax

	desc, err := pck.ReadString()
	r.Description = desc
	_, err = pck.ReadInt() // 0
	score, err := pck.ReadInt()
	r.Score = score
	_, err = pck.ReadInt()
	category, err := pck.ReadInt()
	r.Category = category

	tLen, err := pck.ReadInt()
	tags := make([]string, tLen)
	for i := 0; i < int(tLen); i++ {
		tag, e := pck.ReadString()
		err = e
		tags[i] = tag
	}
	r.Tags = tags

	rawMask, err := pck.ReadInt()
	bitmask := int(rawMask)

	thumbnail := ""
	var gId int32 = 0
	gName := ""
	gBadge := ""
	pTitle := ""
	pDesc := ""
	var pTime int32 = 0

	if Has(bitmask, Thumbnail) {
		thumbnail, err = pck.ReadString()
	}

	// Add data for Group data (if bitmask includes it)
	if Has(bitmask, Guild) {
		gId, err = pck.ReadInt()
		gName, err = pck.ReadString()
		gBadge, err = pck.ReadString()
	}

	// Add data for Room Advertisement (if bitmask includes it)
	if Has(bitmask, Promotion) {
		pTitle, err = pck.ReadString()
		pDesc, err = pck.ReadString()
		pTime, err = pck.ReadInt()
	}

	if Has(bitmask, Owner) {
		r.IsPublic = false
	}

	if Has(bitmask, Pets) {
		r.AllowPets = true
	}

	if Has(bitmask, FeaturedPromotion) {
		r.FeaturedPromotion = true
	}

	r.Thumbnail = thumbnail
	r.GuildID = gId
	r.GuildName = gName
	r.GuildBadge = gBadge
	r.PromotionTitle = pTitle
	r.PromotionDesc = pDesc
	r.PromotionTime = pTime

	return err

}
