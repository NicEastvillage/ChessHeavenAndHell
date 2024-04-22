package main

import "math/rand"

type CreatePieceCmd struct {
	piece             Piece
	anyCaptured       bool
	capturedPiece     uint32
	captureAfterCoord Vec2
}

func NewCreatePieceCmd(sb *Sandbox, typ uint32, color PieceColor, board uint32, coord Vec2) *CreatePieceCmd {
	var cmd = CreatePieceCmd{}
	if IsOffBoard(coord) {
		board = OffBoard
	}
	var captured = sandbox.GetPieceAt(coord, board)
	if captured != nil {
		cmd.anyCaptured = true
		cmd.capturedPiece = captured.Id
		cmd.captureAfterCoord = sb.FindUnoccupiedOffBoardCoord()
		captured.Board = OffBoard
		captured.Coord = cmd.captureAfterCoord
	}
	cmd.piece = *sb.NewPiece(typ, color, board, coord)
	return &cmd
}

func (cmd *CreatePieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.Board = OffBoard
		captured.Coord = cmd.captureAfterCoord
	}
	if cmd.piece.Board != OffBoard {
		ui.board = int32(cmd.piece.Board)
	}
	ui.selection.Deselect()
	ui.tab = TabBoard
}

func (cmd *CreatePieceCmd) undo(sb *Sandbox, ui *UiState) {
	if id, ok := ui.selection.GetSelectedPieceId(); ok && id == cmd.piece.Id {
		ui.selection.Deselect()
	}
	sb.RemovePiece(cmd.piece.Id)
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.Board = cmd.piece.Board
		captured.Coord = cmd.piece.Coord
	}
	if cmd.piece.Board != OffBoard {
		ui.board = int32(cmd.piece.Board)
	}
	ui.selection.Deselect()
	ui.tab = TabBoard
}

type DeletePieceCmd struct {
	piece   Piece
	effects []uint32
}

func NewDeletePieceCmd(sb *Sandbox, ui *UiState, id uint32) *DeletePieceCmd {
	var cmd = DeletePieceCmd{
		piece:   *sb.GetPiece(id),
		effects: sb.GetStatusEffectsOnPiece(id),
	}

	sb.RemovePiece(id)
	ui.selection.Deselect()
	return &cmd
}

func (cmd *DeletePieceCmd) redo(sb *Sandbox, ui *UiState) {
	if id, ok := ui.selection.GetSelectedPieceId(); ok && id == cmd.piece.Id {
		ui.selection.Deselect()
	}
	sb.RemovePiece(cmd.piece.Id)
	if cmd.piece.Board != OffBoard {
		ui.board = int32(cmd.piece.Board)
	}
	ui.tab = TabBoard
}

func (cmd *DeletePieceCmd) undo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
	for _, effect := range cmd.effects {
		sb.NewStatusEffect(cmd.piece.Id, effect)
	}
	if cmd.piece.Board != OffBoard {
		ui.board = int32(cmd.piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece.Id)
	ui.tab = TabBoard
}

type MovePieceCmd struct {
	piece             uint32
	beforeCoord       Vec2
	beforeBoard       uint32
	afterCoord        Vec2
	afterBoard        uint32
	anyCaptured       bool
	capturedPiece     uint32
	captureAfterCoord Vec2
}

func NewMovePieceCmd(sb *Sandbox, id uint32, destCoord Vec2, destBoard uint32) *MovePieceCmd {
	var piece = sb.GetPiece(id)
	if piece.Board == destBoard && piece.Coord == destCoord {
		panic("A piece cannot be moved to where it already is")
	}
	var cmd = MovePieceCmd{
		piece:       id,
		beforeCoord: piece.Coord,
		beforeBoard: piece.Board,
	}
	if IsOffBoard(destCoord) {
		destBoard = OffBoard
	}
	var captured = sandbox.GetPieceAt(destCoord, destBoard)
	if captured != nil {
		cmd.anyCaptured = true
		cmd.capturedPiece = captured.Id
		cmd.captureAfterCoord = sb.FindUnoccupiedOffBoardCoord()
		captured.Board = OffBoard
		captured.Coord = cmd.captureAfterCoord
	}
	cmd.afterBoard = destBoard
	cmd.afterCoord = destCoord
	piece.Coord = destCoord
	piece.Board = destBoard
	return &cmd
}

