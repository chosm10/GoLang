package main

import "fmt"

func f() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	panic("panic")
}

func main() {
	f()
	fmt.Println("Hello, World!")
}
