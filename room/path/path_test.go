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

// TestGetAdjacentTiles check if adjacent tiling on grid is correctly implemented.
func TestGetAdjacentTiles(t *testing.T) {
	layout := &Layout{
		xLen: 3,
		yLen: 3,
		grid: [][]*Tile{
			{&Tile{X: 0, Y: 0}, &Tile{X: 1, Y: 0}, &Tile{X: 2, Y: 0}},
			{&Tile{X: 0, Y: 1}, &Tile{X: 1, Y: 1}, &Tile{X: 2, Y: 1}},
			{&Tile{X: 0, Y: 2}, &Tile{X: 1, Y: 2}, &Tile{X: 2, Y: 2}},
		},
	}

	tile := layout.grid[1][1]

	t.Run("Cardinal directions", func(t *testing.T) {
		adjacent := GetAdjacentTiles(layout, tile, false)
		assert.Len(t, adjacent, 4)
	})

	t.Run("Diagonal directions", func(t *testing.T) {
		adjacent := GetAdjacentTiles(layout, tile, true)
		assert.Len(t, adjacent, 8)
	})

	t.Run("Corner tile without diagonal", func(t *testing.T) {
		adjacent := GetAdjacentTiles(layout, layout.grid[0][0], false)
		assert.Len(t, adjacent, 2)
	})

	t.Run("Corner tile with diagonal", func(t *testing.T) {
		adjacent := GetAdjacentTiles(layout, layout.grid[0][0], true)
		assert.Len(t, adjacent, 3)
	})

	t.Run("Edge tile without diagonal", func(t *testing.T) {
		adjacent := GetAdjacentTiles(layout, layout.grid[1][0], false)
		assert.Len(t, adjacent, 3)
	})

	t.Run("Edge tile with diagonal", func(t *testing.T) {
		adjacent := GetAdjacentTiles(layout, layout.grid[1][0], true)
		assert.Len(t, adjacent, 5)
	})
}
