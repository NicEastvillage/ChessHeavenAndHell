package main

type BoardStyle int32

const (
	BoardStyleHeaven BoardStyle = iota
	BoardStyleEarth
	BoardStyleHell
)

const OffBoard = 99999

type Board struct {
	Id    uint32     `json:"id"`
	Style BoardStyle `json:"style"`
}

func NameOfBoard(board uint32) string {
	switch board {
	case 0:
		return "Heaven"
	case 1:
		return "Earth"
	case 2:
		return "Hell"
	case OffBoard:
		return "Offboard"
	default:
		return "Unnamed Board"
	}
}
