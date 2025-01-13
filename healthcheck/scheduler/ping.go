package scheduler

import (
	"pixels-emulator/core/server"
	"pixels-emulator/healthcheck/message"
	"time"
)

// SchedulePing adds to server scheduling a basic repeating task of pinging all the clients inside connection store.
func SchedulePing() {

	// When rate is zero, client is pinging first (Manually ping in nitro)
	// and handled by pong handler. Here we ping first.
	sv := server.GetServer()
	connStore := sv.ConnStore
	rate := sv.Config.Server.PingRate
	if rate == 0 {
		return
	}

	ping := message.ComposePing()

	task := func() {
		connStore.BroadcastPacket(ping)
	}

	sv.Scheduler.ScheduleRepeatingTask(time.Duration(rate)*time.Second, task)

}
