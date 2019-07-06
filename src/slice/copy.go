package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	//deep copy
	b := a
	b[1] = 9
	fmt.Println(a)
	fmt.Println(b)

	c := []int{4, 5, 6}
	//shallow copy
	d := c
	d[1] = 9
	fmt.Println(c)
	fmt.Println(d)
}
