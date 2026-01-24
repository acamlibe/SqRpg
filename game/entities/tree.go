package entities

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tree struct {
}

func (t *Tree) DrawLocal() {
	var tile int32 = constants.TileSize

	rl.DrawRectangle(0, 0, tile, tile, rl.DarkGreen)
	rl.DrawRectangle(tile/2-2, 2, 4, tile-2, rl.DarkBrown)
}
