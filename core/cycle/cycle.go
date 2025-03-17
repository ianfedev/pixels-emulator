package cycle

// Cycleable is an element which has a renewable lifecycle after a certain
// amount of ticks, (e.g: A room):
type Cycleable interface {
	Cycle()       // Cycle performs the cycle task
	Time() byte   // Time provides the time of every cycle
	Stamp() int64 // Stamp provides the last timestamp of cycle.
	SetStamp()    // SetStamp updates a new timestamp.
}
