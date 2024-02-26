package main

type Undoable interface {
	undo(sb *Sandbox, ui *UiState)
	redo(sb *Sandbox, ui *UiState)
}

type UndoRedoSystem struct {
	history   []Undoable
	doneCount uint32
}

func NewUndoRedoSystem() UndoRedoSystem {
	return UndoRedoSystem{}
}

func (s *UndoRedoSystem) Redo(sb *Sandbox, ui *UiState) {
	if s.doneCount < uint32(len(s.history)) {
		s.history[s.doneCount].redo(sb, ui)
		s.doneCount++
	}
}

func (s *UndoRedoSystem) Undo(sb *Sandbox, ui *UiState) bool {
	if s.doneCount == 0 {
		return false
	}
	s.doneCount--
	s.history[s.doneCount].undo(sb, ui)
	return true
}

func (s *UndoRedoSystem) AppendDone(undoable Undoable) {
	s.history = s.history[:s.doneCount]
	s.history = append(s.history, undoable)
	s.doneCount++
}

type CreatePieceCmd struct {
	piece Piece
}

func NewCreatePieceCmd(sb *Sandbox, typ uint32, color PieceColor, board uint32, coord Vec2) CreatePieceCmd {
	if IsOffBoard(coord) {
		board = OffBoard
	}
	return CreatePieceCmd{
		piece: *sb.NewPiece(typ, color, board, coord),
	}
}

func (cmd *CreatePieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
	if cmd.piece.board != OffBoard {
		ui.board = int32(cmd.piece.board)
	}
	ui.selection.Deselect()
}

func (cmd *CreatePieceCmd) undo(sb *Sandbox, ui *UiState) {
	if id, ok := ui.selection.GetSelectedPieceId(); ok && id == cmd.piece.id {
		ui.selection.Deselect()
	}
	sb.RemovePiece(cmd.piece.id)
	if cmd.piece.board != OffBoard {
		ui.board = int32(cmd.piece.board)
	}
	ui.selection.Deselect()
}

type DeletePieceCmd struct {
	piece   Piece
	effects []uint32
}

func NewDeletePieceCmd(sb *Sandbox, ui *UiState, id uint32) DeletePieceCmd {
	var cmd = DeletePieceCmd{
		piece:   *sb.GetPiece(id),
		effects: sb.GetStatusEffectsOnPiece(id),
	}

	sb.RemovePiece(id)
	ui.selection.Deselect()
	return cmd
}

func (cmd *DeletePieceCmd) redo(sb *Sandbox, ui *UiState) {
	if id, ok := ui.selection.GetSelectedPieceId(); ok && id == cmd.piece.id {
		ui.selection.Deselect()
	}
	sb.RemovePiece(cmd.piece.id)
	if cmd.piece.board != OffBoard {
		ui.board = int32(cmd.piece.board)
	}
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
}

type MovePieceCmd struct {
	piece       uint32
	beforeCoord Vec2
	beforeBoard uint32
	afterCoord  Vec2
	afterBoard  uint32
}

func NewMovePieceCmd(sb *Sandbox, id uint32, destCoord Vec2, destBoard uint32) MovePieceCmd {
	var piece = sb.GetPiece(id)
	var cmd = MovePieceCmd{
		piece:       id,
		beforeCoord: piece.coord,
		beforeBoard: piece.board,
		afterCoord:  destCoord,
		afterBoard:  destBoard,
	}
	piece.coord = destCoord
	if !IsOffBoard(destCoord) {
		piece.board = destBoard
	} else {
		piece.board = OffBoard
		cmd.afterBoard = OffBoard
	}
	return cmd
}

func (cmd *MovePieceCmd) redo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.coord = cmd.afterCoord
	piece.board = cmd.afterBoard
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
}

func (cmd *MovePieceCmd) undo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.coord = cmd.beforeCoord
	piece.board = cmd.beforeBoard
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
}

type CreateStatusEffectCmd struct {
	piece  uint32
	effect uint32
}

