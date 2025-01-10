package ping

import (
	"pixels-emulator/core/config"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/scheduler"
	"time"
)

func ScheduleTask(
	cfg *config.Config,
	sc *scheduler.Scheduler,
	conStore *protocol.ConnectionStore) {

	// When rate is zero, client is pinging first (Manually ping in nitro)
	// and handled by pong handler. Here we ping first.
	if cfg.Server.PingRate == 0 {
		return
	}

	ping := NewPingPacket()

	task := func() {
		conStore.BroadcastPacket(ping)
	}

	(*sc).ScheduleRepeatingTask(time.Duration(cfg.Server.PingRate)*time.Second, task)
}
