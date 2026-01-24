package entities

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
}

func (p Player) DrawLocal() {
	rl.DrawCircle(constants.TileSize/2, constants.TileSize/2, constants.TileSize/2, rl.Blue)
	rl.DrawText("P", constants.TileSize/2-(rl.MeasureText("P", 10)/2), constants.TileSize/2-(rl.MeasureText("P", 10)/2), 10, rl.White)
}
