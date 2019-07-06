package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello, World!")
	}()

	func(s string) {
		fmt.Println(s)
	}("Hello, World!")

	r := func(a int, b int) int {
		return a + b
	}(3, 5)
	fmt.Println(r)
}
