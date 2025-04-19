package room

import (
	"context"
	"go.uber.org/zap"
	"pixels-emulator/room/message"
	"pixels-emulator/room/path"
	"pixels-emulator/user"
)

func (r *Room) Open(p *user.Player, c *path.Coordinate) {

	var err error
	defer func() {
		if err != nil {
			r.logger.Error("Error while opening room session", zap.String("identifier", p.Id), zap.Error(err))
			CloseConnection(p.Conn(), message.Default, "", r.em)
			r.Clear(p.Conn().Identifier())
		}
	}()

	r.ready = true
	if !r.ready {
		r.Transitioning[p.Id] = p
		return
	}

	r.Players[p.Id] = p
	p.Conn().SendPacket(&message.RoomReadyPacket{Room: int32(r.Id), Layout: r.Layout().Slug()})

	// Updates position to the tile on the coordinate provided or the room door.
	var tile path.Coordinate
	door := path.NewCoordinate(
		r.Layout().Door().X(),
		r.Layout().Door().Y(),
		r.Layout().Door().Z(),
		r.Layout().Door().Dir(),
	)

	if c == nil {
		tile = door
	} else {
		tile = *c
	}

	p.Unit().Current = tile
	p.Unit().SetRotation(tile.Dir(), tile.Dir())

	// Prepare player array
	var roomP []*user.Player
	for _, v := range r.Players {
		roomP = append(roomP, v)
	}

	err = SendUnitDetailPacket(context.Background(), r, roomP, p)
	if err != nil {
		return
	}

	// Send new player packet for old players
	for _, online := range roomP {
		err = SendUnitSyncPacket([]*user.Player{p}, online)
		if err != nil {
			return
		}
	}

	// Send online player packet for new player
	err = SendUnitSyncPacket(roomP, p)
	if err != nil {
		return
	}

	// TODO: If enqueued, prevent opening and send to queue.

}
