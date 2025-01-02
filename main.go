package main

import (
	"go.uber.org/zap"
	"pixels-emulator/config"
	log2 "pixels-emulator/log"
)

func main() {

	log, _ := zap.NewDevelopment()

	cfg, err := config.CreateConfig("config.ini", log)

	if err != nil {
		panic(err)
	}

	log2.SetupLogger(cfg)
	zap.L().Info("XDE")
	zap.L().Error("HEHE")

}
