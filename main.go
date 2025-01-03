package main

import (
	"go.uber.org/zap"
	"os"
	"pixels-emulator/config"
	"pixels-emulator/database"
	"pixels-emulator/log"
)

func main() {

	tLog := log.CreateTempLogger()
	tLog.Info("Starting Pixels emulator")

	err := config.CreateDefaultConfig("config.ini", tLog)
	if err != nil {
		tLog.Error("Error while loading configuration", zap.Error(err))
		os.Exit(1)
	}

	cfg, err := config.CreateConfig("config.ini", zap.L())

	if err != nil {
		tLog.Error("Error while loading configuration", zap.Error(err))
		os.Exit(1)
	}

	log.SetupLogger(cfg)
	zap.L().Debug("Logger instantiated")

	_, err = database.SetupDatabase(cfg, zap.L())
	if err != nil {
		tLog.Error("Error while connecting to database", zap.Error(err))
		os.Exit(1)
	}

}
