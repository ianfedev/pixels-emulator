package path

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// newTestLayoutWithHeights creates a layout with specified heights and states.
func newTestLayoutWithHeights(xLen, yLen int, heights map[string]int, states map[string]Status) *Layout {
	layout := &Layout{
		xLen: xLen,
		yLen: yLen,
		grid: make([][]*Tile, xLen),
	}
	for x := 0; x < xLen; x++ {
		layout.grid[x] = make([]*Tile, yLen)
		for y := 0; y < yLen; y++ {
			var state Status = Open
			key := fmt.Sprintf("%d_%d", x, y)
			height := 0
			if h, exists := heights[key]; exists {
				height = h
			}
			if s, exists := states[key]; exists {
				state = s
			}
			layout.grid[x][y] = NewTile(int16(x), int16(y), int16(height), state, true)
		}
	}
	return layout
}

// TestGlyphMapFlatMode verifies that all open tiles are displayed as 'O'.
func TestGlyphMapFlatMode(t *testing.T) {
	layout := newTestLayoutWithHeights(3, 3, nil, nil)
	path := []*Tile{layout.grid[0][0], layout.grid[2][2]}
	output := GlyphMap(layout, path, true)

	expected := "SOO\nOOO\nOOE\n"
	assert.Equal(t, expected, output, "Flat mode visualization incorrect")
}

// TestGlyphMapHeightMode verifies that open tiles display their height value.
func TestGlyphMapHeightMode(t *testing.T) {
	heights := map[string]int{
		"0_0": 5, "1_0": 7, "2_0": 2,
		"0_1": 8, "1_1": 3, "2_1": 6,
		"0_2": 4, "1_2": 1, "2_2": 9,
	}
	layout := newTestLayoutWithHeights(3, 3, heights, nil)
	path := []*Tile{layout.grid[0][0], layout.grid[2][2]}
	output := GlyphMap(layout, path, false)

	expected := "S72\n836\n41E\n"
	assert.Equal(t, expected, output, "Height mode visualization incorrect")
}

// TestGlyphMapWithObstacles verifies that obstacles are displayed correctly.
func TestGlyphMapWithObstacles(t *testing.T) {
	states := map[string]Status{
		"1_0": Blocked,
		"1_1": Blocked,
	}
	layout := newTestLayoutWithHeights(3, 3, nil, states)
	path := []*Tile{layout.grid[0][0], layout.grid[2][2]}
	output := GlyphMap(layout, path, true)

	expected := "S#O\nO#O\nOOE\n"
	assert.Equal(t, expected, output, "Obstacle rendering incorrect")
}

// TestGlyphMapWithEmptyPath verifies that the layout is rendered correctly without a path.
func TestGlyphMapWithEmptyPath(t *testing.T) {
	layout := newTestLayoutWithHeights(3, 3, nil, nil)
	output := GlyphMap(layout, nil, true)

	expected := "OOO\nOOO\nOOO\n"
	assert.Equal(t, expected, output, "Empty path rendering incorrect")
}

// TestGlyphMapWithSpecialTiles verifies that special tile states are displayed correctly.
func TestGlyphMapWithSpecialTiles(t *testing.T) {
	states := map[string]Status{
		"0_0": Invalid,    // X
		"1_0": Sit,        // T
		"2_0": Lay,        // L
		"1_1": Status(99), // Default '?'
	}
	layout := newTestLayoutWithHeights(3, 3, nil, states)
	path := []*Tile{layout.grid[2][2]}
	output := GlyphMap(layout, path, true)

	expected := "XTL\nO?O\nOOS\n"
	assert.Equal(t, expected, output, "Special tile rendering incorrect")
}
