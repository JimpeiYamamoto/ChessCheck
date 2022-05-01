package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	board, err := readFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	checkmate(board)
}
