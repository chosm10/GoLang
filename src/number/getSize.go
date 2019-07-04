package main

import "unsafe"

func main() {
	var num1 int8 = 1
	var num2 int16 = 1
	var num3 int32 = 1
	var num4 int64 = 1

	println(unsafe.Sizeof(num1))
	println(unsafe.Sizeof(num2))
	println(unsafe.Sizeof(num3))
	println(unsafe.Sizeof(num4))
}
