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
	rl.DrawTextureEx(p.texture, rl.Vector2{X: float32(p.coord.x * TILE_SIZE), Y: float32(p.coord.y * TILE_SIZE)}, 0, TILE_SIZE/128., rl.White)
}
