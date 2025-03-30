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
