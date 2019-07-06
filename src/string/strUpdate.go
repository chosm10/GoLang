package main

import "fmt"

func main() {
	var s string = "hello world!\n"

	s = "abcdef"

	println(s)
	fmt.Printf("%c\n", s[1])
	// 기존 문자열에 다시 할당은 가능해도 수정은 불가능함!
	// s[1] = 'z'
}
