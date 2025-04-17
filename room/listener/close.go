package listener

import (
	"context"
	"go.uber.org/zap"
	"pixels-emulator/core/event"
	"pixels-emulator/core/server"
	roomEvent "pixels-emulator/room/event"
)

// ProvideRoomClose encapsulates the event.
func ProvideRoomClose() func(event event.Event) {
	return func(event event.Event) {
		OnRoomConnectionClose(event)
	}
}

// OnRoomConnectionClose handles the connection
func OnRoomConnectionClose(ev event.Event) {

	cEv, valid := ev.(*roomEvent.RoomCloseConnectionEvent)
	if !valid {
		server.GetServer().Logger().Error("event proportioned was not room access, skipping")
		return
	}

	var err error
	defer func() {
		if err != nil {
			server.GetServer().Logger().Error("error during user room access", zap.Error(err))
			if connErr := cEv.Connection.Dispose(); connErr != nil {
				server.GetServer().Logger().Error("Error disposing connection", zap.Error(connErr))
			}
		}
	}()

	rooms, err := server.GetServer().RoomStore().Records().GetAll(context.Background())
	server.GetServer().Logger().Debug("Cleared connection", zap.String("identifier", cEv.Connection.Identifier()))
	if err != nil {
		return
	}

	for _, room := range rooms {
		room.Clear(cEv.Connection.Identifier())
	}

}
