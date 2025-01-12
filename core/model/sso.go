package model

import "pixels-emulator/core/database"

// SSOTicket represents a single sign-on ticket associated with a user.
type SSOTicket struct {

	// BaseModel includes common fields for all models.
	database.BaseModel

	// UserID is the ID of the associated user.
	UserID uint `gorm:"not null"`

	// User is the associated user for the ticket.
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Ticket is the single sign-on ticket string.
	Ticket string `gorm:"type:varchar(255);not null"`
}
