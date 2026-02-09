package game

import (
	"github.com/acamlibe/SqRpg/drawable"
	"github.com/acamlibe/SqRpg/game/tiles"
)

type World struct {
	Tiles    [][]tiles.Tile
	Entities [][]drawable.Drawable
}

func LoadWorld(data [][]int) *World {
	rows := len(data)
	cols := len(data[0])

	tileMap := make([][]tiles.Tile, rows)

	for y := range tileMap {
		tileMap[y] = make([]tiles.Tile, cols)

		for x := range tileMap[y] {
			tileMap[y][x] = getTile(data[y][x])
		}
	}

	return &World{Tiles: tileMap}
}

func getTile(n int) tiles.Tile {
	switch n {
	case 1:
		return &tiles.Water{}
	case 101:
		return &tiles.Tree{}
	}

	return &tiles.Air{}
}
