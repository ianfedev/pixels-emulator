package handler

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/server"
	roomEvent "pixels-emulator/room/event"
	"pixels-emulator/room/message"
)

// RoomEnterHandler manages the connection open access from client to a room.
type RoomEnterHandler struct {
	logger *zap.Logger
	em     event.Manager
}

// Handle process the incoming room search handler.
func (h *RoomEnterHandler) Handle(_ context.Context, raw protocol.Packet, conn protocol.Connection) {

	var err error
	defer func() {
		if err != nil {
			h.logger.Error("error during room join", zap.Error(err))
			if connErr := conn.Dispose(); connErr != nil {
				h.logger.Error("Error disposing connection", zap.Error(connErr))
			}
		}
	}()

	pck, ok := raw.(*message.RoomEnterPacket)
	if !ok {
		err = errors.New("cannot cast navigator search packet, skipping processing")
		return
	}

	h.logger.Debug("Room enter event", zap.Int32("room", pck.RoomId), zap.Bool("password", pck.Password != ""))
	e := roomEvent.NewRoomJoinEvent(conn, pck.RoomId, pck.Password, 0, make(map[string]string))

	h.em.Fire(roomEvent.RoomJoinEventName, e)

}

// NewRoomEnter creates a new handler instance.
func NewRoomEnter() *RoomEnterHandler {
	return &RoomEnterHandler{
		logger: server.GetServer().Logger(),
		em:     server.GetServer().EventManager(),
	}
}
