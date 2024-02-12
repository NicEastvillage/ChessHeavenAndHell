package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	NamePawn          = "Pawn"
	NameKnight        = "Knight"
	NameBishop        = "Bishop"
	NameRook          = "Rook"
	NameQueen         = "Queen"
	NameKing          = "King"
	NameBomber        = "Suicide Bomber"
	NameLeopard       = "Leopard"
	NameChecker       = "Checker"
	NameMountedArcher = "Mounted Archer"
	NameWizard        = "Wizard"
	NameArchbishop    = "Archbishop"
	NameFortress      = "Fortress"
	NameScout         = "Scout"
	NameWarlock       = "Warlock"
	NameCelestial     = "Celestial"
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
	scale uint32
}

func (p *Piece) Render(selection *Selection) {
	var typ = sandbox.GetPieceType(p.typ)
	var tex = typ.texWhite
	if p.color == BLACK {
		tex = typ.texBlack
	}

	var pos = GetWorldOrigo().Add(p.coord.Scale(TileSize))
	var texScale = float32(TileSize) / float32(max(tex.Width, tex.Height))
	var pieceCorner = pos.Add(ONEONE.Scale(TileSize / 2)).Sub(Vec2{int(texScale * float32(tex.Width) / 2), int(texScale * float32(tex.Height) / 2)})
	rl.DrawTextureEx(tex, pieceCorner.ToRlVec(), 0, float32(p.scale)*texScale, rl.White)

	if selection.IsPieceSelected(p.id) {
		rl.DrawRectangleLines(int32(pos.x)+4, int32(pos.y)+4, int32(TileSize*p.scale-8), int32(TileSize*p.scale-8), rl.Blue)
	}
}
