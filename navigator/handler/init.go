package handler

import (
	"context"
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/server"
	"pixels-emulator/navigator/message"
)

// NavigatorInitHandler handles incoming Packet from clients.
// Replies sending a set of packets
type NavigatorInitHandler struct {
	logger *zap.Logger // logger instance for recording packet processing details.
}

// Handle performs logic to handle the packet.
func (h *NavigatorInitHandler) Handle(_ context.Context, packet protocol.Packet, conn protocol.Connection) {

	_, ok := packet.(*message.NavigatorInitPacket)
	if !ok {
		h.logger.Error("cannot cast ping packet, skipping processing")
		return
	}

	h.logger.Debug("Navigator fired by user", zap.String("identifier", conn.Identifier()))

	// INVESTIGATION: Navigation packets useless in Nitro Client.
	// Navigation Settings (518): Stores position of window.
	// Navigator Metadata (3052)
	// Navigator Lifted (3104)
	// Navigator Collapsed Categories (1543)
	// Navigator Searches Event (3984)
	// In Arcturus, there is a weird and maybe buggy way to query both Event and User categories.
	// However, when firing this event, it actually sends all the unused packets with only the event categories
	// It should fire both or just one? I don't know.
	// Categories are provided via UserInfoEvent and expected to be reach only by a request...
	// For now this will remain empty while I figure out this usage.

	navCtx := []string{"official_view", "hotel_view", "roomads_view", "myworld_view"}
	conn.SendPacket(message.NewNavigatorMetaDataPacket(navCtx...))

}

// NewNavigatorInit creates a new handler instance.
func NewNavigatorInit() *NavigatorInitHandler {
	return &NavigatorInitHandler{
		logger: server.GetServer().Logger(),
	}
}
