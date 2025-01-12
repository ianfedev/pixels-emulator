package setup

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"pixels-emulator/auth/sso"
	"pixels-emulator/core/config"
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/healthcheck/pong"

	"pixels-emulator/healthcheck/hello"
)

// Processors generates all the raw packet processing.
func Processors() *registry.ProcessorRegistry {

	pReg := registry.NewProcessorRegistry()

	pReg.Register(hello.PacketCode, func(raw protocol.RawPacket, _ *protocol.Connection) (protocol.Packet, error) {
		return hello.NewPacket(raw)
	})
	pReg.Register(pong.PacketCode, func(raw protocol.RawPacket, _ *protocol.Connection) (protocol.Packet, error) {
		return pong.NewPacket(raw), nil
	})

	pReg.Register(sso.PacketCode, func(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error) {
		return sso.NewPacket(raw)
	})

	return pReg

}

// Handlers generates all the packet handling processing.
func Handlers(
	logger *zap.Logger,
	cfg *config.Config,
	scheduler *scheduler.Scheduler,
	db *gorm.DB,
	cs *protocol.ConnectionStore,
	em *event.Manager,
) *registry.Registry {

	hReg := registry.New()

	hReg.Register(hello.PacketCode, hello.NewPacketHandler(logger))
	hReg.Register(pong.PacketCode, pong.NewPacketHandler(logger, cfg.Server.PingRate, scheduler))

	hReg.Register(sso.PacketCode, sso.NewSSOTicketHandler(logger, db, cs, cfg, em))

	return hReg
}
