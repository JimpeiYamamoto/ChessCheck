package main

import (
	"strings"
	"testing"
)

func boardEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}

func TestReadfile(t *testing.T) {
	tests := map[string][]string{
		"boads/empty": []string{""},
		"boads/pdf_fail": []string{
			"..",
			".K",
		},
		"boads/pdf_success": []string{
			"R...",
			".K..",
			"..P.",
			"....",
		},
	}

	for p, ex := range tests {
		ac, _ := readFile(p)
		if boardEqual(ex, ac) {
			t.Errorf("error: %s\nexpect: rows(%d)\n%s\nactualy: rows(%d)\n%s\n", p, len(ex), strings.Join(ex, "\n"), len(ac), strings.Join(ac, "\n"))
		}
	}
}
