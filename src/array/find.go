package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i])
	}

	b := [5]int{6, 7, 8, 9, 10}
	for i, value := range b {
		fmt.Println("index:", i, "value:", value)
	}
}
