package ui

import (
	"github.com/acamlibe/SqRpg/constants"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sidebar struct {
	Box rl.Rectangle
}

func NewSidebar() *Sidebar {
	return &Sidebar{
		Box: rl.Rectangle{
			X:      float32(constants.ScreenWidth - constants.SidebarWidth),
			Y:      0,
			Width:  float32(constants.SidebarWidth),
			Height: float32(constants.ScreenHeight),
		},
	}
}

func (s *Sidebar) Draw() {
	// Draw sidebar background
	rl.DrawRectangle(int32(s.Box.X), int32(s.Box.Y), int32(s.Box.Width), int32(s.Box.Height), constants.GunmetalColor)
	rl.DrawRectangle(int32(s.Box.X), int32(s.Box.Y+constants.GridPadding), int32(s.Box.Width-constants.GridPadding), int32(s.Box.Height-constants.GridPadding*2), constants.CharcoalGrayColor)

	rl.PushMatrix()
	rl.Translatef(s.Box.X, s.Box.Y+constants.GridPadding, 0)
	// Draw sidebar content relative to top-left
	rl.DrawText("HP: 100", 20, 20, 20, rl.White)
	rl.DrawText("Inventory:", 20, 60, 20, rl.White)
	// Example inventory slots
	slotSize := 40
	for i := range 5 {
		rl.DrawRectangleLines(20, int32(100+i*(slotSize+10)), int32(slotSize), int32(slotSize), rl.LightGray)
	}
	rl.PopMatrix()
}
