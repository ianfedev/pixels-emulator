package listener

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"pixels-emulator/core/event"
	"pixels-emulator/core/server"
	roomEvent "pixels-emulator/room/event"
	"time"
)

// ProvideUserJoin perform user validation when joining a room.
// It must handle all the corresponding logic before a user is connected
// to a room.
func ProvideUserJoin() func(event event.Event) {
	return func(event event.Event) {
		OnUserRoomJoin(event)
	}
}

func OnUserRoomJoin(ev event.Event) {

	var err error
	defer func() {
		if err != nil {
			server.GetServer().Logger().Error("error during user room join", zap.Error(err))
			// TODO: Send user to main.
		}
	}()

	joinEv, valid := ev.(*roomEvent.RoomJoinEvent)
	if !valid {
		err = errors.New("event proportioned was not room join")
		return
	}

	rs := server.GetServer().RoomStore()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	room, err := rs.Read(ctx, string(joinEv.Id))

	// Check if event is cancelled and send to main

	//// Check if is banned and doesnt have permissions (This must differ from normal close)

	// Queue removal: Must find all rooms and remove from queue.

	// Let event proceed
	// - If event has overriding
	// - If user is owner
	// - If user has permissions
	// - If user has rights
	// - Guild rights pending on guild system
	server.GetServer().Logger().Debug("Room", zap.Any("room", room))
	server.GetServer().Logger().Error("Room", zap.Any("err", err))

}
