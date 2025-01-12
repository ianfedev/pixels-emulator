package registry

import (
	"errors"
	"pixels-emulator/core/protocol"
)

// ProcessorRegistry is a registry that maps packet headers (uint16 codes)
// to their respective processors. Each processor is a function that
// converts a raw packet into a structured packet.
//
// This registry allows for dynamic processing of different packet types
// based on their header identifiers.
type ProcessorRegistry struct {
	// processor maps uint16 codes to functions responsible for processing raw packets.
	processor map[uint16]func(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error)
}

// NewProcessorRegistry creates and returns a new ProcessorRegistry instance.
func NewProcessorRegistry() *ProcessorRegistry {
	return &ProcessorRegistry{
		processor: make(map[uint16]func(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error)),
	}
}

// Register adds a handler function to the registry for a specific packet type identified by its uint16 code.
func (pr *ProcessorRegistry) Register(code uint16, handler func(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error)) {
	pr.processor[code] = handler
}

// Handle processes a raw packet by identifying the correct handler using the packet's header code.
// If no handler is found, an error is returned.
func (pr *ProcessorRegistry) Handle(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error) {
	handler, exists := pr.processor[raw.GetHeader()]
	if !exists {
		return nil, errors.New("unprocessable packet received")
	}
	return handler(raw, conn)
}
