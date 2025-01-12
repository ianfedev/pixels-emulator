package grant

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	ok "pixels-emulator/auth/ok"
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
	"strconv"
	"time"
)

// OnAuthGranted performs tasks of authentication granting.
func OnAuthGranted(conStore *protocol.ConnectionStore) func(event event.Event) {
	return func(ev event.Event) {

		var err error
		defer func() {
			if err != nil {
				zap.L().Error("error during authentication handle", zap.Error(err))
			}
		}()

		authEv, valid := ev.(*AuthEvent)
		if !valid {
			err = errors.New("event proportioned was not authentication")
			return
		}

		id := strconv.Itoa(authEv.UserID())
		con, ex := conStore.GetConnection(id)

		if !ex {
			err = errors.New("connection not found")
			return
		}

		if authEv.CancellableEvent.IsCancelled() {
			err = errors.New("connection cancelled by external listener")
			return
		}

		authPack := ok.NewAuthOkPacket()
		(*con).SendPacket(authPack)
		time.AfterFunc(5*time.Second, func() {
			err := (*con).Dispose()
			if err != nil {
				fmt.Println(err)
				return
			}
		})

	}
}
