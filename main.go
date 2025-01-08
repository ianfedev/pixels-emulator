package main

import (
	"go.uber.org/zap"
	"os"
	config2 "pixels-emulator/core/config"
	"pixels-emulator/core/database"
	"pixels-emulator/core/log"
	"pixels-emulator/core/protocol/registry"
	"pixels-emulator/router"
	"strconv"
)

func main() {

	tLog := log.CreateTempLogger()
	tLog.Info("Starting Pixels emulator")

	err := config2.CreateDefaultConfig("config.ini", tLog)
	if err != nil {
		tLog.Error("Error while loading configuration", zap.Error(err))
		os.Exit(1)
	}

	cfg, err := config2.CreateConfig("config.ini", zap.L())

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

	pReg, hReg := registry.SetupProcessors(), registry.SetupHandlers()

	// As the only method of packet receiving, I will not edit this
	// until further needs. Maybe on future this can be rewritten to
	// support other protocols like TCP sockets or something else.
	app, err := router.SetupRouter(zap.L(), pReg, hReg)
	if err != nil || app == nil {
		tLog.Error("Error while setting up HTTP server", zap.Error(err))
		os.Exit(1)
	}

	bind := cfg.Server.IP + ":" + strconv.Itoa(int(cfg.Server.Port))
	zap.L().Info("Starting HTTP server",
		zap.String("server", cfg.Server.IP), zap.Uint16("port", cfg.Server.Port))

	err = app.Listen(bind)
	if err != nil {
		tLog.Error("Error while starting the server", zap.Error(err))
		os.Exit(1)
	}

}
