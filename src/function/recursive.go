package main

import "fmt"

func factorial(num int64) int64 {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {
	fmt.Println(factorial(5))
}
