package tiles

type Air struct {
}

func (a *Air) Walkable() bool {
	return true
}

func (a *Air) DrawLocal() {

}
