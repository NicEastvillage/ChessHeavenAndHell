package main

var UP = Vec2{0, 1}
var DOWN = Vec2{0, -1}
var LEFT = Vec2{-1, 0}
var RIGHT = Vec2{1, 0}
var ZEROZERO = Vec2{0, 0}
var ONEONE = Vec2{1, 1}

type Vec2 struct {
	x, y int
}

func (v Vec2) Add(u Vec2) Vec2 {
	return Vec2{x: v.x + u.x, y: v.y + u.y}
}

func (v Vec2) Sub(u Vec2) Vec2 {
	return Vec2{x: v.x - u.x, y: v.y - u.y}
}

func (v Vec2) Scale(s int) Vec2 {
	return Vec2{x: v.x * s, y: v.y * s}
}

func (v Vec2) ManLength() int {
	return Absi(v.x) + Absi(v.y)
}

func (v Vec2) CompwiseMul(u Vec2) Vec2 {
	return Vec2{x: v.x * u.x, y: v.y * u.y}
}

func (v Vec2) CompwiseMax(u Vec2) Vec2 {
	return Vec2{x: max(v.x, u.x), y: max(v.y, u.y)}
}

func (v Vec2) CompwiseMin(u Vec2) Vec2 {
	return Vec2{x: min(v.x, u.x), y: min(v.y, u.y)}
}

type AARect struct {
	// BottomLeft is included, topRight is excluded
	bottomLeft, topRight Vec2
}

func NewAARect(p1, p2 Vec2) AARect {
	var bottomLeft = Vec2{
		x: min(p1.x, p2.x),
		y: min(p1.y, p2.y),
	}
	var topRight = Vec2{
		x: max(p1.x, p2.x),
		y: max(p1.y, p2.y),
	}
	return AARect{bottomLeft: bottomLeft, topRight: topRight}
}

func NewAARectEmpty() AARect {
	return AARect{bottomLeft: Vec2{x: 0, y: 0}, topRight: Vec2{x: 0, y: 0}}
}

func (r AARect) IsEmpty() bool {
	return r.topRight.x <= r.bottomLeft.x || r.topRight.y <= r.bottomLeft.y
}

func (r AARect) ExpandedToInclude(p Vec2) AARect {
	if r.IsEmpty() {
		return AARect{bottomLeft: p, topRight: p.Add(ONEONE)}
	}
	return AARect{
		bottomLeft: r.bottomLeft.CompwiseMin(p),
		topRight:   r.topRight.CompwiseMax(p.Add(ONEONE)),
	}
}
