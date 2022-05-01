package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func toEmoji(r rune, y, x int) string {
	blackIdx := strings.Index("prnbqk", string(r))
	if blackIdx >= 0 {
		return string(([]rune("♟♜♞♝♛♚"))[blackIdx])
	}
	whiteIdx := strings.Index("PRNBQK", string(r))
	if whiteIdx >= 0 {
		return string(([]rune("♙♖♘♗♕♔"))[whiteIdx])
	}
	return "."
	// return string(([]rune(""))[(y+x)%2])
}

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
	for y, row := range tmpBoard {
		var p string
		for x, r := range row {
			p += toEmoji(r, y, x) + " "
		}
		fmt.Println(p)
	}
	// アンパサン表示用
	// for _, row := range board {
	// 	fmt.Println(string(row))
	// }
	fmt.Println()
	time.Sleep(1 * time.Second)
}

func print() {
	fmt.Print("\033[2J\033[H")

	// board
	printBoard()
}

func printInit() {
	fmt.Print("\033[2J\033[H")
	// board
	printBoard()
	time.Sleep(1 * time.Second)
}
