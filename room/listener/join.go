package listener

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"pixels-emulator/core/database"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/server"
	"pixels-emulator/role"
	roomEvent "pixels-emulator/room/event"
	"strconv"
	"time"
)

const AccessRoomPermissions = "pixels.room.access"

// ProvideUserJoin perform user validation when joining a room.
// It must handle all the corresponding logic before a user is connected
// to a room.
func ProvideUserJoin() func(event event.Event) {
	return func(event event.Event) {
		OnUserRoomJoin(event)
	}
}

// OnUserRoomJoin validates if the user is granted to access.
// This performs the current checks:
//
// - If event has override ability
// - If user is owner of the room
// - If user has permissions given from permissions system
// - If user has user-given permissions
// - If user has guild-related permissions to access (E.g: Is the guild room) and is member.
//
// This is also proceeded by doorbell and password behaviours, which must be handled in
// their corresponding event listeners but primarily fired here.
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = server.GetServer().RoomStore()
	rStore := &database.ModelService[model.Room]{}
	uStore := &database.ModelService[model.User]{}

	// TODO: Remove user from all the rooms.

	if joinEv.IsCancelled() {
		// TODO: Send user to main.
		return
	}

	uid, err := strconv.ParseInt(joinEv.Conn.Identifier(), 10, 32)
	if err != nil {
		return
	}

	uRes := <-uStore.Get(ctx, uint(uid))
	if uRes.Error != nil {
		err = uRes.Error
		return
	}

	rRes := <-rStore.Get(ctx, uint(joinEv.Id))
	if rRes.Error != nil {
		err = rRes.Error
		return
	}

	grantAccess := joinEv.OverrideCheck

	if !grantAccess {
		grantAccess = role.HasPermission(*uRes.Data, AccessRoomPermissions)
	}

	if !grantAccess {
		grantAccess = rRes.Data.OwnerID == uRes.Data.ID
	}

	if !grantAccess {

		q := map[string]interface{}{"room_id": rRes.Data.ID, "user_id": uRes.Data.ID}
		pStore := &database.ModelService[model.RoomPermission]{}
		pRes := <-pStore.FindByQuery(ctx, q)

		if pRes.Error == nil {
			err = pRes.Error
		}

		grantAccess = len(pRes.Data) > 0

	}

	// Check if event is cancelled and send to main

	//// Check if is banned and doesnt have permissions (This must differ from normal close)

	// Queue removal: Must find all rooms and remove from queue.

	// Let event proceed
	// - If event has overriding DONE
	// - If user is owner DONE
	// - If user has permissions DONE
	// - If user has rights DONE
	// - Guild rights pending on guild system
	server.GetServer().Logger().Error("Room", zap.Any("err", err))

	// Doorbelling. This must change from original Arcturus implementation, room must have doorbelling ids and broadcast event
	// This will also have a runnable of 1 minute, runnable must cancel if not in room users. After 1 minute, send message of no one answered.
	// When user with rights join, we should send again after 3 seconds the message.

	// Password: If password is incorrect, send error (Check generic error composer) or open room
	// Fulfill error in case of none of this are correct :)
}
