package main

import (
	"testing"
)

func TestIsSquare(t *testing.T) {
	var board []string

	board = []string{
		"R.P.",
		".K..",
		"..P.",
		"....",
	}
	if !isSquare(board) {
		t.Errorf("aaa")
	}
	board = []string{
		"R.P.",
		".K..",
		"..P..",
		"....",
	}
	if isSquare(board) {
		t.Errorf("bbb")
	}
	board = []string{
		"R",
	}
	if !isSquare(board) {
		t.Errorf("ccc")
	}
	board = []string{
		"R",
		"R",
	}
	if isSquare(board) {
		t.Errorf("ccc")
	}
}
