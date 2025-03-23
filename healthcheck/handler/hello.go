package handler

import (
	"context"
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/core/server"
	"pixels-emulator/healthcheck/message"
)

// HelloHandler handles incoming Packet from clients.
// It logs the version information of the client that sent the packet.
type HelloHandler struct {
	logger *zap.Logger // logger instance for recording packet processing details.
}

// Handle provides simple debug handling to the console showing the operative client version.
func (h *HelloHandler) Handle(_ context.Context, packet protocol.Packet, _ protocol.Connection) {
	incPacket, ok := packet.(*message.HelloPacket)
	if !ok {
		h.logger.Error("cannot cast ping packet, skipping processing")
		return
	}

	h.logger.Debug("Received hello from client", zap.String("client", incPacket.Version))
}

// NewHello creates a new instance of hello handler.
func NewHello() registry.Handler[protocol.Packet] {
	return &HelloHandler{logger: server.GetServer().Logger()}
}
