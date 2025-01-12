package pong

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/healthcheck/ping"
	"time"
)

type PacketHandler struct {
	logger    *zap.Logger
	rate      uint16
	scheduler *scheduler.Scheduler
}

func (h PacketHandler) Handle(packet protocol.Packet, conn *protocol.Connection) {

	_, ok := packet.(*Packet)
	if !ok {
		h.logger.Error("cannot cast ping packet, skipping processing")
		return
	}

	if h.rate > 0 {
		// When rate is zero, client is pinging first (Manually ping in nitro).
		// otherwise, Pixels will be the one pinging first at cronjob.
		return
	}

	(*h.scheduler).ScheduleTaskLater(2*time.Second, func() {
		pingPacket := ping.NewPingPacket()
		(*conn).SendPacket(pingPacket)
	})

}

func NewPacketHandler(logger *zap.Logger, rate uint16, scheduler *scheduler.Scheduler) registry.Handler[protocol.Packet] {
	return PacketHandler{
		logger:    logger,
		rate:      rate,
		scheduler: scheduler,
	}
}
