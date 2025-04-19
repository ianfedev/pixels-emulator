package path

import (
	"github.com/edwingeng/deque"
)

// Request provides an A* pathfinding request between two tiles based on its availability.
type Request struct {
	layout           *Layout            // layout defines the room layout where is request is made.
	basePosition     *Tile              // basePosition defines the origin position from the pathfinding.
	targetPosition   *Tile              // targetPosition defines the target position of the pathfinding.
	openList         deque.Deque        // openList defines the A* open list.
	closedList       map[*Tile]struct{} // closedList defines the A* closed list.
	path             []*Tile            // path defines the selected path
	unit             interface{}        // unit defines the RoomUnit which is requesting the path.
	allowWalkthrough bool               // allowWalkthrough defines if the unit can walk through obstacle tiles.
}

// Node defines a pathfinding node with heuristics.
type Node struct {
	tile    *Tile
	g, h, f int
	parent  *Node
	index   int
}

// Base provides abstract coordinate of base point.
func (pr *Request) Base() Coordinate {
	return NewCoordinate(pr.basePosition.X, pr.basePosition.Y, pr.basePosition.Z, 0)
}

// Target provides abstract coordinate of targeted point.
func (pr *Request) Target() Coordinate {
	return NewCoordinate(pr.targetPosition.X, pr.targetPosition.Y, pr.targetPosition.Z, 0)
}

// CalculatePath finds the shortest path using the A* algorithm.
// It starts from the basePosition and attempts to reach the targetPosition.
func (pr *Request) CalculatePath(diag bool) []*Tile {

	// Validate input to ensure all necessary components exist
	if pr.basePosition == nil || pr.targetPosition == nil || pr.layout == nil {
		return nil
	}

	// Open set holds nodes to be evaluated, closed set tracks visited nodes
	openSet := make(map[*Tile]*Node)
	pr.closedList = make(map[*Tile]struct{})
	pr.openList = deque.NewDeque()

	// Initialize the starting node
	startNode := &Node{
		tile:   pr.basePosition,
		g:      0,
		h:      CalculateCost(int(pr.basePosition.X), int(pr.basePosition.Y), int(pr.targetPosition.X), int(pr.targetPosition.Y), diag),
		parent: nil,
	}
	startNode.f = startNode.g + startNode.h
	openSet[startNode.tile] = startNode
	pr.openList.PushBack(startNode)

	// Main loop: process nodes until a path is found or no nodes remain
	for pr.openList.Len() > 0 {

		// Select the node with the lowest f-score (cost estimate)
		current := pr.openList.PopFront().(*Node)
		pr.removeFromDeque(current)
		pr.closedList[current.tile] = struct{}{}

		// If target is reached, reconstruct and return the path
		if current.tile == pr.targetPosition {
			return pr.reconstructPath(current)
		}

		// Explore adjacent tiles
		for _, neighbor := range GetAdjacentTiles(pr.layout, current.tile, diag) {
			// Skip tiles already processed unless walkthrough is allowed
			if _, found := pr.closedList[neighbor]; found || (!pr.allowWalkthrough && !neighbor.Walkable(pr.allowWalkthrough, current.tile, false)) {
				continue
			}

			// Calculate movement costs
			gScore := current.g + CalculateCost(int(current.tile.X), int(current.tile.Y), int(neighbor.X), int(neighbor.Y), diag)
			hScore := CalculateCost(int(neighbor.X), int(neighbor.Y), int(pr.targetPosition.X), int(pr.targetPosition.Y), diag)
			fScore := gScore + hScore

			// If a better path to the neighbor is found, update its values
			if existing, found := openSet[neighbor]; found && gScore >= existing.g {
				continue
			}

			node := &Node{
				tile:   neighbor,
				g:      gScore,
				h:      hScore,
				f:      fScore,
				parent: current,
			}

			openSet[neighbor] = node
			pr.openList.PushBack(node)
		}
	}

	// No path found, return nil
	return nil
}

// reconstructPath backtracks from the target node to reconstruct the full path.
func (pr *Request) reconstructPath(node *Node) []*Tile {
	var path []*Tile
	for node != nil {
		path = append([]*Tile{node.tile}, path...)
		node = node.parent
	}
	pr.path = path
	return path
}

// removeFromDeque removes a specific node from the deque.
func (pr *Request) removeFromDeque(target *Node) {
	pr.openList.Range(func(i int, v deque.Elem) bool {
		if v.(*Node) == target {
			pr.openList.Replace(i, nil) // Mark as nil (Deque does not support direct deletion)
			return false
		}
		return true
	})
}
