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
	UiButtonFlatH    = 30
	UiButtonNarrowW  = 88
	UiButtonW        = 136
	UiButtonTinyW    = 36
	UiRightMenuWidth = 155
	UiShopWidth      = 660
)

type UiState struct {
	selection       Selection
	clipboard       Clipboard
	shop            Shop
	board           int32
	tab             int32
	color           int32
	renderTexHeaven rl.RenderTexture2D
	renderTexEarth  rl.RenderTexture2D
	renderTexHell   rl.RenderTexture2D
	showShop        bool
}

func NewUiState() UiState {
	return UiState{
		selection:       NewSelection(),
		clipboard:       NewClipboard(),
		shop:            NewShop(),
		board:           int32(1),
		renderTexHeaven: rl.LoadRenderTexture(WindowWidth, WindowHeight),
		renderTexEarth:  rl.LoadRenderTexture(WindowWidth, WindowHeight),
		renderTexHell:   rl.LoadRenderTexture(WindowWidth, WindowHeight),
	}
}

func (s *UiState) Dispose() {
	rl.UnloadRenderTexture(s.renderTexHeaven)
	rl.UnloadRenderTexture(s.renderTexEarth)
	rl.UnloadRenderTexture(s.renderTexHell)
}

func (s *UiState) Update() {
	//var origo = GetBoardOrigo()
	//var previewSourceOrigo = rl.NewVector2(float32(origo.x-TileSize), float32(origo.y-TileSize))

	if s.showShop {
		s.selection.Deselect()
	}

	if s.board != 0 || s.showShop {
		rl.BeginTextureMode(s.renderTexHeaven)
		rl.ClearBackground(rl.RayWhite)
		//rl.Translatef(previewSourceOrigo.X, previewSourceOrigo.Y, 0)
		sandbox.Render(0, true, &s.selection)
		rl.EndTextureMode()
	}
	if s.board != 1 || s.showShop {
		rl.BeginTextureMode(s.renderTexEarth)
		rl.ClearBackground(rl.RayWhite)
		sandbox.Render(1, true, &s.selection)
		rl.EndTextureMode()
	}
	if s.board != 2 || s.showShop {
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
	if rl.IsKeyPressed(rl.KeyUp) && s.board > 0 {
		s.board--
	} else if rl.IsKeyPressed(rl.KeyDown) && s.board < 2 {
		s.board++
	} else if rl.IsKeyPressed(rl.KeyTab) {
		s.board = (s.board + 1) % 3
	}

	s.RenderMoneyWidget()
	if rl.IsKeyPressed(rl.KeyS) {
		s.showShop = !s.showShop
	}
	if s.showShop {
		s.RenderShop()
		return
	}

	var oldTab = s.tab
	s.tab = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-2*UiButtonH-1*int(rg.GetStyle(rg.TOGGLE, rg.GROUP_PADDING))), UiMargin, UiButtonH, UiButtonH), "#149#;#97#", s.tab)
	if rl.IsKeyPressed(rl.KeyT) {
		s.tab = 1 - s.tab
	}
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
	if s.board == index && !s.showShop {
		return
	}

	var previewTex = s.renderTexHeaven.Texture
	switch index {
	case 1:
		previewTex = s.renderTexEarth.Texture
	case 2:
		previewTex = s.renderTexHell.Texture
	}

	var origo = GetBoardOrigo()
	var previewSourceRect = rl.NewRectangle(float32(origo.x-TileSize), float32(origo.y-TileSize), 10*TileSize, -10*TileSize)
	previewSourceRect.Y += float32(previewTex.Height - int32(rl.GetScreenHeight()))
	var previewPlacement = rl.NewRectangle(UiMargin, UiMargin+float32(index)*(UiMargin+previewSourceRect.Width/3), previewSourceRect.Width/3, -previewSourceRect.Height/3)
	var buttonPlacement = rl.NewRectangle(previewPlacement.X+1, previewPlacement.Y+1, previewPlacement.Width-2, previewPlacement.Height-2)

	if rg.Button(buttonPlacement, "") {
		s.board = index
		s.showShop = false
	}
	rl.DrawTexturePro(previewTex, previewSourceRect, previewPlacement, rl.NewVector2(0, 0), 0, rl.White)
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), previewPlacement) {
		rl.DrawRectangleLinesEx(previewPlacement, 1, rl.GetColor(uint(rg.GetStyle(rg.BUTTON, rg.BORDER_COLOR_FOCUSED))))
	}
}

