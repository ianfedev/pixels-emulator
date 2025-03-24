package database

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel includes common fields for all models.
type BaseModel struct {

	// ID is the primary key of the model.
	ID uint `gorm:"primaryKey;autoIncrement"`

	// CreatedAt is the timestamp when the model was created.
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// UpdatedAt is the timestamp when the model was last updated.
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// DeletedAt is the timestamp when the model was soft-deleted.
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
