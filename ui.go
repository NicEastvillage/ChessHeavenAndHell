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
	UiShopTopMargin  = 140
	UiFontSize       = 20
	UiFontSizeBig    = 28
)

const (
	TabBoard int32 = iota
	TabShop
	TabRng
)

type UiState struct {
	selection          Selection
	clipboard          Clipboard
	shop               Shop
	rng                RngStuff
	tab                int32
	board              int32
	mode               int32
	color              int32
	showEffectsOrTypes int32
	renderTexHeaven    rl.RenderTexture2D
	renderTexEarth     rl.RenderTexture2D
	renderTexHell      rl.RenderTexture2D
}

func NewUiState() UiState {
	return UiState{
		selection:       NewSelection(),
		clipboard:       NewClipboard(),
		shop:            NewShop(),
		rng:             NewRngStuff(),
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

	if s.tab != TabBoard {
		s.selection.Deselect()
	}

	if s.board != 0 || s.tab != TabBoard {
		rl.BeginTextureMode(s.renderTexHeaven)
		rl.ClearBackground(rl.RayWhite)
		//rl.Translatef(previewSourceOrigo.X, previewSourceOrigo.Y, 0)
		sandbox.Render(0, true, &s.selection)
		rl.EndTextureMode()
	}
	if s.board != 1 || s.tab != TabBoard {
		rl.BeginTextureMode(s.renderTexEarth)
		rl.ClearBackground(rl.RayWhite)
		sandbox.Render(1, true, &s.selection)
		rl.EndTextureMode()
	}
	if s.board != 2 || s.tab != TabBoard {
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

	s.RenderMoneyWidget(undo)
	if rl.IsKeyPressed(rl.KeyS) {
		if s.tab == TabShop {
			s.tab = TabBoard
		} else {
			s.tab = TabShop
		}
	}
	switch s.tab {
	case TabShop:
		s.RenderShop(undo)
	case TabRng:
		s.RenderRngMenu(undo)
	default:
		s.RenderBoardUi(undo)
	}
}

func (s *UiState) RenderBoardPreview(index int32) {
	if s.board == index && s.tab == TabBoard {
		// Main board. Do not render
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
		s.tab = TabBoard
	}
	rl.DrawTexturePro(previewTex, previewSourceRect, previewPlacement, rl.NewVector2(0, 0), 0, rl.White)
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), previewPlacement) {
		rl.DrawRectangleLinesEx(previewPlacement, 1, rl.GetColor(uint(rg.GetStyle(rg.BUTTON, rg.BORDER_COLOR_FOCUSED))))
	}
}

func (s *UiState) RenderBoardUi(undo *UndoRedoSystem) {
	var oldMode = s.mode
	s.mode = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-2*UiButtonH-1*int(rg.GetStyle(rg.TOGGLE, rg.GROUP_PADDING))), UiMargin, UiButtonH, UiButtonH), "#149#;#97#", s.mode)
	if rl.IsKeyPressed(rl.KeyT) {
		s.mode = 1 - s.mode
	}
	if oldMode != s.mode {
		s.selection.Deselect()
	}

	switch s.selection.selectionType {
	case SelectionTypePiece:
		s.RenderPieceContextMenu(undo)
		s.mode = 0
	case SelectionTypePieceType:
		s.RenderPiecesMode()
		s.mode = 0
	case SelectionTypeCoord:
		s.RenderCoordContextMenu(undo)
		s.mode = 1
	default:
		if s.mode == 0 {
			s.RenderPiecesMode()
		}
	}
}

