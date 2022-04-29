package piscine

func Strlen(str string) (i int) {
	runs := []rune(str)

	for range runs {
		i++
	}
	return i
}
