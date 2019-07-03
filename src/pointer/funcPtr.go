package main

import "fmt"

func hello(a *int) {
	*a = 2
}
func main() {
	var num int = 1
	hello(&num)

	fmt.Println(num)
}