func (cmd *MovePieceCmd) redo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.Coord = cmd.afterCoord
	piece.Board = cmd.afterBoard
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.Board = OffBoard
		captured.Coord = cmd.captureAfterCoord
	}
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *MovePieceCmd) undo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.Coord = cmd.beforeCoord
	piece.Board = cmd.beforeBoard
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.Board = cmd.afterBoard
		captured.Coord = cmd.afterCoord
	}
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type CreateStatusEffectCmd struct {
	piece  uint32
	effect uint32
}

func NewCreateStatusEffectCmd(sb *Sandbox, piece uint32, effect uint32) *CreateStatusEffectCmd {
	sb.NewStatusEffect(piece, effect)
	return &CreateStatusEffectCmd{
		piece:  piece,
		effect: effect,
	}
}

func (cmd *CreateStatusEffectCmd) redo(sb *Sandbox, ui *UiState) {
	sb.NewStatusEffect(cmd.piece, cmd.effect)
	var piece = sb.GetPiece(cmd.piece)
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *CreateStatusEffectCmd) undo(sb *Sandbox, ui *UiState) {
	sb.RemoveStatusEffect(cmd.piece, cmd.effect)
	var piece = sb.GetPiece(cmd.piece)
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type DeleteStatusEffectCmd struct {
	piece  uint32
	effect uint32
}

func NewDeleteStatusEffectCmd(sb *Sandbox, piece uint32, effect uint32) *DeleteStatusEffectCmd {
	sb.RemoveStatusEffect(piece, effect)
	return &DeleteStatusEffectCmd{
		piece:  piece,
		effect: effect,
	}
}

func (cmd *DeleteStatusEffectCmd) redo(sb *Sandbox, ui *UiState) {
	sb.RemoveStatusEffect(cmd.piece, cmd.effect)
	var piece = sb.GetPiece(cmd.piece)
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *DeleteStatusEffectCmd) undo(sb *Sandbox, ui *UiState) {
	sb.NewStatusEffect(cmd.piece, cmd.effect)
	var piece = sb.GetPiece(cmd.piece)
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type IncreasePieceScaleCmd struct {
	piece uint32
}

func NewIncreasePieceScaleCmd(sb *Sandbox, piece uint32) *IncreasePieceScaleCmd {
	sb.GetPiece(piece).Scale++
	return &IncreasePieceScaleCmd{
		piece: piece,
	}
}

func (cmd *IncreasePieceScaleCmd) redo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.Scale++
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *IncreasePieceScaleCmd) undo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.Scale--
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type DecreasePieceScaleCmd struct {
	piece uint32
}

func NewDecreasePieceScaleCmd(sb *Sandbox, piece uint32) *DecreasePieceScaleCmd {
	sb.GetPiece(piece).Scale--
	return &DecreasePieceScaleCmd{
		piece: piece,
	}
}

func (cmd *DecreasePieceScaleCmd) redo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.Scale--
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *DecreasePieceScaleCmd) undo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.Scale++
	if piece.Board != OffBoard {
		ui.board = int32(piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type CreateTileCmd struct {
	board uint32
	coord Vec2
}

func NewCreateTileCmd(sb *Sandbox, board uint32, coord Vec2) *CreateTileCmd {
	sb.NewTile(board, coord)
	return &CreateTileCmd{
		board: board,
		coord: coord,
	}
}

func (cmd *CreateTileCmd) redo(sb *Sandbox, ui *UiState) {
	sb.NewTile(cmd.board, cmd.coord)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
	ui.tab = TabBoard
}

func (cmd *CreateTileCmd) undo(sb *Sandbox, ui *UiState) {
	sb.RemoveTile(cmd.board, cmd.coord)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
	ui.tab = TabBoard
}

type DeleteTileCmd struct {
	board uint32
	coord Vec2
}

func NewDeleteTileCmd(sb *Sandbox, board uint32, coord Vec2) *DeleteTileCmd {
	sb.RemoveTile(board, coord)
	return &DeleteTileCmd{
		board: board,
		coord: coord,
	}
}

func (cmd *DeleteTileCmd) redo(sb *Sandbox, ui *UiState) {
	sb.RemoveTile(cmd.board, cmd.coord)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
	ui.tab = TabBoard
}

func (cmd *DeleteTileCmd) undo(sb *Sandbox, ui *UiState) {
	sb.NewTile(cmd.board, cmd.coord)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
	ui.tab = TabBoard
}

type CreateObstacleCmd struct {
	coord    Vec2
	board    uint32
	obstacle uint32
}

func NewCreateObstacleCmd(sb *Sandbox, coord Vec2, board uint32, obstacle uint32) *CreateObstacleCmd {
	sb.NewObstacle(coord, board, obstacle)
	return &CreateObstacleCmd{
		obstacle: obstacle,
		board:    board,
		coord:    coord,
	}
}

func (cmd *CreateObstacleCmd) redo(sb *Sandbox, ui *UiState) {
	sb.NewObstacle(cmd.coord, cmd.board, cmd.obstacle)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
	ui.tab = TabBoard
}

func (cmd *CreateObstacleCmd) undo(sb *Sandbox, ui *UiState) {
	sb.RemoveObstacle(cmd.coord, cmd.board, cmd.obstacle)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
	ui.tab = TabBoard
}

type DeleteObstacleCmd struct {
	coord    Vec2
	board    uint32
	obstacle uint32
}

func NewDeleteObstacleCmd(sb *Sandbox, coord Vec2, board uint32, obstacle uint32) *DeleteObstacleCmd {
	sb.RemoveObstacle(coord, board, obstacle)
	return &DeleteObstacleCmd{
		obstacle: obstacle,
		board:    board,
		coord:    coord,
	}
}

func (cmd *DeleteObstacleCmd) redo(sb *Sandbox, ui *UiState) {
	sb.RemoveObstacle(cmd.coord, cmd.board, cmd.obstacle)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
	ui.tab = TabBoard
}

func (cmd *DeleteObstacleCmd) undo(sb *Sandbox, ui *UiState) {
	sb.NewObstacle(cmd.coord, cmd.board, cmd.obstacle)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
	ui.tab = TabBoard
}

type PastePieceCmd struct {
	piece             Piece
	effects           []uint32
	anyCaptured       bool
	capturedPiece     uint32
	captureAfterCoord Vec2
}

func NewPastePieceCmd(sb *Sandbox, ui *UiState, coord Vec2, board uint32) *PastePieceCmd {
	var cmd = PastePieceCmd{}
	if IsOffBoard(coord) {
		board = OffBoard
	}
	var captured = sandbox.GetPieceAt(coord, board)
	if captured != nil {
		cmd.anyCaptured = true
		cmd.capturedPiece = captured.Id
		cmd.captureAfterCoord = sb.FindUnoccupiedOffBoardCoord()
		captured.Board = OffBoard
		captured.Coord = cmd.captureAfterCoord
	}
	var piece = sb.NewPiece(ui.clipboard.typ, ui.clipboard.color, board, coord)
	piece.Scale = ui.clipboard.scale
	cmd.piece = *piece
	cmd.effects = make([]uint32, len(ui.clipboard.effects))
	copy(cmd.effects, ui.clipboard.effects)
	for _, effect := range cmd.effects {
		sb.NewStatusEffect(piece.Id, effect)
	}
	ui.selection.SelectPiece(cmd.piece.Id)
	return &cmd
}

func (cmd *PastePieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
	for _, effect := range cmd.effects {
		sb.NewStatusEffect(cmd.piece.Id, effect)
	}
	if cmd.piece.Board != OffBoard {
		ui.board = int32(cmd.piece.Board)
	}
	ui.selection.SelectPiece(cmd.piece.Id)
	ui.tab = TabBoard
}

func (cmd *PastePieceCmd) undo(sb *Sandbox, ui *UiState) {
	if id, ok := ui.selection.GetSelectedPieceId(); ok && id == cmd.piece.Id {
		ui.selection.Deselect()
	}
	sb.RemovePiece(cmd.piece.Id)
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.Board = cmd.piece.Board
		captured.Coord = cmd.piece.Coord
	}
	if cmd.piece.Board != OffBoard {
		ui.board = int32(cmd.piece.Board)
	}
	ui.tab = TabBoard
}

type DuplicatePieceCmd struct {
	originalId uint32
	piece      Piece
	effects    []uint32
}

func NewDuplicatePieceCmd(sb *Sandbox, ui *UiState, id uint32) *DuplicatePieceCmd {
	var original = sb.GetPiece(id)
	var coord = sb.FindUnoccupiedOffBoardCoord()
	var piece = sb.NewPiece(original.Typ, original.Color, OffBoard, coord)
	var effects = sb.GetStatusEffectsOnPiece(id)
	for _, effect := range effects {
		sb.NewStatusEffect(piece.Id, effect)
	}
	ui.selection.SelectPiece(piece.Id)
	return &DuplicatePieceCmd{
		originalId: id,
		piece:      *piece,
		effects:    effects,
	}
}

func (cmd *DuplicatePieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
	for _, effect := range cmd.effects {
		sb.NewStatusEffect(cmd.piece.Id, effect)
	}
	ui.selection.SelectPiece(cmd.piece.Id)
	ui.tab = TabBoard
}

func (cmd *DuplicatePieceCmd) undo(sb *Sandbox, ui *UiState) {
	sb.RemovePiece(cmd.piece.Id)
	ui.selection.SelectPiece(cmd.originalId)
	ui.tab = TabBoard
}

type ChangeColorOfPieceCmd struct {
	piece  uint32
	before PieceColor
	after  PieceColor
}

func NewChangeColorOfPieceCmd(sb *Sandbox, id uint32, color PieceColor) *ChangeColorOfPieceCmd {
	piece := sb.GetPiece(id)
	var before = piece.Color
	piece.Color = color
	return &ChangeColorOfPieceCmd{
		piece:  id,
		before: before,
		after:  color,
	}
}

func (cmd *ChangeColorOfPieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.GetPiece(cmd.piece).Color = cmd.after
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *ChangeColorOfPieceCmd) undo(sb *Sandbox, ui *UiState) {
	sb.GetPiece(cmd.piece).Color = cmd.before
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type ChangeTypeOfPieceCmd struct {
	piece  uint32
	before uint32
	after  uint32
}

func NewChangeTypeOfPieceCmd(sb *Sandbox, id uint32, typ uint32) *ChangeTypeOfPieceCmd {
	piece := sb.GetPiece(id)
	var before = piece.Typ
	piece.Typ = typ
	return &ChangeTypeOfPieceCmd{
		piece:  id,
		before: before,
		after:  typ,
	}
}

func (cmd *ChangeTypeOfPieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.GetPiece(cmd.piece).Typ = cmd.after
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *ChangeTypeOfPieceCmd) undo(sb *Sandbox, ui *UiState) {
	sb.GetPiece(cmd.piece).Typ = cmd.before
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type ChangeMoneyAmountCmd struct {
	whiteBefore int
	whiteAfter  int
	blackBefore int
	blackAfter  int
}

func NewChangeMoneyAmountCmd(sb *Sandbox, whiteNewAmount int, blackNewAmount int) *ChangeMoneyAmountCmd {
	var whiteBefore = sb.Shop.Money[0]
	var blackBefore = sb.Shop.Money[1]
	sb.Shop.Money[0] = whiteNewAmount
	sb.Shop.Money[1] = blackNewAmount
	return &ChangeMoneyAmountCmd{
		whiteBefore: whiteBefore,
		whiteAfter:  whiteNewAmount,
		blackBefore: blackBefore,
		blackAfter:  blackNewAmount,
	}
}

func (cmd *ChangeMoneyAmountCmd) redo(sb *Sandbox, ui *UiState) {
	sb.Shop.Money[0] = cmd.whiteAfter
	sb.Shop.Money[1] = cmd.blackAfter
}

func (cmd *ChangeMoneyAmountCmd) undo(sb *Sandbox, ui *UiState) {
	sb.Shop.Money[0] = cmd.whiteBefore
	sb.Shop.Money[1] = cmd.blackBefore
}

type ChangeShopUnlockCountCmd struct {
	before int
	after  int
}

func NewChangeShopUnlockCountCmd(sb *Sandbox, newCount int) *ChangeShopUnlockCountCmd {
	var before = sb.Shop.UnlockedCount
	sb.Shop.UnlockedCount = newCount
	return &ChangeShopUnlockCountCmd{
		before: before,
		after:  newCount,
	}
}

func (cmd *ChangeShopUnlockCountCmd) redo(sb *Sandbox, ui *UiState) {
	sb.Shop.UnlockedCount = cmd.after
	ui.tab = TabShop
}

func (cmd *ChangeShopUnlockCountCmd) undo(sb *Sandbox, ui *UiState) {
	sb.Shop.UnlockedCount = cmd.before
	ui.tab = TabShop
}

type ShuffleShopCmd struct {
	swaps []int
}

func NewShuffleShopCmd(shop *Shop) *ShuffleShopCmd {
	// Fisher-Yates shuffle, but we store which indices we roll, so we can undo it
	var swaps = make([]int, 0)
	for i := len(shop.Entries) - 1; i >= 1; i-- {
		var j = rand.Intn(i + 1)
		swaps = append(swaps, j)
		shop.Entries[i], shop.Entries[j] = shop.Entries[j], shop.Entries[i]
	}
	return &ShuffleShopCmd{
		swaps: swaps,
	}
}

func (cmd *ShuffleShopCmd) redo(sb *Sandbox, ui *UiState) {
	for i := len(sb.Shop.Entries) - 1; i >= 1; i-- {
		var j = cmd.swaps[len(sb.Shop.Entries)-1-i]
		sb.Shop.Entries[i], sb.Shop.Entries[j] = sb.Shop.Entries[j], sb.Shop.Entries[i]
	}
	ui.tab = TabShop
}

func (cmd *ShuffleShopCmd) undo(sb *Sandbox, ui *UiState) {
	for i := 1; i < len(sb.Shop.Entries); i++ {
		var j = cmd.swaps[len(sb.Shop.Entries)-1-i]
		sb.Shop.Entries[i], sb.Shop.Entries[j] = sb.Shop.Entries[j], sb.Shop.Entries[i]
	}
	ui.tab = TabShop
}

type ChangeShopEntryPriceCmd struct {
	entry       uint32
	priceBefore int
	priceAfter  int
}

func NewChangeShopEntryPriceCmd(shop *Shop, entry uint32, newPrice int) *ChangeShopEntryPriceCmd {
	var e = shop.GetEntry(entry)
	var priceBefore = e.Price
	e.Price = newPrice
	return &ChangeShopEntryPriceCmd{
		entry:       entry,
		priceBefore: priceBefore,
		priceAfter:  newPrice,
	}
}

func (cmd *ChangeShopEntryPriceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.Shop.GetEntry(cmd.entry).Price = cmd.priceAfter
	ui.tab = TabShop
}

func (cmd *ChangeShopEntryPriceCmd) undo(sb *Sandbox, ui *UiState) {
	sb.Shop.GetEntry(cmd.entry).Price = cmd.priceBefore
	ui.tab = TabShop
}

type QuickBuyCmd struct {
	player uint32
	entry  uint32
	unlock bool
}

func NewQuickBuyCmd(shop *Shop, player uint32, entry uint32) *QuickBuyCmd {
	var e = shop.GetEntry(entry)
	shop.Money[player] -= e.Price
	e.Price++
	var unlock = shop.UnlockedCount < len(shop.Entries)
	if unlock {
		shop.UnlockedCount++
	}
	return &QuickBuyCmd{
		player: player,
		entry:  entry,
		unlock: unlock,
	}
}

func (cmd *QuickBuyCmd) redo(sb *Sandbox, ui *UiState) {
	var e = sb.Shop.GetEntry(cmd.entry)
	sb.Shop.Money[cmd.player] -= e.Price
	e.Price++
	if cmd.unlock {
		sb.Shop.UnlockedCount++
	}
	ui.tab = TabShop
}

func (cmd *QuickBuyCmd) undo(sb *Sandbox, ui *UiState) {
	if cmd.unlock {
		sb.Shop.UnlockedCount--
	}
	var e = sb.Shop.GetEntry(cmd.entry)
	e.Price--
	sb.Shop.Money[cmd.player] += e.Price
	ui.tab = TabShop
}
