package main

import "unicode/utf8"

func main() {
	var s1 string = "한글"
	var s2 string = "Hello"

	println(len(s1))
	println(len(s2))
	println(utf8.RuneCountInString(s1))
}
