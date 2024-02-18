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
	return CreatePieceCmd{
		piece: *sb.NewPiece(typ, color, board, coord),
	}
}

func (cmd *CreatePieceCmd) redo(sb *Sandbox, ui *UiState) {
	sb.AddPiece(cmd.piece)
}

func (cmd *CreatePieceCmd) undo(sb *Sandbox, ui *UiState) {
	if id, ok := ui.selection.GetSelectedPieceId(); ok && id == cmd.piece.id {
		ui.selection.Deselect()
	}
	sb.RemovePiece(cmd.piece.id)
}
