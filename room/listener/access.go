package listener

import (
	"context"
	"go.uber.org/zap"
	"pixels-emulator/core/database"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/server"
	"pixels-emulator/room"
	roomEvent "pixels-emulator/room/event"
	"pixels-emulator/room/message"
	"strconv"
	"strings"
	"time"
)

// ProvideUserAccess handles the room loading event when a
// player has been granted to access a room, including room
// loading and any involved action like queue awaiting.
func ProvideUserAccess() func(event event.Event) {
	return func(event event.Event) {
		OnUserRoomJoin(event)
	}
}

// OnUserRoomAccess handles the room loading event when a
// // player has been granted to access a room, including room
// // loading and any involved action like queue awaiting.
func OnUserRoomAccess(ev event.Event) {

	accEv, valid := ev.(*roomEvent.RoomAccessGrantEvent)
	if !valid {
		server.GetServer().Logger().Error("event proportioned was not room access, skipping")
		return
	}

	var err error
	defer func() {
		if err != nil {
			server.GetServer().Logger().Error("error during user room access", zap.Error(err))
			room.CloseConnection(accEv.Conn, message.Default, "")
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rStore := server.GetServer().RoomStore()
	db := server.GetServer().Database()
	rSvc := &database.ModelService[model.Room]{DB: db}
	rRes := <-rSvc.Get(ctx, accEv.Room)
	if rRes.Error != nil {
		err = rRes.Error
		return
	}

	r, lErr := rStore.Records().Read(ctx, strconv.Itoa(int(accEv.Room)))
	if lErr != nil {

		if !strings.Contains(err.Error(), "key not found") {
			err = lErr
			return
		}

		r = room.Load(rRes.Data)
	}

}
