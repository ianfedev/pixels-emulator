package setup

import (
	"go.uber.org/zap"
	"pixels-emulator/core/handler"
	"pixels-emulator/core/protocol"
	"pixels-emulator/healthcheck"
)

// Processors generates all the raw packet processing.
func Processors() *handler.ProcessorRegistry {

	pReg := handler.NewProcessorRegistry()

	pReg.Register(healthcheck.IncomingPingCode, func(raw protocol.RawPacket, _ protocol.Connection) (protocol.Packet, error) {
		return healthcheck.NewPingIncoming(raw)
	})

	return pReg

}

// Handlers generates all the packet handling processing.
func Handlers(logger *zap.Logger) *handler.Registry {

	hReg := handler.New()

	hReg.Register(healthcheck.IncomingPingCode, healthcheck.NewIncomingPingHandler(logger))

	return hReg
}
