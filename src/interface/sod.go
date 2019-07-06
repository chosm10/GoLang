package main

import "fmt"

func hello() <-chan int {

}

func main() {

	var a chan int
	b := make(chan int)
	fmt.Println(a)
	fmt.Println(b)

	v, ok = b
	if ok {
		fmt.Println(v)
	}
}
