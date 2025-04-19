package grant

import (
	"context"
	"errors"
	"go.uber.org/zap"
	authEvent "pixels-emulator/auth/event"
	ok "pixels-emulator/auth/message"
	"pixels-emulator/core/database"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/server"
	"pixels-emulator/user"
	"strconv"
	"time"
)

// UserDatabaseFunc abstracts the function provisioning for testing purposes.
var UserDatabaseFunc = GetUserDatabase

// ProvideAuth performs tasks of authentication granting.
// It should handle the event emitted by the server at low priority to execute login operations.
// This can be cancelled from other sources and prevent the user from login.
func ProvideAuth() func(event event.Event) {
	return func(event event.Event) {
		OnAuthGrantEvent(event)
	}
}

// OnAuthGrantEvent performs tasks of authentication granting.
// It should handle the event emitted by the server at low priority to execute login operations.
// This can be cancelled from other sources and prevent the user from login.
func OnAuthGrantEvent(ev event.Event) {

	connStore := server.GetServer().ConnStore()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var err error
	defer cancel()
	defer func() {
		if err != nil {
			server.GetServer().Logger().Error("error during authentication handle", zap.Error(err))
		}
	}()

	authEv, valid := ev.(*authEvent.AuthGrantEvent)
	if !valid {
		err = errors.New("event proportioned was not authentication")
		return
	}

	id := strconv.Itoa(authEv.UserID())
	con, ex := connStore.GetConnection(id)

	if !ex {
		err = errors.New("connection not found")
		return
	}

	if authEv.CancellableEvent.IsCancelled() {
		err = con.Dispose()
		err = errors.New("connection cancelled by external listener")
		return
	}

	uSvc := UserDatabaseFunc()
	uRes := <-uSvc.Get(ctx, uint(authEv.UserID()))
	if uRes.Error != nil {
		err = uRes.Error
		return
	}

	p := user.Load(uRes.Data, con, server.GetServer().Scheduler(), GetUserDatabase())
	err = server.GetServer().UserStore().Records().Create(ctx, strconv.Itoa(authEv.UserID()), p)
	if err != nil {
		return
	}

	authPack := ok.NewAuthOkPacket()
	con.SendPacket(authPack)

}

func GetUserDatabase() database.DataService[model.User] {
	return &database.ModelService[model.User]{DB: server.GetServer().Database()}
}
