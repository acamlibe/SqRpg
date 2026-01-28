package entities

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Water struct {
}

func (t *Water) DrawLocal() {
	var tile int32 = constants.TileSize

	rl.DrawRectangle(0, 0, tile, tile, rl.SkyBlue)
}
