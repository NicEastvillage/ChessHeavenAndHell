package main

type BoardStyle int32

const (
	BoardStyleEarth BoardStyle = iota
	BoardStyleHeaven
	BoardStyleHell
)

const OffBoard = 99999

type Board struct {
	id    uint32
	style BoardStyle
}
