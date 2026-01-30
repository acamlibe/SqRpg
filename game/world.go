package game

import (
	"github.com/acamlibe/SqRpg/drawable"
	"github.com/acamlibe/SqRpg/game/entities"
)

type World struct {
	Tiles [][]Tile
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

	switch n {
	case 1:
		tile.Walkable = false
		tile.Entities = append(tile.Entities, &entities.Water{})
	case 101:
		tile.Walkable = false
		tile.Entities = append(tile.Entities, &entities.Tree{})
	}

	return tile
}

func (g *World) AddEntity(entity drawable.Drawable, row, col int) {
	g.Tiles[row][col].Entities = append(g.Tiles[row][col].Entities, entity)
}
