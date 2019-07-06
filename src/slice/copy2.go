package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3)

	//deep copy
	copy(b, a)
	fmt.Println(a)
	fmt.Println(b)

	b[1] = 9
	fmt.Println(a)
	fmt.Println(b)
}
