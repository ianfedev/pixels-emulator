package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pixels-emulator/core/model"
)

// TestGetFlatHeights tests if Nitro requested flat height is correctly serialized.
func TestGetFlatHeights(t *testing.T) {
	hMap := &model.HeightMap{
		Slug:          "test_map",
		DoorX:         5,
		DoorY:         5,
		DoorDirection: 2,
		Heightmap:     "xxxx\nx22x\nx01x\nxxxx",
	}

	layout, err := NewLayout(hMap)
	assert.NoError(t, err)

	heights := GetFlatHeights(layout)

	assert.Equal(t, 16, len(heights), "El largo del slice debe ser 9")

	expected := []int16{
		0, 0, 0, 0,
		0, 2, 2, 0,
		0, 0, 1, 0,
		0, 0, 0, 0,
	}

	assert.Equal(t, expected, heights)
}
