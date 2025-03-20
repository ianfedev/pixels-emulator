package listener

import (
	"errors"
	"go.uber.org/zap"
	"pixels-emulator/core/event"
	"pixels-emulator/core/server"
	roomEvent "pixels-emulator/room/event"
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
		}
	}()

	_, valid := ev.(*roomEvent.RoomJoinEvent)
	if !valid {
		err = errors.New("event proportioned was not authentication")
		return
	}

	// Check if event is cancelled and send to main

	//// Check if is banned and doesnt have permissions

	// Queue removal

}
