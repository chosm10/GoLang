package main

func main() {
	var r1 rune = '한'
	var r2 rune = '\ud55c'
	var r3 rune = '\U0000d55c'

	println(r1, r2, r3)
}
