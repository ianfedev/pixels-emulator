package room

import (
	"go.uber.org/zap"
	"pixels-emulator/core/cycle"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/util"
	ev "pixels-emulator/room/event"
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
	logger          *zap.Logger             // logger to logging tools.
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

// Clear removes completely a player from a room.
func (r *Room) Clear(id string) {
	delete(r.Transitioning, id)
	delete(r.Players, id)
	r.Queue.Remove(id)
}

func Load(room *model.Room, logger *zap.Logger, em event.Manager) (*Room, error) {

	q := util.NewQueue[string]()
	cRoom := *room
	l, err := path.NewLayout(&room.Layout)

	if err != nil {
		return nil, err
	}

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
		logger:        logger,
	}

	go func() {
		r.ready = true
		em.Fire(ev.RoomOpenEventName, ev.NewRoomOpenEvent(r.Id, 0, make(map[string]string)))
		for _, p := range r.Transitioning {
			r.Open(p, nil)
		}
		zap.L().Debug("Room opened", zap.Uint("identifier", r.Id))
	}()

	return r, nil

}
