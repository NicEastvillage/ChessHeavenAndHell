package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Piece struct {
	id      uint32
	board   uint32
	coord   Vec2
	texture rl.Texture2D
}

func (p *Piece) Render() {
	var pos = GetWorldOrigo().Add(p.coord.Scale(TileSize))
	rl.DrawTextureEx(p.texture, pos.ToRlVec(), 0, TileSize/128., rl.White)
}