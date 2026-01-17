package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
}

func (p *Player) Draw(x1, y1, x2, y2 int) {
	rl.DrawRectangle(int32(x1), int32(y1), int32(x2-x1), int32(y2-y1), rl.Red)
}
