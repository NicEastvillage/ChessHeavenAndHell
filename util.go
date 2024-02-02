package main

import rl "github.com/gen2brain/raylib-go/raylib"

const TileSize = 100

func Absi(v int) int {
	if v <= 0 {
		return -v
	}
	return 0
}

func GetWorldOrigo() Vec2 {
	return Vec2{rl.GetScreenWidth()/2 - 4*TileSize, rl.GetScreenHeight()/2 - 4*TileSize}
}
