package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a)
	fmt.Println(b)
	a[2] = 10

	fmt.Println(a)
	fmt.Println(b)
}
