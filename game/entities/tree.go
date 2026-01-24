package entities

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tree struct {
}

func (t *Tree) DrawLocal() {
	rl.DrawRectangle(0, 0, constants.TileSize, constants.TileSize, rl.Green)
	rl.DrawRectangle(constants.TileSize/2, 0, 2, constants.TileSize, rl.Brown)
}
