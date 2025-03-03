package handler

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/server"
	"pixels-emulator/navigator/message"
)

// NavigatorSearchHandler performs query result
// logic when requesting from client.
type NavigatorSearchHandler struct {
	logger *zap.Logger // logger instance for recording packet processing details.
}

// Handle performs logic to handle the packet.
func (h *NavigatorSearchHandler) Handle(raw protocol.Packet, _ protocol.Connection) {

	pck, ok := raw.(*message.NavigatorSearchPacket)
	if !ok {
		h.logger.Error("cannot cast ping packet, skipping processing")
		return
	}

	h.logger.Debug(pck.Query)
	h.logger.Debug(pck.View)

}

// NewNavigatorSearch creates a new handler instance.
func NewNavigatorSearch() *NavigatorSearchHandler {
	return &NavigatorSearchHandler{
		logger: server.GetServer().Logger(),
	}
}
