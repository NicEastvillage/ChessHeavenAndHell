package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	coord Vec2
}

func (t *Tile) Render() {
	var color = rl.Beige
	if t.coord.x%2 != t.coord.y%2 {
		color = rl.Brown
	}
	var pos = GetWorldOrigo().Add(t.coord.Scale(TILE_SIZE))
	rl.DrawRectangle(int32(pos.x), int32(pos.y), TILE_SIZE, TILE_SIZE, color)
}
