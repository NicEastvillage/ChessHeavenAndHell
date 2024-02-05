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

func (s *UiState) Render() {
	if rg.Button(rl.NewRectangle(UiMargin, UiMargin, 200, UiButtonH), "Remove random") {
		println("Clicked!")
		if len(sandbox.pieces) > 0 {
			var id = sandbox.pieces[rand.Intn(len(sandbox.pieces))].id
			sandbox.RemovePiece(id)
		}
	}

	rg.Panel(rl.NewRectangle(float32(rl.GetScreenWidth())-UiRightMenuWidth, 0, UiRightMenuWidth, float32(rl.GetScreenHeight())), "")

	s.board = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()/2-(120*3+int(rg.GetStyle(rg.DEFAULT, rg.GROUP_PADDING)))/2), float32(rl.GetScreenHeight()-UiButtonH-UiMargin), 120, UiButtonH), "Heaven;Earth;Hell", s.board)

	if s.selection.selectionType == SelectionTypePiece {
		s.RenderPieceContextMenu()
	} else {
		s.tab = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-2*UiButtonH-2*int(rg.GetStyle(rg.DEFAULT, rg.GROUP_PADDING))), UiMargin+5, UiButtonH, UiButtonH), "#149#;#157#;#97#", s.tab)

		if s.tab == 0 {
			s.RenderPiecesTab()
		} else if s.selection.selectionType == SelectionTypePieceType {
			s.selection.Deselect()
		}
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

func (s *UiState) RenderPieceContextMenu() {
	var selectedPiece, _ = s.selection.GetSelectedPieceId()

	for i, effect := range sandbox.effectTypes {
		var iconPosX = int32(rl.GetScreenWidth()) - 85
		var iconPosY = int32(i*75 + 50)
		rl.DrawTexture(effect.tex, iconPosX, iconPosY, rl.White)

		var effectCount = sandbox.GetStatusEffectCount(selectedPiece, effect.id)
		rl.DrawText(fmt.Sprint(effectCount), iconPosX+5, iconPosY+25, 16, rl.Black)
		if rg.Button(rl.NewRectangle(float32(iconPosX-55), float32(iconPosY), 40, 30), "--") && effectCount > 0 {
			sandbox.RemoveStatusEffect(selectedPiece, effect.id)
		}
		if rg.Button(rl.NewRectangle(float32(iconPosX+35), float32(iconPosY), 40, 30), "++") {
			sandbox.NewStatusEffect(selectedPiece, effect.id)
		}
	}

	var posX = float32(rl.GetScreenWidth() - UiRightMenuWidth + UiMarginSmall)
	var posY = float32(rl.GetScreenHeight() - UiMarginSmall - 40)
	var width = float32(rl.GetScreenWidth() - int(posX) - UiMarginSmall)
	var height float32 = 40
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Remove piece") {
		sandbox.RemovePiece(selectedPiece)
		s.selection.Deselect()
	}
}
