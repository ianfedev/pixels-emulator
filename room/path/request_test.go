package path

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// newTestLayoutWithObstacles creates a layout with dimensions xLen x yLen.
// The obstacles map has keys in the form "x_y" to mark blocked tiles.
func newTestLayoutWithObstacles(xLen, yLen int, obstacles map[string]bool) *Layout {
	layout := &Layout{
		xLen: xLen,
		yLen: yLen,
		grid: make([][]*Tile, xLen),
	}
	for x := 0; x < xLen; x++ {
		layout.grid[x] = make([]*Tile, yLen)
		for y := 0; y < yLen; y++ {
			state := Open
			key := fmt.Sprintf("%d_%d", x, y)
			if obstacles[key] {
				state = Blocked
			}
			layout.grid[x][y] = NewTile(int16(x), int16(y), 3, Status(state), true)
		}
	}
	return layout
}

func TestCalculatePathWithObstacles_Diagonal(t *testing.T) {
	obstacles := map[string]bool{
		"1_1": true,
		"1_2": true,
		"1_3": true,
	}
	layout := newTestLayoutWithObstacles(5, 5, obstacles)
	base := layout.grid[0][0]
	target := layout.grid[4][4]

	req := &Request{
		layout:           layout,
		basePosition:     base,
		targetPosition:   target,
		allowWalkthrough: false,
	}

	path := req.CalculatePath(true)
	fmt.Println(GlyphMap(layout, path, false))
	assert.NotNil(t, path, "Path should not be nil")
	assert.Greater(t, len(path), 0, "Path should contain at least one tile")
	assert.Equal(t, base, path[0], "Path should start at base position")
	assert.Equal(t, target, path[len(path)-1], "Path should end at target position")

	expectedLength := 6
	assert.Equal(t, expectedLength, len(path), fmt.Sprintf("Expected path length %d but got %d", expectedLength, len(path)))
}

func TestCalculatePathWithObstacles_Straight(t *testing.T) {
	obstacles := map[string]bool{
		"2_0": true,
		"2_1": true,
		"2_2": true,
	}
	layout := newTestLayoutWithObstacles(5, 5, obstacles)
	base := layout.grid[0][0]
	target := layout.grid[4][0]

	req := &Request{
		layout:           layout,
		basePosition:     base,
		targetPosition:   target,
		allowWalkthrough: false,
	}

	path := req.CalculatePath(false)
	fmt.Println(GlyphMap(layout, path, false))
	assert.NotNil(t, path, "Path should not be nil")
	assert.Greater(t, len(path), 0, "Path should contain at least one tile")
	assert.Equal(t, base, path[0], "Path should start at base position")
	assert.Equal(t, target, path[len(path)-1], "Path should end at target position")

	assert.LessOrEqual(t, len(path), 12, "Path length should be less or equal to 12")
}

func TestCalculatePath_NoPathAvailable(t *testing.T) {
	// Completely blocked path.
	obstacles := map[string]bool{
		"1_0": true, "1_1": true, "1_2": true, "1_3": true,
		"2_0": true, "2_1": true, "2_2": true, "2_3": true,
	}
	layout := newTestLayoutWithObstacles(4, 4, obstacles)
	base := layout.grid[0][0]
	target := layout.grid[3][3]

	req := &Request{
		layout:           layout,
		basePosition:     base,
		targetPosition:   target,
		allowWalkthrough: false,
	}

	path := req.CalculatePath(true)
	assert.Nil(t, path, "Path should be nil when no valid path exists")
}

func TestCalculatePath_StartEqualsTarget(t *testing.T) {
	layout := newTestLayoutWithObstacles(5, 5, nil)
	base := layout.grid[2][2]
	req := &Request{
		layout:           layout,
		basePosition:     base,
		targetPosition:   base,
		allowWalkthrough: false,
	}

	path := req.CalculatePath(true)
	assert.NotNil(t, path, "Path should not be nil")
	assert.Len(t, path, 1, "Path should contain only the start tile")
}

func TestCalculatePath_LargeGrid(t *testing.T) {
	layout := newTestLayoutWithObstacles(50, 50, nil)
	base := layout.grid[0][0]
	target := layout.grid[49][49]

	req := &Request{
		layout:           layout,
		basePosition:     base,
		targetPosition:   target,
		allowWalkthrough: false,
	}

	path := req.CalculatePath(true)
	assert.NotNil(t, path, "Path should be found in an open 50x50 grid")
	assert.Greater(t, len(path), 0, "Path should contain steps")
	assert.Equal(t, base, path[0], "Path should start at base position")
	assert.Equal(t, target, path[len(path)-1], "Path should end at target position")
}

func TestCalculatePath_CornerToCorner(t *testing.T) {
	layout := newTestLayoutWithObstacles(10, 10, nil)
	base := layout.grid[0][0]
	target := layout.grid[9][9]

	req := &Request{
		layout:           layout,
		basePosition:     base,
		targetPosition:   target,
		allowWalkthrough: false,
	}

	path := req.CalculatePath(true)
	assert.NotNil(t, path, "Path should not be nil")
	assert.Equal(t, base, path[0], "Path should start at base position")
	assert.Equal(t, target, path[len(path)-1], "Path should end at target position")
}

func TestCalculatePath_NarrowPassage(t *testing.T) {
	// Creates a narrow passage from (0,0) to (4,4)
	obstacles := map[string]bool{
		"1_0": true, "1_1": true, "1_2": true, "1_3": true,
		"2_0": true, "2_1": true, "2_2": true, "2_3": true,
		"3_0": true, "3_1": true, "3_2": true, "3_3": true,
	}
	layout := newTestLayoutWithObstacles(5, 5, obstacles)
	base := layout.grid[0][0]
	target := layout.grid[4][4]

	req := &Request{
		layout:           layout,
		basePosition:     base,
		targetPosition:   target,
		allowWalkthrough: false,
	}

	path := req.CalculatePath(true)
	assert.NotNil(t, path, "Path should be found through the narrow passage")
	assert.Greater(t, len(path), 0, "Path should contain steps")
}
