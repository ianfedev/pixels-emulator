package handler

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/server"
	"pixels-emulator/room/message"
)

// GetGuestRoomHandler handles non-owned room join attempt.
type GetGuestRoomHandler struct {
	logger *zap.Logger // Logger for packet processing details.
}

// Handle processes the incoming navigation search packet.
func (h *GetGuestRoomHandler) Handle(raw protocol.Packet, _ protocol.Connection) {

	pck, ok := raw.(*message.GetGuestRoomPacket)
	if !ok {
		h.logger.Error("cannot cast navigator search packet, skipping processing")
		return
	}

	h.logger.Debug("Guest room event", zap.Int32("room", pck.RoomId), zap.Bool("enter", pck.Enter), zap.Bool("forward", pck.Forward))

}

// NewNavigatorSearch creates a new handler instance.
func NewNavigatorSearch() *GetGuestRoomHandler {
	return &GetGuestRoomHandler{
		logger: server.GetServer().Logger(),
	}
}
