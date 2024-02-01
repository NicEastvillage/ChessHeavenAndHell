package main

import rl "github.com/gen2brain/raylib-go/raylib"
import rg "github.com/gen2brain/raylib-go/raygui"

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1600, 980, "Chess - Heaven and Hell")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rg.LoadStyleDefault()
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)

	assets.LoadAll()
	defer assets.UnloadAll()

	var boards = [3]Board{NewBoard(8, 8, BoardStyleHeaven), NewStandardBoardWithPieces(), NewBoard(8, 8, BoardStyleHell)}
	var planeIndex = int32(1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		boards[planeIndex].Render()

		if rg.Button(rl.NewRectangle(20, 20, 120, 32), "Test") {
			println("Clicked!")
		}

		planeIndex = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()/2-(120*3+int(rg.GetStyle(rg.DEFAULT, rg.GROUP_PADDING)))/2), float32(rl.GetScreenHeight()-32-20), 120, 32), "Heaven;Earth;Hell", planeIndex)

		rl.EndDrawing()
	}
}
