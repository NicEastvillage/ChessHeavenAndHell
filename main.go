package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"

	rg "github.com/gen2brain/raylib-go/raygui"
)

const (
	WindowWidth  = 1600
	WindowHeight = 980
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(WindowWidth, WindowHeight, "Chess - Heaven and Hell")
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

	var undo = NewUndoRedoSystem()
	var ui = NewUiState()
	defer ui.Dispose()

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

		ui.Update()
		handleBoardInteraction(&undo, &ui)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		sandbox.Render(uint32(ui.board), false, &ui.selection)
		ui.Render(&undo)
		rl.EndDrawing()
	}
}

func IsCoordUnderUi(coord Vec2) bool {
	return coord.x < -2 || coord.x >= 12 || coord.y < -2 || coord.y >= 10
}

func handleBoardInteraction(undo *UndoRedoSystem, ui *UiState) {
	handleMouseInteraction(undo, ui)

	var ctrlDown = rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyLeftControl)

	if pieceId, ok := ui.selection.GetSelectedPieceId(); ok {
		if rl.IsKeyPressed(rl.KeyDelete) || rl.IsKeyPressed(rl.KeyBackspace) {
			var cmd = NewDeletePieceCmd(&sandbox, ui, pieceId)
			undo.Append(&cmd)
		} else if rl.IsKeyPressed(rl.KeyC) && ctrlDown {
			var piece = sandbox.GetPiece(pieceId)
			ui.clipboard.StorePiece(piece.typ, piece.color, piece.scale, sandbox.GetStatusEffectsOnPiece(pieceId))
		} else if rl.IsKeyPressed(rl.KeyX) && ctrlDown {
			var piece = sandbox.GetPiece(pieceId)
			ui.clipboard.StorePiece(piece.typ, piece.color, piece.scale, sandbox.GetStatusEffectsOnPiece(pieceId))
			var cmd = NewDeletePieceCmd(&sandbox, ui, pieceId)
			undo.Append(&cmd)
		} else if rl.IsKeyReleased(rl.KeyD) && ctrlDown {
			var cmd = NewDuplicatePieceCmd(&sandbox, ui, pieceId)
			undo.Append(&cmd)
		} else if rl.IsKeyPressed(rl.KeyC) {
			var newColor = 1 - sandbox.GetPiece(pieceId).color
			var cmd = NewChangeColorOfPieceCmd(&sandbox, pieceId, newColor)
			undo.Append(&cmd)
		}
	}

	if rl.IsKeyPressed(rl.KeyZ) && ctrlDown {
		undo.Undo(&sandbox, ui)
	} else if rl.IsKeyPressed(rl.KeyY) && ctrlDown {
		undo.Redo(&sandbox, ui)
	} else if rl.IsKeyPressed(rl.KeyV) && ctrlDown {
		if !ui.clipboard.isEmpty {
			var coord = GetHoveredCoord()
			if !IsCoordUnderUi(coord) {
				var cmd = NewPastePieceCmd(&sandbox, ui, coord, uint32(ui.board))
				undo.Append(&cmd)
			}
		}
	}
}

func handleMouseInteraction(undo *UndoRedoSystem, ui *UiState) {
	var coord = GetHoveredCoord()
	if IsCoordUnderUi(coord) {
		return
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		var coord = GetHoveredCoord()
		if ui.tab == 0 {
			var piece = sandbox.GetPieceAtVisual(coord, uint32(ui.board))
			if piece == nil {
				ui.selection.Deselect()
			} else {
				ui.selection.SelectPiece(piece.id)
			}
		} else if ui.tab == 1 {
			ui.selection.SelectCoord(coord)
		}
	} else if id, ok := ui.selection.GetSelectedPieceId(); ok {
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			var piece = sandbox.GetPiece(id)
			if piece.coord != coord || piece.board != uint32(ui.board) {
				var cmd = NewMovePieceCmd(&sandbox, id, coord, uint32(ui.board))
				undo.Append(&cmd)
			}
		}
	} else if id, ok := ui.selection.GetSelectedPieceTypeId(); ok {
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			var cmd = NewCreatePieceCmd(&sandbox, id, PieceColor(ui.color), uint32(ui.board), coord)
			undo.Append(&cmd)
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
	sandbox.RegisterPieceType(NameBomber, assets.texWhiteBomber, assets.texBlackBomber)
	sandbox.RegisterPieceType(NameLeopard, assets.texWhiteLeopard, assets.texBlackLeopard)
	sandbox.RegisterPieceType(NameChecker, assets.texWhiteChecker, assets.texBlackChecker)
	sandbox.RegisterPieceType(NameMountedArcher, assets.texWhiteMountedArcher, assets.texBlackMountedArcher)
	sandbox.RegisterPieceType(NameWizard, assets.texWhiteWizard, assets.texBlackWizard)
	sandbox.RegisterPieceType(NameArchbishop, assets.texWhiteArchbishop, assets.texBlackArchbishop)
	sandbox.RegisterPieceType(NameFortress, assets.texWhiteFortress, assets.texBlackFortress)
	sandbox.RegisterPieceType(NameScout, assets.texWhiteScout, assets.texBlackScout)
	sandbox.RegisterPieceType(NameWarlock, assets.texWhiteWarlock, assets.texBlackWarlock)
	sandbox.RegisterPieceType(NameCelestial, assets.texAngel, assets.texImp)
}

func registerStatusEffectTypes() {
	sandbox.RegisterEffectType(NameBloody, assets.texEffectBlood)
	sandbox.RegisterEffectType(NameExperience, assets.texEffectMedal)
	sandbox.RegisterEffectType(NameCurse, assets.texEffectCurse)
	sandbox.RegisterEffectType(NameForcedMove, assets.texEffectForcedMove)
	sandbox.RegisterEffectType(NamePaid2ndMove, assets.texEffectPaid2ndMove)
	sandbox.RegisterEffectType(NamePortalGun, assets.texEffectPortalGun)
	sandbox.RegisterEffectType(NameStonks, assets.texEffectStonks)
	sandbox.RegisterEffectType(NameStun, assets.texEffectStun)
	sandbox.RegisterEffectType(NameWizardHat, assets.texEffectWizardHat)
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
