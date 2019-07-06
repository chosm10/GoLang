package main

import "fmt"

func main() {
	var a [5]int
	a[2] = 5
	fmt.Println(a)

	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [5]int{6, 7, 8, 9, 10}
	d := [...]int{1, 2, 3}

	fmt.Println(c)
	fmt.Println(d)

	e := [...]string{"Maria", "Scott", "Paul"}
	fmt.Println(e)

	f := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(f)
}
