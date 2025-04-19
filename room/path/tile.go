package path

import (
	"math"
)

type Status int

const (
	Open    = iota // Open defines a walkable tile.
	Blocked        // Blocked defines a no walkable tile.
	Invalid        // Invalid defines an unknown behaviour of the tile.
	Sit            // Sit defines a tile which a unit can be sit in.
	Lay            // Lay defines a tile which a unit can be laying in.
)

const AllowFalling = false    // AllowFalling detects if a unit can fall from a higher Tile. // TODO: Setup this at emulator config
const MaxHeight float64 = 1.1 // MaxHeight defines the maximum height a user can climb a stacked tile.

// Tile defines a single cell of the room map.
type Tile struct {
	X, Y, Z   int16    // X, Y, Z defines coordinates of the tile.
	Units     []string // Units define the list of room units at the tile.
	State     Status   // State defines the current accessibility status of the tile.
	Diagonal  bool     // Diagonal defines if tile allow diagonal.
	stackable bool     // stackable defines if the tile allows stack.
	height    int      // Height defines the actual stackable height of the tile.
}

// UpdateHeight updates the height according to the tile context and value.
func (t *Tile) UpdateHeight(height int) {

	if t.State == Invalid {
		t.height = math.MaxInt16
		t.stackable = false
		return
	}

	if t.height >= 0 && t.height != math.MaxInt16 {
		t.height = height
		t.stackable = true
		return
	}

	t.height = int(t.Z)
	t.stackable = false

}

// Height return the height of the tile.
func (t *Tile) Height() int {
	return t.height
}

// Stackable return if tile is stackable.
func (t *Tile) Stackable() bool {
	return t.stackable && t.State != Invalid
}

// AllowStack updates stack allocation on tile.
func (t *Tile) AllowStack(stackable bool) {
	t.stackable = stackable
}

// RelativeHeight provides the stack height.
func (t *Tile) RelativeHeight() int {

	if t.State == Invalid {
		return math.MaxInt16
	}

	if !t.Stackable() || t.State == Blocked || t.State == Sit {
		return 64 * 256
	}

	return t.Height() * 256

}

// Walkable checks if a tile can be walked on based on its state, height constraints, and presence of units.
func (t *Tile) Walkable(canFall bool, currentAdj *Tile, isFinalDestination bool) bool {
	// If there are units on the tile, it is not walkable.
	if len(t.Units) > 0 {
		return false
	}

	// Calculate height difference
	heightDiff := float64(t.height - currentAdj.height)

	// Prevent movement if the height difference is too large
	if (!canFall && heightDiff < -MaxHeight) || (heightDiff > MaxHeight) {
		return false
	}

	// Open tiles are walkable if height allows.
	if t.State == Open {
		return true
	}

	// Sit and Lay are only walkable if they are the final destination.
	if (t.State == Sit || t.State == Lay) && isFinalDestination {
		return true
	}

	return false
}

func NewTile(x, y, z int16, status Status, stack bool) *Tile {
	return &Tile{
		X:         x,
		Y:         y,
		Z:         z,
		height:    int(z),
		Units:     make([]string, 0),
		State:     status,
		Diagonal:  false,
		stackable: stack,
	}

}
