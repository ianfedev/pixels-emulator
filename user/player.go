package user

import (
	"pixels-emulator/core/cycle"
	"pixels-emulator/core/util"
	"time"
)

// Player defines an ephemeral room which will be
// stored in memory for in-game modifications.
type Player struct {
	cycle.Cycleable                   // Cycleable as the room need to tick every certain amount of time.
	Id              int32             // Id is the identifier of the room
	Queue           util.Queue[int32] // Queue of users pending to enter
	stamp           int64             // stamp is the last timestamp from cycle.
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
	time.Now().UnixMilli()
}
