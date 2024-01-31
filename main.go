package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1600, 980, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	assets.LoadAll()
	defer assets.UnloadAll()

	var board = NewStandardBoardWithPiece()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.Translatef(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2), 0)
		rl.Translatef(-4*TILE_SIZE, -4*TILE_SIZE, 0)
		board.Render()

		rl.EndDrawing()
	}
}
