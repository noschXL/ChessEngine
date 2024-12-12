package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Field struct {
	X uint8
	Y uint8
}

var lastpressed bool = false

var lastSqaure Field = Field{9, 9}

func HandleMouse(board [65]uint8) {
	pressed := rl.IsMouseButtonPressed(rl.MouseButtonLeft)

}
