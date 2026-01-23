package entities

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
}

func (p *Player) DrawLocal() {
	rl.DrawRectangle(0, 0, constants.TileSize, constants.TileSize, rl.Red)
}
