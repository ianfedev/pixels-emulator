package room

import (
	"fmt"
	"pixels-emulator/core/cycle"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/util"
	"pixels-emulator/room/message"
	"pixels-emulator/room/path"
	"pixels-emulator/user"
	"time"
)

// Room defines an ephemeral room which will be
// stored in memory for in-game modifications.
type Room struct {
	cycle.Cycleable                         // Cycleable as the room need to tick every certain amount of time.
	Id              uint                    // Id is the identifier of the room
	Transitioning   map[string]*user.Player // Transitioning is the map of users in process of room rendering.
	Data            model.Room              // Data of retrieved from the database when room was loaded.
	Players         map[string]*user.Player // Players are the connected players in-game.
	Queue           *util.Queue[string]     // Queue of users pending to enter
	lData           model.HeightMap         // lData defines the room layout data on load.
	l               *path.Layout            // l defines the generated ephemeral layout.
	stamp           int64                   // stamp is the last timestamp from cycle
	ready           bool                    // ready defines if room finished loading cycle
	em              event.Manager           // em is an event manager to handle further events.
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

func (r *Room) IsOnline(player *user.Player) bool {
	_, ex := r.Players[player.Id]
	return ex
}

func (r *Room) IsTransitioning(player *user.Player) bool {
	_, ex := r.Transitioning[player.Id]
	return ex
}

func (r *Room) Layout() *path.Layout {
	return r.l
}

func (r *Room) Open(p *user.Player) {

	fmt.Println("LOGGED")
	r.ready = true
	if !r.ready {
		r.Transitioning[p.Id] = p
		return
	}

	r.Players[p.Id] = p
	p.Conn().SendPacket(&message.RoomReadyPacket{Room: int32(r.Id), Layout: r.Layout().Slug()})
	// TODO: If enqueued, prevent opening and send to queue.

}

func Load(room *model.Room, em event.Manager) *Room {

	q := util.NewQueue[string]()
	cRoom := *room
	l := path.NewLayout(&room.Layout)

	r := &Room{
		Id:            room.ID,
		Queue:         q,
		Data:          cRoom,
		stamp:         time.Now().UnixMilli(),
		ready:         false,
		em:            em,
		lData:         room.Layout,
		l:             l,
		Transitioning: make(map[string]*user.Player),
		Players:       make(map[string]*user.Player),
	}

	go func() {
		// Async load simulation
		//time.Sleep(10 * time.Second)
		fmt.Println("WAHAHA")
		r.ready = true
	}()

	return r

}
