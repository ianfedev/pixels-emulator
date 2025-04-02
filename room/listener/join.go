package listener

import (
	"context"
	"go.uber.org/zap"
	"pixels-emulator/core/database"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/server"
	"pixels-emulator/core/util"
	"pixels-emulator/room"
	roomEvent "pixels-emulator/room/event"
	"pixels-emulator/room/message"
	userMsg "pixels-emulator/user/message"
	"strconv"
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

	joinEv, valid := ev.(*roomEvent.RoomJoinEvent)
	if !valid {
		server.GetServer().Logger().Error("event proportioned was not room join, skipping")
		return
	}

	var err error
	defer func() {
		if err != nil {
			server.GetServer().Logger().Error("error during user room join", zap.Error(err))
			room.CloseConnection(joinEv.Conn, message.Default, "")
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rStore := server.GetServer().RoomStore()
	db := server.GetServer().Database()
	rSvc := &database.ModelService[model.Room]{DB: db}
	uSvc := &database.ModelService[model.User]{DB: db}

	if joinEv.IsCancelled() {
		room.CloseConnection(joinEv.Conn, message.Default, "")
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
	rl, err := rStore.Records().GetAll(ctx)
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
		room.CloseConnection(joinEv.Conn, message.Banned, "")
		return
	}

	rs := rRes.Data.State
	accEv := roomEvent.NewRoomAccessGrantEvent(joinEv.Conn, uint(joinEv.Id), 0, make(map[string]string))

	if rel != room.Guest || joinEv.OverrideCheck || rRes.Data.IsPublic || rs == "open" {
		err = server.GetServer().EventManager().Fire(roomEvent.RoomAccessGrantEventName, accEv)
		return
	}

	// INVESTIGATION: Nitro client checks if room is part of a group before
	// prompting password. So, the ideal is not to have password on guild groups.
	// Also, creating a cleaning cronjob will be cool too.
	if rs == "password_protected" {

		u, r := strconv.Itoa(int(uRes.Data.ID)), strconv.Itoa(int(rRes.Data.ID))

		if rStore.Limits().IsFrozen(u, r) {
			room.CloseConnection(joinEv.Conn, message.Default, "exceeded")
			return
		}

		if util.CheckPasswordHash(joinEv.Password, rRes.Data.Password) {
			rStore.Limits().Unfreeze(u + ":" + r)
			err = server.GetServer().EventManager().Fire(roomEvent.RoomAccessGrantEventName, accEv)
			return
		} else {
			if rStore.Limits().RegisterAttempt(u, r) {
				cPck := &message.CloseRoomConnectionPacket{}
				ePck := &userMsg.GenericErrorPacket{Code: userMsg.WrongPasswordCode}
				joinEv.Conn.SendPacket(cPck)
				joinEv.Conn.SendPacket(ePck)
				return
			} else {
				room.CloseConnection(joinEv.Conn, message.Default, "exceeded")
				return
			}
		}

	}

	if false {
		// Doorbelling. This must change from original Arcturus implementation, room must have doorbelling ids and broadcast event
		// This will also have a runnable of 1 minute, runnable must cancel if not in room users. After 1 minute, send message of no one answered.
		// When user with rights join, we should send again after 3 seconds the message.
	}

	room.CloseConnection(joinEv.Conn, message.Default, "")

}
