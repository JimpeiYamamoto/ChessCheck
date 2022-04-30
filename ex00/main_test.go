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
}

func testIsCheckForPawn(t *testing.T) {
	trueTest := [][]string{
		[]string{
			"K..",
			".P.",
			"...",
		},
		[]string{
			"..K",
			".P.",
			"...",
		},
		[]string{
			"PPP",
			"PKP",
			"PPP",
		},
	}
	falseTest := [][]string{
		[]string{
			"P",
		},
		[]string{
			".P",
			"..",
		},
		[]string{
			".K.",
			".P.",
			"...",
		},
		[]string{
			".K.",
			".P.",
			"...",
		},
		[]string{
			"...",
			".PK",
			"...",
		},
		[]string{
			"...",
			".P.",
			"..K",
		},
		[]string{
			"...",
			".P.",
			".K.",
		},
		[]string{
			"...",
			".P.",
			"K..",
		},
		[]string{
			"...",
			"KP.",
			"...",
		},
		[]string{
			"K..",
			"..P",
			"...",
		},
		[]string{
			".P.",
			"...",
			"..K",
		},
		[]string{
			"K..",
			"...",
			"..P",
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

func testIsCheckForBishop(t *testing.T) {
	trueTest := [][]string{
		[]string{
			"K..",
			".B.",
			"...",
		},
		[]string{
			"..K",
			".B.",
			"...",
		},
		[]string{
			"...",
			".B.",
			"..K",
		},
		[]string{
			"...",
			".B.",
			"K..",
		},
		[]string{
			"..K",
			"...",
			"B..",
		},
		[]string{
			"..B",
			"...",
			"K..",
		},
		[]string{
			"K..",
			"...",
			"..B",
		},
		[]string{
			"B..",
			"...",
			"..K",
		},
		[]string{
			"BBB",
			"BKB",
			"BBB",
		},
	}
	falseTest := [][]string{
		[]string{
			"B",
		},
		[]string{
			".B",
			"..",
		},
		[]string{
			".K.",
			".B.",
			"...",
		},
		[]string{
			"...",
			".BK",
			"...",
		},
		[]string{
			"...",
			".B.",
			".K.",
		},
		[]string{
			"...",
			"KB.",
			"...",
		},
		[]string{
			"K..",
			"..B",
			"...",
		},
		[]string{
			".B.",
			"...",
			"..K",
		},
		[]string{
			"...",
			"K.B",
			"...",
		},
		[]string{
			"...",
			"B.K",
			"...",
		},
		[]string{
			".K.",
			"...",
			".B.",
		},
		[]string{
			".B.",
			"...",
			".K.",
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
