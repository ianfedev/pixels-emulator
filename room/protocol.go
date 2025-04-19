package room

import (
	"pixels-emulator/core/event"
	"pixels-emulator/core/protocol"
	"pixels-emulator/room/encode"
	ev "pixels-emulator/room/event"
	"pixels-emulator/room/message"
	"pixels-emulator/room/message/unit"
	"pixels-emulator/room/path"
	"pixels-emulator/user"
	"strings"
)

// CloseConnection is an aux function to send disconnection packets.
func CloseConnection(conn protocol.Connection, reason message.ReasonType, q string, em event.Manager) {
	cPck := &message.CloseRoomConnectionPacket{}
	rPck := &message.DenyRoomConnectionPacket{Type: reason, QueryHolder: q}
	conn.SendPacket(rPck)
	conn.SendPacket(cPck)
	em.Fire(ev.RoomCloseConnectionEventName, ev.NewRoomCloseConnectionEvent(conn, 0, make(map[string]string)))
}

func SendHeightMapPackets(conn protocol.Connection, h int32, l *path.Layout) {

	s, _, y := l.GetSizes()
	rl := strings.ReplaceAll(l.RawMap(), "\\r\\n", "\r")

	fPck := &message.FloorHeightMapRequestPacket{
		Scale:      true, // INVESTIGATION: What does this scale really means?
		WallHeight: h,
		Layout:     rl,
	}
	hPck := &message.HeightMapRequestPacket{
		Width:   int32(y),
		Total:   int32(s),
		Heights: path.GetFlatHeights(l),
	}

	conn.SendPacket(fPck)
	conn.SendPacket(hPck)

}

// SendUnitSyncPacket sends the essential unit data and position to the target player about a group of units.
func SendUnitSyncPacket(origin []*user.Player, target *user.Player) error {

	units := make([]encode.UnitMessage, len(origin))
	for i, p := range origin {
		enc, err := EncodeUnit(p.Unit())
		if err != nil {
			return err
		}
		units[i] = *enc
	}

	target.Conn().SendPacket(&unit.UpdateStatusPacket{Units: units})
	return nil

}
