package util

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateTestLogger creates a zap logger with a buffer for capturing log output.
func CreateTestLogger() (*zap.Logger, *bytes.Buffer) {
	var buf bytes.Buffer
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.AddSync(&buf), zapcore.DebugLevel)
	logger := zap.New(core)
	return logger, &buf
}

// MockAsyncResponse creates the mocking of an async response for model services.
func MockAsyncResponse[T any](data T, err error) <-chan struct {
	Data  T
	Error error
} {
	ch := make(chan struct {
		Data  T
		Error error
	}, 1)

	go func() {
		ch <- struct {
			Data  T
			Error error
		}{Data: data, Error: err}
		close(ch)
	}()

	return ch
}
