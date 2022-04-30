package main

import (
	"strings"
	"testing"
)

func TestIsOnlyKing(t *testing.T) {
	trueTest := [][]string{
		[]string{
			"R.P.",
			".K..",
			"..P.",
			"....",
		},
		[]string{
			"Kk",
			"kk",
		},
		[]string{
			"K",
		},
	}
	falseTest := [][]string{
		[]string{
			".K",
			".K",
		},
		[]string{
			".K",
			"K.",
		},
		[]string{
			"KK.",
			"...",
			"...",
		},
		[]string{
			"k",
		},
		[]string{
			"K.K",
			"...",
			"...",
		},
		[]string{
			"K..",
			"...",
			"K..",
		},
	}
	for _, v := range trueTest {
		if isOnlyKing(v) != true {
			t.Errorf("error: expected true\n%s", strings.Join(v[:], "\n"))
		}
	}
	for _, v := range falseTest {
		if isOnlyKing(v) != false {
			t.Errorf("error: expected false\n%s", strings.Join(v[:], "\n"))
		}
	}
}

func TestIsSquare(t *testing.T) {
	tests := []struct {
		board    []string
		expected bool
	}{
		{[]string{
			"R.P.",
			".K..",
			"..P.",
			"....",
		}, true},
		{[]string{
			"R",
		}, true},
		{[]string{
			"R",
			"R",
		}, false},
		{[]string{
			"",
		}, false},
		{[]string{}, false},
	}

	for i, v := range tests {
		if isSquare(v.board) != v.expected {
			t.Errorf("error: index=%d", i)
		}
	}
}

func TestIsCheck(t *testing.T) {
	testIsCheckForPawn(t)
	testIsCheckForBishop(t)
	testIsCheckForRook(t)
	testIsCheckForQueen(t)
	testIsCheckMix(t)
}

func testIsCheckForPawn(t *testing.T) {
	// trueTest := [][]byte{
	// 	[]byte("PPPPP"),
	// 	[]byte("PPPPP"),
	// 	[]byte("PPKPP"),
	// 	[]byte("PPPPP"),
	// 	[]byte("PPPPP"),
	// }
	trueTest := [][]byte{
		[]byte("....."),
		[]byte("....."),
		[]byte("..K.."),
		[]byte(".P.P."),
		[]byte("....."),
	}
	falseTest := [][]byte{
		[]byte("PPPPP"),
		[]byte("PPPPP"),
		[]byte("PPKPP"),
		[]byte("P.P.P"),
		[]byte("PPPPP"),
	}
	testIsCheckHelper(t, trueTest, true, 'P')
	testIsCheckHelper(t, falseTest, false, 'P')
}

func testIsCheckForBishop(t *testing.T) {
	// trueTest := [][]byte{
	// 	[]byte("BBBBB"),
	// 	[]byte("BBBBB"),
	// 	[]byte("BBKBB"),
	// 	[]byte("BBBBB"),
	// 	[]byte("BBBBB"),
	// }
	trueTest := [][]byte{
		[]byte("B...B"),
		[]byte(".B.B."),
		[]byte("..K.."),
		[]byte(".B.B."),
		[]byte("B...B"),
	}
	falseTest := [][]byte{
		[]byte(".BBB."),
		[]byte("B.B.B"),
		[]byte("BBKBB"),
		[]byte("B.B.B"),
		[]byte(".BBB."),
	}
	testIsCheckHelper(t, trueTest, true, 'B')
	testIsCheckHelper(t, falseTest, false, 'B')
}

func testIsCheckForRook(t *testing.T) {
	// trueTest := [][]byte{
	// 	[]byte("RRRRR"),
	// 	[]byte("RRRRR"),
	// 	[]byte("RRKRR"),
	// 	[]byte("RRRRR"),
	// 	[]byte("RRRRR"),
	// }
	trueTest := [][]byte{
		[]byte("..R.."),
		[]byte("..R.."),
		[]byte("RRKRR"),
		[]byte("..R.."),
		[]byte("..R.."),
	}
	falseTest := [][]byte{
		[]byte("RR.RR"),
		[]byte("RR.RR"),
		[]byte("..K.."),
		[]byte("RR.RR"),
		[]byte("RR.RR"),
	}
	testIsCheckHelper(t, trueTest, true, 'R')
	testIsCheckHelper(t, falseTest, false, 'R')
}

func testIsCheckForQueen(t *testing.T) {
	// trueTest := [][]byte{
	// 	[]byte("QQQQQ"),
	// 	[]byte("QQQQQ"),
	// 	[]byte("QQKQQ"),
	// 	[]byte("QQQQQ"),
	// 	[]byte("QQQQQ"),
	// }
	trueTest := [][]byte{
		[]byte("Q.Q.Q"),
		[]byte(".QQQ."),
		[]byte("QQKQQ"),
		[]byte(".QQQ."),
		[]byte("Q.Q.Q"),
	}
	falseTest := [][]byte{
		[]byte(".Q.Q."),
		[]byte("Q...Q"),
		[]byte("..K.."),
		[]byte("Q...Q"),
		[]byte(".Q.Q."),
	}
	testIsCheckHelper(t, trueTest, true, 'Q')
	testIsCheckHelper(t, falseTest, false, 'Q')
}

func testIsCheckMix(t *testing.T) {
	trueTest := [][]string{
		[]string{
			"RB..",
			".K..",
			"..P.",
			"Q...",
		},
		[]string{
			"PB..",
			".K.R",
			"....",
			"Q...",
		},
		[]string{
			"RP..",
			".K..",
			"....",
			"Q..B",
		},
		[]string{
			"RBP.",
			".K..",
			"....",
			"あQ..",
		},
		[]string{
			"RBP.",
			".K..",
			"....",
			"♘Q..",
		},
	}
	falseTest := [][]string{
		[]string{
			"K...",
			"....",
			"P...",
			"Q...",
		},
		[]string{
			"...K",
			"....",
			".R..",
			"Q...",
		},
		[]string{
			"QB.K",
			"....",
			"....",
			"....",
		},
		[]string{
			"Q...",
			".R..",
			"....",
			"...K",
		},
		[]string{
			"Q...",
			"....",
			"P...",
			"K...",
		},
		[]string{
			"...Q",
			"....",
			".R..",
			"K...",
		},
		[]string{
			"KB.Q",
			"....",
			"....",
			"....",
		},
		[]string{
			"K...",
			".R..",
			"....",
			"...Q",
		},
	}
	for _, v := range trueTest {
		if isCheck(v) != true {
			t.Errorf("error: expected true\n%s", strings.Join(v[:], "\n"))
		}
	}
	for _, v := range falseTest {
		if isCheck(v) != false {
			t.Errorf("error: expected false\n%s", strings.Join(v[:], "\n"))
		}
	}
}

func testIsCheckHelper(t *testing.T, board [][]byte, expected bool, piece byte) {
	for y := range board {
		for x := range board {
			if board[y][x] != piece {
				continue
			}
			tmp := [][]byte{
				[]byte("....."),
				[]byte("....."),
				[]byte("..K.."),
				[]byte("....."),
				[]byte("....."),
			}
			tmp[y][x] = piece
			tmpBoard := []string{
				string(tmp[0]),
				string(tmp[1]),
				string(tmp[2]),
				string(tmp[3]),
				string(tmp[4]),
			}
			if isCheck(tmpBoard) != expected {
				if expected {
					t.Errorf("error: expected true\n%s", strings.Join(tmpBoard[:], "\n"))
				} else {
					t.Errorf("error: expected false\n%s", strings.Join(tmpBoard[:], "\n"))
				}
			}
		}
	}
}
