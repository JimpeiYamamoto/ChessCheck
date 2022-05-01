package main

func main() {
	moves := getPgn()
	// for i, v := range moves {
	// 	fmt.Println(i, v)
	// }
	printInit()

	for _, mv2 := range moves {
		for _, mv := range mv2 {
			move(mv)
			print()
			turn = WHITE + BLACK - turn
			// return
		}
	}
}
