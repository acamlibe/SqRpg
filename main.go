package main

import (
	"github.com/acamlibe/SqRpg/constants"
	"github.com/acamlibe/SqRpg/game"
	"github.com/acamlibe/SqRpg/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(constants.ScreenWidth, constants.ScreenHeight, "SqRpg")
	rl.SetTargetFPS(constants.FPS)
	defer rl.CloseWindow()

	gameWindow := ui.NewGameWindow()
	sidebar := ui.NewSidebar()

	game := game.NewGame(int((gameWindow.Box.Height-constants.GridPadding*2)/constants.TileSize), int((gameWindow.Box.Width-constants.GridPadding*2)/constants.TileSize))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(constants.BackgroundColor)

		gameWindow.Draw(game.Grid)
		sidebar.Draw()

		rl.EndDrawing()
	}
}
