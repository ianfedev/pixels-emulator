package handler

import (
	"context"
	"go.uber.org/zap"
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/server"
	navEvent "pixels-emulator/navigator/event"
	"pixels-emulator/navigator/message"
)

// NavigatorSearchHandler handles client search queries.
type NavigatorSearchHandler struct {
	logger *zap.Logger // Logger for packet processing details.
	em     event.Manager
}

// Handle processes the incoming navigation search packet.
func (h *NavigatorSearchHandler) Handle(_ context.Context, raw protocol.Packet, conn protocol.Connection) {

	pck, ok := raw.(*message.NavigatorSearchPacket)
	if !ok {
		h.logger.Error("cannot cast navigator search packet, skipping processing")
		return
	}

	ev := navEvent.NewNavigatorQueryEvent(pck.View, pck.Query, conn, 0, nil)
	h.em.Fire(navEvent.NavigatorQueryEventName, ev)

}

// NewNavigatorSearch creates a new handler instance.
func NewNavigatorSearch() *NavigatorSearchHandler {
	return &NavigatorSearchHandler{
		logger: server.GetServer().Logger(),
		em:     server.GetServer().EventManager(),
	}
}