func (s *UiState) RenderPiecesTab() {
	s.color = rg.ToggleSlider(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-UiButtonW), 2*UiMargin+UiButtonH, UiButtonW, UiButtonH), "White;Black", s.color)
	if rl.IsKeyPressed(rl.KeyC) {
		s.color = 1 - s.color
	}

	for i := 0; i < len(sandbox.pieceTypes); i++ {
		if rg.Toggle(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-UiButtonW), float32(3*UiMargin+2*UiButtonH+i*(UiMarginSmall+UiButtonH)), UiButtonW, UiButtonH), sandbox.GetPieceType(uint32(i)).name, s.selection.IsPieceTypeSelected(uint32(i))) {
			s.selection.SelectPieceType(uint32(i))
		}
	}
}

func (s *UiState) RenderPieceContextMenu(undo *UndoRedoSystem) {
	var selectedPieceId, _ = s.selection.GetSelectedPieceId()
	piece := sandbox.GetPiece(selectedPieceId)

	var spinnerX = float32(rl.GetScreenWidth() - 150)
	var spinnerY = float32(UiMargin + UiMarginBig + UiButtonH)

	{
		var pieceScale = piece.scale
		var change = SpinnerWithIcon(spinnerX, spinnerY, fmt.Sprint(pieceScale), assets.texPieceScale)
		if change < 0 && pieceScale > 1 {
			var cmd = NewDecreasePieceScaleCmd(&sandbox, selectedPieceId)
			undo.Append(&cmd)
		}
		if change > 0 {
			var cmd = NewIncreasePieceScaleCmd(&sandbox, selectedPieceId)
			undo.Append(&cmd)
		}
	}

	for i := range sandbox.effectTypes {
		var effect = &sandbox.effectTypes[i]
		var effectCount = sandbox.GetStatusEffectCount(selectedPieceId, effect.id)
		var change = SpinnerWithIcon(spinnerX, spinnerY+float32(i*55)+55, fmt.Sprint(effectCount), effect.tex)
		if change < 0 && effectCount > 0 {
			var cmd = NewDeleteStatusEffectCmd(&sandbox, selectedPieceId, effect.id)
			undo.Append(&cmd)
		}
		if change > 0 {
			var cmd = NewCreateStatusEffectCmd(&sandbox, selectedPieceId, effect.id)
			undo.Append(&cmd)
		}
	}

	var posX = float32(rl.GetScreenWidth() - UiButtonW - UiMargin)
	var posY = float32(rl.GetScreenHeight() - UiMargin - UiButtonH)
	var width = float32(UiButtonW)
	var height float32 = UiButtonH
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Remove piece") {
		var cmd = NewDeletePieceCmd(&sandbox, s, selectedPieceId)
		undo.Append(&cmd)
	}

	posY -= UiButtonH + UiMargin
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Change color") {
		var newColor = 1 - piece.color
		var cmd = NewChangeColorOfPieceCmd(&sandbox, selectedPieceId, newColor)
		undo.Append(&cmd)
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
			undo.Append(&cmd)
		}
		if change > 0 {
			var cmd = NewCreateObstacleCmd(&sandbox, coord, uint32(s.board), obt.id)
			undo.Append(&cmd)
		}
	}

	var posX = float32(rl.GetScreenWidth() - UiButtonW - UiMargin)
	var posY = float32(rl.GetScreenHeight() - UiMargin - UiButtonH)
	var width = float32(UiButtonW)
	var height float32 = UiButtonH
	if rg.Button(rl.NewRectangle(posX, posY-UiMarginSmall-UiButtonH, width, height), "Add tile") {
		if sandbox.GetTile(uint32(s.board), coord) == nil {
			var cmd = NewCreateTileCmd(&sandbox, uint32(s.board), coord)
			undo.Append(&cmd)
		}
	}
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Remove tile") {
		var cmd = NewDeleteTileCmd(&sandbox, uint32(s.board), coord)
		undo.Append(&cmd)
	}
}

func SpinnerWithIcon(x float32, y float32, text string, tex rl.Texture2D) int {
	const (
		spacing  = UiMarginSmall
		iconSize = 32
	)

	var texScale = iconSize / float32(tex.Height)
	rl.DrawTextureEx(tex, rl.NewVector2(x+UiButtonTinyW+spacing+iconSize/2-texScale*float32(tex.Width/2), y+UiButtonFlatH/2-texScale*float32(tex.Height)/2-3), 0, texScale, rl.White)
	rl.DrawText(text, int32(x+UiButtonTinyW+spacing+13), int32(y+UiButtonFlatH/2+iconSize/2-3), 16, rl.Black)
	var res = 0
	if rg.Button(rl.NewRectangle(x, y, UiButtonTinyW, UiButtonFlatH), "--") {
		res--
	}
	if rg.Button(rl.NewRectangle(x+UiButtonTinyW+iconSize+2*spacing, y, UiButtonTinyW, UiButtonFlatH), "++") {
		res++
	}
	return res
}

