package main

var selection = Selection{typ: SelectionTypeNone}

type SelectionType int32

const (
	SelectionTypeNone SelectionType = iota
	SelectionTypePiece
)

type Selection struct {
	typ SelectionType
	id  uint32
}

func (s *Selection) Deselect() {
	s.typ = SelectionTypeNone
}

func (s *Selection) SelectPiece(id uint32) {
	s.typ = SelectionTypePiece
	s.id = id
}

func (s *Selection) IsPieceSelected(id uint32) bool {
	return s.typ == SelectionTypePiece && s.id == id
}