func (s *UiState) RenderPiecesMode() {
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

	s.showEffectsOrTypes = rg.ToggleSlider(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-UiButtonW), 2*UiMargin+UiButtonH, UiButtonW, UiButtonH), "Effect;Types", s.showEffectsOrTypes)
	if rl.IsKeyPressed(rl.KeyC) {
		s.color = 1 - s.color
	}

	if s.showEffectsOrTypes == 1 {
		// Show types
		for i := 0; i < len(sandbox.pieceTypes); i++ {
			if rg.Toggle(rl.NewRectangle(float32(rl.GetScreenWidth()-UiMargin-UiButtonW), float32(3*UiMargin+2*UiButtonH+i*(UiMarginSmall+UiButtonH)), UiButtonW, UiButtonH), sandbox.GetPieceType(uint32(i)).name, s.selection.IsPieceTypeSelected(uint32(i))) {
				if piece.typ != sandbox.GetPieceType(uint32(i)).id {
					undo.Append(NewChangeTypeOfPieceCmd(&sandbox, piece.id, sandbox.GetPieceType(uint32(i)).id))
				}
			}
		}
	} else {
		// Show effects
		var spinnerX = float32(rl.GetScreenWidth() - 150)
		var spinnerY = float32(3*UiMargin + 2*UiButtonH)

		{
			var pieceScale = piece.scale
			var change = SpinnerWithIcon(spinnerX, spinnerY, fmt.Sprint(pieceScale), assets.texPieceScale)
			if change < 0 && pieceScale > 1 {
				undo.Append(NewDecreasePieceScaleCmd(&sandbox, selectedPieceId))
			}
			if change > 0 {
				undo.Append(NewIncreasePieceScaleCmd(&sandbox, selectedPieceId))
			}
		}

		for i := range sandbox.effectTypes {
			var effect = &sandbox.effectTypes[i]
			var effectCount = sandbox.GetStatusEffectCount(selectedPieceId, effect.id)
			var change = SpinnerWithIcon(spinnerX, spinnerY+float32(i*55)+55, fmt.Sprint(effectCount), effect.tex)
			if change < 0 && effectCount > 0 {
				undo.Append(NewDeleteStatusEffectCmd(&sandbox, selectedPieceId, effect.id))
			}
			if change > 0 {
				undo.Append(NewCreateStatusEffectCmd(&sandbox, selectedPieceId, effect.id))
			}
		}
	}

	var posX = float32(rl.GetScreenWidth() - UiButtonW - UiMargin)
	var posY = float32(rl.GetScreenHeight() - UiMargin - UiButtonH)
	var width = float32(UiButtonW)
	var height float32 = UiButtonH
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Remove piece") {
		undo.Append(NewDeletePieceCmd(&sandbox, s, selectedPieceId))
	}

	posY -= UiButtonH + UiMargin
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Change color") {
		var newColor = 1 - piece.color
		undo.Append(NewChangeColorOfPieceCmd(&sandbox, selectedPieceId, newColor))
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
			undo.Append(NewDeleteObstacleCmd(&sandbox, coord, uint32(s.board), obt.id))
		}
		if change > 0 {
			undo.Append(NewCreateObstacleCmd(&sandbox, coord, uint32(s.board), obt.id))
		}
	}

	var posX = float32(rl.GetScreenWidth() - UiButtonW - UiMargin)
	var posY = float32(rl.GetScreenHeight() - UiMargin - UiButtonH)
	var width = float32(UiButtonW)
	var height float32 = UiButtonH
	if rg.Button(rl.NewRectangle(posX, posY-UiMarginSmall-UiButtonH, width, height), "Add tile") {
		if sandbox.GetTile(uint32(s.board), coord) == nil {
			undo.Append(NewCreateTileCmd(&sandbox, uint32(s.board), coord))
		}
	}
	if rg.Button(rl.NewRectangle(posX, posY, width, height), "Remove tile") {
		undo.Append(NewDeleteTileCmd(&sandbox, uint32(s.board), coord))
	}
}

func SpinnerWithIcon(x float32, y float32, text string, tex rl.Texture2D) int {
	const (
		spacing  = UiMarginSmall
		iconSize = 32
	)

	var texScale = iconSize / float32(tex.Height)
	rl.DrawTextureEx(tex, rl.NewVector2(x+UiButtonTinyW+spacing+iconSize/2-texScale*float32(tex.Width/2), y+UiButtonFlatH/2-texScale*float32(tex.Height)/2-3), 0, texScale, rl.White)
	rl.DrawTextEx(assets.fontComicSansMs, text, rl.NewVector2(x+UiButtonTinyW+spacing+13, y+UiButtonFlatH/2+iconSize/2-3), 20, 1, rl.Black)
	var res = 0
	if rg.Button(rl.NewRectangle(x, y, UiButtonTinyW, UiButtonFlatH), "--") {
		res--
	}
	if rg.Button(rl.NewRectangle(x+UiButtonTinyW+iconSize+2*spacing, y, UiButtonTinyW, UiButtonFlatH), "++") {
		res++
	}
	return res
}

