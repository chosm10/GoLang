package main

func main() {
	var s1 string = "Hello, world!\n"
	var s2 string = "안녕하세요\n"
	var s3 string = "\ud55c\uae00"
	var s4 string = "\U0000d55c\U0000ae00"
	var s5 string = "\xed\x95\x9c\xea\xb8\x80"
	var s6 string = `안녕하세요
	Hello, world!`

	println(s1)
	println(s2)
	println(s3)
	println(s4)
	println(s5)
	println(s6)
}
