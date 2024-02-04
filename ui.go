package main

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

const (
	UiMargin      = 20
	UiMarginSmall = 10
	UiMarginBig   = 40
	UiButtonH     = 36
)

type UIState struct {
	board            int32
	tab              int32
	color            int32
	anyPieceSelected bool
	piece            uint32
}

func NewUIState() UIState {
	return UIState{
		board: int32(1),
	}
}

func (s *UIState) Render() {
	if rg.Button(rl.NewRectangle(UiMargin, UiMargin, 200, UiButtonH), "Remove random") {
		println("Clicked!")
		if len(sandbox.pieces) > 0 {
			var id = sandbox.pieces[rand.Intn(len(sandbox.pieces))].id
			sandbox.RemovePiece(id)
		}
	}

	s.board = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()/2-(120*3+int(rg.GetStyle(rg.DEFAULT, rg.GROUP_PADDING)))/2), float32(rl.GetScreenHeight()-UiButtonH-UiMargin), 120, UiButtonH), "Heaven;Earth;Hell", s.board)

	s.tab = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-2*UiButtonH-2*int(rg.GetStyle(rg.DEFAULT, rg.GROUP_PADDING))), UiMargin, UiButtonH, UiButtonH), "#149#;#157#;#97#", s.tab)

	if s.tab == 0 {
		s.RenderPiecesTab()
	} else {
		s.anyPieceSelected = false
	}
}

func (s *UIState) RenderPiecesTab() {
	s.color = rg.ToggleSlider(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-130), 2*UiMargin+UiButtonH, 130, UiButtonH), "White;Black", s.color)

	for i := 0; i < len(sandbox.pieceTypes); i++ {
		if rg.Toggle(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-130), float32(3*UiMargin+2*UiButtonH+i*(UiMarginSmall+UiButtonH)), 130, UiButtonH), sandbox.GetPieceType(uint32(i)).name, s.anyPieceSelected && s.piece == uint32(i)) {
			s.anyPieceSelected = true
			s.piece = uint32(i)
		}
	}
}
