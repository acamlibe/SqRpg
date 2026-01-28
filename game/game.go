package game

import (
	"bufio"
	"os"
	"strconv"
	"strings"

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
		Player: &entities.Player{},
	}

	// g.World.AddEntity(g.Player, 10, 10)
	// g.World.AddEntity(&entities.Tree{}, 0, 0)

	return g
}

func (g *Game) Input() {
}

func (g *Game) Update() {
	dt := rl.GetFrameTime()
	g.Player.AnimTime += dt

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
		return [][]int{} // Return an empty 2D slice if dimensions don't match
	}

	// Create a 2D slice of appropriate size
	// a := make([][]int, m)
	// for i := range a {
	//     a[i] = make([]int, n)
	// }
	// A more direct way to create and initialize the structure:
	// Use a single underlying array for cache efficiency if needed,
	// but this slice of slices approach is more idiomatic and simpler.

	twoD := make([][]int, m)
	for i := 0; i < m; i++ {
		// Slice the original array to get the elements for the current row
		// The slice for row i goes from index i*n to (i+1)*n
		twoD[i] = original[i*n : (i+1)*n]
	}

	return twoD
}

func strArrToIntArr(original []string) []int {
	d := make([]int, len(original))

	for i := range original {
		num, _ := strconv.Atoi(original[i])
		d[i] = num
	}

	return d
}
