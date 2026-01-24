package entities

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	trunkWidth    = 10
	trunkHeight   = 10
	foliageRadius = 6
)

type Tree struct {
}

func (t *Tree) DrawLocal() {
	var tile int32 = constants.TileSize

	var trunkX int32 = tile/2 - trunkWidth/2
	var trunkY int32 = tile - trunkHeight

	var foliageX int32 = tile / 2
	var foliageY int32 = tile / 2

	//rl.DrawRectangle(0, 0, tile, tile, rl.DarkGreen)
	rl.DrawRectangle(trunkX, trunkY, trunkWidth, trunkHeight, rl.DarkBrown)
	//rl.DrawCircle(foliageX, foliageY, foliageRadius, rl.Lime)

	rl.DrawCircle(foliageX+3, foliageY, 4, rl.Lime)
	rl.DrawCircle(foliageX-2, foliageY, 4, rl.Lime)
	rl.DrawCircle(foliageX, foliageY-4, 4, rl.Lime)
}
