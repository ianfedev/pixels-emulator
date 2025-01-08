package registry

import (
	"pixels-emulator/core/protocol"
	"pixels-emulator/server/ping"
)

// SetupProcessors generates all the raw packet processing.
func SetupProcessors() *ProcessorRegistry {

	pReg := NewProcessorRegistry()

	pReg.Register(ping.ClientPacketCode, func(raw protocol.RawPacket, _ protocol.Connection) (protocol.Packet, error) {
		return ping.NewPingPacket(raw)
	})

	return pReg

}

// SetupHandlers generates all the packet handling processing.
func SetupHandlers() *HandlerRegistry {

	hReg := NewHandlerRegistry()

	return hReg
}
