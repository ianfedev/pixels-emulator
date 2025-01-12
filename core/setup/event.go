package setup

import (
	"pixels-emulator/auth/grant"
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
)

func Event(ev *event.Manager, conStore *protocol.ConnectionStore) {
	ev.AddListener(grant.AuthEventName, grant.OnAuthGranted(conStore), 10)
}
