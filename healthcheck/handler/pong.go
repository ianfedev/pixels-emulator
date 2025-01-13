package handler

import (
	"go.uber.org/zap"
	"pixels-emulator/core/config"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/core/server"
	"pixels-emulator/healthcheck/message"
	"time"
)

// PongHandler handles incoming pong packet from the client.
// It should be used when pinging is in charge of client.
type PongHandler struct {
	logger    *zap.Logger
	scheduler scheduler.Scheduler
	cfg       *config.Config
}

// Handle executes the packet processing, responding with a ping 2 seconds after sent if no ping rate is set, assuming client will handle it.
func (h PongHandler) Handle(packet protocol.Packet, conn *protocol.Connection) {

	rate := h.cfg.Server.PingRate

	_, ok := packet.(*message.PongPacket)
	if !ok {
		h.logger.Error("cannot cast ping packet, skipping processing")
		return
	}

	if rate > 0 {
		// When rate is zero, client is pinging first (Manually ping in nitro).
		// otherwise, Pixels will be the one pinging first at cronjob.
		return
	}

	h.scheduler.ScheduleTaskLater(2*time.Second, func() {
		pingPacket := message.ComposePing()
		(*conn).SendPacket(pingPacket)
	})

}

// NewPong creates a new instance of pong handler.
func NewPong() registry.Handler[protocol.Packet] {
	sv := server.GetServer()
	return PongHandler{
		logger:    sv.Logger,
		cfg:       sv.Config,
		scheduler: sv.Scheduler,
	}
}
