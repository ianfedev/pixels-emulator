package encode

import (
	"pixels-emulator/core/protocol"
	"strconv"
	"strings"
)

// UnitType defines an encodable type of unit.
type UnitType int32

const (
	User     = 1 // User represents a default playable character
	Pet      = 2 // Pet represents a room pet
	Bot      = 3 // Bot represents a room permanent bot
	Rentable = 4 // Rentable represents an ephemeral bot
)

// UnitDetail defines essential information for units to be encoded
// in nitro client.
type UnitDetail struct {
	protocol.Encodable
	Id                       int32    // Id defines the unit identifier
	Username                 string   // Username defines the name of the unit
	Custom                   string   // Custom defines custom information depending on the unit.
	Figure                   string   // Figure defines the look string.
	RoomIndex                int32    // RoomIndex defines the room id of the unit
	UnitX, UnitY, UnitZ, Rot int32    // UnitX, UnitY, UnitZ and Rot defines the unit position and rotation.
	Type                     UnitType // Type defines the unit type to parse further data
}

// Encode transforms unit details into binary packet.
func (u *UnitDetail) Encode(pck *protocol.RawPacket) {
	pck.AddInt(u.Id)
	pck.AddString(u.Username)
	pck.AddString(u.Custom)
	pck.AddString(u.Figure)
	pck.AddInt(u.RoomIndex)
	pck.AddInt(u.UnitX)
	pck.AddInt(u.UnitY)
	pck.AddString(strconv.Itoa(int(u.UnitZ)))
	pck.AddInt(u.Rot)
}

// Decode parses unit detail from binary packet.
func (u *UnitDetail) Decode(pck *protocol.RawPacket) error {
	var err error

	if u.Id, err = pck.ReadInt(); err != nil {
		return err
	}

	if u.Username, err = pck.ReadString(); err != nil {
		return err
	}

	if u.Custom, err = pck.ReadString(); err != nil {
		return err
	}

	if u.Figure, err = pck.ReadString(); err != nil {
		return err
	}

	if u.RoomIndex, err = pck.ReadInt(); err != nil {
		return err
	}

	if u.UnitX, err = pck.ReadInt(); err != nil {
		return err
	}

	if u.UnitY, err = pck.ReadInt(); err != nil {
		return err
	}

	var zStr string
	if zStr, err = pck.ReadString(); err != nil {
		return err
	}

	if z, convErr := strconv.Atoi(zStr); convErr != nil {
		return convErr
	} else {
		u.UnitZ = int32(z)
	}

	if u.Rot, err = pck.ReadInt(); err != nil {
		return err
	}

	return nil
}

// PlayerDetail defines the types to be parsed as subtype of details.
type PlayerDetail struct {
	Gender         string // Gender defines if character is M or F
	GroupId        int32  // GroupId defines the id of the player group.
	GroupName      string // GroupName defines the name of the player group.
	SwimFigure     string // SwimFigure INVESTIGATION
	ActivityPoints int32  // ActivityPoints define the sum of level points at achievement track.
	Moderator      bool   // Moderator defines if user has moderation access.
}

// Encode transforms player details into binary data.
func (p *PlayerDetail) Encode(pck *protocol.RawPacket) {
	pck.AddString(strings.ToUpper(p.Gender))
	pck.AddInt(p.GroupId)

	if p.GroupId == 0 {
		pck.AddInt(-1)
	} else {
		pck.AddInt(1)
	}

	pck.AddString(p.GroupName)
	pck.AddString("")
	pck.AddInt(p.ActivityPoints)
	pck.AddBoolean(p.Moderator)
}

// Decode parses player detail data from the raw packet.
func (p *PlayerDetail) Decode(pck *protocol.RawPacket) error {
	var err error

	if p.Gender, err = pck.ReadString(); err != nil {
		return err
	}
	p.Gender = strings.ToUpper(p.Gender)

	if p.GroupId, err = pck.ReadInt(); err != nil {
		return err
	}

	if _, err = pck.ReadInt(); err != nil {
		return err
	}

	if p.GroupName, err = pck.ReadString(); err != nil {
		return err
	}

	if p.SwimFigure, err = pck.ReadString(); err != nil {
		return err
	}

	if p.ActivityPoints, err = pck.ReadInt(); err != nil {
		return err
	}

	if p.Moderator, err = pck.ReadBoolean(); err != nil {
		return err
	}

	return nil
}

