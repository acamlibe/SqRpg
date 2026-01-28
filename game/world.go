package game

import (
	"github.com/acamlibe/SqRpg/constants"
	"github.com/acamlibe/SqRpg/drawable"
	"github.com/acamlibe/SqRpg/game/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type World struct {
	Tiles   [][]Tile
	CameraX int
	CameraY int
}

type Tile struct {
	Entities []drawable.Drawable
	Walkable bool
}

func NewWorld(rows, cols int) *World {
	tileMap := make([][]Tile, rows)

	for y := range tileMap {
		tileMap[y] = make([]Tile, cols)

		for x := range tileMap[y] {
			tileMap[y][x] = Tile{Walkable: true}
		}
	}

	return &World{Tiles: tileMap}
}

func LoadWorld(data [][]int) *World {
	rows := len(data)
	cols := len(data[0])

	world := NewWorld(rows, cols)

	for row := range data {
		for col := range data[row] {
			world.Tiles[row][col] = getEntity(data[row][col])
		}
	}

	return world
}

func getEntity(n int) Tile {
	tile := Tile{}

	if n == 101 {
		tile.Entities = append(tile.Entities, &entities.Water{})
	}

	return tile
}

func (g *World) AddEntity(entity drawable.Drawable, row, col int) {
	g.Tiles[row][col].Entities = append(g.Tiles[row][col].Entities, entity)
}

func (g *World) DrawLocal() {
	for rowIdx, row := range g.Tiles {
		if rowIdx != g.CameraY {
			continue
		}

		for colIdx, tile := range row {
			if colIdx != g.CameraX {
				continue
			}

			g.drawTile(rowIdx, colIdx, &tile)
		}
	}
}

func (g *World) drawTile(row, col int, tile *Tile) {
	rl.PushMatrix()
	rl.Translatef(float32(col*constants.TileSize), float32(row*constants.TileSize), 0)

	tileX, tileY := 0, 0
	tileSize := constants.TileSize

	rl.DrawRectangle(int32(tileX), int32(tileY), int32(tileSize), int32(tileSize), rl.Black)

	for _, entity := range tile.Entities {
		entity.DrawLocal()
	}

	// Draw subtle outline
	rl.DrawRectangleLines(
		int32(tileX), int32(tileY),
		int32(tileSize), int32(tileSize),
		rl.Color{R: 60, G: 60, B: 60, A: 80}, // very subtle dark line, semi-transparent
	)

	rl.PopMatrix()
}
