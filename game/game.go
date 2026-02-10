package game

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/acamlibe/SqRpg/constants"
	"github.com/acamlibe/SqRpg/drawable"
	"github.com/acamlibe/SqRpg/game/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	World  *World
	Camera *Camera
	Player *entities.Player
}

func NewGame(rows, cols int) *Game {
	w, _ := loadMap()
	g := &Game{
		World:  LoadWorld(w),
		Camera: NewCamera(rows, cols),
		Player: &entities.Player{
			X: 6,
			Y: 8,
		},
	}

	g.movePlayer(g.Player.Y, g.Player.X)

	return g
}

func (g *Game) Input() {
	if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) {
		g.movePlayer(g.Player.Y, g.Player.X+1)
	}
	if rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) {
		g.movePlayer(g.Player.Y+1, g.Player.X)
	}
	if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) {
		g.movePlayer(g.Player.Y, g.Player.X-1)
	}
	if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
		g.movePlayer(g.Player.Y-1, g.Player.X)
	}
}

func (g *Game) Update() {
	dt := rl.GetFrameTime()
	g.Player.Update(dt)

}

func (g *Game) DrawLocal() {
	for rowIdx, row := range g.World.Tiles[g.Camera.WorldY : g.Camera.WorldY+g.Camera.Rows] {
		for colIdx, tile := range row[g.Camera.WorldX : g.Camera.WorldX+g.Camera.Cols] {
			worldRow := g.Camera.WorldY + rowIdx
			worldCol := g.Camera.WorldX + colIdx

			g.draw(rowIdx, colIdx, true, tile)

			if g.Player.X == worldCol && g.Player.Y == worldRow {
				g.draw(rowIdx, colIdx, false, g.Player)
			}
		}
	}

}

func (g *Game) movePlayer(row, col int) {
	worldRows := len(g.World.Tiles)
	worldCols := len(g.World.Tiles[0])

	// bounds check
	if row < 0 || col < 0 || row >= worldRows || col >= worldCols {
		return
	}

	// walkable check
	if !g.World.Tiles[row][col].Walkable() {
		return
	}

	// move player
	g.Player.Y = row
	g.Player.X = col

	// clamp camera to world bounds
	maxCamX := worldCols - g.Camera.Cols
	maxCamY := worldRows - g.Camera.Rows

	// center camera on player
	camX := min(max(g.Player.X-g.Camera.Cols/2, 0), maxCamX)
	camY := min(max(g.Player.Y-g.Camera.Rows/2, 0), maxCamY)

	g.Camera.WorldX = camX
	g.Camera.WorldY = camY
}

func (g *Game) draw(row, col int, clearBackground bool, drawable drawable.Drawable) {
	if drawable == nil {
		return
	}

	rl.PushMatrix()
	rl.Translatef(float32(col*constants.TileSize), float32(row*constants.TileSize), 0)

	tileX, tileY := 0, 0
	tileSize := constants.TileSize

	if clearBackground {
		rl.DrawRectangle(int32(tileX), int32(tileY), int32(tileSize), int32(tileSize), rl.Black)
	}

	drawable.DrawLocal()

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
	for i := range m {
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
