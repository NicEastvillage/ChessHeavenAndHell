package main

import (
	"fmt"
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
	selection       Selection
	clipboard       Clipboard
	board           int32
	tab             int32
	color           int32
	renderTexHeaven rl.RenderTexture2D
	renderTexEarth  rl.RenderTexture2D
	renderTexHell   rl.RenderTexture2D
}

func NewUiState() UiState {
	return UiState{
		selection:       NewSelection(),
		clipboard:       NewClipboard(),
		board:           int32(1),
		renderTexHeaven: rl.LoadRenderTexture(WindowWidth, WindowHeight),
		renderTexEarth:  rl.LoadRenderTexture(WindowWidth, WindowHeight),
		renderTexHell:   rl.LoadRenderTexture(WindowWidth, WindowHeight),
	}
}

func (s *UiState) Update() {
	if s.board != 0 {
		rl.BeginTextureMode(s.renderTexHeaven)
		rl.ClearBackground(rl.RayWhite)
		sandbox.Render(0, true, &s.selection)
		rl.EndTextureMode()
	}
	if s.board != 1 {
		rl.BeginTextureMode(s.renderTexEarth)
		rl.ClearBackground(rl.RayWhite)
		sandbox.Render(1, true, &s.selection)
		rl.EndTextureMode()
	}
	if s.board != 2 {
		rl.BeginTextureMode(s.renderTexHell)
		rl.ClearBackground(rl.RayWhite)
		sandbox.Render(2, true, &s.selection)
		rl.EndTextureMode()
	}
}

func (s *UiState) Render(undo *UndoRedoSystem) {

	s.RenderBoardPreview(0)
	s.RenderBoardPreview(1)
	s.RenderBoardPreview(2)

	var oldTab = s.tab
	s.tab = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-2*UiButtonH-1*int(rg.GetStyle(rg.TOGGLE, rg.GROUP_PADDING))), UiMargin, UiButtonH, UiButtonH), "#149#;#97#", s.tab)
	if oldTab != s.tab {
		s.selection.Deselect()
	}

	switch s.selection.selectionType {
	case SelectionTypePiece:
		s.RenderPieceContextMenu(undo)
		s.tab = 0
	case SelectionTypePieceType:
		s.RenderPiecesTab()
		s.tab = 0
	case SelectionTypeCoord:
		s.RenderCoordContextMenu(undo)
		s.tab = 1
	default:
		if s.tab == 0 {
			s.RenderPiecesTab()
		}
	}
}

func (s *UiState) RenderBoardPreview(index int32) {
	if s.board == index {
		return
	}

	var origo = GetBoardOrigo()
	var previewSourceRect = rl.NewRectangle(float32(origo.x-TileSize), float32(origo.y-TileSize), 10*TileSize, -10*TileSize)
	var previewPlacement = rl.NewRectangle(UiMargin, UiMargin+float32(index)*(UiMargin+previewSourceRect.Width/3), previewSourceRect.Width/3, -previewSourceRect.Height/3)
	var buttonPlacement = rl.NewRectangle(previewPlacement.X+1, previewPlacement.Y+1, previewPlacement.Width-2, previewPlacement.Height-2)

	var previewTex = s.renderTexHeaven.Texture
	switch index {
	case 1:
		previewTex = s.renderTexEarth.Texture
	case 2:
		previewTex = s.renderTexHell.Texture
	}

	if rg.Button(buttonPlacement, "") {
		s.board = index
	}
	rl.DrawTexturePro(previewTex, previewSourceRect, previewPlacement, rl.NewVector2(0, 0), 0, rl.White)
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), previewPlacement) {
		rl.DrawRectangleLinesEx(previewPlacement, 1, rl.GetColor(uint(rg.GetStyle(rg.BUTTON, rg.BORDER_COLOR_FOCUSED))))
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
			var cmd = NewDeleteStatusEffectCmd(&sandbox, selectedPiece, effect.id)
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
		var cmd = NewDeletePieceCmd(&sandbox, s, selectedPiece)
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
			var cmd = NewDeleteObstacleCmd(&sandbox, coord, uint32(s.board), obt.id)
			undo.AppendDone(&cmd)
		}
		if change > 0 {
			var cmd = NewCreateObstacleCmd(&sandbox, coord, uint32(s.board), obt.id)
			undo.AppendDone(&cmd)
		}
	}

	var posX = float32(rl.GetScreenWidth() - 130 - UiMargin)
	var posY = float32(rl.GetScreenHeight() - UiMargin - UiButtonH)
	var width = float32(130)
	var height float32 = UiButtonH
	if rg.Button(rl.NewRectangle(posX, posY-UiMarginSmall-UiButtonH, width, height), "Add tile") {
		if sandbox.GetTile(uint32(s.board), coord) == nil {
			var cmd = NewCreateTileCmd(&sandbox, uint32(s.board), coord)
			undo.AppendDone(&cmd)
		}
	}
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Remove tile") {
		var cmd = NewDeleteTileCmd(&sandbox, uint32(s.board), coord)
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
