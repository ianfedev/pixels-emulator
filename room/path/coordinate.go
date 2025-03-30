package path

// Direction represents the possible movement directions in an 8-way grid.
type Direction int16

const (
	North     Direction = 0
	NorthEast Direction = 1
	East      Direction = 2
	SouthEast Direction = 3
	South     Direction = 4
	SouthWest Direction = 5
	West      Direction = 6
	NorthWest Direction = 7
)

// Coordinate represents a position in the grid with immutable X, Y, Z coordinates and a movement direction.
type Coordinate struct {
	x, y, z int16     // Grid coordinates (private for immutability)
	dir     Direction // Movement direction (private for immutability)
}

// NewCoordinate creates a new immutable Coordinate instance.
func NewCoordinate(x, y, z int16, dir Direction) Coordinate {
	return Coordinate{x: x, y: y, z: z, dir: dir}
}

// X returns the X coordinate.
func (c Coordinate) X() int16 {
	return c.x
}

// Y returns the Y coordinate.
func (c Coordinate) Y() int16 {
	return c.y
}

// Z returns the Z coordinate.
func (c Coordinate) Z() int16 {
	return c.z
}

// Dir returns the movement direction.
func (c Coordinate) Dir() Direction {
	return c.dir
}
