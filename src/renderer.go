package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func ResizeAndBlurTexture(width int32, height int32, texture rl.Texture2D) rl.Texture2D {
	newTexture := rl.LoadRenderTexture(width, height)
	rl.BeginTextureMode(newTexture)
	rl.DrawTexturePro(
		texture,
		rl.Rectangle{X: 0, Y: 0, Width: float32(texture.Width), Height: float32(texture.Height)},
		rl.Rectangle{X: 0, Y: 0, Width: float32(width), Height: float32(height)},
		rl.NewVector2(0, 0),
		0,
		color.RGBA{255, 255, 255, 255},
	)
	rl.EndTextureMode()
	//img := rl.LoadImageFromTexture(newTexture.Texture)
	//rl.ImageBlurGaussian(img, 5)
	return newTexture.Texture

}

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

			piecetype := GetType(data) - 1
			piececolor := GetColor(data) >> 4

			if IsEmpty(data) {
				break
			}

			imgWidth := PIECETEXTURE.Width / 6
			imgHeight := PIECETEXTURE.Height / 2

			x := int(imgWidth) * int(piecetype) // EMPTY = 0 but the king is at pos 0 and so on
			y := int(imgHeight) * int(piececolor)

			rl.DrawTextureRec(PIECETEXTURE, rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(imgWidth), Height: float32(imgHeight)}, rl.Vector2{X: float32(j * 100), Y: float32(i * 100)}, rl.White)
		}
	}
}
