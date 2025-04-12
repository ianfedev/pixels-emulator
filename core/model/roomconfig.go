package model

// RoomConfiguration represents all configuration settings for a room.
// It includes design, lighting, permissions, chat settings, and other options.
type RoomConfiguration struct {
	// RoomID is both the primary key and the foreign key referencing the associated room.
	RoomID uint `gorm:"primaryKey"`

	// FloorPaper specifies the design or material used for the floor.
	FloorPaper string `gorm:"type:varchar(255)"`

	// WallPaper specifies the design or material used for the walls.
	WallPaper string `gorm:"type:varchar(255)"`

	// LandscapePaper specifies the background or landscape design of the room.
	LandscapePaper string `gorm:"type:varchar(255)"`

	// WallThickness indicates the thickness of the room's walls.
	WallThickness float64

	// WallHeight indicates the height of the room's walls.
	WallHeight float64

	// FloorThickness indicates the thickness of the room's floor.
	FloorThickness float64

	// LightData stores the configuration data for the room's lighting in JSON format.
	LightData string `gorm:"type:json"`

	// AllowPets indicates whether pets are allowed in the room.
	AllowPets bool `gorm:"not null;default:true"`

	// AllowPetsFeed indicates whether pets are allowed to feed in the room.
	AllowPetsFeed bool `gorm:"not null;default:true"`

	// AllowWalkThrough indicates whether users can walk through the room.
	AllowWalkThrough bool `gorm:"not null;default:true"`

	// AllowHideWall indicates whether users can hide the room's walls.
	AllowHideWall bool `gorm:"not null;default:true"`

	// MoveDiagonally indicates whether diagonal movement is allowed in the room.
	MoveDiagonally bool `gorm:"not null;default:false"`

	// ChatMode represents the mode of chat in the room. (0: FreeFlow, 0: 1: One by one)
	ChatMode int `gorm:"type:varchar(50);not null;default:0"`

	// ChatWeight specifies the priority or weight assigned to the chat. (0: Wide, 1: Normal, 2: Thin)
	ChatWeight int `gorm:"not null;default:1"`

	// ChatSpeed indicates the speed at which chat messages are transmitted. (0: Fast, 1: Normal, 2: Slow)
	ChatSpeed int `gorm:"not null;default:1"`

	// ChatHearingDistance indicates the distance at which chat can be heard.
	ChatHearingDistance int `gorm:"not null;default:10"`

	// ChatProtection indicates whether chat protection (such as moderation) is enabled. (0: Strict, 1: Normal, 2: Loose)
	ChatProtection int `gorm:"not null;default:1"`

	// RollerSpeed indicates the speed of moving objects (rollers) in the room.
	RollerSpeed float64 `gorm:"not null;default:1.0"`

	// TradeMode represents the trading mode in the room (e.g., "open", "friends_only", "closed").
	TradeMode string `gorm:"type:varchar(50);not null"`

	// Room is the associated room for this configuration.
	Room *Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RoomID"`
}