func (s *UiState) RenderShop() {
	const fontSize = 20
	const unlockButtonW = 88
	const incButtonH = 30
	var posX = float32(rl.GetScreenWidth()/2 - UiShopWidth/2)
	var posY = float32(180)
	rl.DrawText("Shop", int32(posX), int32(posY-incButtonH-UiMarginSmall), fontSize, rl.Black)
	for i := 0; i < len(s.shop.entries); i++ {
		var entry = &s.shop.entries[i]
		var posX = posX
		if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, incButtonH), "$W") {
			*s.shop.WhiteMoney() -= entry.price
			entry.price++
			s.shop.unlockedCount++
		}
		posX += UiButtonH + UiMarginSmall
		if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, incButtonH), "$B") {
			*s.shop.BlackMoney() -= entry.price
			entry.price++
			s.shop.unlockedCount++
		}
		posX += UiButtonH + UiMarginSmall
		if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, incButtonH), "++") {
			entry.price++
		}
		posX += UiButtonH + UiMarginSmall
		if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, incButtonH), "--") && entry.price > 0 {
			entry.price--
		}
		posX += UiButtonH + UiMargin
		var text = fmt.Sprint(entry.price, ": ", entry.description)
		var color = rl.Black
		if i >= s.shop.unlockedCount {
			color = rl.LightGray
		}
		rl.DrawText(text, int32(posX), int32(posY+incButtonH/2-fontSize/2), fontSize, color)
		posY += incButtonH + UiMarginSmall
	}
	posY += UiMarginSmall
	if rg.Button(rl.NewRectangle(posX, posY, unlockButtonW, UiButtonH), "Unlock") && s.shop.unlockedCount < len(s.shop.entries) {
		s.shop.unlockedCount++
	}
	if rg.Button(rl.NewRectangle(posX+unlockButtonW+UiMarginSmall, posY, unlockButtonW, UiButtonH), "Lock") && s.shop.unlockedCount > 0 {
		s.shop.unlockedCount--
	}
	if rg.Button(rl.NewRectangle(posX+2*unlockButtonW+2*UiMarginSmall, posY, unlockButtonW, UiButtonH), "Shuffle") {
		s.shop.Shuffle()
	}
}

func (s *UiState) RenderMoneyWidget() {
	const fontSize = 20
	const unlockButtonW = 88
	const incButtonH = 30
	var posX = float32(rl.GetScreenWidth()/2 - UiShopWidth/2)
	var posY = float32(UiMargin)

	// Row 1 (White)
	if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, incButtonH), "++") {
		*s.shop.WhiteMoney()++
	}
	if rg.Button(rl.NewRectangle(posX+UiButtonH+UiMarginSmall, posY, UiButtonH, incButtonH), "--") {
		*s.shop.WhiteMoney()--
	}
	rl.DrawText(fmt.Sprint("White money: ", *s.shop.WhiteMoney()), int32(posX+2*UiButtonH+UiMarginSmall+UiMargin), int32(posY+incButtonH/2-fontSize/2), fontSize, rl.Black)
	s.showShop = rg.Toggle(rl.NewRectangle(posX+UiShopWidth-unlockButtonW, posY, unlockButtonW, UiButtonH), "Shop", s.showShop)
	if rg.Button(rl.NewRectangle(posX+UiShopWidth-2*unlockButtonW-UiMarginSmall, posY, unlockButtonW, UiButtonH), "++Both") {
		*s.shop.WhiteMoney()++
		*s.shop.BlackMoney()++
	}
	posY += incButtonH + UiMarginSmall

	// Row 2 (Black)
	if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, incButtonH), "++") {
		*s.shop.BlackMoney()++
	}
	if rg.Button(rl.NewRectangle(posX+UiButtonH+UiMarginSmall, posY, UiButtonH, incButtonH), "--") {
		*s.shop.BlackMoney()--
	}
	rl.DrawText(fmt.Sprint("Black money: ", *s.shop.BlackMoney()), int32(posX+2*UiButtonH+UiMarginSmall+UiMargin), int32(posY+incButtonH/2-fontSize/2), fontSize, rl.Black)
	posY += incButtonH + UiMarginSmall
}
