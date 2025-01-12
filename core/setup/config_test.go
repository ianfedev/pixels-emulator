package setup_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"pixels-emulator/core/config"
	"pixels-emulator/core/setup"
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

	cfg := &config.Config{
		Server: config.ServerConfig{
			IP:          "127.0.0.1",
			Port:        8080,
			Environment: "PRODUCTION",
		},
		Database: config.DatabaseConfig{
			Database: "prod_db",
			Password: "prod_secret",
			User:     "admin",
			Host:     "db.prod.example.com",
			Port:     5432,
		},
	}

	setup.CheckSecurityAlerts(cfg, logger)
	verifySecurityLog(t, buf.String(), []string{"Password", "User", "Host", "Database", "Port"})
}

// TestCheckStruct test if the environment variable message is logged at console.
func TestCheckStruct(t *testing.T) {
	logger, buf := createTestLogger(t)

	cfg := &config.Config{
		Server: config.ServerConfig{
			IP:          "127.0.0.1",
			Port:        8080,
			Environment: "PRODUCTION",
		},
		Database: config.DatabaseConfig{
			Database: "prod_db",
			Password: "prod_secret",
			User:     "admin",
			Host:     "db.prod.example.com",
			Port:     5432,
		},
	}

	// Run the checkStruct function directly
	setup.CheckStruct(reflect.ValueOf(cfg), "PRODUCTION", logger)
	verifySecurityLog(t, buf.String(), []string{"Password", "User", "Host", "Database", "Port"})
}

// TestCheckStructNonStruct verifies that checkStruct correctly handles non-struct values.
func TestCheckStructNonStruct(t *testing.T) {
	logger, buf := createTestLogger(t)

	// Pass a non-struct value, such as an integer
	nonStructValue := 42
	setup.CheckStruct(reflect.ValueOf(nonStructValue), "PRODUCTION", logger)

	// Ensure that no log entries are created since it's not a struct
	if buf.String() != "" {
		t.Errorf("Expected no log output for non-struct value, got: %s", buf.String())
	}
}

// TestCreateConfig_Success tests successful configuration loading.
func TestCreateConfig_Success(t *testing.T) {
	logger, _ := createTestLogger(t)

	tempFile, err := os.CreateTemp("", "test_config_*.ini")
	assert.NoError(t, err)

	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Error removing temp file: %v", err)
		}
	}()

	// Write valid config content to the temporary file
	configContent := `[server]
ip=192.168.1.1
port=8080
environment=PRODUCTION

[database]
host=localhost
port=5432
database=test_db
user=test_user
password=test_password

[logging]
console_color=false
json=true
level=DEBUG`
	err = os.WriteFile(tempFile.Name(), []byte(configContent), 0644)
	if err != nil {
		t.Logf("Error writing temp config on test: %v", err)
		return
	}
	cfg, err := setup.Config(tempFile.Name(), logger)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "192.168.1.1", cfg.Server.IP)
	assert.Equal(t, "PRODUCTION", cfg.Server.Environment)
	assert.Equal(t, "DEBUG", cfg.Logging.Level)
}

// TestCreateConfig_FileNotFound tests when the configuration file is missing.
func TestCreateConfig_FileNotFound(t *testing.T) {
	logger, _ := createTestLogger(t)
	_, err := setup.Config("nonexistent_config.ini", logger)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error reading config file")
}

// TestCreateConfig_InvalidContent tests when the configuration file has invalid content.
func TestCreateConfig_InvalidContent(t *testing.T) {
	logger, _ := createTestLogger(t)

	tempFile, err := os.CreateTemp("", "test_config_*.ini")
	assert.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Error removing temp file: %v", err)
		}
	}()

	// Write invalid config content to the temporary file
	err = os.WriteFile(tempFile.Name(), []byte("[invalid_section\nkey: value"), 0644)
	if err != nil {
		t.Logf("Error writing temp config on test: %v", err)
		return
	}

	_, err = setup.Config(tempFile.Name(), logger)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error reading config file")
}

// TestCreateConfigEnvOverrides tests environment variable overrides.
func TestCreateConfigEnvOverrides(t *testing.T) {
	logger, _ := createTestLogger(t)

	tempFile, err := os.CreateTemp("", "test_config_*.ini")
	assert.NoError(t, err)
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Error removing temp file: %v", err)
		}
	}()

	// Write minimal valid config
	configContent := `[server]
ip=192.168.1.1
port=8080
environment=DEVELOPMENT`

	err = os.WriteFile(tempFile.Name(), []byte(configContent), 0644)
	err = os.Setenv("SERVER_IP", "10.0.0.1")

	if err != nil {
		t.Errorf("Error setting environment for ENV testing %d", err)
	}

	defer func() {
		if err := os.Unsetenv("SERVER_IP"); err != nil {
			t.Logf("Error removing test environment, please clear it on your PC: %v", err)
		}
	}()

	cfg, err := setup.Config(tempFile.Name(), logger)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "10.0.0.1", cfg.Server.IP) // Environment variable override
	assert.Equal(t, uint16(8080), cfg.Server.Port)
	assert.Equal(t, "DEVELOPMENT", cfg.Server.Environment)
}
