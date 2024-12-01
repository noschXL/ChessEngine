package main

import "strings"

func IsEmpty(data uint8) bool {
	return data == EMPTY
}

func GetType(piece uint8) uint8 {
	return piece & typemask
}

func GetColor(piece uint8) uint8 {
	return piece & colormask
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
