package registry

import (
	"fmt"
	"pixels-emulator/core/protocol"
)

// Handler is a generic interface that represents a registry for a specific packet type.
// T is the type of the packet that the registry is responsible for.
type Handler[T protocol.Packet] interface {
	// Handle processes the given packet of type T.
	//
	// Parameters:
	//   packet: The packet of type T to be handled.
	//   conn: The connection instance representing the sender of the packet.
	Handle(packet T, conn protocol.Connection)
}

// Registry holds a collection of registered handlers, where each registry is associated
// with a specific packet type identified by a uint16 ID.
type Registry struct {
	// handlers maps packet types (uint16) to their corresponding registry (Handler[protocol.Packet]).
	handlers map[uint16]Handler[protocol.Packet]
}

// New creates and returns a new instance of Registry with an empty handlers map.
//
// Returns:
//
//	*Registry: A new Registry instance with no registered handlers.
func New() *Registry {
	return &Registry{
		handlers: make(map[uint16]Handler[protocol.Packet]),
	}
}

// Register adds a new registry to the registry, associating it with a specific packet type.
//
// Parameters:
//
//	packetType: A uint16 identifier for the packet type that the registry will process.
//	registry: An instance of Handler[protocol.Packet] responsible for processing packets of the specified type.
func (r *Registry) Register(packetType uint16, handler Handler[protocol.Packet]) {
	r.handlers[packetType] = handler
}

// Handle processes the provided packet by finding the appropriate registry
// based on the packet type. If no registry is found for the packet type,
// it returns an error.
//
// Parameters:
//
//	packet: The packet to be processed. It must implement the protocol.Packet interface.
//	conn: The connection instance representing the sender of the packet.
//
// Returns:
//
//	error: An error if no registry is registered for the packet type,
//	       or if the packet is of the wrong type for the registry.
//
// Example:
//
//	packet := SomePacket{}
//	conn := SomeConnection{}
//	err := registry.Handle(packet, conn)
//	if err != nil {
//	    fmt.Println("Error handling packet:", err)
//	}
func (r *Registry) Handle(packet protocol.Packet, conn protocol.Connection) error {
	// Find the registry for the given packet type
	handler, exists := r.handlers[packet.Id()]
	if !exists {
		return fmt.Errorf("no registry registered for packet type: %d", packet.Id())
	}

	// Use reflection to handle the packet if the registry exists
	switch p := packet.(type) {
	case protocol.Packet:
		handler.Handle(p, conn)
	default:
		return fmt.Errorf("packet is of the wrong type for the registry")
	}

	return nil
}
