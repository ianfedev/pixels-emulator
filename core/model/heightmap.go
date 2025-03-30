package model

import "pixels-emulator/core/database"

// HeightMap defines a pre-configured height map which
// can be used as room layout.
type HeightMap struct {
	database.BaseModel

	// Slug is the identifier of the map.
	Slug string `gorm:"type:varchar(255);not null;unique"`

	// DoorX represents the X coordinate of the door.
	DoorX int `gorm:"not null;default:2"`

	// DoorY represents the Y coordinate of the door.
	DoorY int `gorm:"not null;default:2"`

	// DoorDirection represents the direction of the door.
	DoorDirection int `gorm:"not null;default:2"`

	// Heightmap stores the layout of the map as a string.
	Heightmap string `gorm:"type:text;not null"`

	// Essential defines whether the layout is available for all players.
	Essential bool `gorm:"default:false"`
}
