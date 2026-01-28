package game

type Camera struct {
	Rows   int
	Cols   int
	WorldX int
	WorldY int
}

func NewCamera(rows, cols int) *Camera {
	return &Camera{Rows: rows, Cols: cols}
}
