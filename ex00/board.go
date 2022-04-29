package main

import "piscine"

type Board []string

func (board Board) len() int {
	return piscine.Len([]string(board))
}

func (board Board) at(y, x int) rune {
	len := board.len()
	if !piscine.IsRange(x, 0, len) || !piscine.IsRange(y, 0, len) {
		return 0
	}

	row := board[y]
	return []rune(row)[x]
}
