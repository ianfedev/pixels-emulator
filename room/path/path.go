package path

import (
	"math"
)

const (
	CostStraight = 10 // Cost of moving in straight directions.
	CostDiagonal = 14 // Cost of moving diagonally.
)

// Directions for every rotation (8 directions).
var directions = [8][2]int{
	{0, -1},  // 0 → North
	{1, -1},  // 1 → North-east
	{1, 0},   // 2 → East
	{1, 1},   // 3 → South-east
	{0, 1},   // 4 → South
	{-1, 1},  // 5 → South-west
	{-1, 0},  // 6 → West
	{-1, -1}, // 7 → Northwest
}

// Cardinal directions (4 directions: N, E, S, W)
var cardinalDirections = [][2]int{
	{0, -1}, // North
	{1, 0},  // East
	{0, 1},  // South
	{-1, 0}, // West
}

// GetTileInFront obtains the Tile in the direction with an optional offset.
func GetTileInFront(layout *Layout, tile *Tile, rotation Direction, offset int) *Tile {
	if tile == nil || layout == nil || offset < 0 {
		return nil
	}

	// Adjusts rotation from 0-7.
	rotation = rotation % 8
	dx, dy := directions[rotation][0], directions[rotation][1]

	// Calculate new position.
	x, y := int(tile.X)+dx*offset, int(tile.Y)+dy*offset

	// Prevent out-of-bounds access.
	if x < 0 || y < 0 || x >= layout.xLen || y >= layout.yLen {
		return nil
	}
	return layout.grid[x][y]
}

// GetAdjacentTiles returns all adjacent tiles based on the allowDiagonal parameter.
// If allowDiagonal is true, returns neighbors in 8 directions; otherwise, only in 4 cardinal directions.
func GetAdjacentTiles(layout *Layout, tile *Tile, allowDiagonal bool) []*Tile {
	if tile == nil || layout == nil {
		return nil
	}

	var adjacentTiles []*Tile
	var dirs [][2]int

	if allowDiagonal {
		dirs = directions[:]
	} else {
		dirs = cardinalDirections
	}

	for _, dir := range dirs {
		x, y := int(tile.X)+dir[0], int(tile.Y)+dir[1]
		if layout.TileExists(x, y) {
			adjacentTiles = append(adjacentTiles, layout.grid[x][y])
		}
	}

	return adjacentTiles
}

// CalculateCost computes the movement cost between two points.
// Uses diagonal cost if allowed; otherwise, applies Manhattan distance.
func CalculateCost(x1, y1, x2, y2 int, allowDiagonal bool) int {
	dx := int(math.Abs(float64(x2 - x1)))
	dy := int(math.Abs(float64(y2 - y1)))

	if allowDiagonal {
		return CostDiagonal*int(math.Min(float64(dx), float64(dy))) + CostStraight*int(math.Abs(float64(dx-dy)))
	}
	return CostStraight * (dx + dy)
}
