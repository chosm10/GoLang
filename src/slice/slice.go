package main

import "fmt"

func main() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b)

	fmt.Println(len(a))
	fmt.Println(cap(a))
}
