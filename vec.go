package main

import rl "github.com/gen2brain/raylib-go/raylib"

var UP = Vec2{0, -1}
var DOWN = Vec2{0, 1}
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

func (v Vec2) ToRlVec() rl.Vector2 {
	return rl.NewVector2(float32(v.x), float32(v.y))
}
