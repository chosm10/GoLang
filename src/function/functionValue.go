package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	var hello func(a int, b int) int = sum
	world := hello

	fmt.Println(hello(1, 2))
	fmt.Println(world(1, 2))

	f := []func(int, int) int{sum, diff}
	fmt.Println(f[0](2, 5))
	fmt.Println(f[1](5, 1))

	function := map[string]func(int, int) int{
		"sum":  sum,
		"diff": diff,
	}
	fmt.Println(function["sum"](1, 2))
	fmt.Println(function["diff"](5, 1))
}
