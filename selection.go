package main

type SelectionType int32

const (
	SelectionTypeNone SelectionType = iota
	SelectionTypePiece
	SelectionTypePieceType
	SelectionTypeCoord
)

type Selection struct {
	selectionType SelectionType
	id            uint32
	coord         Vec2
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

func (s *Selection) SelectCoord(coord Vec2) {
	s.selectionType = SelectionTypeCoord
	s.coord = coord
}

func (s *Selection) IsCoordSelected(coord Vec2) bool {
	return s.selectionType == SelectionTypeCoord && s.coord == coord
}

func (s *Selection) GetSelectedCoord() (Vec2, bool) {
	if s.selectionType != SelectionTypeCoord {
		return s.coord, false
	}
	return s.coord, true
}

func (s *Selection) HasSelection() bool {
	return s.selectionType != SelectionTypeNone
}
