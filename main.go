package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)
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

	sandbox = NewSandbox()
	var planeIndex = int32(1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		sandbox.Render(uint32(planeIndex))

		if rg.Button(rl.NewRectangle(20, 20, 200, 36), "Remove random") {
			println("Clicked!")
			if len(sandbox.pieces) > 0 {
				var id = sandbox.pieces[rand.Intn(len(sandbox.pieces))].id
				sandbox.RemovePiece(id)
			}
		}

		planeIndex = rg.ToggleGroup(rl.NewRectangle(float32(rl.GetScreenWidth()/2-(120*3+int(rg.GetStyle(rg.DEFAULT, rg.GROUP_PADDING)))/2), float32(rl.GetScreenHeight()-36-20), 120, 36), "Heaven;Earth;Hell", planeIndex)

		rl.EndDrawing()
	}
}
