package registry

import (
	"errors"
	"pixels-emulator/core/protocol"
)

type ProcessorRegistry struct {
	processor map[uint16]func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error)
}

// NewProcessorRegistry creates an empty handler registry.
func NewProcessorRegistry() *ProcessorRegistry {
	return &ProcessorRegistry{
		processor: make(map[uint16]func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error)),
	}
}

// Register adds a new processor to the registry.
func (pr *ProcessorRegistry) Register(code uint16, handler func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error)) {
	pr.processor[code] = handler
}

// Handle finds the processor corresponding to the packet and registers it.
func (pr *ProcessorRegistry) Handle(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error) {
	handler, exists := pr.processor[raw.GetHeader()]
	if !exists {
		return nil, errors.New("unprocessable packet received")
	}
	return handler(raw, conn)
}
