package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var UP = Vec2{0, -1}
var DOWN = Vec2{0, 1}
var LEFT = Vec2{-1, 0}
var RIGHT = Vec2{1, 0}
var ZEROZERO = Vec2{0, 0}
var ONEONE = Vec2{1, 1}

type Vec2 struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (v Vec2) Add(u Vec2) Vec2 {
	return Vec2{X: v.X + u.X, Y: v.Y + u.Y}
}

func (v Vec2) Sub(u Vec2) Vec2 {
	return Vec2{X: v.X - u.X, Y: v.Y - u.Y}
}

func (v Vec2) Scale(s int) Vec2 {
	return Vec2{X: v.X * s, Y: v.Y * s}
}

func (v Vec2) ManLength() int {
	return Absi(v.X) + Absi(v.Y)
}

func (v Vec2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vec2) Normalized() Vec2 {
	var length = v.Length()
	return Vec2{
		X: int(float32(v.X) / length),
		Y: int(float32(v.Y) / length),
	}
}

func (v Vec2) CompwiseModulo(d int) Vec2 {
	return Vec2{X: v.X % d, Y: v.Y % d}
}

func (v Vec2) CompwiseMul(u Vec2) Vec2 {
	return Vec2{X: v.X * u.X, Y: v.Y * u.Y}
}

func (v Vec2) CompwiseMax(u Vec2) Vec2 {
	return Vec2{X: max(v.X, u.X), Y: max(v.Y, u.Y)}
}

func (v Vec2) CompwiseMin(u Vec2) Vec2 {
	return Vec2{X: min(v.X, u.X), Y: min(v.Y, u.Y)}
}

func (v Vec2) ToRlVec() rl.Vector2 {
	return rl.NewVector2(float32(v.X), float32(v.Y))
}
