package main

import (
	"go.uber.org/zap"
	"pixels-emulator/config"
)

func main() {

	log, _ := zap.NewDevelopment()

	cfg, err := config.CreateConfig("config.ini", log)

	if err != nil {
		panic(err)
	}

	log.Info(cfg.Server.Environment)

}
