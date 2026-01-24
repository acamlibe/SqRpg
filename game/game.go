package game

type Game struct {
	Grid *Grid
}

func NewGame(rows, cols int) *Game {
	return &Game{Grid: NewGrid(rows, cols)}
}

func (g *Game) Input() {

}

func (g *Game) Update() {

}
