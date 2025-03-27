package model

import "pixels-emulator/core/database"

// Room represents a room entity with its basic information.
type Room struct {
	// BaseModel includes common fields for all models.
	database.BaseModel

	// Name is the name of the room.
	Name string `gorm:"type:varchar(255);not null"`

	// Description provides detailed information about the room.
	Description string `gorm:"type:text"`

	// Password restricts access to the room if set.
	Password string `gorm:"type:varchar(100)"`

	// State represents the current state of the room (e.g., "open", "closed", "password_protected").
	State string `gorm:"type:varchar(50);not null"`

	// UsersMax specifies the maximum number of users allowed in the room.
	UsersMax int `gorm:"not null"`

	// Tags are keywords or labels associated with the room.
	Tags string `gorm:"type:text"`

	// IsPublic indicates whether the room is publicly accessible.
	IsPublic bool `gorm:"not null;default:true"`

	// Configuration holds the one-to-one relationship to the room's configuration settings.
	Configuration RoomConfiguration `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// OwnerID represents the user who owns the room.
	OwnerID uint `gorm:"not null"`

	// Owner is the user who owns the room.
	Owner User `gorm:"foreignKey:OwnerID"`

	// TODO: Correlation Room model, owner, guild, category, votes, staff picks, mute permissions, ban permissions, poll, promotions.
}

// RoomPermission represents a user's permission to a specific room.
type RoomPermission struct {
	// BaseModel includes common fields for all models.
	database.BaseModel

	// RoomID is the ID of the room.
	RoomID uint `gorm:"not null;index;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`

	// UserID is the ID of the user receiving the permission.
	UserID uint `gorm:"not null;index;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`

	// Unique constraint to prevent duplicate entries.
	UniqueIndex string `gorm:"uniqueIndex:room_user_unique"`
}
