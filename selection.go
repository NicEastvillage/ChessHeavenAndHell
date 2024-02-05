package main

type SelectionType int32

const (
	SelectionTypeNone SelectionType = iota
	SelectionTypePiece
	SelectionTypePieceType
)

type Selection struct {
	selectionType SelectionType
	id            uint32
}

func (s *Selection) Deselect() {
	s.selectionType = SelectionTypeNone
}

func (s *Selection) SelectPiece(id uint32) {
	s.selectionType = SelectionTypePiece
	s.id = id
}

func (s *Selection) IsPieceSelected(id uint32) bool {
	return s.selectionType == SelectionTypePiece && s.id == id
}

func (s *Selection) GetSelectedPieceId() (uint32, bool) {
	if s.selectionType != SelectionTypePiece {
		return s.id, false
	}
	return s.id, true
}

func (s *Selection) SelectPieceType(id uint32) {
	s.selectionType = SelectionTypePieceType
	s.id = id
}

func (s *Selection) IsPieceTypeSelected(id uint32) bool {
	return s.selectionType == SelectionTypePieceType && s.id == id
}

func (s *Selection) GetSelectedPieceTypeId() (uint32, bool) {
	if s.selectionType != SelectionTypePieceType {
		return s.id, false
	}
	return s.id, true
}

func (s *Selection) HasSelection() bool {
	return s.selectionType != SelectionTypeNone
}
