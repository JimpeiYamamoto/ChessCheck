package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func getPgn() [][]string {
	// 文字列にする
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Trim(scanner.Text(), " ")
		if strings.HasPrefix(input, "[") {
			continue
		}
		if len(input) == 0 {
			continue
		}
		str += input + " "
	}
	str = strings.ReplaceAll(str, "1/2-1/2", "")
	str = strings.ReplaceAll(str, "1-0", "")
	str = strings.ReplaceAll(str, "0-1", "")
	// checkmating
	str = strings.ReplaceAll(str, "#", "")
	// check
	str = strings.ReplaceAll(str, "+", "")
	// capture
	str = strings.ReplaceAll(str, "x", "")
	// asta
	str = strings.ReplaceAll(str, "*", "")
	str = strings.Trim(str, " ")
	// fmt.Println(str)

	// '1. 2. ...'を正規表現でsplit
	reg := `[0-9]+\.`
	arr1 := regexp.MustCompile(reg).Split(str, -1)
	var arr []string
	for _, v := range arr1 {
		arr = append(arr, strings.Trim(v, " "))
	}
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	// "e2 e4" -> {"e2", "e4"}
	arr2 := [][]string{}
	for _, v := range arr {
		arr2 = append(arr2, strings.Split(v, " "))
	}
	return arr2[1:]
}
