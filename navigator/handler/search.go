package handler

import (
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
func (h *NavigatorSearchHandler) Handle(raw protocol.Packet, _ protocol.Connection) {

	pck, ok := raw.(*message.NavigatorSearchPacket)
	if !ok {
		h.logger.Error("cannot cast navigator search packet, skipping processing")
		return
	}

	queryParams := map[string]string{"query": pck.Query}
	ev := navEvent.NewNavigatorQueryEvent(pck.View, queryParams, 0, nil)
	err := h.em.Fire(navEvent.NavigatorQueryEventName, ev)

	if err != nil {
		h.logger.Error("error while broadcasting navigation event", zap.Error(err))
	}

}

// NewNavigatorSearch creates a new handler instance.
func NewNavigatorSearch() *NavigatorSearchHandler {
	return &NavigatorSearchHandler{
		logger: server.GetServer().Logger(),
		em:     server.GetServer().EventManager(),
	}
}
