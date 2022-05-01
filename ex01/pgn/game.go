package main

const (
	SIZE = 8
)

// ターン
const (
	WHITE = iota
	BLACK = iota
)

// 駒
const (
	P = 'P'
	K = 'K'
	Q = 'Q'
	N = 'N'
	R = 'R'
	B = 'B'
	// アンパッサン
	E = 'E'
	// 空
	EMPTY = '.'
)

var board [][]byte
var turn int
var cnt int
