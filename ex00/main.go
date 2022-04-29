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

func isCheck() bool {
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
	bl := isCheck()
	output(bl)
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
