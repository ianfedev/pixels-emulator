package listener

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"pixels-emulator/core/database"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/server"
	"pixels-emulator/room"
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

	rStore := server.GetServer().RoomStore()
	db := server.GetServer().Database()
	rSvc := &database.ModelService[model.Room]{DB: db}
	uSvc := &database.ModelService[model.User]{DB: db}

	if joinEv.IsCancelled() {
		// TODO: Send user to main.
		return
	}

	uid, err := strconv.ParseInt(joinEv.Conn.Identifier(), 10, 32)
	if err != nil {
		return
	}

	uRes := <-uSvc.Get(ctx, uint(uid))
	if uRes.Error != nil {
		err = uRes.Error
		return
	}

	// Removes from every queue the user.
	// TODO: Check again original event if missing behaviour
	rl, err := rStore.GetAll(ctx)
	for _, r := range rl {
		r.Queue.Remove(strconv.Itoa(int(uRes.Data.ID)))
	}

	rRes := <-rSvc.Get(ctx, uint(joinEv.Id))
	if rRes.Error != nil {
		err = rRes.Error
		return
	}

	rel, rErr := room.VerifyUserRoomRelationship(ctx, db, *rRes.Data, *uRes.Data)
	if rErr != nil {
		err = rErr
		return
	}

	if rel == room.Restriction {
		// TODO: Kick user and send to home screen.
		return
	}

	if rel != room.Guest || joinEv.OverrideCheck || rRes.Data.IsPublic || rRes.Data.State == "open" {
		// TODO: Further handling
		return
	}

	// Doorbelling. This must change from original Arcturus implementation, room must have doorbelling ids and broadcast event
	// This will also have a runnable of 1 minute, runnable must cancel if not in room users. After 1 minute, send message of no one answered.
	// When user with rights join, we should send again after 3 seconds the message.

	// Password: If password is incorrect, send error (Check generic error composer) or open room
	// Fulfill error in case of none of this are correct :)

	// TODO: Close connection

}
