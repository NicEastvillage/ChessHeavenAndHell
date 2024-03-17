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
		cmd.capturedPiece = captured.id
		cmd.captureAfterCoord = sb.FindUnoccupiedOffBoardCoord()
		captured.board = OffBoard
		captured.coord = cmd.captureAfterCoord
	}
	cmd.piece = *sb.NewPiece(typ, color, board, coord)
	return &cmd
}

func (cmd *CreatePieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.board = OffBoard
		captured.coord = cmd.captureAfterCoord
	}
	if cmd.piece.board != OffBoard {
		ui.board = int32(cmd.piece.board)
	}
	ui.selection.Deselect()
	ui.tab = TabBoard
}

func (cmd *CreatePieceCmd) undo(sb *Sandbox, ui *UiState) {
	if id, ok := ui.selection.GetSelectedPieceId(); ok && id == cmd.piece.id {
		ui.selection.Deselect()
	}
	sb.RemovePiece(cmd.piece.id)
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.board = cmd.piece.board
		captured.coord = cmd.piece.coord
	}
	if cmd.piece.board != OffBoard {
		ui.board = int32(cmd.piece.board)
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
	if id, ok := ui.selection.GetSelectedPieceId(); ok && id == cmd.piece.id {
		ui.selection.Deselect()
	}
	sb.RemovePiece(cmd.piece.id)
	if cmd.piece.board != OffBoard {
		ui.board = int32(cmd.piece.board)
	}
	ui.tab = TabBoard
}

func (cmd *DeletePieceCmd) undo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
	for _, effect := range cmd.effects {
		sb.NewStatusEffect(cmd.piece.id, effect)
	}
	if cmd.piece.board != OffBoard {
		ui.board = int32(cmd.piece.board)
	}
	ui.selection.SelectPiece(cmd.piece.id)
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
	if piece.board == destBoard && piece.coord == destCoord {
		panic("A piece cannot be moved to where it already is")
	}
	var cmd = MovePieceCmd{
		piece:       id,
		beforeCoord: piece.coord,
		beforeBoard: piece.board,
	}
	if IsOffBoard(destCoord) {
		destBoard = OffBoard
	}
	var captured = sandbox.GetPieceAt(destCoord, destBoard)
	if captured != nil {
		cmd.anyCaptured = true
		cmd.capturedPiece = captured.id
		cmd.captureAfterCoord = sb.FindUnoccupiedOffBoardCoord()
		captured.board = OffBoard
		captured.coord = cmd.captureAfterCoord
	}
	cmd.afterBoard = destBoard
	cmd.afterCoord = destCoord
	piece.coord = destCoord
	piece.board = destBoard
	return &cmd
}

func (cmd *MovePieceCmd) redo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.coord = cmd.afterCoord
	piece.board = cmd.afterBoard
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.board = OffBoard
		captured.coord = cmd.captureAfterCoord
	}
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *MovePieceCmd) undo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.coord = cmd.beforeCoord
	piece.board = cmd.beforeBoard
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.board = cmd.afterBoard
		captured.coord = cmd.afterCoord
	}
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
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
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *CreateStatusEffectCmd) undo(sb *Sandbox, ui *UiState) {
	sb.RemoveStatusEffect(cmd.piece, cmd.effect)
	var piece = sb.GetPiece(cmd.piece)
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
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
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *DeleteStatusEffectCmd) undo(sb *Sandbox, ui *UiState) {
	sb.NewStatusEffect(cmd.piece, cmd.effect)
	var piece = sb.GetPiece(cmd.piece)
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type IncreasePieceScaleCmd struct {
	piece uint32
}

func NewIncreasePieceScaleCmd(sb *Sandbox, piece uint32) *IncreasePieceScaleCmd {
	sb.GetPiece(piece).scale++
	return &IncreasePieceScaleCmd{
		piece: piece,
	}
}

func (cmd *IncreasePieceScaleCmd) redo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.scale++
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *IncreasePieceScaleCmd) undo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.scale--
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type DecreasePieceScaleCmd struct {
	piece uint32
}

func NewDecreasePieceScaleCmd(sb *Sandbox, piece uint32) *DecreasePieceScaleCmd {
	sb.GetPiece(piece).scale--
	return &DecreasePieceScaleCmd{
		piece: piece,
	}
}