func (s *UiState) RenderShop(undo *UndoRedoSystem) {
	var posX = float32(rl.GetScreenWidth()/2 - UiShopWidth/2)
	var posY = float32(UiShopTopMargin)
	rl.DrawTextEx(assets.fontComicSansMsBig, "Shop", rl.NewVector2(posX, posY), UiFontSizeBig, 1, rl.Black)
	posY += UiButtonFlatH + UiMarginSmall
	for i := 0; i < len(s.shop.entries); i++ {
		var entry = &s.shop.entries[i]
		var posX = posX
		if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, UiButtonFlatH), "$W") {
			undo.Append(NewQuickBuyCmd(&s.shop, 0, entry.id))
		}
		posX += UiButtonH + UiMarginSmall
		if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, UiButtonFlatH), "$B") {
			undo.Append(NewQuickBuyCmd(&s.shop, 1, entry.id))
		}
		posX += UiButtonH + UiMarginSmall
		if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, UiButtonFlatH), "++") {
			undo.Append(NewChangeShopEntryPriceCmd(&s.shop, entry.id, entry.price+1))
		}
		posX += UiButtonH + UiMarginSmall
		if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, UiButtonFlatH), "--") && entry.price > 0 {
			undo.Append(NewChangeShopEntryPriceCmd(&s.shop, entry.id, entry.price-1))
		}
		posX += UiButtonH + UiMargin
		var text = fmt.Sprint(entry.price, ": ", entry.description)
		var color = rl.Black
		if i >= s.shop.unlockedCount {
			color = rl.LightGray
		}
		rl.DrawTextEx(assets.fontComicSansMsBig, text, rl.NewVector2(posX, posY+UiButtonFlatH/2-UiFontSizeBig/2), UiFontSizeBig, 1, color)
		posY += UiButtonFlatH + UiMarginSmall
	}
	posY += UiMarginSmall
	if rg.Button(rl.NewRectangle(posX, posY, UiButtonNarrowW, UiButtonH), "Unlock") && s.shop.unlockedCount < len(s.shop.entries) {
		undo.Append(NewChangeShopUnlockCountCmd(s, s.shop.unlockedCount+1))
	}
	if rg.Button(rl.NewRectangle(posX+UiButtonNarrowW+UiMarginSmall, posY, UiButtonNarrowW, UiButtonH), "Lock") && s.shop.unlockedCount > 0 {
		undo.Append(NewChangeShopUnlockCountCmd(s, s.shop.unlockedCount-1))
	}
	if rg.Button(rl.NewRectangle(posX+2*UiButtonNarrowW+2*UiMarginSmall, posY, UiButtonNarrowW, UiButtonH), "Shuffle") {
		undo.Append(NewShuffleShopCmd(&s.shop))
	}
}

func (s *UiState) RenderMoneyWidget(undo *UndoRedoSystem) {
	var posX = float32(rl.GetScreenWidth()/2 - UiShopWidth/2)
	var posY = float32(UiMargin)

	// Row 1 (White)
	if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, UiButtonFlatH), "++") {
		undo.Append(NewChangeMoneyAmountCmd(s, *s.shop.WhiteMoney()+1, *s.shop.BlackMoney()))
	}
	if rg.Button(rl.NewRectangle(posX+UiButtonH+UiMarginSmall, posY, UiButtonH, UiButtonFlatH), "--") {
		undo.Append(NewChangeMoneyAmountCmd(s, *s.shop.WhiteMoney()-1, *s.shop.BlackMoney()))
	}
	rl.DrawTextEx(assets.fontComicSansMsBig, fmt.Sprint("White money: ", *s.shop.WhiteMoney()), rl.NewVector2(posX+2*UiButtonH+UiMarginSmall+UiMargin, posY+UiButtonFlatH/2-UiFontSizeBig/2), UiFontSizeBig, 1, rl.Black)
	var inRngTab = rg.Toggle(rl.NewRectangle(posX+UiShopWidth-UiButtonNarrowW, posY, UiButtonNarrowW, UiButtonH), "Rng", s.tab == TabRng)
	var inShopTab = rg.Toggle(rl.NewRectangle(posX+UiShopWidth-2*UiButtonNarrowW-UiMarginSmall, posY, UiButtonNarrowW, UiButtonH), "Shop", s.tab == TabShop)
	if inRngTab && inShopTab {
		if s.tab == TabShop {
			s.tab = TabRng
		} else {
			s.tab = TabShop
		}
	} else if inShopTab {
		s.tab = TabShop
	} else if inRngTab {
		s.tab = TabRng
	} else {
		s.tab = TabBoard
	}
	if rg.Button(rl.NewRectangle(posX+UiShopWidth-3*UiButtonNarrowW-2*UiMarginSmall, posY, UiButtonNarrowW, UiButtonH), "++Both") {
		undo.Append(NewChangeMoneyAmountCmd(s, *s.shop.WhiteMoney()+1, *s.shop.BlackMoney()+1))
	}
	posY += UiButtonFlatH + UiMarginSmall

	// Row 2 (Black)
	if rg.Button(rl.NewRectangle(posX, posY, UiButtonH, UiButtonFlatH), "++") {
		undo.Append(NewChangeMoneyAmountCmd(s, *s.shop.WhiteMoney(), *s.shop.BlackMoney()+1))
	}
	if rg.Button(rl.NewRectangle(posX+UiButtonH+UiMarginSmall, posY, UiButtonH, UiButtonFlatH), "--") {
		undo.Append(NewChangeMoneyAmountCmd(s, *s.shop.WhiteMoney(), *s.shop.BlackMoney()-1))
	}
	rl.DrawTextEx(assets.fontComicSansMsBig, fmt.Sprint("Black money: ", *s.shop.BlackMoney()), rl.NewVector2(posX+2*UiButtonH+UiMarginSmall+UiMargin, posY+UiButtonFlatH/2-UiFontSizeBig/2), UiFontSizeBig, 1, rl.Black)
	posY += UiButtonFlatH + UiMarginSmall
}

