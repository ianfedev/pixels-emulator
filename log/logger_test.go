package log

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"testing"
)

// TestCreateTempLogger tests the CreateTempLogger function
func TestCreateTempLogger(t *testing.T) {

	logger := CreateTempLogger()
	assert.NotNil(t, logger, "Logger should not be nil")

	expectedLevel := zapcore.DebugLevel
	actualLevel := logger.Core().Enabled(expectedLevel)
	assert.True(t, actualLevel, "Logger should be enabled for the Debug level")

}
