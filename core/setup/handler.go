package setup

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/healthcheck/pong"

	"pixels-emulator/healthcheck/hello"
)

// Processors generates all the raw packet processing.
func Processors() *registry.ProcessorRegistry {

	pReg := registry.NewProcessorRegistry()

	pReg.Register(hello.PacketCode, func(raw protocol.RawPacket, _ protocol.Connection) (protocol.Packet, error) {
		return hello.NewPacket(raw)
	})
	pReg.Register(pong.PacketCode, func(raw protocol.RawPacket, _ protocol.Connection) (protocol.Packet, error) {
		return pong.NewPacket(raw), nil
	})

	return pReg

}

// Handlers generates all the packet handling processing.
func Handlers(logger *zap.Logger) *registry.Registry {

	hReg := registry.New()

	hReg.Register(hello.PacketCode, hello.NewPacketHandler(logger))
	hReg.Register(pong.PacketCode, pong.NewPacketHandler(logger))

	return hReg
}
