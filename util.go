package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const TileSize = 80

func Absi(v int) int {
	if v <= 0 {
		return -v
	}
	return 0
}

func GetBoardOrigo() Vec2 {
	return Vec2{rl.GetScreenWidth()/2 - 4*TileSize, rl.GetScreenHeight()/2 - 4*TileSize}
}

func GetHoveredCoord() Vec2 {
	var mousef = rl.GetMousePosition()
	var origof = GetBoardOrigo().ToRlVec()
	var coordf = rl.Vector2Scale(rl.Vector2Subtract(mousef, origof), 1./TileSize)
	return Vec2{X: int(math.Floor(float64(coordf.X))), Y: int(math.Floor(float64(coordf.Y)))}
}
