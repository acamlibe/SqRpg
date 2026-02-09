package tiles

import "github.com/acamlibe/SqRpg/drawable"

type Tile interface {
	drawable.Drawable
	Walkable() bool
}
