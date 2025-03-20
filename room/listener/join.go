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

	// Doorbelling. This must change from original Arcturus implementation, room must have doorbelling ids and broadcast event
	// This will also have a runnable of 1 minute, runnable must cancel if not in room users. After 1 minute, send message of no one answered.
	// When user with rights join, we should send again after 3 seconds the message.

	// Password: If password is incorrect, send error (Check generic error composer) or open room
	// Fulfill error in case of none of this are correct :)
}
