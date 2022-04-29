package main

import (
	"piscine"
)

func isSquare(board []string) bool {
	len := piscine.Strslen(board)
	for _, row := range board {
		if len != piscine.Strlen(row) {
			return false
		}
	}
	return true
}

func checkmate(board []string) {
	if !isSquare(board) {
		return
	}
}

func main() {
	board := []string{
		"R.P.",
		".K..",
		"..P.",
		"....",
	}
	checkmate(board)
}
