package game

import (
	"github.com/acamlibe/SqRpg/constants"
	"github.com/acamlibe/SqRpg/game/entities"
	"github.com/acamlibe/SqRpg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	Tiles [][]Tile
}

type Tile struct {
	Entity Drawable
}

type Drawable interface {
	DrawLocal()
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

func (g *Grid) Draw() {
	for rowIdx, row := range g.Tiles {
		for colIdx, tile := range row {
			if tile.Entity != nil {
				g.drawTile(rowIdx, colIdx, tile)
			}
		}
	}
}

func (g *Grid) drawTile(row, col int, tile Tile) {
	rl.PushMatrix()
	rl.Translatef(float32(col*constants.TileSize), float32(row*constants.TileSize), 0)
	tile.Entity.DrawLocal()
	rl.PopMatrix()
}
