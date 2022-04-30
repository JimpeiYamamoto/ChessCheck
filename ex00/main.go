package main

import (
	"piscine"
)

func isSquare(board []string) bool {
	len := piscine.Strslen(board)
	if len == 0 {
		return false
	}
	for _, row := range board {
		if len != piscine.Strlen(row) {
			return false
		}
	}
	return true
}

func isOnlyKing(board []string) bool {
	var nbKing int

	for _, v := range board {
		for _, r := range v {
			if r == 'K' {
				nbKing++
			}
		}
	}
	return nbKing == 1
}

func isValid(board []string) bool {
	if !isSquare(board) {
		return false
	}
	if !isOnlyKing(board) {
		return false
	}
	return true
}

func isCheck(board Board) bool {
	for y, row := range board {
		for x, cell := range row {
			if Piece(cell).isCollideKing(board, y, x) {
				return true
			}
		}
	}
	return false
}

func output(isCheck bool) {
	if isCheck {
		piscine.PrintStr("Success\n")
	} else {
		piscine.PrintStr("Fail\n")
	}
}

func checkmate(board []string) {
	if !isValid(board) {
		return
	}
	bl := isCheck(board)
	output(bl)
}

func main() {
	board := []string{
		"R...",
		".K..",
		"..P.",
		"....",
	}
	checkmate(board)
}