func (s *UiState) RenderRngMenu(undo *UndoRedoSystem) {
	var posX = float32(rl.GetScreenWidth()/2 - UiShopWidth/2)
	var posY = float32(UiShopTopMargin)

	rl.DrawTextEx(assets.fontComicSansMsBig, "Chaos", rl.NewVector2(posX, posY), UiFontSizeBig, 1, rl.Black)
	posY += UiButtonFlatH + UiMarginSmall
	for i := 0; i < len(s.rng.chaosShown); i++ {
		var text = fmt.Sprint("- ", s.rng.chaosShown[i])
		rl.DrawTextEx(assets.fontComicSansMsBig, text, rl.NewVector2(posX, posY+UiButtonFlatH/2-UiFontSizeBig/2), UiFontSizeBig, 1, rl.Black)
		posY += UiButtonFlatH + UiMarginSmall
	}
	posY += UiMarginSmall
	if rg.Button(rl.NewRectangle(posX, posY, UiButtonNarrowW, UiButtonH), "Reroll") {
		s.rng.RerollChaosShown()
	}
	posY += 3 * UiMarginBig

	rl.DrawTextEx(assets.fontComicSansMsBig, "RNG Utils", rl.NewVector2(posX, posY), UiFontSizeBig, 1, rl.Black)
	posY += UiButtonFlatH + UiMarginSmall

	if rg.Button(rl.NewRectangle(posX, posY, UiButtonNarrowW, UiButtonH), "Reroll") {
		s.rng.RerollPiece(&sandbox)
	}
	rl.DrawTextEx(assets.fontComicSansMsBig, "Random piece: "+s.rng.piece, rl.NewVector2(posX+UiButtonNarrowW+UiMarginSmall, posY+UiButtonH/2-UiFontSizeBig/2), UiFontSizeBig, 1, rl.Black)
	posY += UiButtonH + UiMarginSmall

	if rg.Button(rl.NewRectangle(posX, posY, UiButtonNarrowW, UiButtonH), "Reroll") {
		s.rng.RerollPlane()
	}
	rl.DrawTextEx(assets.fontComicSansMsBig, "Random plane: "+s.rng.plane, rl.NewVector2(posX+UiButtonNarrowW+UiMarginSmall, posY+UiButtonH/2-UiFontSizeBig/2), UiFontSizeBig, 1, rl.Black)
	posY += UiButtonH + UiMarginSmall

	if rg.Button(rl.NewRectangle(posX, posY, UiButtonNarrowW, UiButtonH), "Reroll") {
		s.rng.RerollTile()
	}
	rl.DrawTextEx(assets.fontComicSansMsBig, "Random tile: "+s.rng.tile, rl.NewVector2(posX+UiButtonNarrowW+UiMarginSmall, posY+UiButtonH/2-UiFontSizeBig/2), UiFontSizeBig, 1, rl.Black)
	posY += UiButtonH + UiMarginSmall

	if rg.Button(rl.NewRectangle(posX, posY, UiButtonNarrowW, UiButtonH), "Reroll") {
		s.rng.RerollUnoccupiedTile(&sandbox)
	}
	rl.DrawTextEx(assets.fontComicSansMsBig, "Random unoccupied tile: "+s.rng.unoccupiedTile, rl.NewVector2(posX+UiButtonNarrowW+UiMarginSmall, posY+UiButtonH/2-UiFontSizeBig/2), UiFontSizeBig, 1, rl.Black)
	posY += UiButtonH + UiMarginSmall

	if rg.Button(rl.NewRectangle(posX, posY, UiButtonNarrowW, UiButtonH), "Reroll") {
		s.rng.RerollEmptyTile(&sandbox)
	}
	rl.DrawTextEx(assets.fontComicSansMsBig, "Random empty tile: "+s.rng.emptyTile, rl.NewVector2(posX+UiButtonNarrowW+UiMarginSmall, posY+UiButtonH/2-UiFontSizeBig/2), UiFontSizeBig, 1, rl.Black)
	posY += UiButtonH + UiMarginSmall
}
