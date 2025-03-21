package grant

import (
	"errors"
	"go.uber.org/zap"
	authEvent "pixels-emulator/auth/event"
	ok "pixels-emulator/auth/message"
	"pixels-emulator/core/event"
	"pixels-emulator/core/server"
	"strconv"
)

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

	var err error
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

	authPack := ok.NewAuthOkPacket()
	con.SendPacket(authPack)

}
