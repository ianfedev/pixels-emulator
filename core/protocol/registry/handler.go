package registry

import (
	"fmt"
	"pixels-emulator/core/protocol"
)

// Handler is a generic interface that represents a handler for a specific packet type.
// T is the type of the packet that the handler is responsible for.
type Handler[T protocol.Packet] interface {
	// Handle processes the given packet of type T.
	// The implementation of this method will define how the packet is handled.
	Handle(packet T)
}

// HandlerRegistry holds a collection of registered handlers, where each handler is associated
// with a specific packet type identified by a string.
type HandlerRegistry struct {
	// handlers maps packet types (string) to their corresponding handler (Handler[any]).
	handlers map[uint16]Handler[protocol.Packet]
}

// NewHandlerRegistry creates and returns a new instance of HandlerRegistry with an empty handlers map.
func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		handlers: make(map[uint16]Handler[protocol.Packet]),
	}
}

// Register adds a new handler to the registry, associating it with a specific packet type.
// packetType is the identifier for the packet type the handler will process.
// handler is the handler instance responsible for handling packets of the specified type.
func (r *HandlerRegistry) Register(packetType uint16, handler Handler[protocol.Packet]) {
	r.handlers[packetType] = handler
}

// Handle processes the provided packet by finding the appropriate handler
// based on the packet type. If no handler is found for the packet type,
// it returns an error. This method is intended to be called when processing packets
// and dispatching them to their respective handlers.
func (r *HandlerRegistry) Handle(packet protocol.Packet) error {
	// Find the handler for the given packet type
	handler, exists := r.handlers[packet.GetId()]
	if !exists {
		return fmt.Errorf("no handler registered for packet type: %d", packet.GetId())
	}

	// Use reflection to handle the packet if the handler exists
	switch p := packet.(type) {
	case protocol.Packet:
		handler.Handle(p)
	default:
		return fmt.Errorf("packet is of the wrong type for the handler")
	}

	return nil
}
