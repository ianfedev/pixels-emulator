package room

import (
	"context"
	"pixels-emulator/core/protocol"
	"pixels-emulator/room/message"
	"pixels-emulator/room/path"
	"pixels-emulator/user"
	"strings"
)

// CloseConnection is an aux function to send disconnection packets.
func CloseConnection(conn protocol.Connection, reason message.ReasonType, q string) {
	cPck := &message.CloseRoomConnectionPacket{}
	rPck := &message.DenyRoomConnectionPacket{Type: reason, QueryHolder: q}
	conn.SendPacket(rPck)
	conn.SendPacket(cPck)
}

// GetUserRoom provides the related user room, if queuing, transitioning or in-game.
func GetUserRoom(ctx context.Context, rs Store, p *user.Player) (*Room, error) {

	room, err := rs.Records().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, r := range room {

		if r.Queue.Contains(p.Id) {
			return r, nil
		}

		if r.IsTransitioning(p) {
			return r, nil
		}

		if r.IsOnline(p) {
			return r, nil
		}

	}

	return nil, nil

}

func SendHeightMapPackets(conn protocol.Connection, h int32, l *path.Layout) {

	s, _, y := l.GetSizes()

	fPck := &message.FloorHeightMapRequestPacket{
		WallHeight:  h,
		RelativeMap: strings.ReplaceAll(l.RawMap(), "\r\n", "\r"),
	}
	hPck := &message.HeightMapRequestPacket{
		Width:   int32(y),
		Total:   int32(s),
		Heights: path.GetFlatHeights(l),
	}

	conn.SendPacket(fPck)
	conn.SendPacket(hPck)

}
