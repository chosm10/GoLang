package main

func main() {
	var num1 int = 4
	var num2 int = 2

	println(num1 + num2)
	println(num1 - num2)
	println(num1 * num2)
	println(num1 / num2)
	println(num1 % num2)
	// println(num1 << num2)
	// println(num1 >> num2)
	println(^num1)

	var num3 float32 = 2
	// println(num1 + num3)
	println(float32(num1) + num3)
	println(num1 + int(num3))

	var num4 uint16 = 10
	var num5 uint32 = 80000

	println(num4 + uint16(num5))
}
