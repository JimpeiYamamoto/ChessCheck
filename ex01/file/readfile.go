package main

import (
	"io/ioutil"
	"strings"
)

func readFile(filepath string) (board []string, err error) {
	var bytes []byte
	bytes, err = ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	board = strings.Split(string(bytes), "\n")

	if len(board[len(board)-1]) == 0 {
		board = board[:len(board)-1]
	}
	return
}
