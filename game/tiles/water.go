package tiles

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Water struct {
}

func (t *Water) Walkable() bool {
	return false
}

func (w *Water) DrawLocal() {
	tile := int32(constants.TileSize)

	// Base water color
	rl.DrawRectangle(0, 0, tile, tile, rl.Color{
		R: 70, G: 130, B: 180, A: 255,
	})

	// Soft wave highlights
	for y := int32(3); y < tile; y += 6 {
		rl.DrawLine(
			2, y,
			tile-3, y,
			rl.Color{R: 180, G: 220, B: 255, A: 50},
		)
	}
}
