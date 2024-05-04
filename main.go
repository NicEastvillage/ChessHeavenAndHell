package main

import (
	"fmt"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"runtime/debug"
)

const (
	WindowWidth  = 1600
	WindowHeight = 980
)

// Commit Git hash for current commit
var Commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}

	return ""
}()

func main() {
	fmt.Printf("Starting Chess Heaven and Hell v1.0/%s\n", Commit[:7])
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(WindowWidth, WindowHeight, "Chess - Heaven and Hell")
	defer rl.CloseWindow()
	rl.SetExitKey(rl.KeyNull)
	var exit = false

	rl.SetTargetFPS(60)

	assets.LoadAll()
	defer assets.UnloadAll()

	rg.LoadStyleDefault()
	rg.SetFont(assets.fontComicSansMs)
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)

	registerPieceTypes()
	registerStatusEffectTypes()
	registerObstacleTypes()
	setupBoard(0, BoardStyleHeaven, false)
	setupBoard(1, BoardStyleEarth, true)
	setupBoard(2, BoardStyleHell, false)
	sandbox.NewObstacle(Vec2{1, 2}, 0, sandbox.GetObstacleTypeByName(NameCoin).Id)
	sandbox.NewObstacle(Vec2{1, 2}, 0, sandbox.GetObstacleTypeByName(NameCoin).Id)
	sandbox.NewObstacle(Vec2{6, 5}, 0, sandbox.GetObstacleTypeByName(NameCoin).Id)
	sandbox.NewObstacle(Vec2{6, 5}, 0, sandbox.GetObstacleTypeByName(NameCoin).Id)
	sandbox.NewObstacle(Vec2{2, 5}, 0, sandbox.GetObstacleTypeByName(NameChaosOrb).Id)
	sandbox.NewObstacle(Vec2{5, 2}, 0, sandbox.GetObstacleTypeByName(NameChaosOrb).Id)
	sandbox.NewObstacle(Vec2{6, 2}, 2, sandbox.GetObstacleTypeByName(NameFire).Id)
	sandbox.NewObstacle(Vec2{1, 5}, 2, sandbox.GetObstacleTypeByName(NameFire).Id)
	sandbox.NewObstacle(Vec2{2, 2}, 2, sandbox.GetObstacleTypeByName(NameChaosOrb).Id)
	sandbox.NewObstacle(Vec2{5, 5}, 2, sandbox.GetObstacleTypeByName(NameChaosOrb).Id)

	var undo = NewUndoRedoSystem()
	var ui = NewUiState()
	defer ui.Dispose()

	for !exit {

		if rl.WindowShouldClose() {
			exit = AskAboutSaveBeforeExit(&sandbox, &undo)
		} else {
			CheckSavingAndLoading(&sandbox, &undo)
		}

		ui.Update()
		handleBoardInteraction(&undo, &ui)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		if ui.tab == TabBoard {
			sandbox.Render(uint32(ui.board), false, &ui.selection)
		}
		ui.Render(&undo)
		rl.EndDrawing()
	}
}

func IsCoordUnderUi(coord Vec2) bool {
	return coord.X < -2 || coord.X >= 12 || coord.Y < -2 || coord.Y >= 10
}

