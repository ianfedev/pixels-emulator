package main

import (
	"go.uber.org/zap"
	"os"
	config2 "pixels-emulator/core/config"
	"pixels-emulator/core/database"
	"pixels-emulator/core/event"
	"pixels-emulator/core/log"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/setup"
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

	cfg, err := setup.Config("config.ini", zap.L())

	if err != nil {
		tLog.Error("Error while loading configuration", zap.Error(err))
		os.Exit(1)
	}

	log.SetupLogger(cfg)
	zap.L().Debug("Logger instantiated")

	db, err := database.SetupDatabase(cfg, zap.L())
	if err != nil {
		tLog.Error("Error while connecting to database", zap.Error(err))
		os.Exit(1)
	}

	err = setup.ModelMigration(zap.L(), db)
	if err != nil {
		tLog.Error("Error while generating model migrations", zap.Error(err))
		os.Exit(1)
	}

	conStore := protocol.NewConnectionStore()
	cron := setup.Cron(cfg, conStore)
	(*cron).Start()

	pReg := setup.Processors()

	em := event.NewManager()
	setup.Event(em, conStore)

	hReg := setup.Handlers(zap.L(), cfg, cron, db, conStore, em)

	zap.L().Info("Starting scheduler")

	// As the only method of packet receiving, I will not edit this
	// until further needs. Maybe on future this can be rewritten to
	// support other protocols like TCP sockets or something else.
	app, err := setup.Router(zap.L(), pReg, hReg, conStore)
	if err != nil || app == nil {
		tLog.Error("Error while setting up HTTP healthcheck", zap.Error(err))
		os.Exit(1)
	}

	bind := cfg.Server.IP + ":" + strconv.Itoa(int(cfg.Server.Port))
	zap.L().Info("Starting HTTP healthcheck",
		zap.String("healthcheck", cfg.Server.IP), zap.Uint16("port", cfg.Server.Port))

	err = app.Listen(bind)
	if err != nil {
		tLog.Error("Error while starting the healthcheck", zap.Error(err))
		os.Exit(1)
	}

}
