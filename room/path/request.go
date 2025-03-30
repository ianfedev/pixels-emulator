package path

import "pixels-emulator/core/util"

// Request provides an A* pathfinding request between two tiles based on its availability.
type Request struct {
	layout           *Layout        // layout defines the room layout where is request is made.
	basePosition     *Tile          // basePosition defines the origin position from the pathfinding.
	targetPosition   *Tile          // targetPosition defines the target position of the pathfinding.
	openList         util.Dequeue   // openList defines the A* open list.
	closedList       map[*Tile]bool // closedList defines the A* closed list.
	path             []*Tile        // path defines the selected path
	unit             interface{}    // unit defines the RoomUnit which is requesting the path.
	allowWalkthrough bool           // allowWalkthrough defines if the unit can walk through obstacle tiles.
}

// Node defines a pathfinding node with heuristics.
type Node struct {
	tile    *Tile
	g, h, f int
	parent  *Node
	index   int
}

func (pr *Request) CalculatePath() []*Tile {

	return nil
}
