package main

import "fmt"

func main() {
	var s1 string = "한글"
	var s2 string = "한글"
	var s3 string = "GO"
	var s4 string = "Hello"

	println(s1 == s2)
	println(s1 + s2 + s3)
	println("안녕하세요" + s1)
	fmt.Printf("%c\n", s4[1])
}
