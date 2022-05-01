package piscine

func Len(a interface{}) int {
	len := 0
	switch val := a.(type) {
	case []string:
		len = Strslen(val)
	case string:
		len = Strlen(val)
	}
	return len
}
