package encode

// RoomDataMask holds the constants for bitmask values.
type RoomDataMask struct{}

// Define bitmask values for different properties in RoomData.
const (
	Thumbnail         = 1 << iota // 1
	Guild                         // 2
	Promotion                     // 4
	Owner                         // 8
	Pets                          // 16
	FeaturedPromotion             // 32
)

// Enable sets the given flag in the mask.
func Enable(mask, flag int) int {
	return mask | flag
}

// Has checks if the given flag is set in the mask.
func Has(mask, flag int) bool {
	return mask&flag != 0
}
