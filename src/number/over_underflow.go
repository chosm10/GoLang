package main

import "math"

func main() {
	var num1 uint8 = math.MaxUint8
	var num2 uint16 = math.MaxUint16
	var num3 uint32 = math.MaxUint32
	var num4 uint64 = math.MaxUint64

	println(num1, num2, num3, num4)
	// num1 += 1
	// num2 += 1
	// num3 += 1
	// num4 += 1
	// println(num1, num2, num3, num4)
	num1 = 0
	num2 = 0
	num3 = 0
	num4 = 0
	println(num1-1, num2-1, num3-1, num4-1)

}
