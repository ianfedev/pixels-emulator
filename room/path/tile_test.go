package path

import (
	"github.com/stretchr/testify/assert"
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
