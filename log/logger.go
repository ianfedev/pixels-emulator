package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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

// SetupLogger sets up the logger based on the configuration.
func SetupLogger(cfg *config.Config) {
	if err := os.MkdirAll("log", 0755); err != nil {
		zap.S().Fatal("Failed to create log directory: ", err)
	}

	logLevel := zapcore.InfoLevel
	_ = logLevel.UnmarshalText([]byte(cfg.Logging.Level))
	fileEncoder := zapcore.NewConsoleEncoder(getDefaultEncoderConfig(false))
	if cfg.Logging.JSON {
		fileEncoder = zapcore.NewJSONEncoder(getDefaultEncoderConfig(false))
	}

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(getDefaultEncoderConfig(cfg.Logging.ConsoleColor)), zapcore.Lock(os.Stdout), logLevel),
		zapcore.NewCore(fileEncoder, getLogWriter("log/log.log"), logLevel),
		zapcore.NewCore(fileEncoder, getLogWriter("log/error.log"), zapcore.ErrorLevel),
	)

	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
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
