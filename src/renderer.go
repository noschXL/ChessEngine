package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawBoard(board [65]uint8) {

}
func DrawSquares() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			currcolor := BLACKRGB
			if (i+j)%2 == 0 {
				currcolor = WHITERGB
			}
			rl.DrawRectangle(int32(i)*100, int32(j)*100, 100, 100, currcolor)
		}
	}
}

func DrawPieces(board [65]uint8, PIECETEXTURE rl.Texture2D) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i == 64 {
				break
			}

			data := board[i*8+j]

			piecetype := GetType(data)
			piececolor := uint8(math.Abs(float64(GetColor(data)>>3) - 1))

			if IsEmpty(data) {
				break
			}

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
				X:      float32(int32(j)*100) + 10,
				Y:      float32(int32(i)*100) + 10,
				Width:  100,
				Height: 100,
			}

			origin := rl.Vector2{
				X: 10,
				Y: 10,
			}

			rl.DrawTexturePro(PIECETEXTURE, texture_rect, dest_rect, origin, 0, rl.RayWhite)
		}
	}
}
