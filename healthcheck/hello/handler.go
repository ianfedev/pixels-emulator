package hello

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
)

// PacketHandler handles incoming ClientHello packets from clients.
// It logs the version information of the client that sent the packet.
type PacketHandler struct {
	logger *zap.Logger // Logger instance for recording packet processing details.
}

// Handle processes a received Packet. It verifies the packet type and
// logs the client's version information. If the packet cannot be cast to a
// Packet, an error is logged and the packet is skipped.
//
// Parameters:
//
//	packet: The packet to process. It should be of type *pPacket.Packet.
//	_: The connection associated with the packet (not used in this handler).
//
// Behavior:
//   - Logs an error if the packet cannot be cast to a Packet.
//   - Logs debug information containing the client version if the packet is valid.
func (h PacketHandler) Handle(packet protocol.Packet, _ protocol.Connection) {
	// Attempt to cast the incoming packet to a Packet
	incPacket, ok := packet.(*Packet)
	if !ok {
		h.logger.Error("cannot cast ping packet, skipping processing")
		return
	}

	// Log the version of the client sending the packet
	h.logger.Debug("Received hello from client", zap.String("client", incPacket.Version))
}

// NewPacketHandler creates a new instance of PacketHandler and
// returns it as a generic packet handler.
//
// Parameters:
//
//	logger: A *zap.Logger instance for logging packet handling information.
//
// Returns:
//
//	registry.Handler[protocol.Packet]: A handler that processes Packet packets.
//
// Example:
//
//	logger := zap.NewExample()
//	handler := NewPacketHandler(logger)
//	registry := registry.New()
//	registry.Register(PacketCode, handler)
func NewPacketHandler(logger *zap.Logger) registry.Handler[protocol.Packet] {
	return PacketHandler{logger: logger}
}
