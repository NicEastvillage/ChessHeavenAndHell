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

	registerPieceTypes()
	registerStatusEffectTypes()
	registerObstacleTypes()
	setupBoard(0, BoardStyleHeaven, false)
	setupBoard(1, BoardStyleEarth, true)
	setupBoard(2, BoardStyleHell, false)

	var ui = NewUIState()

	for i := 0; i < 20; i++ {
		var piece = sandbox.pieces[rand.Intn(len(sandbox.pieces))].id
		var effect = sandbox.effectTypes[rand.Intn(len(sandbox.effectTypes))].id
		sandbox.NewStatusEffect(piece, effect)
	}
	for x := 0; x < 8; x++ {
		for i := 0; i < x; i++ {
			var obstType = sandbox.obstacleTypes[rand.Intn(len(sandbox.obstacleTypes))].id
			sandbox.NewObstacle(Vec2{x: x, y: 4}, 0, obstType)
		}
	}

	for !rl.WindowShouldClose() {

		handleBoardInteraction(&ui)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		sandbox.Render(uint32(ui.board))
		ui.Render()

		rl.EndDrawing()
	}
}

func handleBoardInteraction(ui *UIState) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		var coord = GetHoveredCoord()
		var piece = sandbox.GetPieceAt(coord)
		if piece == nil {
			selection.Deselect()
		} else {
			selection.SelectPiece(piece.id)
			ui.anyPieceSelected = false
		}
	} else if id, ok := selection.GetSelectedPieceId(); ok {
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			var coord = GetHoveredCoord()
			var piece = sandbox.GetPiece(id)
			piece.coord = coord
		} else if rl.IsKeyPressed(rl.KeyDelete) || rl.IsKeyPressed(rl.KeyBackspace) {
			sandbox.RemovePiece(id)
		}
	} else if ui.anyPieceSelected {
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			var coord = GetHoveredCoord()
			sandbox.NewPiece(ui.piece, PieceColor(ui.color), uint32(ui.board), coord)
		}
	}
}

func registerPieceTypes() {
	sandbox.RegisterPieceType(NamePawn, assets.texWhitePawn, assets.texBlackPawn)
	sandbox.RegisterPieceType(NameKnight, assets.texWhiteKnight, assets.texBlackKnight)
	sandbox.RegisterPieceType(NameBishop, assets.texWhiteBishop, assets.texBlackBishop)
	sandbox.RegisterPieceType(NameRook, assets.texWhiteRook, assets.texBlackRook)
	sandbox.RegisterPieceType(NameQueen, assets.texWhiteQueen, assets.texBlackQueen)
	sandbox.RegisterPieceType(NameKing, assets.texWhiteKing, assets.texBlackKing)
}

func registerStatusEffectTypes() {
	sandbox.RegisterEffectType(NameBloody, assets.texEffectBlood)
	sandbox.RegisterEffectType(NameExperience, assets.texEffectMedal)
	sandbox.RegisterEffectType(NameCurse, assets.texEffectCurse)
}

func registerObstacleTypes() {
	sandbox.RegisterObstacleType(NameChaosOrb, assets.texObstacleChaosOrb)
	sandbox.RegisterObstacleType(NameCoin, assets.texObstacleCoin)
	sandbox.RegisterObstacleType(NameIce, assets.texObstacleIce)
	sandbox.RegisterObstacleType(NameFire, assets.texObstacleFire)
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
	sandbox.NewPieceFromName(NameRook, BLACK, boardId, Vec2{0, 0})
	sandbox.NewPieceFromName(NameKnight, BLACK, boardId, Vec2{1, 0})
	sandbox.NewPieceFromName(NameBishop, BLACK, boardId, Vec2{2, 0})
	sandbox.NewPieceFromName(NameQueen, BLACK, boardId, Vec2{3, 0})
	sandbox.NewPieceFromName(NameKing, BLACK, boardId, Vec2{4, 0})
	sandbox.NewPieceFromName(NameBishop, BLACK, boardId, Vec2{5, 0})
	sandbox.NewPieceFromName(NameKnight, BLACK, boardId, Vec2{6, 0})
	sandbox.NewPieceFromName(NameRook, BLACK, boardId, Vec2{7, 0})
	for x := 0; x < 8; x++ {
		sandbox.NewPieceFromName(NamePawn, BLACK, boardId, Vec2{x, 1})
	}
	sandbox.NewPieceFromName(NameRook, WHITE, boardId, Vec2{0, 7})
	sandbox.NewPieceFromName(NameKnight, WHITE, boardId, Vec2{1, 7})
	sandbox.NewPieceFromName(NameBishop, WHITE, boardId, Vec2{2, 7})
	sandbox.NewPieceFromName(NameQueen, WHITE, boardId, Vec2{3, 7})
	sandbox.NewPieceFromName(NameKing, WHITE, boardId, Vec2{4, 7})
	sandbox.NewPieceFromName(NameBishop, WHITE, boardId, Vec2{5, 7})
	sandbox.NewPieceFromName(NameKnight, WHITE, boardId, Vec2{6, 7})
	sandbox.NewPieceFromName(NameRook, WHITE, boardId, Vec2{7, 7})
	for x := 0; x < 8; x++ {
		sandbox.NewPieceFromName(NamePawn, WHITE, boardId, Vec2{x, 6})
	}
}
