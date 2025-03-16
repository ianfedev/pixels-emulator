package ephemeral

import (
	authHandler "pixels-emulator/auth/handler"
	authMsg "pixels-emulator/auth/message"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/server"
	healthHandler "pixels-emulator/healthcheck/handler"
	healthMsg "pixels-emulator/healthcheck/message"
	navigatorHandler "pixels-emulator/navigator/handler"
	navigatorMsg "pixels-emulator/navigator/message"
	roomHandler "pixels-emulator/room/handler"
	roomMsg "pixels-emulator/room/message"
	guestRoomMsg "pixels-emulator/room/message/guest"
)

// Processors generates all the raw packet processing.
func Processors() {

	pReg := server.GetServer().PacketProcessors()

	pReg.Register(healthMsg.HelloCode, func(raw protocol.RawPacket, _ protocol.Connection) (protocol.Packet, error) {
		return healthMsg.ComposeHello(raw)
	})
	pReg.Register(healthMsg.PongCode, func(raw protocol.RawPacket, _ protocol.Connection) (protocol.Packet, error) {
		return healthMsg.ComposePong(raw), nil
	})

	pReg.Register(authMsg.AuthTicketCode, func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error) {
		return authMsg.ComposeTicket(raw)
	})

	pReg.Register(navigatorMsg.NavigatorInitCode, func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error) {
		return navigatorMsg.ComposeNavigatorInit(raw), nil
	})
	pReg.Register(navigatorMsg.NavigatorSearchCode, func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error) {
		return navigatorMsg.ComposeNavigatorSearch(raw)
	})

	pReg.Register(roomMsg.RoomEnterCode, func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error) {
		return roomMsg.ComposeRoomEnterPacket(raw)
	})
	pReg.Register(guestRoomMsg.GetGuestRoomCode, func(raw protocol.RawPacket, conn protocol.Connection) (protocol.Packet, error) {
		return guestRoomMsg.ComposeGuestRoomPacket(raw)
	})

}

// Handlers generates all the packet handling processing.
func Handlers() {

	hReg := server.GetServer().PacketHandlers()

	hReg.Register(healthMsg.HelloCode, healthHandler.NewHello())
	hReg.Register(healthMsg.PongCode, healthHandler.NewPong())

	hReg.Register(authMsg.AuthTicketCode, authHandler.NewAuthTicket())

	hReg.Register(navigatorMsg.NavigatorInitCode, navigatorHandler.NewNavigatorInit())
	hReg.Register(navigatorMsg.NavigatorSearchCode, navigatorHandler.NewNavigatorSearch())

	hReg.Register(roomMsg.RoomEnterCode, roomHandler.NewRoomEnter())
	hReg.Register(guestRoomMsg.GetGuestRoomCode, roomHandler.NewNavigatorSearch())

}
