package game

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/acamlibe/SqRpg/constants"
	"github.com/acamlibe/SqRpg/game/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	World  *World
	Camera *Camera
	Player *entities.Player

	playerX int
	playerY int
}

func NewGame(rows, cols int) *Game {
	w, _ := loadMap()
	g := &Game{
		World:  LoadWorld(w),
		Camera: NewCamera(rows, cols),
		Player: &entities.Player{},

		playerX: 6,
		playerY: 6,
	}

	g.movePlayer(g.playerY, g.playerX, true)

	return g
}

func (g *Game) Input() {
	if rl.IsKeyPressed(rl.KeyD) {
		g.playerX++
		g.movePlayer(g.playerY, g.playerX, false)
	}
}

func (g *Game) Update() {
	dt := rl.GetFrameTime()
	g.Player.AnimTime += dt

}

func (g *Game) DrawLocal() {
	for rowIdx, row := range g.World.Tiles[g.Camera.WorldY : g.Camera.WorldY+g.Camera.Rows] {
		for colIdx, tile := range row[g.Camera.WorldX : g.Camera.WorldX+g.Camera.Cols] {
			g.drawTile(rowIdx, colIdx, &tile)
		}
	}
}

func (g *Game) movePlayer(row, col int, initial bool) {
	if row < 0 || col < 0 || row >= len(g.World.Tiles) || col >= len(g.World.Tiles[row]) {
		return
	}

	if !initial && !g.World.Tiles[row][col].Walkable {
		return
	}

	if !initial {
		entities := g.World.Tiles[g.playerY][g.playerX].Entities
		filtered := entities[:0]
		for _, e := range entities {
			if e != g.Player {
				filtered = append(filtered, e)
			}
		}
		g.World.Tiles[g.playerY][g.playerX].Entities = filtered
	}

	g.playerY = row
	g.playerX = col

	g.World.Tiles[g.playerY][g.playerX].Entities = append(g.World.Tiles[g.playerY][g.playerX].Entities, g.Player)
}

func (g *Game) drawTile(row, col int, tile *Tile) {
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

func loadMap() ([][]int, error) {
	// 1. Open the file
	file, err := os.Open("map.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close() // Ensure the file is closed

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	dataStr := strings.Split(line, ",")

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	data := strArrToIntArr(dataStr)

	return convert1Dto2D(data, 100, 100), nil
}

func convert1Dto2D(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	twoD := make([][]int, m)
	for i := 0; i < m; i++ {
		twoD[i] = original[i*n : (i+1)*n]
	}

	return twoD
}

func strArrToIntArr(original []string) []int {
	d := make([]int, len(original))

	for i := range original {
		num, _ := strconv.Atoi(strings.TrimSpace(original[i]))
		d[i] = num
	}

	return d
}
