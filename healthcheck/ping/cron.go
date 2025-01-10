package ping

import (
	"pixels-emulator/core/config"
	"pixels-emulator/core/scheduler"
	"time"
)

func ScheduleTask(
	cfg config.Config,
	s scheduler.Scheduler) {

	if cfg.Server.PingRate == 0 {
		return
	}

	task := func() {
		// Generate ping to all active connections
	}

	s.ScheduleRepeatingTask(time.Duration(cfg.Server.PingRate)*time.Second, task)
}
