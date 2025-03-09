package ephemeral

import (
	authEvent "pixels-emulator/auth/event"
	authListener "pixels-emulator/auth/grant"
	"pixels-emulator/core/server"
	navEvent "pixels-emulator/navigator/event"
	navListener "pixels-emulator/navigator/listener"
)

func Event() {
	em := server.GetServer().EventManager()
	em.AddListener(authEvent.AuthGrantEventName, authListener.ProvideAuth(), 10)
	em.AddListener(navEvent.NavigatorQueryEventName, navListener.ProvideSearch(), 10)
}
