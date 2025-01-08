package healthcheck

import (
	"go.uber.org/zap"
	"pixels-emulator/core/handler"
	"pixels-emulator/core/protocol"
)

// IncomingPingHandler receives the ping packet and handles a simple pong for
// the client which received it.
type IncomingPingHandler struct {
	logger *zap.Logger
}

func (h IncomingPingHandler) Handle(packet protocol.Packet, conn protocol.Connection) {

	incPacket, ok := packet.(*PingIncomingPacket)
	if !ok {
		h.logger.Error("cannot cast ping packet, skipping processing")
		return
	}

	h.logger.Debug("Received hello from client", zap.String("client", incPacket.Version))

}

func NewIncomingPingHandler(logger *zap.Logger) handler.Handler[protocol.Packet] {
	return IncomingPingHandler{logger: logger}
}
