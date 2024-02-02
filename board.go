package main

type BoardStyle int32

const (
	BoardStyleEarth BoardStyle = iota
	BoardStyleHeaven
	BoardStyleHell
)

type Board struct {
	id    uint32
	style BoardStyle
}