func NewCreateStatusEffectCmd(sb *Sandbox, piece uint32, effect uint32) CreateStatusEffectCmd {
	sb.NewStatusEffect(piece, effect)
	return CreateStatusEffectCmd{
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
}

func (cmd *CreateStatusEffectCmd) undo(sb *Sandbox, ui *UiState) {
	sb.RemoveStatusEffect(cmd.piece, cmd.effect)
	var piece = sb.GetPiece(cmd.piece)
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
}

type DeleteStatusEffectCmd struct {
	piece  uint32
	effect uint32
}

func NewDeleteStatusEffectCmd(sb *Sandbox, piece uint32, effect uint32) DeleteStatusEffectCmd {
	sb.RemoveStatusEffect(piece, effect)
	return DeleteStatusEffectCmd{
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
}

func (cmd *DeleteStatusEffectCmd) undo(sb *Sandbox, ui *UiState) {
	sb.NewStatusEffect(cmd.piece, cmd.effect)
	var piece = sb.GetPiece(cmd.piece)
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
}

type IncreasePieceScaleCmd struct {
	piece uint32
}

func NewIncreasePieceScaleCmd(sb *Sandbox, piece uint32) IncreasePieceScaleCmd {
	sb.GetPiece(piece).scale++
	return IncreasePieceScaleCmd{
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
}

func (cmd *IncreasePieceScaleCmd) undo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.scale--
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
}

type DecreasePieceScaleCmd struct {
	piece uint32
}

func NewDecreasePieceScaleCmd(sb *Sandbox, piece uint32) DecreasePieceScaleCmd {
	sb.GetPiece(piece).scale--
	return DecreasePieceScaleCmd{
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
}

func (cmd *DecreasePieceScaleCmd) undo(sb *Sandbox, ui *UiState) {
	var piece = sb.GetPiece(cmd.piece)
	piece.scale++
	if piece.board != OffBoard {
		ui.board = int32(piece.board)
	}
	ui.selection.SelectPiece(cmd.piece)
}

type CreateTileCmd struct {
	board uint32
	coord Vec2
}

func NewCreateTileCmd(sb *Sandbox, board uint32, coord Vec2) CreateTileCmd {
	sb.NewTile(board, coord)
	return CreateTileCmd{
		board: board,
		coord: coord,
	}
}

func (cmd *CreateTileCmd) redo(sb *Sandbox, ui *UiState) {
	sb.NewTile(cmd.board, cmd.coord)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
}

func (cmd *CreateTileCmd) undo(sb *Sandbox, ui *UiState) {
	sb.RemoveTile(cmd.board, cmd.coord)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
}

type DeleteTileCmd struct {
	board uint32
	coord Vec2
}

func NewDeleteTileCmd(sb *Sandbox, board uint32, coord Vec2) DeleteTileCmd {
	sb.RemoveTile(board, coord)
	return DeleteTileCmd{
		board: board,
		coord: coord,
	}
}

func (cmd *DeleteTileCmd) redo(sb *Sandbox, ui *UiState) {
	sb.RemoveTile(cmd.board, cmd.coord)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
}

func (cmd *DeleteTileCmd) undo(sb *Sandbox, ui *UiState) {
	sb.NewTile(cmd.board, cmd.coord)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
}

type CreateObstacleCmd struct {
	coord    Vec2
	board    uint32
	obstacle uint32
}

func NewCreateObstacleCmd(sb *Sandbox, coord Vec2, board uint32, obstacle uint32) CreateObstacleCmd {
	sb.NewObstacle(coord, board, obstacle)
	return CreateObstacleCmd{
		obstacle: obstacle,
		board:    board,
		coord:    coord,
	}
}

func (cmd *CreateObstacleCmd) redo(sb *Sandbox, ui *UiState) {
	sb.NewObstacle(cmd.coord, cmd.board, cmd.obstacle)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
}

func (cmd *CreateObstacleCmd) undo(sb *Sandbox, ui *UiState) {
	sb.RemoveObstacle(cmd.coord, cmd.board, cmd.obstacle)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
}

type DeleteObstacleCmd struct {
	coord    Vec2
	board    uint32
	obstacle uint32
}

func NewDeleteObstacleCmd(sb *Sandbox, coord Vec2, board uint32, obstacle uint32) DeleteObstacleCmd {
	sb.RemoveObstacle(coord, board, obstacle)
	return DeleteObstacleCmd{
		obstacle: obstacle,
		board:    board,
		coord:    coord,
	}
}

func (cmd *DeleteObstacleCmd) redo(sb *Sandbox, ui *UiState) {
	sb.RemoveObstacle(cmd.coord, cmd.board, cmd.obstacle)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
}

func (cmd *DeleteObstacleCmd) undo(sb *Sandbox, ui *UiState) {
	sb.NewObstacle(cmd.coord, cmd.board, cmd.obstacle)
	ui.board = int32(cmd.board)
	ui.selection.SelectCoord(cmd.coord)
}
