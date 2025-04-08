package handler

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/server"
	"pixels-emulator/room"
	"pixels-emulator/room/message"
	"pixels-emulator/user"
)

// FurnitureRequestHandler INVESTIGATION: Why does this exist?
type FurnitureRequestHandler struct {
	logger    *zap.Logger // logger for packet processing details.
	roomStore room.Store  // roomStore is the room list to check user transitioning room.
	userStore user.Store  // userStore is the user store to check user related conn.
}

// Handle processes the incoming navigation search packet.
func (h *FurnitureRequestHandler) Handle(ctx context.Context, raw protocol.Packet, conn protocol.Connection) {

	_, ok := raw.(*message.RoomFurnitureAliasPacket)
	if !ok {
		h.logger.Error("cannot cast navigator search packet, skipping processing")
		return
	}

	var err error
	defer func() {
		if err != nil {
			server.GetServer().Logger().Error("error during user furniture handling", zap.Error(err), zap.String("identifier", conn.Identifier()))
			room.CloseConnection(conn, message.Default, "")
		}
	}()

	users, err := h.userStore.Records().GetAll(ctx)
	if err != nil {
		return
	}

	var p *user.Player
	for _, u := range users {
		if u.Id == conn.Identifier() {
			p = u
			break
		}
	}

	if p == nil {
		err = errors.New("connection player not found")
		return
	}

	r, err := room.GetUserRoom(ctx, h.roomStore, p)
	if err != nil {
		return
	}

	if r == nil {
		err = errors.New("room transitioning not found")
		return
	}

	room.SendHeightMapPackets(conn, int32(r.Data.Configuration.WallHeight), r.Layout())

}

// NewFurnitureRequest creates a new handler instance.
func NewFurnitureRequest() *FurnitureRequestHandler {
	return &FurnitureRequestHandler{
		logger:    server.GetServer().Logger(),
		roomStore: server.GetServer().RoomStore(),
		userStore: server.GetServer().UserStore(),
	}
}
