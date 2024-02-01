package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Piece struct {
	coord   Vec2
	texture rl.Texture2D
}

func NewPiece(x int, y int, texture rl.Texture2D) Piece {
	return Piece{coord: Vec2{x: x, y: y}, texture: texture}
}

func (p *Piece) Render() {
	var pos = GetWorldOrigo().Add(p.coord.Scale(TILE_SIZE))
	rl.DrawTextureEx(p.texture, pos.ToRlVec(), 0, TILE_SIZE/128., rl.White)
}
