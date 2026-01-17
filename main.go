package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	Tiles [][]Tile
}

type Tile struct {
	X int
	Y int

	Entity Drawable
}

type Drawable interface {
	Draw(x1, y1, x2, y2 int)
}

const (
	screenWidth  = 1280
	screenHeight = 900

	fps = 60

	tileSize = 20
)

var (
	running = rl.WindowShouldClose()
)

var (
	backgroundColor = rl.Black
)

var (
	player Player
)

var grid [][]Tile

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Pong")
	rl.SetTargetFPS(fps)

	width := (screenWidth - 400) / tileSize
	height := screenHeight / tileSize

	grid := make([][]Tile, height)

	for y := range grid {
		grid[y] = make([]Tile, width)

		for x := range grid[y] {
			t := Tile{X: x, Y: y}

			if RandBool() {
				t.Entity = Tree{}
			}

			grid[y][x] = t
		}
	}
}

func input() {

}

func update() {
	running = !rl.WindowShouldClose()

}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(backgroundColor)
	draw()

	rl.EndDrawing()
}

func draw() {
	drawUI()
	drawGame()
}

func drawUI() {

}

func drawGame() {

}

func quit() {

}

func main() {
	defer rl.CloseWindow()

	for running {
		input()
		update()
		render()
	}

	quit()
}
