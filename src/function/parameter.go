package main

import "fmt"

func sum(n ...int) int {
	total := 0

	for _, value := range n {
		total += value
	}

	return total
}
4
func main() {
	r := sum(1, 2, 3, 4, 5)
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(r)
	fmt.Println(sum(arr...))
}
