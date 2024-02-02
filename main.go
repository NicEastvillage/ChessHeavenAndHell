package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)
import rg "github.com/gen2brain/raylib-go/raygui"

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1600, 980, "Chess - Heaven and Hell")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rg.LoadStyleDefault()
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)

	assets.LoadAll()
	defer assets.UnloadAll()

	registerStatusEffectTypes()
	setupBoard(0, BoardStyleHeaven, false)
	setupBoard(1, BoardStyleEarth, true)
	setupBoard(2, BoardStyleHell, false)

	var planeIndex = int32(1)

	for i := 0; i < 20; i++ {
		var piece = sandbox.pieces[rand.Intn(len(sandbox.pieces))].id
		var effect = sandbox.effectTypes[rand.Intn(len(sandbox.effectTypes))].id
		sandbox.NewStatusEffect(piece, effect)
	}

	for !rl.WindowShouldClose() {

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			var hoveredCoord = GetHoveredCoord()
			var piece = sandbox.GetPieceAt(hoveredCoord)
			if piece == nil {
				selection.Deselect()
			} else {
				selection.SelectPiece(piece.id)
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		sandbox.Render(uint32(planeIndex))

		if rg.Button(rl.NewRectangle(20, 20, 200, 36), "Remove random") {
			println("Clicked!")
			if len(sandbox.pieces) > 0 {
				var id = sandbox.pieces[rand.Intn(len(sandbox.pieces))].id
				sandbox.RemovePiece(id)
			}
		}

		planeIndex = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()/2-(120*3+int(rg.GetStyle(rg.DEFAULT, rg.GROUP_PADDING)))/2), float32(rl.GetScreenHeight()-36-20), 120, 36), "Heaven;Earth;Hell", planeIndex)

		rl.EndDrawing()
	}
}

func registerStatusEffectTypes() {
	sandbox.RegisterEffectType(assets.texEffectBlood)
	sandbox.RegisterEffectType(assets.texEffectMedal)
	sandbox.RegisterEffectType(assets.texEffectCurse)
}

func setupBoard(boardId uint32, style BoardStyle, withPieces bool) {
	var board = sandbox.GetBoard(boardId)
	board.id = boardId
	board.style = style
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			sandbox.NewTile(board.id, Vec2{x, y})
		}
	}
	if !withPieces {
		return
	}
	sandbox.NewPiece(boardId, Vec2{0, 0}, assets.texBlackRook)
	sandbox.NewPiece(boardId, Vec2{1, 0}, assets.texBlackKnight)
	sandbox.NewPiece(boardId, Vec2{2, 0}, assets.texBlackBishop)
	sandbox.NewPiece(boardId, Vec2{3, 0}, assets.texBlackQueen)
	sandbox.NewPiece(boardId, Vec2{4, 0}, assets.texBlackKing)
	sandbox.NewPiece(boardId, Vec2{5, 0}, assets.texBlackBishop)
	sandbox.NewPiece(boardId, Vec2{6, 0}, assets.texBlackKnight)
	sandbox.NewPiece(boardId, Vec2{7, 0}, assets.texBlackRook)
	for x := 0; x < 8; x++ {
		sandbox.NewPiece(boardId, Vec2{x, 1}, assets.texBlackPawn)
	}
	sandbox.NewPiece(boardId, Vec2{0, 7}, assets.texWhiteRook)
	sandbox.NewPiece(boardId, Vec2{1, 7}, assets.texWhiteKnight)
	sandbox.NewPiece(boardId, Vec2{2, 7}, assets.texWhiteBishop)
	sandbox.NewPiece(boardId, Vec2{3, 7}, assets.texWhiteQueen)
	sandbox.NewPiece(boardId, Vec2{4, 7}, assets.texWhiteKing)
	sandbox.NewPiece(boardId, Vec2{5, 7}, assets.texWhiteBishop)
	sandbox.NewPiece(boardId, Vec2{6, 7}, assets.texWhiteKnight)
	sandbox.NewPiece(boardId, Vec2{7, 7}, assets.texWhiteRook)
	for x := 0; x < 8; x++ {
		sandbox.NewPiece(boardId, Vec2{x, 6}, assets.texWhitePawn)
	}
}
