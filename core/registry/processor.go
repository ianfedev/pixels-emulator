package registry

import (
	"errors"
	"pixels-emulator/core/protocol"
	"strconv"
)

// ProcessorRegistry defines the interface for a registry that maps packet headers to processors.
type ProcessorRegistry interface {
	// Register adds a processor function for the specified uint16 header code.
	Register(code uint16, handler func(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error))

	// Handle processes a raw packet using the corresponding processor for its header.
	Handle(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error)
}

// MapProcessorRegistry is an implementation of ProcessorRegistry using a map for storage.
type MapProcessorRegistry struct {
	// processor maps uint16 codes to functions responsible for processing raw packets.
	processor map[uint16]func(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error)
}

// NewProcessor creates and returns a new MapProcessorRegistry instance.
func NewProcessor() ProcessorRegistry {
	return &MapProcessorRegistry{
		processor: make(map[uint16]func(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error)),
	}
}

// Register adds a handler function to the registry for a specific packet type identified by its uint16 code.
func (mpr *MapProcessorRegistry) Register(code uint16, handler func(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error)) {
	mpr.processor[code] = handler
}

// Handle processes a raw packet by identifying the correct handler using the packet's header code.
// If no handler is found, an error is returned.
func (mpr *MapProcessorRegistry) Handle(raw protocol.RawPacket, conn *protocol.Connection) (protocol.Packet, error) {
	handler, exists := mpr.processor[raw.GetHeader()]
	if !exists {
		return nil, errors.New("unprocessable packet received " + strconv.Itoa(int(raw.GetHeader())))
	}
	return handler(raw, conn)
}
