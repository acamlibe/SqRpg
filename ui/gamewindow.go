package ui

import (
	"github.com/acamlibe/SqRpg/constants"
	"github.com/acamlibe/SqRpg/drawable"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameWindow struct {
	Box rl.Rectangle
}

func NewGameWindow() GameWindow {
	return GameWindow{
		Box: rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  float32(constants.ScreenWidth - constants.SidebarWidth),
			Height: float32(constants.ScreenHeight),
		},
	}
}

func (w *GameWindow) Draw(drawable drawable.Drawable) {
	// Draw background for the window
	rl.DrawRectangle(int32(w.Box.X), int32(w.Box.Y), int32(w.Box.Width), int32(w.Box.Height), constants.GunmetalColor)

	rl.PushMatrix()
	rl.Translatef(w.Box.X+constants.GridPadding, w.Box.Y+constants.GridPadding, 0)
	drawable.DrawLocal()
	rl.PopMatrix()
}
