package piscine

import "ft"

func PrintStr(s string) {
	runes := []rune(s)

	for _, r := range runes {
		ft.PrintRune(r)
	}
}
