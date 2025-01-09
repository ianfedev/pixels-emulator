package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"pixels-emulator/core/config"
	"time"
)

// GetDSN constructs and returns the Data Source Name (DSN) for connecting to the database.
func GetDSN(cfg *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)
}

// SetupDatabase initializes the database connection using the provided configuration and logger.
// It sets up connection pooling parameters and returns a *gorm.DB instance for interacting with the database.
func SetupDatabase(cfg *config.Config, log *zap.Logger) (*gorm.DB, error) {

	dsn := GetDSN(cfg)

	zLog := zapgorm2.New(zap.L())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: zLog})
	if err != nil {
		return nil, err
	}

	pool, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql pool: %w", err)
	}

	oc, ic, ml := 100, 10, time.Duration(5)*time.Minute
	pool.SetMaxOpenConns(oc)
	pool.SetMaxIdleConns(ic)
	pool.SetConnMaxLifetime(ml)

	log.Info("Database connection established successfully")
	log.Debug(
		"Established connection parameters",
		zap.Int("open_conn", oc), zap.Int("max_idle_conn", ic), zap.Int64("max_lifetime", ml.Milliseconds()))

	return db, nil

}
