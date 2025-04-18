package setup

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"pixels-emulator/core/model"
)

// ModelMigration setups model migrations.
func ModelMigration(logger *zap.Logger, db *gorm.DB) error {
	logger.Info("Performing database migrations")
	return db.AutoMigrate(
		&model.HeightMap{},
		&model.NavigatorDisplay{},
		&model.User{},
		&model.SSOTicket{},
		&model.Room{},
		&model.RoomConfiguration{},
		&model.RoomPermission{},
		&model.Role{},
		&model.RolePermission{},
	)
}
