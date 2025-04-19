package encode

import (
	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/protocol"
	"testing"
)

// TestUnitDetailEncodeDecode validates UnitDetail encoding and decoding.
func TestUnitDetailEncodeDecode(t *testing.T) {
	original := &UnitDetail{
		Id:        7,
		Username:  "TestUser",
		Custom:    "TestCustom",
		Figure:    "hd-180-1.ch-210-66",
		RoomIndex: 1,
		UnitX:     2,
		UnitY:     3,
		UnitZ:     4,
		Rot:       6,
		Type:      User,
	}

	pck := protocol.NewPacket(0)
	original.Encode(&pck)

	data := pck.ToBytes()
	parsed, err := protocol.FromBytes(data)
	assert.NoError(t, err)

	result := &UnitDetail{}
	err = result.Decode(parsed)
	assert.NoError(t, err)

	assert.Equal(t, original.Id, result.Id)
	assert.Equal(t, original.Username, result.Username)
	assert.Equal(t, original.Custom, result.Custom)
	assert.Equal(t, original.Figure, result.Figure)
	assert.Equal(t, original.RoomIndex, result.RoomIndex)
	assert.Equal(t, original.UnitX, result.UnitX)
	assert.Equal(t, original.UnitY, result.UnitY)
	assert.Equal(t, original.UnitZ, result.UnitZ)
	assert.Equal(t, original.Rot, result.Rot)
}

// TestPlayerDetailEncodeDecode validates PlayerDetail encoding and decoding.
func TestPlayerDetailEncodeDecode(t *testing.T) {
	original := &PlayerDetail{
		Gender:         "m",
		GroupId:        5,
		GroupName:      "TestGroup",
		SwimFigure:     "",
		ActivityPoints: 42,
		Moderator:      true,
	}

	pck := protocol.NewPacket(0)
	original.Encode(&pck)

	data := pck.ToBytes()
	parsed, err := protocol.FromBytes(data)
	assert.NoError(t, err)

	result := &PlayerDetail{}
	err = result.Decode(parsed)
	assert.NoError(t, err)

	assert.Equal(t, "M", result.Gender)
	assert.Equal(t, original.GroupId, result.GroupId)
	assert.Equal(t, original.GroupName, result.GroupName)
	assert.Equal(t, original.SwimFigure, result.SwimFigure)
	assert.Equal(t, original.ActivityPoints, result.ActivityPoints)
	assert.Equal(t, original.Moderator, result.Moderator)
}

// TestPetDetailEncodeDecode validates PetDetail encoding and decoding.
func TestPetDetailEncodeDecode(t *testing.T) {
	original := &PetDetail{
		SubType:         1,
		OwnerId:         42,
		OwnerName:       "Coco",
		Rarity:          3,
		Saddle:          true,
		Riding:          false,
		Breed:           true,
		Harvest:         true,
		Revive:          false,
		BreedPermission: true,
		Level:           5,
		Posture:         "sit",
	}

	pck := protocol.NewPacket(0)
	original.Encode(&pck)

	data := pck.ToBytes()
	parsed, err := protocol.FromBytes(data)
	assert.NoError(t, err)

	result := &PetDetail{}
	err = result.Decode(parsed)
	assert.NoError(t, err)

	assert.Equal(t, original.SubType, result.SubType)
	assert.Equal(t, original.OwnerId, result.OwnerId)
	assert.Equal(t, original.OwnerName, result.OwnerName)
	assert.Equal(t, original.Rarity, result.Rarity)
	assert.Equal(t, original.Saddle, result.Saddle)
	assert.Equal(t, original.Riding, result.Riding)
	assert.Equal(t, original.Breed, result.Breed)
	assert.Equal(t, original.Harvest, result.Harvest)
	assert.Equal(t, original.Revive, result.Revive)
	assert.Equal(t, original.BreedPermission, result.BreedPermission)
	assert.Equal(t, original.Level, result.Level)
	assert.Equal(t, original.Posture, result.Posture)
}

// TestRentableBotDetailEncodeDecode validates RentableBotDetail encoding and decoding.
func TestRentableBotDetailEncodeDecode(t *testing.T) {
	original := &RentableBotDetail{
		Gender:    "F",
		OwnerId:   101,
		OwnerName: "Admin",
		Skills:    []int16{10, 20, 30},
	}

	pck := protocol.NewPacket(0)
	original.Encode(&pck)

	data := pck.ToBytes()
	parsed, err := protocol.FromBytes(data)
	assert.NoError(t, err)

	result := &RentableBotDetail{}
	err = result.Decode(parsed)
	assert.NoError(t, err)

	assert.Equal(t, original.Gender, result.Gender)
	assert.Equal(t, original.OwnerId, result.OwnerId)
	assert.Equal(t, original.OwnerName, result.OwnerName)
	assert.Equal(t, original.Skills, result.Skills)
}
