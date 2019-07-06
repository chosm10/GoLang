package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[0:5]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a[0:3])
	fmt.Println(a[1:5])
	fmt.Println(a[2:5])
	fmt.Println(a[3:])
	fmt.Println(a[:3])
	fmt.Println(a[:])
	fmt.Println(a[0:len(a)])

	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3]
	d[0] = 9
	fmt.Println(c)
	fmt.Println(d)

	e := []int{1, 2, 3, 4, 5, 6, 7, 8}
	f := e[0:6:8]
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(len(f), cap(f))
}
