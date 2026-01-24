package entities

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tree struct {
}

func (t *Tree) DrawLocal() {
	var tile int32 = constants.TileSize

	var trunkX int32 = tile/2 - tile/6
	var trunkY int32 = tile - tile/2

	var foliageX int32 = tile / 2
	var foliageY int32 = tile / 2

	// trunk
	trunkWidth := tile / 4
	rl.DrawRectangle(trunkX, trunkY, trunkWidth, tile/2, rl.DarkBrown)

	// trunk lines
	darkerDarkBrown := rl.ColorBrightness(rl.DarkBrown, -1)
	rl.DrawLine(trunkX+trunkWidth/4, trunkY+tile/4, trunkX+trunkWidth/4, tile, darkerDarkBrown)
	rl.DrawLine(trunkX+trunkWidth/2, trunkY+tile/4, trunkX+trunkWidth/2, tile, darkerDarkBrown)
	rl.DrawLine(trunkX+(trunkWidth/2+trunkWidth/4), trunkY+tile/4, trunkX+(trunkWidth/2+trunkWidth/4), tile, darkerDarkBrown)

	// foliage
	rl.DrawCircle(foliageX, foliageY-tile/6, float32(tile/5), rl.Lime)
	rl.DrawCircle(foliageX-tile/6, foliageY+tile/16, float32(tile/5), rl.Lime)
	rl.DrawCircle(foliageX+tile/10, foliageY+tile/18, float32(tile/5), rl.Lime)

	// apple
	rl.DrawCircle(foliageX+tile/16, foliageY, float32(tile/16), rl.Red)
	rl.DrawCircle(foliageX-tile/16, foliageY-tile/8, float32(tile/16), rl.Red)
	rl.DrawCircle(foliageX-tile/8, foliageY+tile/8, float32(tile/16), rl.Red)
}
