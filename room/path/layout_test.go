package path

import (
	"pixels-emulator/core/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewLayout verifies that a Layout is correctly created.
func TestNewLayout(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_map",
		DoorX:         5,
		DoorY:         5,
		DoorDirection: 2,
		Heightmap:     "xxxx\nx22x\nx00x\nxxxx",
	}

	layout, err := NewLayout(hMap)

	assert.NoError(t, err)
	assert.NotNil(t, layout)
	assert.Equal(t, "test_map", layout.Slug())
	assert.Equal(t, 5, int(layout.Door().X()))
	assert.Equal(t, 5, int(layout.Door().Y()))
	assert.Equal(t, East, layout.Door().Dir())
}

// TestTileExists verifies tile existence checks.
func TestTileExists(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_map",
		DoorX:         1,
		DoorY:         1,
		DoorDirection: 0,
		Heightmap:     "xxx\r\nx2x\r\nxxx",
	}

	layout, err := NewLayout(hMap)

	assert.NoError(t, err)
	assert.True(t, layout.TileExists(1, 1))
	assert.False(t, layout.TileExists(5, 5)) // Out of bounds
	assert.False(t, layout.TileExists(-1, -1))
}

// TestLayout_GetSizes verifies size of the tiles
func TestLayout_GetSizes(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_map",
		DoorX:         1,
		DoorY:         1,
		DoorDirection: 0,
		Heightmap:     "xxx\r\nx2x\r\nxxx",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)

	tot, x, y := layout.GetSizes()
	assert.Equal(t, tot, 9)
	assert.Equal(t, x, 3)
	assert.Equal(t, y, 3)

}

// TestDoorTile verifies door tile retrieval.
func TestDoorTile(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_map",
		DoorX:         1,
		DoorY:         1,
		DoorDirection: 0,
		Heightmap:     "xxx\r\nx2x\r\nxxx",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)

	doorTile := layout.DoorTile()
	assert.NotNil(t, doorTile)
	assert.Equal(t, int16(1), doorTile.X)
	assert.Equal(t, int16(1), doorTile.Y)
}

// TestSetupDoor ensures the door tile is set up correctly.
func TestSetupDoor(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_map",
		DoorX:         1,
		DoorY:         1,
		DoorDirection: 2,
		Heightmap:     "xxxx\r\nx22x\r\nx00x\r\nxxxx",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)
	layout.setupDoor()

	doorTile := layout.DoorTile()
	assert.False(t, doorTile.Stackable()) // Door tile should not allow stacking
}

// TestLetterToHeight ensures letters are correctly converted to heights.
func TestLetterToHeight(t *testing.T) {
	h, err := letterToHeight("A")
	assert.NoError(t, err)
	assert.Equal(t, int16(10), h)

	h, err = letterToHeight("Z")
	assert.NoError(t, err)
	assert.Equal(t, int16(35), h)

	h, err = letterToHeight("!")
	assert.Error(t, err)
	assert.Equal(t, int16(0), h)
	h, err = letterToHeight("aa")

	assert.Error(t, err)
}

// TestGenerateGrid_EmptyOrMalformedRow tests when a malformed row is present.
func TestGenerateGrid_EmptyOrMalformedRow(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_malformed",
		DoorX:         1,
		DoorY:         1,
		DoorDirection: 0,
		Heightmap:     "xxxx\r\n\r\nx2x\r\nxxxx",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)

	assert.Equal(t, 4, layout.yLen)
	assert.Equal(t, 4, layout.xLen)

	assert.Nil(t, layout.grid[0][1])
	assert.Nil(t, layout.grid[1][1])
	assert.Nil(t, layout.grid[2][1])
	assert.Nil(t, layout.grid[3][1])
}

// TestGenerateGrid_EmptyHeight tests when a height character is empty.
func TestGenerateGrid_EmptyHeight(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_empty_height",
		DoorX:         1,
		DoorY:         1,
		DoorDirection: 0,
		Heightmap:     "xxxx\r\nx xx\r\nxxxx",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)

	assert.Equal(t, int16(0), layout.grid[1][1].Z)
}

// TestGenerateGrid_LetterHeightValid tests when a letter is correctly converted to height.
func TestGenerateGrid_LetterHeightValid(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_letter_height",
		DoorX:         1,
		DoorY:         1,
		DoorDirection: 0,
		Heightmap:     "xxxx\r\nxAxx\r\nxxxx",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)

	assert.Equal(t, int16(10), layout.grid[1][1].Z)
}

func TestGenerateGrid_LetterHeightInvalid(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_invalid_letter",
		DoorX:         1,
		DoorY:         1,
		DoorDirection: 0,
		Heightmap:     "xxxx\r\nx!xx\r\nxxxx",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)
	assert.Equal(t, int16(0), layout.grid[1][1].Z)
}

// TestSetupDoor_NoDoorTile ensures setupDoor does nothing if the door tile is nil.
func TestSetupDoor_NoDoorTile(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_no_door_tile",
		DoorX:         10,
		DoorY:         10,
		DoorDirection: 2,
		Heightmap:     "xxxx\r\nx22x\r\nx00x\r\nxxxx",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)
	layout.setupDoor()

	assert.Nil(t, layout.DoorTile())
}

// TestSetupDoor_NoDoorFrontTile ensures setupDoor does nothing if the front tile is nil.
func TestSetupDoor_NoDoorFrontTile(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_no_door_front_tile",
		DoorX:         1,
		DoorY:         1,
		DoorDirection: 6,
		Heightmap:     "x",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)
	layout.setupDoor()

	doorTile := layout.DoorTile()
	assert.Nil(t, doorTile)
}
