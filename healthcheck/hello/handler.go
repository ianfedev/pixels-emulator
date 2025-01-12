package hello

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
)

// PacketHandler handles incoming Packet from clients.
// It logs the version information of the client that sent the packet.
type PacketHandler struct {
	logger *zap.Logger // logger instance for recording packet processing details.
}

// Handle processes the incoming packet and logs client version information.
//
// If the packet is of the expected type (*Packet), the handler will log the version of the client that sent the packet.
// If the packet cannot be cast to the expected type, an error will be logged.
func (h PacketHandler) Handle(packet protocol.Packet, _ *protocol.Connection) {
	incPacket, ok := packet.(*Packet)
	if !ok {
		h.logger.Error("cannot cast ping packet, skipping processing")
		return
	}

	h.logger.Debug("Received hello from client", zap.String("client", incPacket.Version))
}

// NewPacketHandler creates a new PacketHandler with a provided logger.
//
// Returns:
//
//	PacketHandler: A new instance of PacketHandler.
func NewPacketHandler(logger *zap.Logger) registry.Handler[protocol.Packet] {
	return PacketHandler{logger: logger}
}
