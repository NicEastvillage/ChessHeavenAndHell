package main

import rl "github.com/gen2brain/raylib-go/raylib"
import rg "github.com/gen2brain/raylib-go/raygui"

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1600, 980, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rg.LoadStyleDefault()
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)

	assets.LoadAll()
	defer assets.UnloadAll()

	var board = NewStandardBoardWithPiece()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if rg.Button(rl.NewRectangle(20, 20, 120, 32), "Test") {
			println("Clicked!")
		}

		board.Render()

		rl.EndDrawing()
	}
}
