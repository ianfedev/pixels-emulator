package pong

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/healthcheck/ping"
	"time"
)

type PacketHandler struct {
	logger *zap.Logger
}

func (h PacketHandler) Handle(packet protocol.Packet, conn protocol.Connection) {

	// Attempt to cast the incoming packet to a Packet
	_, ok := packet.(*Packet)
	if !ok {
		h.logger.Error("cannot cast ping packet, skipping processing")
		return
	}

	time.AfterFunc(2*time.Second, func() {
		pingPacket := ping.NewPingPacket()
		conn.SendPacket(pingPacket)
	})

}

func NewPacketHandler(logger *zap.Logger) registry.Handler[protocol.Packet] {
	return PacketHandler{
		logger: logger,
	}
}
