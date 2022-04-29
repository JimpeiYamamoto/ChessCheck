package main

type Piece rune

const (
	KING   = 'K'
	PAWN   = 'P'
	BISHOP = 'B'
	ROOK   = 'R'
	QUEEN  = 'Q'
)

func (p Piece) isPiece() bool {
	const pieces = "KPBRQ"

	for _, v := range pieces {
		if v == rune(p) {
			return true
		}
	}
	return false
}

func (p Piece) isCollideKing(board Board, y, x int) bool {
	m := map[Piece]func(Board, int, int) bool{
		PAWN: isCollideByPawn,
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
