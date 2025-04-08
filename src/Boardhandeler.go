package main

import (
	"math"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func IsEmpty(data uint8) bool {
	return data == EMPTY
}

func GetType(piece uint8) uint8 {
	return piece & typemask
}

func GetColor(piece uint8) uint8 {
	return piece & colormask
}

func LinearInterpolation(a float32, b float32, t float32) float32 {
	return a + (b-a)*t
}

func DrawMove(board [65]uint8, move Move, PIECETEXTURE rl.Texture2D, drawmove bool) {
	movingPiece := board[move.from]

	board[move.from] = EMPTY

	if drawmove {
		for t := range 60 {
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			DrawSquares()
			DrawPieces(board, PIECETEXTURE)

			progress := float32(t) / 59.0

			oldx := float32((move.from % 8) * 100)
			oldy := float32((move.from / 8) * 100)
			newx := float32((move.to % 8) * 100)
			newy := float32((move.to / 8) * 100)

			currentX := LinearInterpolation(oldx, newx, progress)
			currentY := LinearInterpolation(oldy, newy, progress)

			piecetype := GetType(movingPiece)
			piececolor := uint8(math.Abs(float64(GetColor(movingPiece)>>3) - 1))

			imgsize := PIECETEXTURE.Height / 2

			x := imgsize * int32(piecetype-1)
			y := imgsize * int32(piececolor)

			texture_rect := rl.Rectangle{
				X:      float32(x),
				Y:      float32(y),
				Width:  float32(imgsize),
				Height: float32(imgsize),
			}
			dest_rect := rl.Rectangle{
				X:      currentX + 10,
				Y:      currentY + 10,
				Width:  100,
				Height: 100,
			}

			rl.DrawTexturePro(PIECETEXTURE, texture_rect, dest_rect, rl.Vector2{X: 50, Y: 50}, 0, rl.White)

			rl.EndDrawing()
			rl.WaitTime(16.0 / 1000)
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		DrawSquares()
		DrawPieces(board, PIECETEXTURE)
		rl.EndDrawing()
	}

	board[move.to] = movingPiece

}

func LoadFEN(fen string) [65]uint8 {
	var board [65]uint8
	parts := strings.Fields(fen)
	println(parts[0])
	println(parts[1])
	println(parts[2])
	println(parts[3])
	println(parts[4])
	println(parts[5])
	j := int32(0)
	for _, char := range parts[0] {
		piece := uint8(0)
		color := uint8(0)
		if string(char) == strings.ToLower(string(char)) {
			color = BLACK
		} else {
			color = WHITE
		}

		println(color)

		switch strings.ToLower(string(char)) {
		case "k":
			piece = KING
		case "q":
			piece = QUEEN
		case "r":
			piece = ROOK
		case "b":
			piece = BISHOP
		case "n":
			piece = KNIGHT
		case "p":
			piece = PAWN
		case "/":
			continue
		default:
			j += int32(char) - 48
			continue
		}

		board[j] = piece | color

		j++

	}

	return board
}
