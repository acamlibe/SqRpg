package game

import (
	"container/heap"
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

	path          []gridPos
	pathStep      int
	pathTargetRow int
	pathTargetCol int
}

func NewGame(rows, cols int) *Game {
	g := &Game{
		Grid:             NewGrid(rows, cols),
		Player:           &entities.Player{},
		moveCooldown:     0,
		moveCooldownTime: 0.2, // move every 0.25 seconds (4 moves/sec)
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
		g.path = nil
		g.pathStep = 0
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

	if g.targetRow == g.playerRow && g.targetCol == g.playerCol {
		g.path = nil
		g.pathStep = 0
		return
	}

	if g.path == nil || g.pathStep >= len(g.path) || g.pathTargetRow != g.targetRow || g.pathTargetCol != g.targetCol || !g.isPlayerOnPath() {
		g.path = g.findPathAStar(g.playerRow, g.playerCol, g.targetRow, g.targetCol)
		g.pathStep = 1
		g.pathTargetRow = g.targetRow
		g.pathTargetCol = g.targetCol
	}

	if g.path == nil || g.pathStep >= len(g.path) {
		return
	}

	next := g.path[g.pathStep]
	if g.movePlayerTo(next.Row, next.Col) {
		g.pathStep++
		g.moveCooldown = g.moveCooldownTime // reset cooldown
	}
}

func (g *Game) movePlayerTo(row, col int) bool {
	tileMap := g.Grid.Tiles
	tile := &tileMap[row][col]

	if !tile.Walkable {
		return false
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
	return true
}

func (g *Game) isPlayerOnPath() bool {
	if len(g.path) == 0 {
		return false
	}
	step := g.pathStep - 1
	if step < 0 || step >= len(g.path) {
		return false
	}
	return g.path[step].Row == g.playerRow && g.path[step].Col == g.playerCol
}

type gridPos struct {
	Row int
	Col int
}

type pathNode struct {
	Pos   gridPos
	GCost int
	FCost int
	Index int
}

type nodePriorityQueue []*pathNode

func (pq nodePriorityQueue) Len() int { return len(pq) }
func (pq nodePriorityQueue) Less(i, j int) bool {
	if pq[i].FCost == pq[j].FCost {
		return pq[i].GCost < pq[j].GCost
	}
	return pq[i].FCost < pq[j].FCost
}
func (pq nodePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
func (pq *nodePriorityQueue) Push(x interface{}) {
	n := x.(*pathNode)
	n.Index = len(*pq)
	*pq = append(*pq, n)
}
func (pq *nodePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (g *Game) findPathAStar(startRow, startCol, goalRow, goalCol int) []gridPos {
	if g.Grid == nil || g.Grid.Tiles == nil {
		return nil
	}
	rows := len(g.Grid.Tiles)
	if rows == 0 {
		return nil
	}
	cols := len(g.Grid.Tiles[0])
	if cols == 0 {
		return nil
	}

	if startRow == goalRow && startCol == goalCol {
		return []gridPos{{Row: startRow, Col: startCol}}
	}

	if goalRow < 0 || goalRow >= rows || goalCol < 0 || goalCol >= cols {
		return nil
	}
	if !g.Grid.Tiles[goalRow][goalCol].Walkable {
		return nil
	}

	start := gridPos{Row: startRow, Col: startCol}
	goal := gridPos{Row: goalRow, Col: goalCol}

	startKey := posKey(start.Row, start.Col, cols)
	goalKey := posKey(goal.Row, goal.Col, cols)

	open := &nodePriorityQueue{}
	heap.Init(open)
	startNode := &pathNode{Pos: start, GCost: 0, FCost: manhattan(start, goal)}
	heap.Push(open, startNode)

	cameFrom := make(map[int]int, rows*cols)
	gScore := map[int]int{startKey: 0}
	openNodes := map[int]*pathNode{startKey: startNode}

	for open.Len() > 0 {
		current := heap.Pop(open).(*pathNode)
		currentKey := posKey(current.Pos.Row, current.Pos.Col, cols)
		delete(openNodes, currentKey)

		if currentKey == goalKey {
			return reconstructPath(cameFrom, currentKey, startKey, cols)
		}

		neighbors := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range neighbors {
			nr := current.Pos.Row + d[0]
			nc := current.Pos.Col + d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if !g.Grid.Tiles[nr][nc].Walkable {
				continue
			}

			neighborKey := posKey(nr, nc, cols)
			candidateG := gScore[currentKey] + 1

			prevG, ok := gScore[neighborKey]
			if !ok || candidateG < prevG {
				cameFrom[neighborKey] = currentKey
				gScore[neighborKey] = candidateG
				f := candidateG + manhattan(gridPos{Row: nr, Col: nc}, goal)

				if existing, exists := openNodes[neighborKey]; exists {
					existing.GCost = candidateG
					existing.FCost = f
					heap.Fix(open, existing.Index)
				} else {
					n := &pathNode{Pos: gridPos{Row: nr, Col: nc}, GCost: candidateG, FCost: f}
					heap.Push(open, n)
					openNodes[neighborKey] = n
				}
			}
		}
	}

	return nil
}

func reconstructPath(cameFrom map[int]int, currentKey, startKey, cols int) []gridPos {
	path := []gridPos{keyToPos(currentKey, cols)}
	for currentKey != startKey {
		prev, ok := cameFrom[currentKey]
		if !ok {
			return nil
		}
		currentKey = prev
		path = append(path, keyToPos(currentKey, cols))
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func manhattan(a, b gridPos) int {
	return abs(a.Row-b.Row) + abs(a.Col-b.Col)
}

func posKey(row, col, cols int) int {
	return row*cols + col
}

func keyToPos(key, cols int) gridPos {
	return gridPos{Row: key / cols, Col: key % cols}
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

	g.targetRow = playerRow
	g.targetCol = playerCol

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
