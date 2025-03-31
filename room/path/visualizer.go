package path

import (
	"fmt"
	"strings"
)

// GlyphMap returns an ASCII representation of the layout with the given path marked.
// The start tile (first in path) is marked with 'S' and the target tile (last in path) with 'E'.
// Other tiles in the path are marked with '*'. Open tiles can display height or remain 'O' if flatMode is true.
func GlyphMap(l *Layout, path []*Tile, flatMode bool) string {
	pathSet := make(map[*Tile]bool)
	for _, t := range path {
		pathSet[t] = true
	}

	var startTile, targetTile *Tile
	if len(path) > 0 {
		startTile = path[0]
		targetTile = path[len(path)-1]
	}

	var sb strings.Builder
	for y := 0; y < l.yLen; y++ {
		for x := 0; x < l.xLen; x++ {
			tile := l.grid[x][y]
			if tile == startTile {
				sb.WriteString("S")
			} else if tile == targetTile {
				sb.WriteString("E")
			} else if pathSet[tile] {
				sb.WriteString("*")
			} else {
				switch tile.State {
				case Open:
					if flatMode {
						sb.WriteString("O")
					} else {
						fmt.Println(tile.height)
						sb.WriteString(fmt.Sprintf("%d", tile.Height()))
					}
				case Blocked:
					sb.WriteString("#")
				case Invalid:
					sb.WriteString("X")
				case Sit:
					sb.WriteString("T")
				case Lay:
					sb.WriteString("L")
				default:
					sb.WriteString("?")
				}
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
