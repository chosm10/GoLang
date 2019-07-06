package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}
func world() {
	fmt.Println("world")
}

func helloworld() {
	defer func() {
		fmt.Println("Hello")
	}()

	func() {
		fmt.Println("World")
	}()
}
func main() {
	defer hello()
	world()
	world()
	world()
	helloworld()

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
