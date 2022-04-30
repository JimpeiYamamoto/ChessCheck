package main

import (
	"fmt"
	"regexp"
)

func printBoard() {
	reg := regexp.MustCompile(`[eE]`)

	tmpBoard := []string{
		reg.ReplaceAllString(string(board[0]), "."),
		reg.ReplaceAllString(string(board[1]), "."),
		reg.ReplaceAllString(string(board[2]), "."),
		reg.ReplaceAllString(string(board[3]), "."),
		reg.ReplaceAllString(string(board[4]), "."),
		reg.ReplaceAllString(string(board[5]), "."),
		reg.ReplaceAllString(string(board[6]), "."),
		reg.ReplaceAllString(string(board[7]), "."),
	}
	for _, row := range tmpBoard {
		fmt.Println(row)
	}
	// アンパサン表示用
	// for _, row := range board {
	// 	fmt.Println(string(row))
	// }
	fmt.Println()
}

func print() {
	// turn count
	if turn == WHITE {
		fmt.Printf("%d.\n", cnt)
		cnt++
	}

	// board
	printBoard()
}

func printInit() {
	// board
	printBoard()
}
