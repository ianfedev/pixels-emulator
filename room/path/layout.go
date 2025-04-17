package path

import (
	"errors"
	"pixels-emulator/core/model"
	"strconv"
	"strings"
)

// Layout defines the spatial grid of the room
// composed by an array of X,Y,Z positioned Tile list
// and a heightmap which is used for pathfinding.
type Layout struct {

	// Data from database storage to prevent modification after load.
	slug    string      // slug is the identifier of the map.
	doorPos *Coordinate // doorPos defines the coordinates of the door.
	hMap    string      // hMap stores the layout of the map as a string.

	// Data from processing
	size, xLen, yLen int // size defines the quantity of elements and length from the layout.

	grid [][]*Tile // grid provides the tile grid created from the hMap.

}

// Slug provides the layout slug.
func (l *Layout) Slug() string {
	return l.slug
}

// Door provides the door direction.
func (l *Layout) Door() *Coordinate {
	return l.doorPos
}

// GetSizes provides total tiles, x and y length.
func (l *Layout) GetSizes() (int, int, int) {
	return l.size, l.xLen, l.yLen
}

// RawMap provides the heightmap in raw state.
func (l *Layout) RawMap() string {
	return l.hMap
}

// generateGrid creates a set of Tile from the heightmap provided.
func (l *Layout) generateGrid() {

	// Creates an array with all Y tiles to further processing
	c := strings.ReplaceAll(l.hMap, "\n", "")
	yRawGrid := strings.Split(c, "\r")
	l.size = 0
	l.xLen = len(yRawGrid[0])
	l.yLen = len(yRawGrid)

	l.grid = make([][]*Tile, l.xLen)
	for i := 0; i < l.xLen; i++ {
		l.grid[i] = make([]*Tile, l.yLen)
	}

	for y, yRow := range yRawGrid {

		l.grid[y] = make([]*Tile, l.yLen)

		if len(yRow) == 0 || yRow == "\r" || len(yRow) != l.xLen {
			// Prevent bad formatted lines or empty ones.
			continue
		}

		for x := range len(yRow) {

			xRow := string(yRow[x])
			var state Status = Open
			var h int16 = 0

			if xRow == "x" {
				state = Invalid
			} else {

				n, err := strconv.ParseInt(xRow, 10, 16)

				if err == nil { // If int is parseable set as height
					h = int16(n)
				} else { // Check if is exponential
					exL, err := letterToHeight(xRow)
					if err != nil {
						h = 0
					} else {
						h = exL
					}
				}
			}

			l.size++
			l.grid[x][y] = NewTile(int16(x), int16(y), h, state, true)
		}

	}

	if l.DoorTile() != nil {
		l.setupDoor()
	}

}

// TileExists checks if a Tile exists at the given coordinates.
func (l *Layout) TileExists(x, y int) bool {
	if x < 0 || y < 0 || x >= l.xLen || y >= l.yLen {
		return false
	}
	return l.grid[x][y] != nil
}

// GetTile provides a tile at an specific coordinate.
func (l *Layout) GetTile(x, y int) *Tile {
	return l.grid[x][y]
}

// DoorTile obtains the door tile of the layout.
func (l *Layout) DoorTile() *Tile {
	if l.doorPos == nil || !l.TileExists(int(l.doorPos.X()), int(l.doorPos.Y())) {
		return nil
	}
	return l.grid[l.doorPos.X()][l.doorPos.Y()]
}

// setupDoor configure correctly the door tile to prevent stacking and validate non-blocking behaviour.2
func (l *Layout) setupDoor() {
	doorTile := l.DoorTile()
	if doorTile == nil {
		return
	}

	doorTile.AllowStack(false)

	doorFrontTile := GetTileInFront(l, doorTile, l.doorPos.Dir(), 0)
	if doorFrontTile == nil || !l.TileExists(int(doorFrontTile.X), int(doorFrontTile.Y)) {
		return
	}

	// Verifies if front tile is invalid
	frontTile := l.grid[doorFrontTile.X][doorFrontTile.Y]
	if frontTile.State != Invalid {
		// Adjustment if results are different
		if l.doorPos.Z() != frontTile.Z || l.grid[l.doorPos.X()][l.doorPos.Y()].State != frontTile.State {
			l.doorPos.z = frontTile.Z
			l.grid[l.doorPos.X()][l.doorPos.Y()].State = Open
		}
	}
}

// letterToHeight transform string to heightmap.
func letterToHeight(letter string) (int16, error) {
	if len(letter) != 1 {
		return 0, errors.New("input must be a single letter")
	}

	index := strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", strings.ToUpper(letter))
	if index == -1 {
		return 0, errors.New("invalid letter")
	}

	return int16(10 + index), nil
}

// NewLayout provides a clean layout
func NewLayout(hMap *model.HeightMap) (*Layout, error) {

	if hMap == nil {
		return nil, errors.New("layout not found")
	}

	door := NewCoordinate(int16(hMap.DoorX), int16(hMap.DoorY), 0, Direction(hMap.DoorDirection))
	l := &Layout{
		slug:    hMap.Slug,
		doorPos: &door,
		hMap:    strings.ReplaceAll(hMap.Heightmap, "\n", ""),
		size:    -1,
		xLen:    0,
		yLen:    0,
	}
	l.generateGrid()
	return l, nil
}
