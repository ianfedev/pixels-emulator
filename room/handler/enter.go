package handler

import (
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/server"
	"pixels-emulator/room/message"
)

// RoomEnterHandler manages the connection open access from client to a room.
type RoomEnterHandler struct {
	logger *zap.Logger
}

// Handle process the incoming room search handler.
func (h *RoomEnterHandler) Handle(raw protocol.Packet, conn protocol.Connection) {

	pck, ok := raw.(*message.RoomEnterPacket)
	if !ok {
		h.logger.Error("cannot cast navigator search packet, skipping processing")
		return
	}

	h.logger.Debug("Room enter event", zap.Int32("room", pck.RoomId), zap.Bool("password", pck.Password != ""))

}

// NewRoomEnter creates a new handler instance.
func NewRoomEnter() *RoomEnterHandler {
	return &RoomEnterHandler{
		logger: server.GetServer().Logger(),
	}
}
