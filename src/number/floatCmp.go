package main

import "math"

func main() {
	var num float64 = 10.0

	for i := 0; i < 10; i++ {
		num -= 0.1
	}
	println(num)
	println(num == 9.0)

	const epsilon = 1e-14
	println(math.Abs(num-9.0) <= epsilon)
}
