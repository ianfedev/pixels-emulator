package room

import (
	"pixels-emulator/core/protocol"
	"pixels-emulator/room/message"
)

// CloseConnection is an aux function to send disconnection packets.
func CloseConnection(conn protocol.Connection, reason message.ReasonType, q string) {
	cPck := &message.CloseRoomConnectionPacket{}
	rPck := &message.DenyRoomConnectionPacket{Type: reason, QueryHolder: q}
	conn.SendPacket(rPck)
	conn.SendPacket(cPck)
}
