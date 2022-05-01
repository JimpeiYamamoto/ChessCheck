package main

import (
	"strings"
)

type Move struct {
	p     byte
	str   string
	from  string
	to    string
	castl string
	promo string
	// debug string
}

func collide(piece byte, x, y, dx, dy int) (int, int) {
	x = x + dx
	y = y + dy
	for (0 <= x && x < SIZE) && (0 <= y && y < SIZE) {
		if board[y][x] != EMPTY {
			if board[y][x] == piece {
				return x, y
			} else {
				return -1, -1
			}
		}
		x = x + dx
		y = y + dy
	}
	return -1, -1
}

func candidate(piece byte, fromX, fromY, toX, toY int) (int, int) {
	if piece == 'K' || piece == 'k' {
		for _, dy := range []int{-1, 0, 1} {
			for _, dx := range []int{-1, 0, 1} {
				// 自分自身
				if dy == 0 && dx == 0 {
					continue
				}
				// 範囲外
				if toX+dx < 0 || SIZE <= toX+dx {
					continue
				}
				if toY+dy < 0 || SIZE <= toY+dy {
					continue
				}
				if piece == board[toY+dy][toX+dx] {
					return toX + dx, toY + dy
				}
			}
		}
	}
	if piece == 'Q' || piece == 'q' {
		dx := []int{1, -1, 0, 0, 1, 1, -1, -1}
		dy := []int{0, 0, 1, -1, 1, -1, 1, -1}
		for i := range dx {
			x, y := collide(piece, toX, toY, dx[i], dy[i])
			if x < 0 || y < 0 {
				continue
			}
			if fromX >= 0 && x != fromX {
				continue
			}
			if fromY >= 0 && y != fromY {
				continue
			}
			if board[y][x] == piece {
				return x, y
			}
		}
	}
	if piece == 'R' || piece == 'r' {
		dx := []int{1, -1, 0, 0}
		dy := []int{0, 0, 1, -1}
		for i := range dx {
			x, y := collide(piece, toX, toY, dx[i], dy[i])
			if x < 0 || y < 0 {
				continue
			}
			if fromX >= 0 && x != fromX {
				continue
			}
			if fromY >= 0 && y != fromY {
				continue
			}
			if board[y][x] == piece {
				return x, y
			}
		}
	}
	if piece == 'N' || piece == 'n' {
		dx := []int{2, 2, -2, -2, 1, -1, 1, -1}
		dy := []int{1, -1, 1, -1, 2, 2, -2, -2}
		for i := range dx {
			x := toX + dx[i]
			y := toY + dy[i]
			if x < 0 || SIZE <= x || y < 0 || SIZE <= y {
				continue
			}
			if fromX >= 0 && x != fromX {
				continue
			}
			if fromY >= 0 && y != fromY {
				continue
			}
			if board[y][x] == piece {
				return x, y
			}
		}
	}
	if piece == 'B' || piece == 'b' {
		dx := []int{1, 1, -1, -1}
		dy := []int{1, -1, 1, -1}
		for i := range dx {
			x, y := collide(piece, toX, toY, dx[i], dy[i])
			if x < 0 || y < 0 {
				continue
			}
			if fromX >= 0 && x != fromX {
				continue
			}
			if fromY >= 0 && y != fromY {
				continue
			}
			if board[y][x] == piece {
				return x, y
			}
		}
	}
	if piece == 'p' {
		if board[toY][toX] != EMPTY {
			dx := []int{-1, 1}
			dy := []int{-1, -1}
			for i := range dx {
				x, y := toX+dx[i], toY+dy[i]
				if x < 0 || SIZE <= x || y < 0 || SIZE <= y {
					continue
				}
				if fromX >= 0 && x != fromX {
					continue
				}
				if fromY >= 0 && y != fromY {
					continue
				}
				if board[y][x] == piece {
					return x, y
				}
			}
		} else {
			dx := []int{0, 0}
			dy := []int{-1, -2}
			for i := range dx {
				x, y := toX+dx[i], toY+dy[i]
				if x < 0 || SIZE <= x || y < 0 || SIZE <= y {
					continue
				}
				if fromX >= 0 && x != fromX {
					continue
				}
				if fromY >= 0 && y != fromY {
					continue
				}
				if board[y][x] == piece {
					return x, y
				}
			}
		}
	}
	if piece == 'P' {
		if board[toY][toX] != EMPTY {
			dx := []int{-1, 1}
			dy := []int{1, 1}
			for i := range dx {
				x, y := toX+dx[i], toY+dy[i]
				if x < 0 || SIZE <= x || y < 0 || SIZE <= y {
					continue
				}
				if fromX >= 0 && x != fromX {
					continue
				}
				if fromY >= 0 && y != fromY {
					continue
				}
				if board[y][x] == piece {
					return x, y
				}
			}
		} else {
			dx := []int{0, 0}
			dy := []int{1, 2}
			for i := range dx {
				x, y := toX+dx[i], toY+dy[i]
				if x < 0 || SIZE <= x || y < 0 || SIZE <= y {
					continue
				}
				if fromX >= 0 && x != fromX {
					continue
				}
				if fromY >= 0 && y != fromY {
					continue
				}
				if board[y][x] == piece {
					return x, y
				}
			}
		}
	}
	return -1, -1
}