// PetDetail defines the types to be parsed as pet details.
type PetDetail struct {
	protocol.Encodable
	SubType         int32  // SubType of the pet
	OwnerId         int32  // OwnerId of the pet
	OwnerName       string // OwnerName of the pet
	Rarity          int32  // Rarity of the pet
	Saddle          bool   // Saddle if pet has saddle
	Riding          bool   // Riding if user is riding the pet
	Breed           bool   // Breed if pet can breed
	Harvest         bool   // Harvest if pet can harvest
	Revive          bool   // Revive if pet can revive after death
	BreedPermission bool   // BreedPermission if owner has breeding permission
	Level           int32  // Level of the pet
	Posture         string // Posture status of the pet
}

func (p *PetDetail) Encode(pck *protocol.RawPacket) {
	pck.AddInt(p.SubType)
	pck.AddInt(p.OwnerId)
	pck.AddString(p.OwnerName)
	pck.AddInt(p.Rarity)
	pck.AddBoolean(p.Saddle)
	pck.AddBoolean(p.Riding)
	pck.AddBoolean(p.Breed)
	pck.AddBoolean(p.Harvest)
	pck.AddBoolean(p.Revive)
	pck.AddBoolean(p.BreedPermission)
	pck.AddInt(p.Level)
	pck.AddString(p.Posture)
}

// Decode parses pet detail data from the raw packet.
func (p *PetDetail) Decode(pck *protocol.RawPacket) error {
	var err error

	if p.SubType, err = pck.ReadInt(); err != nil {
		return err
	}
	if p.OwnerId, err = pck.ReadInt(); err != nil {
		return err
	}
	if p.OwnerName, err = pck.ReadString(); err != nil {
		return err
	}
	if p.Rarity, err = pck.ReadInt(); err != nil {
		return err
	}
	if p.Saddle, err = pck.ReadBoolean(); err != nil {
		return err
	}
	if p.Riding, err = pck.ReadBoolean(); err != nil {
		return err
	}
	if p.Breed, err = pck.ReadBoolean(); err != nil {
		return err
	}
	if p.Harvest, err = pck.ReadBoolean(); err != nil {
		return err
	}
	if p.Revive, err = pck.ReadBoolean(); err != nil {
		return err
	}
	if p.BreedPermission, err = pck.ReadBoolean(); err != nil {
		return err
	}
	if p.Level, err = pck.ReadInt(); err != nil {
		return err
	}
	if p.Posture, err = pck.ReadString(); err != nil {
		return err
	}

	return nil
}

type RentableBotDetail struct {
	protocol.Encodable
	Gender    string  // Gender of the bot
	OwnerId   int32   // OwnerId of the bot
	OwnerName string  // OwnerName of the bot
	Skills    []int16 // Skills list of the bot
}

func (b *RentableBotDetail) Encode(pck *protocol.RawPacket) {
	pck.AddString(b.Gender)
	pck.AddInt(b.OwnerId)
	pck.AddString(b.OwnerName)
	pck.AddInt(int32(len(b.Skills)))
	for _, skill := range b.Skills {
		pck.AddShort(skill)
	}
}

// Decode parses rentable bot detail data from the raw packet.
func (b *RentableBotDetail) Decode(pck *protocol.RawPacket) error {
	var err error

	if b.Gender, err = pck.ReadString(); err != nil {
		return err
	}
	if b.OwnerId, err = pck.ReadInt(); err != nil {
		return err
	}
	if b.OwnerName, err = pck.ReadString(); err != nil {
		return err
	}
	var skillCount int32
	if skillCount, err = pck.ReadInt(); err != nil {
		return err
	}

	b.Skills = make([]int16, skillCount)
	for i := int32(0); i < skillCount; i++ {
		if b.Skills[i], err = pck.ReadShort(); err != nil {
			return err
		}
	}

	return nil
}
