package model

import "pixels-emulator/core/database"

// User represents a user in the system.
type User struct {

	// BaseModel includes common fields for all models.
	database.BaseModel

	// Username is the user's full name.
	Username string `gorm:"type:varchar(255);not null;unique"`

	// Motto is the user's personal motto or quote.
	Motto string `gorm:"type:varchar(255)"`

	// Look is the user's avatar appearance or look.
	Look string `gorm:"type:varchar(255)"`

	// Gender represents the user's gender ('F' or 'M').
	Gender string `gorm:"type:char(1);not null"`

	// Credits is the user's balance of credits.
	Credits int `gorm:"not null;default:0"`

	// Pixels is the user's balance of pixels.
	Pixels int `gorm:"not null;default:0"`

	// Duckets is the user's balance of duckets.
	Duckets int `gorm:"not null;default:0"`

	// SSOTickets are the user's associated single sign-on tickets.
	SSOTickets []SSOTicket `gorm:"foreignKey:UserID"`

	// Roles define the user roles
	Roles []Role `gorm:"many2many:user_roles"`
}
