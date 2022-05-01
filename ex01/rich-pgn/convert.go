package main

import "strings"

// e3 = {4, 5}
func str2xy(str string) (int, int) {
	x := -1
	y := -1
	if len(str) == 2 {
		x = int(str[0] - 'a')
		y = int('8' - str[1])
	}
	if len(str) == 1 {
		if 'a' <= str[0] && str[0] <= 'h' {
			x = int(str[0] - 'a')
		}
		if '1' <= str[0] && str[0] <= '8' {
			y = int('8' - str[0])
		}
	}
	return x, y
}

// {4, 5} = e3
func xy2str(x, y int) string {
	str := string([]byte{'a' + byte(x), '8' - byte(y)})
	return str
}

func str2Move(str string) Move {
	var mv Move

	// プロモーション
	if strings.Contains(str, "=") {
		mv.promo = str[len(str)-1:]
		str = str[:len(str)-2]
	}

	if strings.Contains(str, "O-O") {
		// キャスリング
		mv.castl = strings.Trim(str, " ")
		str = ""
	} else if strings.Contains("KQBNR", str[0:1]) {
		mv.str = string(str[0])
		mv.p = str[0]
		str = str[1:]
		if turn == BLACK {
			mv.str = strings.ToLower(mv.str)
			mv.p = byte(mv.str[0])
		}
	} else {
		mv.p = 'P'
		mv.str = "P"
		if turn == BLACK {
			mv.str = strings.ToLower(mv.str)
			mv.p = byte(mv.str[0])
		}
	}

	len := len(str)
	if len > 2 {
		mv.from = str[:len-2]
		str = str[len-2:]
	}
	mv.to = str

	return mv
}
