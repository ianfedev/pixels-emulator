package path

import (
	"math"
)

// GetFlatHeights create a flat array with every height in Nitro way (Y first, X then).
func GetFlatHeights(l *Layout) []int16 {

	heights := make([]int16, l.size)

	i := 0
	for y := 0; y < l.yLen; y++ {
		for x := 0; x < l.xLen; x++ {

			if !l.TileExists(x, y) {
				heights[i] = math.MaxInt16
			} else { // You can force this to be non-stackable but, I do not intend to create a non-stackable emulator.
				heights[i] = int16(l.GetTile(x, y).Height())
			}
			i++

		}
	}

	return heights

}
