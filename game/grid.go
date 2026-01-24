package game

import (
	"github.com/acamlibe/SqRpg/constants"
	"github.com/acamlibe/SqRpg/drawable"
	"github.com/acamlibe/SqRpg/game/entities"
	"github.com/acamlibe/SqRpg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	Tiles [][]Tile
}

type Tile struct {
	Entity drawable.Drawable
}

func NewGrid(rows, cols int) Grid {
	tileMap := make([][]Tile, rows)

	for y := range tileMap {
		tileMap[y] = make([]Tile, cols)

		for x := range tileMap[y] {
			t := Tile{}

			if utils.RandChance(10) {
				t.Entity = entities.Tree{}
			}

			tileMap[y][x] = t
		}
	}

	return Grid{Tiles: tileMap}
}

func (g *Grid) DrawLocal() {
	for rowIdx, row := range g.Tiles {
		for colIdx, tile := range row {
			g.drawTile(rowIdx, colIdx, tile)
		}
	}
}

func (g *Grid) drawTile(row, col int, tile Tile) {
	rl.PushMatrix()
	rl.Translatef(float32(col*constants.TileSize), float32(row*constants.TileSize), 0)

	if tile.Entity != nil {
		tile.Entity.DrawLocal()
	} else {
		// Empty space
		rl.DrawRectangle(0, 0, constants.TileSize, constants.TileSize, rl.Black)
	}

	rl.PopMatrix()
}
