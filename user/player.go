package user

import (
	"pixels-emulator/core/cycle"
	"pixels-emulator/core/model"
	"pixels-emulator/core/protocol"
	"strconv"
	"time"
)

// Player defines an ephemeral room which will be
// stored in memory for in-game modifications.
type Player struct {
	cycle.Cycleable                     // Cycleable as the room need to tick every certain amount of time.
	Id              string              // Id is the identifier of the room
	stamp           int64               // stamp is the last timestamp from cycle.
	conn            protocol.Connection // conn defines the connection of the player.
}

func (r *Player) Cycle() {

}

func (r *Player) Time() byte {
	return 0
}

func (r *Player) Stamp() int64 {
	return r.stamp
}

func (r *Player) SetStamp() {
	r.stamp = time.Now().UnixMilli()
}

// Conn returns the user connection.
func (r *Player) Conn() protocol.Connection {
	return r.conn
}

// Load ephemeral user from record.
func Load(user *model.User, conn protocol.Connection) *Player {
	return &Player{
		Id:    strconv.Itoa(int(user.ID)),
		stamp: time.Now().UnixMilli(),
		conn:  conn,
	}
}
