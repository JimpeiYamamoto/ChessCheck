package main

type Piece rune

const (
	KING   = 'K'
	PAWN   = 'P'
	BISHOP = 'B'
	ROOK   = 'R'
	QUEEN  = 'Q'
	KNIGHT = 'N'
)

func (p Piece) isPiece() bool {
	const pieces = "KPBRQN"

	for _, v := range pieces {
		if v == rune(p) {
			return true
		}
	}
	return false
}

func (p Piece) isCollideKing(board Board, y, x int) bool {
	m := map[Piece]func(Board, int, int) bool{
		PAWN:   isCollideByPawn,
		BISHOP: isCollideByBishop,
		ROOK:   isCollideByRook,
		QUEEN:  isCollideByQueen,
		KNIGHT: isCollideByKnight,
	}

	if p.isPiece() && p != KING {
		return m[p](board, y, x)
	}
	return false
}

func isCollideByPawn(b Board, y, x int) bool {
	if b.at(y-1, x-1) == KING || b.at(y-1, x+1) == KING {
		return true
	}
	return false
}

func isCollideByKnight(b Board, y, x int) bool {
	dx := []int{-2, -1, -2, -1, +2, +1, +2, +1}
	dy := []int{-1, -2, +1, +2, -1, -2, +1, -2}
	for i := range dy {
		if Piece(b.at(y+dy[i], x+dx[i])) == KING {
			return true
		}
	}
	return false
}

func checkDirection(b Board, y, x, dy, dx int) bool {
	for y, x = y+dy, x+dx; b.isInBoard(y, x); y, x = y+dy, x+dx {
		if Piece(b.at(y, x)).isPiece() {
			return b.at(y, x) == KING
		}
	}
	return false
}

func isCollideByBishop(b Board, y, x int) bool {
	dx := []int{-1, 1, 1, -1}
	dy := []int{-1, -1, 1, 1}
	for i := range dx {
		if checkDirection(b, y, x, dx[i], dy[i]) {
			return true
		}
	}
	return false
}

func isCollideByRook(b Board, y, x int) bool {
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}
	for i := range dx {
		if checkDirection(b, y, x, dx[i], dy[i]) {
			return true
		}
	}
	return false
}

func isCollideByQueen(b Board, y, x int) bool {
	return isCollideByBishop(b, y, x) || isCollideByRook(b, y, x)
}
