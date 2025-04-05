package ephemeral

import (
	authEvent "pixels-emulator/auth/event"
	authListener "pixels-emulator/auth/grant"
	"pixels-emulator/core/server"
	navEvent "pixels-emulator/navigator/event"
	navListener "pixels-emulator/navigator/listener"
	roomEvent "pixels-emulator/room/event"
	roomListener "pixels-emulator/room/listener"
)

func Event() {
	em := server.GetServer().EventManager()
	em.AddListener(authEvent.AuthGrantEventName, authListener.ProvideAuth(), 10)
	em.AddListener(navEvent.NavigatorQueryEventName, navListener.ProvideSearch(), 10)
	em.AddListener(roomEvent.RoomJoinEventName, roomListener.ProvideUserJoin(), 10)
	em.AddListener(roomEvent.RoomLoadRequestEventName, roomListener.ProvideRoomLoadRequest(), 10)
}
