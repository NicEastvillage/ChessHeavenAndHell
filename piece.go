package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	PAWN   = "Pawn"
	KNIGHT = "Knight"
	BISHOP = "Bishop"
	ROOK   = "Rook"
	QUEEN  = "Queen"
	KING   = "King"
)

type PieceColor uint32

const (
	WHITE PieceColor = iota
	BLACK
)

type PieceType struct {
	id       uint32
	name     string
	texWhite rl.Texture2D
	texBlack rl.Texture2D
}

type Piece struct {
	id    uint32
	typ   uint32
	color PieceColor
	board uint32
	coord Vec2
}

func (p *Piece) Render() {
	var typ = sandbox.GetPieceType(p.typ)
	var tex = typ.texWhite
	if p.color == BLACK {
		tex = typ.texBlack
	}

	var pos = GetWorldOrigo().Add(p.coord.Scale(TileSize))
	rl.DrawTextureEx(tex, pos.ToRlVec(), 0, TileSize/128., rl.White)

	if selection.IsPieceSelected(p.id) {
		rl.DrawRectangleLines(int32(pos.x)+4, int32(pos.y)+4, TileSize-8, TileSize-8, rl.Blue)
	}
}
