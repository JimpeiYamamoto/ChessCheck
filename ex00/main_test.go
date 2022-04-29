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
