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
	processor map[uint16]func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error)
}

// NewProcessorRegistry creates and returns an empty ProcessorRegistry instance.
//
// Returns:
//
//	*ProcessorRegistry: A new registry with no registered processors.
func NewProcessorRegistry() *ProcessorRegistry {
	return &ProcessorRegistry{
		processor: make(map[uint16]func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error)),
	}
}

// Register adds a new processor to the registry, associating it with a specific packet code.
//
// Parameters:
//
//	code: A uint16 identifier representing the packet type.
//	registry: A function that processes raw packets into structured packets. It takes a
//	         raw packet and a connection as inputs and returns a structured packet or an error.
func (pr *ProcessorRegistry) Register(code uint16, handler func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error)) {
	pr.processor[code] = handler
}

// Handle finds and invokes the processor corresponding to the raw packet's header.
//
// Parameters:
//
//	raw: The raw packet to be processed. It must implement the protocol.RawPacket interface
//	     and provide access to its header via the GetHeader method.
//	conn: The connection instance representing the sender of the raw packet.
//
// Returns:
//
//	protocol.Packet: The structured packet produced by the processor.
//	error: An error if no processor is registered for the packet's header or if the
//	       processor fails to process the raw packet.
//
// Example:
//
//	registry := NewProcessorRegistry()
//	registry.Register(0x01, someProcessorFunc)
//	rawPacket := SomeRawPacket{}
//	conn := SomeConnection{}
//	packet, err := registry.Handle(rawPacket, conn)
//	if err != nil {
//	    fmt.Println("Error processing packet:", err)
//	}
func (pr *ProcessorRegistry) Handle(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error) {
	handler, exists := pr.processor[raw.GetHeader()]
	if !exists {
		return nil, errors.New("unprocessable packet received")
	}
	return handler(raw, conn)
}
