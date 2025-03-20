package room

import (
	genstack "github.com/markekraus/genstack/pkg"
	"pixels-emulator/core/cycle"
	"time"
)

// Room defines an ephemeral room which will be
// stored in memory for in-game modifications.
type Room struct {
	cycle.Cycleable                     // Cycleable as the room need to tick every certain amount of time.
	Queue           genstack.Stack[int] // Queue of users pending to enter
	stamp           int64               // stamp is the last timestamp from cycle.
}

func (r *Room) Cycle() {
}

func (r *Room) Time() byte {
	return 0
}

func (r *Room) Stamp() int64 {
	return r.stamp
}

func (r *Room) SetStamp() {
	time.Now().UnixMilli()
}
