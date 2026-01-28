package game

import (
	"github.com/acamlibe/SqRpg/game/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Grid   *Grid
	Player *entities.Player
}

func NewGame(rows, cols int) *Game {
	g := &Game{
		Grid:   NewGrid(rows, cols),
		Player: &entities.Player{},
	}

	g.Grid.AddEntity(g.Player, 10, 10)
	g.Grid.AddEntity(&entities.Tree{}, 0, 0)

	return g
}

func (g *Game) Input() {
}

func (g *Game) Update() {
	dt := rl.GetFrameTime()
	g.Player.AnimTime += dt

}
