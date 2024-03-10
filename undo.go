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

func (s *UndoRedoSystem) Append(undoable Undoable) {
	s.history = s.history[:s.doneCount]
	s.history = append(s.history, undoable)
	s.doneCount++
}
