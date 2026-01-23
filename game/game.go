package game

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Grid Grid
}

func NewGame() Game {
	width := (constants.ScreenWidth - 400) / constants.TileSize
	height := constants.ScreenHeight / constants.TileSize

	return Game{Grid: NewGrid(width, height)}
}

func (g *Game) Run() {
	for !rl.WindowShouldClose() {
		g.input()
		g.update()
		g.render()
	}
}

func (g *Game) input() {

}

func (g *Game) update() {

}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(constants.BackgroundColor)
	g.draw()
	rl.EndDrawing()
}

func (g *Game) draw() {
	g.Grid.Draw()
}
