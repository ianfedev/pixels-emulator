package log

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"pixels-emulator/config"
	"testing"
)

// cleanLogsDirectory clears the log folder at the end of the test.
func cleanLogsDirectory(t *testing.T) {
	t.Cleanup(func() {
		err := os.RemoveAll("logs")
		assert.NoError(t, err, "Failed to clean up the logs directory")
	})
}

// checkValidJSON validates that the content of the given buffer is in JSON format.
func checkValidJSON(buffer *bytes.Buffer) error {
	var js map[string]interface{}
	if err := json.NewDecoder(buffer).Decode(&js); err != nil {
		return err
	}
	return nil
}

// generateTestLog setups a logger for all the corresponding tests.
func generateTestLog(t *testing.T, color bool, json bool, environment string) (string, []*bytes.Buffer) {

	cfg := &config.Config{
		Server: config.ServerConfig{
			Environment: environment,
		},
		Logging: config.LoggingConfig{
			ConsoleColor: color,
			JSON:         json,
			Level:        "INFO",
		},
	}

	buffers := SetupLogger(cfg)
	consoleBuffer := buffers[0]

	logger := zap.L()
	assert.NotNil(t, logger, "Logger should not be nil")

	logger.Info("Test message")

	return consoleBuffer.String(), buffers

}

// TestCreateTempLogger tests the CreateTempLogger function
func TestCreateTempLogger(t *testing.T) {

	cleanLogsDirectory(t)
	logger := CreateTempLogger()
	assert.NotNil(t, logger, "Logger should not be nil")

	expectedLevel := zapcore.DebugLevel
	actualLevel := logger.Core().Enabled(expectedLevel)
	assert.True(t, actualLevel, "Logger should be enabled for the Debug level")
}

// TestSetupLogger_ColorBasic tests the SetupLogger function with colors and non JSON logging.
func TestSetupLogger_ColorBasic(t *testing.T) {
	cleanLogsDirectory(t)
	logMsg, _ := generateTestLog(t, true, false, "TEST")
	assert.Contains(t, logMsg, "Test message", "The log message should be in the console buffer")
	assert.Contains(t, logMsg, "\x1b[", "The log message should contain color codes")
}

// TestSetupLogger_Colorless tests the SetupLogger function without colors and non JSON logging.
func TestSetupLogger_Colorless(t *testing.T) {
	cleanLogsDirectory(t)
	logMsg, _ := generateTestLog(t, false, false, "TEST")
	assert.Contains(t, logMsg, "Test message", "The log message should be in the console buffer")
	assert.NotContains(t, logMsg, "\x1b[", "The log message should not contain color codes")
}

// TestSetupLogger_Files tests the SetupLogger function logging at files.
func TestSetupLogger_Files(t *testing.T) {

	cleanLogsDirectory(t)
	_, buff := generateTestLog(t, true, false, "TEST")
	zap.L().Error("error msg")
	info, err := buff[1].String(), buff[2].String()

	assert.Contains(t, info, "error msg", "The log message should be in the file buffer")
	assert.NotContains(t, info, "\x1b[", "The log message should not contain color codes in file")

	assert.Contains(t, err, "error msg", "The error message should be in the file buffer")
	assert.NotContains(t, err, "\x1b[", "The error message should not contain color codes")
	assert.NotContains(t, err, "Test message", "The info message should not log into error file")

}

// TestSetupLogger_FilesJSON tests the SetupLogger function logging at files with JSON format.
func TestSetupLogger_FilesJSON(t *testing.T) {

	cleanLogsDirectory(t)
	_, buff := generateTestLog(t, true, true, "TEST")
	zap.L().Error("error msg")

	// Retrieve the contents of the buffers for the log and error files
	info, err := buff[1].String(), buff[2].String()

	// Check that the log message appears in the console buffer
	assert.Contains(t, info, "error msg", "The log message should be in the file buffer")
	assert.NotContains(t, info, "\x1b[", "The log message should not contain color codes in the file buffer")

	assert.Contains(t, err, "error msg", "The error message should be in the file buffer")
	assert.NotContains(t, err, "\x1b[", "The error message should not contain color codes")

	// Check that the log file is in JSON format (for both log and error files)
	assert.Nil(t, checkValidJSON(buff[1]))
	assert.Nil(t, checkValidJSON(buff[2]))
}

// TestSetupLogger_EmptyBuffers test if buffers provided are empty when environment is not test.
func TestSetupLogger_EmptyBuffers(t *testing.T) {
	cleanLogsDirectory(t)
	cfg := &config.Config{
		Server: config.ServerConfig{
			Environment: "FOO",
		},
	}

	buffers := SetupLogger(cfg)
	assert.Nil(t, buffers)
}
