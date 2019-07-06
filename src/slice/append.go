package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	a = append(a, 4, 5, 6)
	fmt.Println(a)

	b := []int{7, 8, 9}
	b = append(b, a...)
	fmt.Println(b)
}
