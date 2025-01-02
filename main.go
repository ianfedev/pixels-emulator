package main

import (
	"go.uber.org/zap"
	"os"
	"pixels-emulator/config"
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

	cfg, err := config.CreateConfig("config.ini", tLog)

	if err != nil {
		tLog.Error("Error while loading configuration", zap.Error(err))
		os.Exit(1)
	}

	log.SetupLogger(cfg)
	tLog.Debug("Logger instantiated")

}
