package database_test

import (
	"pixels-emulator/core/config"
	"pixels-emulator/core/database"
	"testing"
)

func TestGetDSN(t *testing.T) {
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			User:     "test_user",
			Password: "test_password",
			Host:     "localhost",
			Port:     3306,
			Database: "test_db",
		},
	}

	expectedDSN := "test_user:test_password@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	actualDSN := database.GetDSN(cfg)

	if actualDSN != expectedDSN {
		t.Errorf("Expected DSN %s, got %s", expectedDSN, actualDSN)
	}
}

func TestLoggerSetUp(t *testing.T) {

}
