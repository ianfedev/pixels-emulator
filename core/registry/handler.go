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
	Handle(packet T, conn *protocol.Connection)
}

// Registry holds a collection of registered handlers, where each registry is associated
// with a specific packet type identified by a uint16 ID.
type Registry struct {
	// handlers maps packet types (uint16) to their corresponding registry (Handler[protocol.Packet]).
	handlers map[uint16]Handler[protocol.Packet]
}

// New creates and returns a new Registry instance.
func New() *Registry {
	return &Registry{
		handlers: make(map[uint16]Handler[protocol.Packet]),
	}
}

// Register registers a handler for a specific packet type identified by its uint16 ID.
func (r *Registry) Register(packetType uint16, handler Handler[protocol.Packet]) {
	r.handlers[packetType] = handler
}

// Handle processes a given packet, invoking the appropriate handler for its type.
// If no handler is found, an error is returned.
func (r *Registry) Handle(packet protocol.Packet, conn *protocol.Connection) error {
	handler, exists := r.handlers[packet.Id()]
	if !exists {
		return fmt.Errorf("no registry registered for packet type: %d", packet.Id())
	}

	switch p := packet.(type) {
	case protocol.Packet:
		handler.Handle(p, conn)
	default:
		return fmt.Errorf("packet is of the wrong type for the registry")
	}

	return nil
}
