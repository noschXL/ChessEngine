package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	board := LoadFEN(STARTFEN)

	rl.InitWindow(800, 800, "Chess Engine")
	pieceimg := rl.LoadImage(AppendFilePath(GetPath(), []string{"img", "pieces.png"}))
	rl.ImageResize(pieceimg, 600, 200)
	piecetexture := rl.LoadTextureFromImage(pieceimg)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		HandleMouse(board)

		rl.BeginDrawing()
		DrawSquares()
		DrawPieces(board, piecetexture)

		rl.EndDrawing()
	}
}
