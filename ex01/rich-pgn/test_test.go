package main

import (
	"strings"
	"testing"
)

func TestXy2Str(t *testing.T) {
	tests := []struct {
		x        int
		y        int
		expected string
	}{
		{0, 0, "a8"},
		{7, 7, "h1"},
	}

	for i, v := range tests {
		actual := xy2str(v.x, v.y)
		if strings.Compare(actual, v.expected) != 0 {
			t.Errorf("case %d: case={%d, %d}, actual=%s, expected=%s\n", i, v.x, v.y, actual, v.expected)
		}
	}
}

func TestStr2Xy(t *testing.T) {
	tests := []struct {
		str       string
		expectedX int
		expectedY int
	}{
		{"a8", 0, 0},
		{"h1", 7, 7},
		{"h", 7, -1},
		{"a", 0, -1},
		{"1", -1, 7},
		{"8", -1, 0},
		{"", -1, -1},
	}

	for i, v := range tests {
		actualX, actualY := str2xy(v.str)
		if actualX != v.expectedX || actualY != v.expectedY {
			t.Errorf("case %d: case=%s, actual={%d, %d}, expected={%d, %d}\n", i, v.str, actualX, actualY, v.expectedX, v.expectedY)
		}
	}
}
