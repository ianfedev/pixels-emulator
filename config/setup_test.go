package config

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"testing"
)

// createTestLogger creates a logger that writes logs to a buffer for verification.
func createTestLogger(t *testing.T) (*zap.Logger, *bytes.Buffer) {
	t.Helper()
	buf := &bytes.Buffer{}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(buf),
		zap.InfoLevel,
	)
	logger := zap.New(core)
	return logger, buf
}

// verifySecurityLog checks if the expected security alert is logged for specific fields.
func verifySecurityLog(t *testing.T, logOutput string, fields []string) {
	t.Helper()
	assert.Contains(t, logOutput, "Security Alert: Sensitive data detected in configuration")
	for _, field := range fields {
		assert.Contains(t, logOutput, field)
	}
}

// TestCheckSecurityAlerts test if the environment variable message is logged at console.
func TestCheckSecurityAlerts(t *testing.T) {
	logger, buf := createTestLogger(t)

	config := &Config{
		Server: ServerConfig{
			IP:          "127.0.0.1",
			Port:        8080,
			Environment: "PRODUCTION",
		},
		Database: DatabaseConfig{
			Database: "prod_db",
			Password: "prod_secret",
			User:     "admin",
			Host:     "db.prod.example.com",
			Port:     5432,
		},
	}

	CheckSecurityAlerts(config, logger)
	verifySecurityLog(t, buf.String(), []string{"Password", "User", "Host", "Database", "Port"})
}

// TestCheckStruct test if the environment variable message is logged at console.
func TestCheckStruct(t *testing.T) {
	logger, buf := createTestLogger(t)

	config := &Config{
		Server: ServerConfig{
			IP:          "127.0.0.1",
			Port:        8080,
			Environment: "PRODUCTION",
		},
		Database: DatabaseConfig{
			Database: "prod_db",
			Password: "prod_secret",
			User:     "admin",
			Host:     "db.prod.example.com",
			Port:     5432,
		},
	}

	// Run the checkStruct function directly
	checkStruct(reflect.ValueOf(config), "PRODUCTION", logger)
	verifySecurityLog(t, buf.String(), []string{"Password", "User", "Host", "Database", "Port"})
}

// TestCheckStructNonStruct verifies that checkStruct correctly handles non-struct values.
func TestCheckStructNonStruct(t *testing.T) {
	logger, buf := createTestLogger(t)

	// Pass a non-struct value, such as an integer
	nonStructValue := 42
	checkStruct(reflect.ValueOf(nonStructValue), "PRODUCTION", logger)

	// Ensure that no log entries are created since it's not a struct
	if buf.String() != "" {
		t.Errorf("Expected no log output for non-struct value, got: %s", buf.String())
	}
}
