package game

import (
	"math/rand"
	"time"

	"github.com/acamlibe/SqRpg/game/entities"
)

type Game struct {
	Grid   *Grid
	Player *entities.Player
}

func NewGame(rows, cols int) *Game {
	g := &Game{Grid: NewGrid(rows, cols), Player: &entities.Player{}}
	g.generateWorld()
	return g
}

func (g *Game) Input() {

}

func (g *Game) Update() {

}

func (g *Game) generateWorld() {
	if g.Grid == nil || g.Grid.Tiles == nil {
		return
	}

	tiles := *g.Grid.Tiles
	rows := len(tiles)
	if rows == 0 {
		return
	}
	cols := len(tiles[0])
	if cols == 0 {
		return
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Clear any existing entities
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			tiles[y][x].Entity = nil
		}
	}

	// Place player randomly
	playerRow := rng.Intn(rows)
	playerCol := rng.Intn(cols)
	tiles[playerRow][playerCol].Entity = g.Player

	// Small grove around player (but keep immediate neighbors open)
	for y := playerRow - 2; y <= playerRow+2; y++ {
		for x := playerCol - 2; x <= playerCol+2; x++ {
			if y < 0 || y >= rows || x < 0 || x >= cols {
				continue
			}
			if y == playerRow && x == playerCol {
				continue
			}
			if abs(y-playerRow) <= 1 && abs(x-playerCol) <= 1 {
				continue
			}
			if rng.Float32() < 0.4 {
				tiles[y][x].Entity = &entities.Tree{}
			}
		}
	}

	// Forest seeds + random walks for clustered, procedural feel
	seedCount := max(3, (rows*cols)/150)
	baseWalk := max(10, (rows*cols)/40)

	for i := 0; i < seedCount; i++ {
		sy, sx := rng.Intn(rows), rng.Intn(cols)
		walkLen := baseWalk + rng.Intn(baseWalk/2+1)

		for step := 0; step < walkLen; step++ {
			if tiles[sy][sx].Entity == nil {
				tiles[sy][sx].Entity = &entities.Tree{}
			}

			switch rng.Intn(4) {
			case 0:
				sx++
			case 1:
				sx--
			case 2:
				sy++
			case 3:
				sy--
			}

			if sy < 0 {
				sy = 0
			}
			if sy >= rows {
				sy = rows - 1
			}
			if sx < 0 {
				sx = 0
			}
			if sx >= cols {
				sx = cols - 1
			}
		}
	}

	// Simple smoothing pass to make blobs feel more organic
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if y == playerRow && x == playerCol {
				continue
			}
			if tiles[y][x].Entity == nil {
				neighbors := countTreeNeighbors(tiles, y, x)
				if neighbors >= 4 && rng.Float32() < 0.6 {
					tiles[y][x].Entity = &entities.Tree{}
				}
			}
		}
	}
}

func countTreeNeighbors(tiles [][]Tile, row, col int) int {
	rows := len(tiles)
	cols := len(tiles[0])
	count := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			ny := row + dy
			nx := col + dx
			if ny < 0 || ny >= rows || nx < 0 || nx >= cols {
				continue
			}
			if _, ok := tiles[ny][nx].Entity.(*entities.Tree); ok {
				count++
			}
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
