package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	board := LoadFEN(STARTFEN)

	rl.InitWindow(800, 800, "Chess Engine")
	PIECETEXTURE := rl.LoadTexture(AppendFilePath(GetPath(), []string{"img", "pieces.png"}))

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		DrawSquares()
		DrawPieces(board, PIECETEXTURE)

		rl.EndDrawing()
	}
}