func (cmd *DecreasePieceScaleCmd) redo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.scale--
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *DecreasePieceScaleCmd) undo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.scale++
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
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
		cmd.capturedPiece = captured.id
		cmd.captureAfterCoord = sb.FindUnoccupiedOffBoardCoord()
		captured.board = OffBoard
		captured.coord = cmd.captureAfterCoord
	}
	var piece = sb.NewPiece(ui.clipboard.typ, ui.clipboard.color, board, coord)
	piece.scale = ui.clipboard.scale
	cmd.piece = *piece
	cmd.effects = make([]uint32, len(ui.clipboard.effects))
	copy(cmd.effects, ui.clipboard.effects)
	for _, effect := range cmd.effects {
		sb.NewStatusEffect(piece.id, effect)
	}
	ui.selection.SelectPiece(cmd.piece.id)
	return &cmd
}

func (cmd *PastePieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
	for _, effect := range cmd.effects {
		sb.NewStatusEffect(cmd.piece.id, effect)
	}
	if cmd.piece.board != OffBoard {
		ui.board = int32(cmd.piece.board)
	}
	ui.selection.SelectPiece(cmd.piece.id)
	ui.tab = TabBoard
}

func (cmd *PastePieceCmd) undo(sb *Sandbox, ui *UiState) {
	if id, ok := ui.selection.GetSelectedPieceId(); ok && id == cmd.piece.id {
		ui.selection.Deselect()
	}
	sb.RemovePiece(cmd.piece.id)
	if cmd.anyCaptured {
		var captured = sb.GetPiece(cmd.capturedPiece)
		captured.board = cmd.piece.board
		captured.coord = cmd.piece.coord
	}
	if cmd.piece.board != OffBoard {
		ui.board = int32(cmd.piece.board)
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
	var piece = sb.NewPiece(original.typ, original.color, OffBoard, coord)
	var effects = sb.GetStatusEffectsOnPiece(id)
	for _, effect := range effects {
		sb.NewStatusEffect(piece.id, effect)
	}
	ui.selection.SelectPiece(piece.id)
	return &DuplicatePieceCmd{
		originalId: id,
		piece:      *piece,
		effects:    effects,
	}
}

func (cmd *DuplicatePieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
	for _, effect := range cmd.effects {
		sb.NewStatusEffect(cmd.piece.id, effect)
	}
	ui.selection.SelectPiece(cmd.piece.id)
	ui.tab = TabBoard
}

func (cmd *DuplicatePieceCmd) undo(sb *Sandbox, ui *UiState) {
	sb.RemovePiece(cmd.piece.id)
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
	var before = piece.color
	piece.color = color
	return &ChangeColorOfPieceCmd{
		piece:  id,
		before: before,
		after:  color,
	}
}

func (cmd *ChangeColorOfPieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.GetPiece(cmd.piece).color = cmd.after
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

func (cmd *ChangeColorOfPieceCmd) undo(sb *Sandbox, ui *UiState) {
	sb.GetPiece(cmd.piece).color = cmd.before
	ui.selection.SelectPiece(cmd.piece)
	ui.tab = TabBoard
}

type ChangeMoneyAmountCmd struct {
	whiteBefore int
	whiteAfter  int
	blackBefore int
	blackAfter  int
}

func NewChangeMoneyAmountCmd(ui *UiState, whiteNewAmount int, blackNewAmount int) *ChangeMoneyAmountCmd {
	var whiteBefore = ui.shop.money[0]
	var blackBefore = ui.shop.money[1]
	ui.shop.money[0] = whiteNewAmount
	ui.shop.money[1] = blackNewAmount
	return &ChangeMoneyAmountCmd{
		whiteBefore: whiteBefore,
		whiteAfter:  whiteNewAmount,
		blackBefore: blackBefore,
		blackAfter:  blackNewAmount,
	}
}

func (cmd *ChangeMoneyAmountCmd) redo(sb *Sandbox, ui *UiState) {
	ui.shop.money[0] = cmd.whiteAfter
	ui.shop.money[1] = cmd.blackAfter
}

func (cmd *ChangeMoneyAmountCmd) undo(sb *Sandbox, ui *UiState) {
	ui.shop.money[0] = cmd.whiteBefore
	ui.shop.money[1] = cmd.blackBefore
}

type ChangeShopUnlockCountCmd struct {
	before int
	after  int
}

func NewChangeShopUnlockCountCmd(ui *UiState, newCount int) *ChangeShopUnlockCountCmd {
	var before = ui.shop.unlockedCount
	ui.shop.unlockedCount = newCount
	return &ChangeShopUnlockCountCmd{
		before: before,
		after:  newCount,
	}
}

func (cmd *ChangeShopUnlockCountCmd) redo(sb *Sandbox, ui *UiState) {
	ui.shop.unlockedCount = cmd.after
	ui.tab = TabShop
}

func (cmd *ChangeShopUnlockCountCmd) undo(sb *Sandbox, ui *UiState) {
	ui.shop.unlockedCount = cmd.before
	ui.tab = TabShop
}

type ShuffleShopCmd struct {
	swaps []int
}

func NewShuffleShopCmd(shop *Shop) *ShuffleShopCmd {
	// Fisher-Yates shuffle, but we store which indices we roll, so we can undo it
	var swaps = make([]int, 0)
	for i := len(shop.entries) - 1; i >= 1; i-- {
		var j = rand.Intn(i + 1)
		swaps = append(swaps, j)
		shop.entries[i], shop.entries[j] = shop.entries[j], shop.entries[i]
	}
	return &ShuffleShopCmd{
		swaps: swaps,
	}
}

func (cmd *ShuffleShopCmd) redo(sb *Sandbox, ui *UiState) {
	for i := len(ui.shop.entries) - 1; i >= 1; i-- {
		var j = cmd.swaps[len(ui.shop.entries)-1-i]
		ui.shop.entries[i], ui.shop.entries[j] = ui.shop.entries[j], ui.shop.entries[i]
	}
	ui.tab = TabShop
}

func (cmd *ShuffleShopCmd) undo(sb *Sandbox, ui *UiState) {
	for i := 1; i < len(ui.shop.entries); i++ {
		var j = cmd.swaps[len(ui.shop.entries)-1-i]
		ui.shop.entries[i], ui.shop.entries[j] = ui.shop.entries[j], ui.shop.entries[i]
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
	var priceBefore = e.price
	e.price = newPrice
	return &ChangeShopEntryPriceCmd{
		entry:       entry,
		priceBefore: priceBefore,
		priceAfter:  newPrice,
	}
}

func (cmd *ChangeShopEntryPriceCmd) redo(sb *Sandbox, ui *UiState) {
	ui.shop.GetEntry(cmd.entry).price = cmd.priceAfter
	ui.tab = TabShop
}

func (cmd *ChangeShopEntryPriceCmd) undo(sb *Sandbox, ui *UiState) {
	ui.shop.GetEntry(cmd.entry).price = cmd.priceBefore
	ui.tab = TabShop
}

type QuickBuyCmd struct {
	player uint32
	entry  uint32
	unlock bool
}

func NewQuickBuyCmd(shop *Shop, player uint32, entry uint32) *QuickBuyCmd {
	var e = shop.GetEntry(entry)
	shop.money[player] -= e.price
	e.price++
	var unlock = shop.unlockedCount < len(shop.entries)
	if unlock {
		shop.unlockedCount++
	}
	return &QuickBuyCmd{
		player: player,
		entry:  entry,
		unlock: unlock,
	}
}

func (cmd *QuickBuyCmd) redo(sb *Sandbox, ui *UiState) {
	var e = ui.shop.GetEntry(cmd.entry)
	ui.shop.money[cmd.player] -= e.price
	e.price++
	if cmd.unlock {
		ui.shop.unlockedCount++
	}
	ui.tab = TabShop
}

func (cmd *QuickBuyCmd) undo(sb *Sandbox, ui *UiState) {
	if cmd.unlock {
		ui.shop.unlockedCount--
	}
	var e = ui.shop.GetEntry(cmd.entry)
	e.price--
	ui.shop.money[cmd.player] += e.price
	ui.tab = TabShop
}
