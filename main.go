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

	game := game.NewGame(13, 13)

	for !rl.WindowShouldClose() {
		game.Input()
		game.Update()

		rl.BeginDrawing()
		rl.ClearBackground(constants.BackgroundColor)

		gameWindow.Draw(game)
		sidebar.Draw()

		rl.EndDrawing()
	}
}
