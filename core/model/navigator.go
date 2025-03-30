package model

import "pixels-emulator/core/database"

// NavigatorDisplay represents an element in the navigator.
type NavigatorDisplay struct {
	database.BaseModel

	// Name is the display name of the element.
	Name string `gorm:"type:varchar(255);not null"`

	// Realm defines the category of the navigator.
	Realm string `gorm:"type:enum('official_view','hotel_view','roomads_view','myworld_view');not null"`

	// DisplayType deter`gorm:"type:enum('list','thumbnails');not null"`mines how the element is displayed.
	DisplayType string

	// OrderType defines the sorting method.
	OrderType string `gorm:"type:enum('activity','order');not null"`

	// Overridable determines if hidden groups show all rooms.
	Overridable bool `gorm:"default:false"`

	// Priority manages order between categories.
	Priority int `gorm:"default:0"`

	// Filter defines how rooms are selected for this display.
	Filter string `gorm:"type:varchar(50);not null"`
}
