package model

import (
	"pixels-emulator/core/database"
)

// Role defines a group of permissions which can be assigned to a user.
type Role struct {

	// BaseModel includes common fields for all models.
	database.BaseModel

	// Name defines the identifier of the group.
	Name string `gorm:"size:100;not null;unique"`

	// Color defines a hexadecimal code for the rank color.
	Color string `gorm:"size:7;not null"`

	// Priority defines group priority order. Lower is higher.
	Priority int `gorm:"not null"`

	// Permissions define the role of the permission.
	Permissions []RolePermission `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}

// RolePermission represents a many-to-many relationship between roles and permissions.
type RolePermission struct {

	// ID defines the common id model. (This is not necessary to be time stamped).
	ID uint `gorm:"primaryKey;autoIncrement"`

	// RoleID defines the parent id role.
	RoleID uint `gorm:"not null;index"`

	// Permission defines the dotted style permission.
	Permission string `gorm:"size:255;not null"`
}
