package main

import (
	"fmt"
	"math/rand"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	UiMargin         = 20
	UiMarginSmall    = 10
	UiMarginBig      = 40
	UiButtonH        = 36
	UiRightMenuWidth = 155
)

type UiState struct {
	selection Selection
	board     int32
	tab       int32
	color     int32
}

func NewUiState() UiState {
	return UiState{
		board: int32(1),
	}
}

func (s *UiState) Render(undo *UndoRedoSystem) {
	if rg.Button(rl.NewRectangle(UiMargin, UiMargin, 200, UiButtonH), "Remove random") {
		println("Clicked!")
		if len(sandbox.pieces) > 0 {
			var id = sandbox.pieces[rand.Intn(len(sandbox.pieces))].id
			sandbox.RemovePiece(id)
		}
	}

	s.board = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()/2-(120*3+int(rg.GetStyle(rg.DEFAULT, rg.GROUP_PADDING)))/2), float32(rl.GetScreenHeight()-UiButtonH-UiMargin), 120, UiButtonH), "Heaven;Earth;Hell", s.board)

	s.tab = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-2*UiButtonH-2*int(rg.GetStyle(rg.DEFAULT, rg.GROUP_PADDING))), UiMargin+5, UiButtonH, UiButtonH), "#149#;#157#;#97#", s.tab)

	if s.tab != 0 && s.selection.selectionType == SelectionTypePiece {
		s.selection.Deselect()
	} else if s.tab != 0 && s.selection.selectionType == SelectionTypePieceType {
		s.selection.Deselect()
	} else if s.tab != 1 && s.selection.selectionType == SelectionTypeCoord {
		s.selection.Deselect()
	}

	switch s.selection.selectionType {
	case SelectionTypePiece:
		s.RenderPieceContextMenu(undo)
	case SelectionTypePieceType:
		s.RenderPiecesTab()
	case SelectionTypeCoord:
		s.RenderCoordContextMenu(undo)
	default:
		s.RenderPiecesTab()
	}
}

func (s *UiState) RenderPiecesTab() {
	s.color = rg.ToggleSlider(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-130), 2*UiMargin+UiButtonH, 130, UiButtonH), "White;Black", s.color)

	for i := 0; i < len(sandbox.pieceTypes); i++ {
		if rg.Toggle(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-130), float32(3*UiMargin+2*UiButtonH+i*(UiMarginSmall+UiButtonH)), 130, UiButtonH), sandbox.GetPieceType(uint32(i)).name, s.selection.IsPieceTypeSelected(uint32(i))) {
			s.selection.SelectPieceType(uint32(i))
		}
	}
}

func (s *UiState) RenderPieceContextMenu(undo *UndoRedoSystem) {
	var selectedPiece, _ = s.selection.GetSelectedPieceId()

	var spinnerX = float32(rl.GetScreenWidth() - 150)
	var spinnerY = float32(UiMargin + UiMarginBig + UiButtonH)

	{
		var pieceScale = sandbox.GetPiece(selectedPiece).scale
		var change = SpinnerWithIcon(spinnerX, spinnerY, fmt.Sprint(pieceScale), assets.texPieceScale)
		if change < 0 && pieceScale > 1 {
			var cmd = NewDecreasePieceScaleCmd(&sandbox, selectedPiece)
			undo.AppendDone(&cmd)
		}
		if change > 0 {
			var cmd = NewIncreasePieceScaleCmd(&sandbox, selectedPiece)
			undo.AppendDone(&cmd)
		}
	}

	for i := range sandbox.effectTypes {
		var effect = &sandbox.effectTypes[i]
		var effectCount = sandbox.GetStatusEffectCount(selectedPiece, effect.id)
		var change = SpinnerWithIcon(spinnerX, spinnerY+float32(i*55)+55, fmt.Sprint(effectCount), effect.tex)
		if change < 0 && effectCount > 0 {
			var cmd = NewRemoveStatusEffectCmd(&sandbox, selectedPiece, effect.id)
			undo.AppendDone(&cmd)
		}
		if change > 0 {
			var cmd = NewCreateStatusEffectCmd(&sandbox, selectedPiece, effect.id)
			undo.AppendDone(&cmd)
		}
	}

	var posX = float32(rl.GetScreenWidth() - 130 - UiMargin)
	var posY = float32(rl.GetScreenHeight() - UiMargin - UiButtonH)
	var width = float32(130)
	var height float32 = UiButtonH
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Remove piece") {
		var cmd = NewRemovePieceCmd(&sandbox, s, selectedPiece)
		undo.AppendDone(&cmd)
	}
}

func (s *UiState) RenderCoordContextMenu(undo *UndoRedoSystem) {
	var coord, _ = s.selection.GetSelectedCoord()

	var spinnerX = float32(rl.GetScreenWidth() - 150)
	var spinnerY = float32(UiMargin + UiMarginBig + UiButtonH)

	for i := range sandbox.obstacleTypes {
		var obt = &sandbox.obstacleTypes[i]
		var obCount = sandbox.GetObstacleCount(coord, uint32(s.board), obt.id)
		var change = SpinnerWithIcon(spinnerX, spinnerY+float32(i*55)+55, fmt.Sprint(obCount), obt.tex)
		if change < 0 && obCount > 0 {
			var cmd = NewRemoveObstacleCmd(&sandbox, coord, uint32(s.board), obt.id)
			undo.AppendDone(&cmd)
		}
		if change > 0 {
			var cmd = NewAddObstacleCmd(&sandbox, coord, uint32(s.board), obt.id)
			undo.AppendDone(&cmd)
		}
	}

	var posX = float32(rl.GetScreenWidth() - 130 - UiMargin)
	var posY = float32(rl.GetScreenHeight() - UiMargin - UiButtonH)
	var width = float32(130)
	var height float32 = UiButtonH
	if rg.Button(rl.NewRectangle(posX, posY-UiMarginSmall-UiButtonH, width, height), "Add tile") {
		if sandbox.GetTile(uint32(s.board), coord) == nil {
			var cmd = NewAddTileCmd(&sandbox, uint32(s.board), coord)
			undo.AppendDone(&cmd)
		}
	}
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Remove tile") {
		var cmd = NewRemoveTileCmd(&sandbox, uint32(s.board), coord)
		undo.AppendDone(&cmd)
	}
}

func SpinnerWithIcon(x float32, y float32, text string, tex rl.Texture2D) int {
	const (
		buttonW  = 40
		buttonH  = 30
		spacing  = UiMarginSmall
		iconSize = 32
	)

	var texScale = iconSize / float32(tex.Height)
	rl.DrawTextureEx(tex, rl.NewVector2(x+buttonW+spacing+iconSize/2-texScale*float32(tex.Width/2), y+buttonH/2-texScale*float32(tex.Height)/2-3), 0, texScale, rl.White)
	rl.DrawText(text, int32(x+buttonW+spacing+13), int32(y+buttonH/2+iconSize/2-3), 16, rl.Black)
	var res = 0
	if rg.Button(rl.NewRectangle(x, y, buttonW, buttonH), "--") {
		res--
	}
	if rg.Button(rl.NewRectangle(x+buttonW+iconSize+2*spacing, y, buttonW, buttonH), "++") {
		res++
	}
	return res
}
