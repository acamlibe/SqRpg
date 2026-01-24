package game

import (
	"math/rand"
	"slices"
	"time"

	"github.com/acamlibe/SqRpg/constants"
	"github.com/acamlibe/SqRpg/game/entities"
	"github.com/acamlibe/SqRpg/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Grid   *Grid
	Player *entities.Player

	playerTile *Tile
	playerRow  int
	playerCol  int

	moveCooldown     float32 // seconds until next allowed move
	moveCooldownTime float32 // max cooldown duration

	targetRow int
	targetCol int
}

func NewGame(rows, cols int) *Game {
	g := &Game{
		Grid:             NewGrid(rows, cols),
		Player:           &entities.Player{},
		moveCooldown:     0,
		moveCooldownTime: 0.25, // move every 0.25 seconds (4 moves/sec)
	}
	g.generateWorld()
	return g
}

func (g *Game) Input() {
	row := g.playerRow
	col := g.playerCol

	tileSize := constants.TileSize
	mouseX := int(rl.GetMouseX())
	mouseY := int(rl.GetMouseY())

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		col = (mouseX - constants.GridPadding) / tileSize
		row = (mouseY - constants.GridPadding) / tileSize

		row = utils.IntMin(utils.IntMax(row, 0), len(g.Grid.Tiles)-1)
		col = utils.IntMin(utils.IntMax(col, 0), len(g.Grid.Tiles[0])-1)

		g.targetRow = row
		g.targetCol = col
	}
}

func (g *Game) Update() {
	dt := rl.GetFrameTime()
	g.Player.AnimTime += dt

	if g.moveCooldown > 0 {
		g.moveCooldown -= dt
	}

	if g.moveCooldown > 0 {
		return // too soon to move
	}

	nextRow := g.playerRow
	nextCol := g.playerCol

	if g.targetRow != nextRow || g.targetCol != nextCol {
		if nextRow < g.targetRow {
			nextRow++
		} else if nextRow > g.targetRow {
			nextRow--
		} else if nextCol < g.targetCol {
			nextCol++
		} else if nextCol > g.targetCol {
			nextCol--
		}

		g.movePlayerTo(nextRow, nextCol)
		g.moveCooldown = g.moveCooldownTime // reset cooldown
	}
}

func (g *Game) movePlayerTo(row, col int) {
	tileMap := g.Grid.Tiles
	tile := &tileMap[row][col]

	if !tile.Walkable {
		return
	}

	g.playerRow = row
	g.playerCol = col

	if g.playerTile != nil && g.playerTile.Entities != nil {
		for i, entity := range g.playerTile.Entities {
			if entity == g.Player {
				g.playerTile.Entities = slices.Delete(g.playerTile.Entities, i, i+1)
				break
			}
		}
	}

	tile.Entities = append(tile.Entities, g.Player)

	g.playerTile = tile
}

func (g *Game) generateWorld() {
	if g.Grid == nil || g.Grid.Tiles == nil {
		return
	}

	tiles := g.Grid.Tiles
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
	for y := range rows {
		for x := range cols {
			tiles[y][x].Entities = nil
		}
	}

	// Place player randomly
	playerRow := rng.Intn(rows)
	playerCol := rng.Intn(cols)

	g.playerRow = playerRow
	g.playerCol = playerCol

	g.movePlayerTo(playerRow, playerCol)

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
				tiles[y][x].Walkable = false
				tiles[y][x].Entities = append(tiles[y][x].Entities, &entities.Tree{})
			}
		}
	}

	// Forest seeds + random walks for clustered, procedural feel
	seedCount := max(3, (rows*cols)/150)
	baseWalk := max(10, (rows*cols)/40)

	for range seedCount {
		sy, sx := rng.Intn(rows), rng.Intn(cols)
		walkLen := baseWalk + rng.Intn(baseWalk/2+1)

		for range walkLen {
			if tiles[sy][sx].Entities == nil {
				tiles[sy][sx].Walkable = false
				tiles[sy][sx].Entities = append(tiles[sy][sx].Entities, &entities.Tree{})
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
	for y := range rows {
		for x := range cols {
			if y == playerRow && x == playerCol {
				continue
			}
			if tiles[y][x].Entities == nil {
				neighbors := countTreeNeighbors(tiles, y, x)
				if neighbors >= 4 && rng.Float32() < 0.6 {
					tiles[y][x].Walkable = false
					tiles[y][x].Entities = append(tiles[y][x].Entities, &entities.Tree{})
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

			for _, entity := range tiles[ny][nx].Entities {
				if _, ok := entity.(*entities.Tree); ok {
					count++
				}
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
