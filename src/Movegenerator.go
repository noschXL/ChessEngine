//MoveGenerator.go
// this File contains all the move generation

package main

// Pseudomoves are moves that do not look for check
func GetPseudoMoves(board [65]uint8, from uint8) []Move {
	piece := board[from]
	piecetype := GetType(piece)
	piececolor := GetColor(piece)
	println(piececolor, piecetype)

	var moves []Move

	return moves
}
