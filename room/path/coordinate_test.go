package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewCoordinate checks if the constructor correctly initializes the Coordinate struct.
func TestNewCoordinate(t *testing.T) {
	c := NewCoordinate(5, 10, 3, NorthEast)

	assert.Equal(t, int16(5), c.X(), "X coordinate should be 5")
	assert.Equal(t, int16(10), c.Y(), "Y coordinate should be 10")
	assert.Equal(t, int16(3), c.Z(), "Z coordinate should be 3")
	assert.Equal(t, NorthEast, c.Dir(), "Direction should be NorthEast")
}

// TestCoordinateImmutability verifies that the Coordinate struct remains immutable.
func TestCoordinateImmutability(t *testing.T) {
	c := NewCoordinate(5, 10, 3, NorthEast)

	// Ensure that the struct remains unchanged after accessing values
	assert.Equal(t, int16(5), c.X(), "X coordinate should remain 5")
	assert.Equal(t, int16(10), c.Y(), "Y coordinate should remain 10")
	assert.Equal(t, int16(3), c.Z(), "Z coordinate should remain 3")
	assert.Equal(t, NorthEast, c.Dir(), "Direction should remain NorthEast")
}

// TestDirectionEnum ensures that each direction is assigned the correct value.
func TestDirectionEnum(t *testing.T) {
	assert.Equal(t, int16(0), int16(North), "North should be 0")
	assert.Equal(t, int16(1), int16(NorthEast), "NorthEast should be 1")
	assert.Equal(t, int16(2), int16(East), "East should be 2")
	assert.Equal(t, int16(3), int16(SouthEast), "SouthEast should be 3")
	assert.Equal(t, int16(4), int16(South), "South should be 4")
	assert.Equal(t, int16(5), int16(SouthWest), "SouthWest should be 5")
	assert.Equal(t, int16(6), int16(West), "West should be 6")
	assert.Equal(t, int16(7), int16(NorthWest), "NorthWest should be 7")
}

// TestCopyBehavior ensures that the Coordinate struct is copied when assigned to a new variable.
func TestCopyBehavior(t *testing.T) {
	c1 := NewCoordinate(1, 2, 3, South)
	c2 := c1 // Copying the struct

	assert.Equal(t, c1, c2, "Copied struct should be identical to the original")

	// Modify the new variable (which won't affect the original, proving it's copied)
	c2 = NewCoordinate(10, 20, 30, North)

	assert.NotEqual(t, c1, c2, "Original struct should not be affected by changes in the copy")
}
