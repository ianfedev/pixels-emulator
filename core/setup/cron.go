package setup

import (
	"pixels-emulator/core/config"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/healthcheck/ping"
)

func Cron(cfg *config.Config, conStore *protocol.ConnectionStore) *scheduler.Scheduler {

	sc := scheduler.NewCronScheduler()
	ping.ScheduleTask(cfg, &sc, conStore)

	return &sc

}
