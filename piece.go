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
	Id       uint32
	Name     string
	TexWhite rl.Texture2D
	TexBlack rl.Texture2D
}

type Piece struct {
	Id    uint32
	Typ   uint32
	Color PieceColor
	Board uint32
	Coord Vec2
	Scale uint32
}

func (p *Piece) Render() {
	var typ = sandbox.GetPieceType(p.Typ)
	var tex = typ.TexWhite
	if p.Color == BLACK {
		tex = typ.TexBlack
	}

	var pos = GetBoardOrigo().Add(p.Coord.Scale(TileSize))
	var texScale = float32(TileSize) / float32(max(tex.Width, tex.Height))
	var pieceCorner = pos.Add(ONEONE.Scale(TileSize / 2)).Sub(Vec2{int(texScale * float32(tex.Width) / 2), int(texScale * float32(tex.Height) / 2)})
	rl.DrawTextureEx(tex, pieceCorner.ToRlVec(), 0, float32(p.Scale)*texScale, rl.White)
}

func (p *Piece) RenderCrossPlaneIndicator() {
	var pos = GetBoardOrigo().Add(p.Coord.Scale(TileSize))
	var offset = TileSize / 2
	rl.DrawCircle(int32(pos.X+offset), int32(pos.Y+offset), 12, rl.Blue)
}
