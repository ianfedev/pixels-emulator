package grant

import (
	"errors"
	"go.uber.org/zap"
	event2 "pixels-emulator/auth/event"
	ok "pixels-emulator/auth/message"
	"pixels-emulator/core/event"
	"pixels-emulator/core/server"
	"strconv"
)

// OnAuthGranted performs tasks of authentication granting.
// It should handle the event emitted by the server at low priority to execute login operations.
// This can be cancelled from other sources and prevent the user from login.
func OnAuthGranted() func(event event.Event) {

	connStore := server.GetServer().ConnStore()

	return func(ev event.Event) {

		var err error
		defer func() {
			if err != nil {
				zap.L().Error("error during authentication handle", zap.Error(err))
			}
		}()

		authEv, valid := ev.(*event2.AuthGrantEvent)
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
			_ = con.Dispose()
			err = errors.New("connection cancelled by external listener")
			return
		}

		authPack := ok.NewAuthOkPacket()
		con.SendPacket(authPack)

	}

}
