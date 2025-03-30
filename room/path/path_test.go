package path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTileInFront(t *testing.T) {
	layout := &Layout{
		xLen: 5,
		yLen: 5,
		grid: makeGrid(5, 5),
	}
	tile := layout.grid[2][2] // Central tile

	tests := []struct {
		name      string
		rotation  Direction
		offset    int
		expectNil bool
		expectX   int
		expectY   int
	}{
		{"North in bounds", 0, 1, false, 2, 1},
		{"South out of bounds", 4, 3, true, 0, 0},
		{"East in bounds", 2, 2, false, 4, 2},
		{"West out of bounds", 6, 3, true, 0, 0},
		{"Negative offset", 0, -1, true, 0, 0},
		{"Zero offset (same tile)", 0, 0, false, 2, 2},
		{"Diagonal NE in bounds", 1, 1, false, 3, 1},
		{"Diagonal SW out of bounds", 5, 3, true, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tileInFront := GetTileInFront(layout, tile, tt.rotation, tt.offset)
			if tt.expectNil {
				assert.Nil(t, tileInFront, "Expected nil")
			} else {
				assert.NotNil(t, tileInFront, "Expected a tile, got nil")
				assert.Equal(t, tt.expectX, int(tileInFront.X), "Unexpected X coordinate")
				assert.Equal(t, tt.expectY, int(tileInFront.Y), "Unexpected Y coordinate")
			}
		})
	}
}

func makeGrid(x, y int) [][]*Tile {
	grid := make([][]*Tile, x)
	for i := range grid {
		grid[i] = make([]*Tile, y)
		for j := range grid[i] {
			grid[i][j] = &Tile{X: int16(uint8(i)), Y: int16(uint8(j))}
		}
	}
	return grid
}

// TestGetAdjacentTiles checks if GetAdjacentTiles returns correct neighbors.
func TestGetAdjacentTiles(t *testing.T) {
	layout := &Layout{
		xLen: 3,
		yLen: 3,
		grid: [][]*Tile{
			{NewTile(0, 0, 0, Open, true), NewTile(1, 0, 0, Open, true), NewTile(2, 0, 0, Open, true)},
			{NewTile(0, 1, 0, Open, true), NewTile(1, 1, 0, Open, true), NewTile(2, 1, 0, Open, true)},
			{NewTile(0, 2, 0, Open, true), NewTile(1, 2, 0, Open, true), NewTile(2, 2, 0, Open, true)},
		},
	}
	tile := layout.grid[1][1]
	adjacent := GetAdjacentTiles(layout, tile)
	assert.Len(t, adjacent, 8, "Expected 8 adjacent tiles")

	nilAdjacent := GetAdjacentTiles(nil, nil)
	assert.Nil(t, nilAdjacent, "Expected nil result for nil inputs")
}

// TestGetAdjacentTilesEdge checks if GetAdjacentTiles handles edge cases.
func TestGetAdjacentTilesEdge(t *testing.T) {
	layout := &Layout{
		xLen: 3,
		yLen: 3,
		grid: [][]*Tile{
			{NewTile(0, 0, 0, Open, true), nil, NewTile(2, 0, 0, Open, true)},
			{NewTile(0, 1, 0, Open, true), NewTile(1, 1, 0, Open, true), NewTile(2, 1, 0, Open, true)},
			{NewTile(0, 2, 0, Open, true), NewTile(1, 2, 0, Open, true), NewTile(2, 2, 0, Open, true)},
		},
	}
	tile := layout.grid[0][0]
	adjacent := GetAdjacentTiles(layout, tile)
	assert.Less(t, len(adjacent), 8, "Expected fewer than 8 adjacent tiles due to missing neighbors")
}

// TestCalculateCost verifies movement cost calculations in different scenarios.
func TestCalculateCost(t *testing.T) {
	tests := []struct {
		name           string
		x1, y1, x2, y2 int
		allowDiagonal  bool
		expectedCost   int
	}{
		{"Straight horizontal", 0, 0, 3, 0, false, 30},
		{"Straight vertical", 0, 0, 0, 4, false, 40},
		{"Diagonal equal steps", 0, 0, 3, 3, true, 42},
		{"Diagonal with extra horizontal", 0, 0, 4, 3, true, 52},
		{"Diagonal with extra vertical", 0, 0, 3, 4, true, 52},
		{"No movement", 2, 2, 2, 2, true, 0},
		{"No movement without diagonal", 2, 2, 2, 2, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cost := CalculateCost(tt.x1, tt.y1, tt.x2, tt.y2, tt.allowDiagonal)
			assert.Equal(t, tt.expectedCost, cost)
		})
	}
}
