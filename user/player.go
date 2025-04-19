package user

import (
	"pixels-emulator/core/model"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/scheduler"
	"pixels-emulator/room/unit"
	"strconv"
	"time"
)

// Player defines an ephemeral room which will be
// stored in memory for in-game modifications.
type Player struct {
	Id    string              // Id is the identifier of the room
	stamp int64               // stamp is the last timestamp from cycle.
	conn  protocol.Connection // conn defines the connection of the player.
	cr    scheduler.Scheduler // cr defines the scheduler for movement.
	unit  *unit.Unit          // unit defines the player unit
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
func Load(user *model.User, conn protocol.Connection, cs scheduler.Scheduler) *Player {
	id := strconv.Itoa(int(user.ID))
	return &Player{
		Id:    id,
		stamp: time.Now().UnixMilli(),
		conn:  conn,
		unit:  unit.NewUnit(int32(user.ID)),
	}
}