func handleBoardInteraction(undo *UndoRedoSystem, ui *UiState) {
	handleMouseInteraction(undo, ui)

	var ctrlDown = rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyLeftControl)

	if pieceId, ok := ui.selection.GetSelectedPieceId(); ok && ui.tab == TabBoard {
		if rl.IsKeyPressed(rl.KeyDelete) || rl.IsKeyPressed(rl.KeyBackspace) {
			undo.Append(NewDeletePieceCmd(&sandbox, ui, pieceId))
		} else if rl.IsKeyPressed(rl.KeyC) && ctrlDown {
			var piece = sandbox.GetPiece(pieceId)
			ui.clipboard.StorePiece(piece.Typ, piece.Color, piece.Scale, sandbox.GetStatusEffectsOnPiece(pieceId))
		} else if rl.IsKeyPressed(rl.KeyX) && ctrlDown {
			var piece = sandbox.GetPiece(pieceId)
			ui.clipboard.StorePiece(piece.Typ, piece.Color, piece.Scale, sandbox.GetStatusEffectsOnPiece(pieceId))
			undo.Append(NewDeletePieceCmd(&sandbox, ui, pieceId))
		} else if rl.IsKeyReleased(rl.KeyD) && ctrlDown {
			undo.Append(NewDuplicatePieceCmd(&sandbox, ui, pieceId))
		} else if rl.IsKeyPressed(rl.KeyC) {
			var newColor = 1 - sandbox.GetPiece(pieceId).Color
			undo.Append(NewChangeColorOfPieceCmd(&sandbox, pieceId, newColor))
		}
	}

	if rl.IsKeyPressed(rl.KeyZ) && ctrlDown {
		undo.Undo(&sandbox, ui)
	} else if rl.IsKeyPressed(rl.KeyY) && ctrlDown {
		undo.Redo(&sandbox, ui)
	} else if rl.IsKeyPressed(rl.KeyV) && ctrlDown {
		if !ui.clipboard.isEmpty && ui.tab == TabBoard {
			var coord = GetHoveredCoord()
			if !IsCoordUnderUi(coord) {
				undo.Append(NewPastePieceCmd(&sandbox, ui, coord, uint32(ui.board)))
			}
		}
	}
}

func handleMouseInteraction(undo *UndoRedoSystem, ui *UiState) {
	if ui.tab != TabBoard {
		return
	}

	var coord = GetHoveredCoord()
	if IsCoordUnderUi(coord) {
		return
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		var coord = GetHoveredCoord()
		if ui.mode == 0 {
			var piece = sandbox.GetPieceAtVisual(coord, uint32(ui.board))
			if piece == nil {
				ui.selection.Deselect()
			} else {
				ui.selection.SelectPiece(piece.Id)
			}
		} else if ui.mode == 1 {
			ui.selection.SelectCoord(coord)
		}
	} else if id, ok := ui.selection.GetSelectedPieceId(); ok {
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			var piece = sandbox.GetPiece(id)
			if piece.Coord != coord || piece.Board != uint32(ui.board) {
				undo.Append(NewMovePieceCmd(&sandbox, id, coord, uint32(ui.board)))
			}
		}
	} else if id, ok := ui.selection.GetSelectedPieceTypeId(); ok {
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			undo.Append(NewCreatePieceCmd(&sandbox, id, PieceColor(ui.color), uint32(ui.board), coord))
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
	sandbox.RegisterEffectType(NameBloody, RenderStyleBottom, assets.texEffectBlood)
	sandbox.RegisterEffectType(NameExperience, RenderStyleBottom, assets.texEffectMedal)
	sandbox.RegisterEffectType(NameCurse, RenderStyleBottom, assets.texEffectCurse)
	sandbox.RegisterEffectType(NameForcedMove, RenderStyleBottom, assets.texEffectForcedMove)
	sandbox.RegisterEffectType(NamePaid2ndMove, RenderStyleBottom, assets.texEffectPaid2ndMove)
	sandbox.RegisterEffectType(NamePortalGun, RenderStyleBottom, assets.texEffectPortalGun)
	sandbox.RegisterEffectType(NameStonks, RenderStyleBottom, assets.texEffectStonks)
	sandbox.RegisterEffectType(NameStun, RenderStyleStun, assets.texEffectStun)
	sandbox.RegisterEffectType(NameWizardHat, RenderStyleBottom, assets.texEffectWizardHat)
}

func registerObstacleTypes() {
	sandbox.RegisterObstacleType(NameChaosOrb, assets.texObstacleChaosOrb)
	sandbox.RegisterObstacleType(NameCoin, assets.texObstacleCoin)
	sandbox.RegisterObstacleType(NameIce, assets.texObstacleIce)
	sandbox.RegisterObstacleType(NameFire, assets.texObstacleFire)
}

func setupBoard(boardId uint32, style BoardStyle, withPieces bool) {
	var board = sandbox.GetBoard(boardId)
	board.Id = boardId
	board.Style = style
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			sandbox.NewTile(board.Id, Vec2{x, y})
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
