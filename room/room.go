package room

import (
	"fmt"
	"pixels-emulator/core/cycle"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/util"
	"pixels-emulator/room/message"
	"pixels-emulator/user"
	"strconv"
	"time"
)

// Room defines an ephemeral room which will be
// stored in memory for in-game modifications.
type Room struct {
	cycle.Cycleable                    // Cycleable as the room need to tick every certain amount of time.
	Id              uint               // Id is the identifier of the room
	Queue           *util.Queue[int32] // Queue of users pending to enter
	l               model.HeightMap    // l defines the room layout.
	stamp           int64              // stamp is the last timestamp from cycle
	ready           bool               // ready defines if room finished loading cycle
	em              event.Manager      // em is an event manager to handle further events.
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

func (r *Room) Ready() bool {
	return r.ready
}

func (r *Room) Open(p *user.Player) {

	fmt.Println("LOGGED")
	r.ready = true
	if !r.ready {
		r.Queue.Enqueue(strconv.Itoa(int(p.Id)), int32(p.Id))
		return
	}

	p.Conn().SendPacket(&message.RoomReadyPacket{Room: int32(r.Id), Layout: r.l.Heightmap})
	// TODO: If enqueued, prevent opening and send to queue.

}

func Load(room *model.Room, em event.Manager) *Room {

	q := util.NewQueue[int32]()

	r := &Room{
		Id:    room.ID,
		Queue: q,
		stamp: time.Now().UnixMilli(),
		ready: false,
		em:    em,
		l:     room.Layout,
	}

	go func() {
		// Async load simulation
		//time.Sleep(10 * time.Second)
		fmt.Println("WAHAHA")
		r.ready = true
	}()

	return r

}
