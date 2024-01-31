package main

import rl "github.com/gen2brain/raylib-go/raylib"

var assets = AssetManager{}

type AssetManager struct {
	texWhitePawn   rl.Texture2D
	texWhiteKnight rl.Texture2D
	texWhiteBishop rl.Texture2D
	texWhiteRook   rl.Texture2D
	texWhiteQueen  rl.Texture2D
	texWhiteKing   rl.Texture2D
	texBlackPawn   rl.Texture2D
	texBlackKnight rl.Texture2D
	texBlackBishop rl.Texture2D
	texBlackRook   rl.Texture2D
	texBlackQueen  rl.Texture2D
	texBlackKing   rl.Texture2D
}

func (am *AssetManager) LoadAll() {
	am.texWhitePawn = rl.LoadTexture("Assets/white_pawn.png")
	am.texWhiteKnight = rl.LoadTexture("Assets/white_knight.png")
	am.texWhiteBishop = rl.LoadTexture("Assets/white_bishop.png")
	am.texWhiteRook = rl.LoadTexture("Assets/white_rook.png")
	am.texWhiteQueen = rl.LoadTexture("Assets/white_queen.png")
	am.texWhiteKing = rl.LoadTexture("Assets/white_king.png")
	am.texBlackPawn = rl.LoadTexture("Assets/black_pawn.png")
	am.texBlackKnight = rl.LoadTexture("Assets/black_knight.png")
	am.texBlackBishop = rl.LoadTexture("Assets/black_bishop.png")
	am.texBlackRook = rl.LoadTexture("Assets/black_rook.png")
	am.texBlackQueen = rl.LoadTexture("Assets/black_queen.png")
	am.texBlackKing = rl.LoadTexture("Assets/black_king.png")
}

func (am *AssetManager) UnloadAll() {
	rl.UnloadTexture(am.texWhitePawn)
	rl.UnloadTexture(am.texWhiteKnight)
	rl.UnloadTexture(am.texWhiteBishop)
	rl.UnloadTexture(am.texWhiteRook)
	rl.UnloadTexture(am.texWhiteQueen)
	rl.UnloadTexture(am.texWhiteKing)
	rl.UnloadTexture(am.texBlackPawn)
	rl.UnloadTexture(am.texBlackKnight)
	rl.UnloadTexture(am.texBlackBishop)
	rl.UnloadTexture(am.texBlackRook)
	rl.UnloadTexture(am.texBlackQueen)
	rl.UnloadTexture(am.texBlackKing)
}
