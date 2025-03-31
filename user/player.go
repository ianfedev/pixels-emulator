package user

import (
	"pixels-emulator/core/cycle"
	"pixels-emulator/core/model"
	"time"
)

// Player defines an ephemeral room which will be
// stored in memory for in-game modifications.
type Player struct {
	cycle.Cycleable       // Cycleable as the room need to tick every certain amount of time.
	Id              uint  // Id is the identifier of the room
	stamp           int64 // stamp is the last timestamp from cycle.
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

// Load ephemeral user from record.
func Load(user *model.User) *Player {
	return &Player{
		Id:    user.ID,
		stamp: time.Now().UnixMilli(),
	}
}
