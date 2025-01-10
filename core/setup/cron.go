package setup

import (
	"pixels-emulator/core/config"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/healthcheck/ping"
)

func Cron(cfg *config.Config) *scheduler.CronScheduler {

	s := scheduler.NewCronScheduler()
	ping.ScheduleTask(*cfg, s)

	return s

}
