package main

import (
	"testing"
)

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