// fromを作成
func updateFrom(mv Move) Move {
	if len(mv.castl) > 0 {
		return mv
	}
	if len(mv.from) == 2 {
		return mv
	}
	fromX, fromY := str2xy(mv.from)
	toX, toY := str2xy(mv.to)
	canX, canY := candidate(mv.p, fromX, fromY, toX, toY)
	if canX < 0 || canY < 0 {
		panic("panic: candidate")
	}

	// fmt.Printf("%+v\n", struct {
	// 	x int
	// 	y int
	// }{canX, canY})
	mv.from = xy2str(canX, canY)
	return mv
}

func actMove(mv Move) {
	// キャスリングの場合
	if len(mv.castl) > 0 {
		var x int
		var y int
		if turn == WHITE {
			// right
			if strings.Compare(mv.castl, "O-O") == 0 {
				x, y = str2xy("g1")
				board[y][x] = 'K'
				x, y = str2xy("f1")
				board[y][x] = 'R'
				x, y = str2xy("e1")
				board[y][x] = EMPTY
				x, y = str2xy("h1")
				board[y][x] = EMPTY
			} else {
				x, y = str2xy("c1")
				board[y][x] = 'K'
				x, y = str2xy("d1")
				board[y][x] = 'R'
				x, y = str2xy("e1")
				board[y][x] = EMPTY
				x, y = str2xy("a1")
				board[y][x] = EMPTY
			}
		} else {
			// right
			if strings.Compare(mv.castl, "O-O") == 0 {
				x, y = str2xy("g8")
				board[y][x] = 'k'
				x, y = str2xy("f8")
				board[y][x] = 'r'
				x, y = str2xy("e8")
				board[y][x] = EMPTY
				x, y = str2xy("h8")
				board[y][x] = EMPTY
			} else {
				x, y = str2xy("c8")
				board[y][x] = 'k'
				x, y = str2xy("d8")
				board[y][x] = 'r'
				x, y = str2xy("e8")
				board[y][x] = EMPTY
				x, y = str2xy("a8")
				board[y][x] = EMPTY
			}
		}
		return
	}

	fromX, fromY := str2xy(mv.from)
	toX, toY := str2xy(mv.to)
	// アンパッサンを取る場合、ポーンも削除
	if board[toY][toX] == 'E' {
		board[toY-1][toX] = EMPTY
	}
	if board[toY][toX] == 'e' {
		board[toY+1][toX] = EMPTY
	}
	// TODO
	board[fromY][fromX] = EMPTY
	board[toY][toX] = mv.p

	// プロモーション
	if len(mv.promo) > 0 {
		if turn == BLACK {
			mv.promo = strings.ToLower(mv.promo)
		}
		board[toY][toX] = mv.promo[0]
	}
	// 前のアンパッサン削除
	for y := range board {
		for x := range board {
			if board[y][x] == 'E' || board[y][x] == 'e' {
				board[y][x] = EMPTY
			}
		}
	}

	// アンパサン追加
	if mv.p == 'P' {
		if toY-fromY == -2 {
			_, y := str2xy("3")
			board[y][toX] = 'e'
		}
	}
	if mv.p == 'p' {
		if toY-fromY == 2 {
			_, y := str2xy("6")
			board[y][toX] = 'e'
		}
	}
}

func move(str string) {
	// fmt.Println(str)
	// 棋譜をstructに変換
	mv := str2Move(str)
	// fmt.Printf("%+v\n", mv)
	// fromを埋める
	mv = updateFrom(mv)
	// fmt.Printf("%+v\n", mv)
	// structをもとに移動
	actMove(mv)
}
