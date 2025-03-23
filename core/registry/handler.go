package registry

import (
	"context"
	"fmt"
	"pixels-emulator/core/protocol"
)

// HandlerRegistry defines the interface for a registry that maps packet types to handlers.
type HandlerRegistry interface {
	// Register adds a handler for a specific packet type identified by its uint16 ID.
	Register(packetType uint16, handler Handler[protocol.Packet])

	// Handle processes a packet by invoking the appropriate handler for its type.
	Handle(ctx context.Context, packet protocol.Packet, conn protocol.Connection) error
}

// Handler is a generic interface that represents a registry for a specific packet type.
// T is the type of the packet that the registry is responsible for.
type Handler[T protocol.Packet] interface {
	// Handle processes the given packet of type T with the provided context.
	Handle(ctx context.Context, packet T, conn protocol.Connection)
}

// MapHandlerRegistry is an implementation of HandlerRegistry using a map for storage.
type MapHandlerRegistry struct {
	// handlers maps packet types (uint16) to their corresponding handler (Handler[protocol.Packet]).
	handlers map[uint16]Handler[protocol.Packet]
}

// New creates and returns a new MapHandlerRegistry instance.
func New() HandlerRegistry {
	return &MapHandlerRegistry{
		handlers: make(map[uint16]Handler[protocol.Packet]),
	}
}

// Register registers a handler for a specific packet type identified by its uint16 ID.
func (r *MapHandlerRegistry) Register(packetType uint16, handler Handler[protocol.Packet]) {
	r.handlers[packetType] = handler
}

// Handle processes a given packet, invoking the appropriate handler for its type.
// If no handler is found, an error is returned.
func (r *MapHandlerRegistry) Handle(ctx context.Context, packet protocol.Packet, conn protocol.Connection) error {
	handler, exists := r.handlers[packet.Id()]
	if !exists {
		return fmt.Errorf("no handler registered for packet type: %d", packet.Id())
	}

	handler.Handle(ctx, packet, conn)

	return nil
}
