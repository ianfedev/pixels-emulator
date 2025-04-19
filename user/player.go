package user

import (
	"context"
	"pixels-emulator/core/database"
	"pixels-emulator/core/model"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/room/unit"
	"strconv"
)

// Player defines an ephemeral room which will be
// stored in memory for in-game modifications.
type Player struct {
	Id string // Id is the identifier of the room

	// Private dependencies
	conn protocol.Connection              // conn defines the connection of the player.
	cr   scheduler.Scheduler              // cr defines the scheduler for movement.
	svc  database.DataService[model.User] // svc defines the user service to query.
	unit *unit.Unit                       // unit defines the player unit
}

func (p *Player) Record(ctx context.Context) <-chan struct {
	Data  *model.User
	Error error
} {
	result := make(chan struct {
		Data  *model.User
		Error error
	}, 1)
	id, err := strconv.ParseInt(p.Id, 10, 32)
	if err != nil {
		result <- struct {
			Data  *model.User
			Error error
		}{nil, err}
	}
	return p.svc.Get(ctx, uint(id))
}

func (p *Player) Move() {

}

func (p *Player) Unit() *unit.Unit {
	return p.unit
}

// Conn returns the user connection.
func (p *Player) Conn() protocol.Connection {
	return p.conn
}

// Load ephemeral user from record.
func Load(
	user *model.User,
	conn protocol.Connection,
	cr scheduler.Scheduler,
	svc database.DataService[model.User],
) *Player {
	id := strconv.Itoa(int(user.ID))
	return &Player{
		Id:   id,
		conn: conn,
		unit: unit.NewUnit(int32(user.ID)),
		cr:   cr,
		svc:  svc,
	}
}
