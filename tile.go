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
	rl.DrawRectangle(int32(t.coord.x*TILE_SIZE), int32(t.coord.y*TILE_SIZE), TILE_SIZE, TILE_SIZE, color)
}
