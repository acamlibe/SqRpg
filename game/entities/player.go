package entities

import (
	"math"

	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	AnimTime float32
}

func (p *Player) DrawLocal() {
	var tile int32 = constants.TileSize

	bob := float32(math.Sin(float64(p.AnimTime)*3.0)) * 1.2
	armSwing := float32(math.Sin(float64(p.AnimTime)*4.0)) * 1.1

	var headX int32 = tile / 2
	var headY int32 = tile/4 + int32(bob)
	var bodyX int32 = tile / 2
	var bodyY int32 = headY + tile/8

	// head
	rl.DrawCircleLines(headX, headY, float32(tile/8), rl.SkyBlue)
	// body
	rl.DrawLine(bodyX, bodyY, bodyX, tile-tile/6, rl.SkyBlue)

	// arms
	rl.DrawLine(bodyX, tile-tile/3, tile/4+int32(armSwing), tile-tile/2, rl.SkyBlue)
	rl.DrawLine(bodyX, tile-tile/3, tile-tile/4+int32(armSwing), tile-tile/2, rl.SkyBlue)

	// legs
	rl.DrawLine(bodyX, tile-tile/6, tile/4, tile, rl.SkyBlue)
	rl.DrawLine(bodyX, tile-tile/6, tile-tile/4, tile, rl.SkyBlue)
}
