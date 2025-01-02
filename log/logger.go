package log

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"pixels-emulator/config"
)

// CreateTempLogger creates a temporary logger.
func CreateTempLogger() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig = getDefaultEncoderConfig(false)
	cfg.DisableCaller, cfg.DisableStacktrace = true, true
	return zap.Must(cfg.Build())
}

// getDefaultEncoderConfig returns a default zapcore.EncoderConfig.
func getDefaultEncoderConfig(useColor bool) zapcore.EncoderConfig {
	cfg := zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		LineEnding:   zapcore.DefaultLineEnding,
	}
	if useColor {
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	return cfg
}

// SetupLogger sets up the logger based on the configuration and uses buffers only in the TEST environment.
func SetupLogger(cfg *config.Config) []*bytes.Buffer {

	if err := os.MkdirAll("logs", 0755); err != nil {
		zap.S().Fatal("Failed to create log directory: ", err)
	}

	logLevel := zapcore.InfoLevel
	_ = (logLevel).UnmarshalText([]byte(cfg.Logging.Level))

	consoleEncoder := zapcore.NewConsoleEncoder(getDefaultEncoderConfig(cfg.Logging.ConsoleColor))
	fileEncoder := zapcore.NewConsoleEncoder(getDefaultEncoderConfig(false))
	if cfg.Logging.JSON {
		fileEncoder = zapcore.NewJSONEncoder(getDefaultEncoderConfig(false))
	}

	var infoWriter, errorWriter, consoleWriter zapcore.WriteSyncer
	var infoBuffer, errorBuffer, consoleBuffer *bytes.Buffer

	if cfg.Server.Environment == "TEST" {
		infoBuffer = new(bytes.Buffer)
		errorBuffer = new(bytes.Buffer)
		consoleBuffer = new(bytes.Buffer)

		infoWriter = zapcore.AddSync(io.MultiWriter(getLogWriter("logs/log.log"), infoBuffer))
		errorWriter = zapcore.AddSync(io.MultiWriter(getLogWriter("logs/error.log"), errorBuffer))
		consoleWriter = zapcore.AddSync(io.MultiWriter(os.Stdout, consoleBuffer))
	} else {
		infoWriter = zapcore.AddSync(getLogWriter("logs/log.log"))
		errorWriter = zapcore.AddSync(getLogWriter("logs/error.log"))
		consoleWriter = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleWriter, logLevel),
		zapcore.NewCore(fileEncoder, infoWriter, logLevel),
		zapcore.NewCore(fileEncoder, errorWriter, zapcore.ErrorLevel),
	)

	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))

	// Return buffers only if TEST environment; otherwise, return nil slices.
	if cfg.Server.Environment == "TEST" {
		return []*bytes.Buffer{consoleBuffer, infoBuffer, errorBuffer}
	}

	return nil

}

// getLogWriter returns a lumberjack log writer.
func getLogWriter(path string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    50,
		MaxBackups: 5,
		MaxAge:     14,
		Compress:   true,
	})
}
