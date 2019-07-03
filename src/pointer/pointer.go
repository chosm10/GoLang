package main

import "fmt"

func main() {
	var numPtr *int
	fmt.Println(numPtr)

	var ptr *int = new(int)
	fmt.Println(ptr)

	*ptr = 1
	fmt.Println(*ptr)

	var address int = 1
	var reference *int = &address
	fmt.Println(&address)
	fmt.Println(reference)
}
