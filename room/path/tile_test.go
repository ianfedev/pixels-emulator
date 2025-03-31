package path

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestWalkable(t *testing.T) {
	adjacent := NewTile(1, 1, 0, Open, true)
	t.Run("Walkable Open Tile", func(t *testing.T) {
		tile := NewTile(2, 2, 0, Open, true)
		assert.True(t, tile.Walkable(false, adjacent, false))
	})

	t.Run("Blocked Tile", func(t *testing.T) {
		tile := NewTile(2, 2, 0, Blocked, true)
		assert.False(t, tile.Walkable(false, adjacent, false))
	})

	t.Run("Invalid Tile", func(t *testing.T) {
		tile := NewTile(2, 2, 0, Invalid, true)
		assert.False(t, tile.Walkable(false, adjacent, false))
	})

	t.Run("Sit Tile Not Destination", func(t *testing.T) {
		tile := NewTile(2, 2, 0, Sit, true)
		assert.False(t, tile.Walkable(false, adjacent, false))
	})

	t.Run("Sit Tile As Destination", func(t *testing.T) {
		tile := NewTile(2, 2, 0, Sit, true)
		assert.True(t, tile.Walkable(false, adjacent, true))
	})

	t.Run("Lay Tile Not Destination", func(t *testing.T) {
		tile := NewTile(2, 2, 0, Lay, true)
		assert.False(t, tile.Walkable(false, adjacent, false))
	})

	t.Run("Lay Tile As Destination", func(t *testing.T) {
		tile := NewTile(2, 2, 0, Lay, true)
		assert.True(t, tile.Walkable(false, adjacent, true))
	})

	t.Run("Tile With Units", func(t *testing.T) {
		tile := NewTile(2, 2, 0, Open, true)
		tile.Units = append(tile.Units, struct{}{})
		assert.False(t, tile.Walkable(false, adjacent, false))
	})

	t.Run("Too High To Climb", func(t *testing.T) {
		tile := NewTile(2, 2, 2, Open, true)
		adj := NewTile(1, 1, 0, Open, true)
		assert.False(t, tile.Walkable(false, adj, false))
	})

	t.Run("Fall Allowed", func(t *testing.T) {
		tile := NewTile(2, 2, -2, Open, true)
		assert.True(t, tile.Walkable(true, adjacent, false))
	})
}

// TestRelativeHeight verifies that RelativeHeight returns the expected values.
func TestRelativeHeight(t *testing.T) {
	t.Run("Invalid Tile", func(t *testing.T) {
		tile := NewTile(1, 1, 0, Invalid, false)
		assert.Equal(t, math.MaxInt16, tile.RelativeHeight())
	})

	t.Run("Blocked Tile", func(t *testing.T) {
		tile := NewTile(1, 1, 0, Blocked, false)
		assert.Equal(t, 64*256, tile.RelativeHeight())
	})

	t.Run("Sit Tile", func(t *testing.T) {
		tile := NewTile(1, 1, 0, Sit, false)
		assert.Equal(t, 64*256, tile.RelativeHeight())
	})

	t.Run("Non-Stackable Open Tile", func(t *testing.T) {
		tile := NewTile(1, 1, 2, Open, false)
		assert.Equal(t, 64*256, tile.RelativeHeight())
	})

	t.Run("Stackable Open Tile", func(t *testing.T) {
		tile := NewTile(1, 1, 3, Open, true)
		assert.Equal(t, 3*256, tile.RelativeHeight())
	})

	t.Run("Stackable Lay Tile", func(t *testing.T) {
		tile := NewTile(1, 1, 5, Lay, true)
		assert.Equal(t, 5*256, tile.RelativeHeight())
	})
}

// TestUpdateHeight verifies that UpdateHeight updates the height correctly.
func TestUpdateHeight(t *testing.T) {
	t.Run("Invalid Tile", func(t *testing.T) {
		tile := NewTile(1, 1, 0, Invalid, true)
		tile.UpdateHeight(5)
		assert.Equal(t, math.MaxInt16, tile.Height())
		assert.False(t, tile.Stackable())
	})

	t.Run("Valid Tile with Positive Height", func(t *testing.T) {
		tile := NewTile(1, 1, 0, Open, true)
		tile.UpdateHeight(5)
		assert.Equal(t, 5, tile.Height())
		assert.True(t, tile.Stackable())
	})

	t.Run("Valid Tile with Zero Height", func(t *testing.T) {
		tile := NewTile(1, 1, 0, Open, true)
		tile.UpdateHeight(0)
		assert.Equal(t, 0, tile.Height())
		assert.True(t, tile.Stackable())
	})

	t.Run("Tile with MaxInt16 Height", func(t *testing.T) {
		tile := NewTile(1, 1, 0, Open, true)
		tile.UpdateHeight(math.MaxInt16)
		assert.Equal(t, math.MaxInt16, tile.Height())
		assert.True(t, tile.Stackable())
	})

	t.Run("Tile with Negative Height", func(t *testing.T) {
		tile := NewTile(1, 1, 5, Open, true)
		tile.UpdateHeight(-3)
		assert.Equal(t, -3, tile.Height())
		assert.True(t, tile.Stackable())
	})

	t.Run("Tile Keeps Initial Height if Unchanged", func(t *testing.T) {
		tile := NewTile(1, 1, 7, Open, true)
		tile.UpdateHeight(tile.Height())
		assert.Equal(t, 7, tile.Height())
		assert.True(t, tile.Stackable())
	})
}
