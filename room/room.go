package room

import (
	"pixels-emulator/core/cycle"
	"pixels-emulator/core/model"
	"pixels-emulator/core/util"
	"pixels-emulator/user"
	"time"
)

// Room defines an ephemeral room which will be
// stored in memory for in-game modifications.
type Room struct {
	cycle.Cycleable                    // Cycleable as the room need to tick every certain amount of time.
	Id              uint               // Id is the identifier of the room
	Queue           *util.Queue[int32] // Queue of users pending to enter
	stamp           int64              // stamp is the last timestamp from cycle.
	ready           bool               // ready defines if room finished loading cycle.
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
	r.stamp = time.Now().UnixMilli()
}

func Open(p *user.Player) {

}

func Load(room *model.Room) *Room {

	q := util.NewQueue[int32]()
	return &Room{
		Id:    room.ID,
		Queue: q,
		stamp: time.Now().UnixMilli(),
		ready: false,
	}

}
