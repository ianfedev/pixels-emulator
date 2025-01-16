package ephemeral

import (
	"pixels-emulator/auth/event"
	"pixels-emulator/auth/grant"
	"pixels-emulator/core/server"
)

func Event() {
	em := server.GetServer().EventManager()
	em.AddListener(event.AuthGrantEventName, grant.ProvideAuth(), 10)
}
