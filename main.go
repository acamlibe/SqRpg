package main

import (
	"github.com/acamlibe/SqRpg/constants"
	"github.com/acamlibe/SqRpg/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(constants.ScreenWidth, constants.ScreenHeight, "Pong")
	rl.SetTargetFPS(constants.FPS)
	defer rl.CloseWindow()

	game := game.NewGame()
	game.Run()
}
