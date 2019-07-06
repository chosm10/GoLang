package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func main() {
	hello()
	fmt.Println(sum(1, 2))
	fmt.Println(diff(5, 1))
	q, p := sumAndDiff(5, 2)
	fmt.Println(q, p)
	sum, diff := SumAndDiff(10, 2)
	fmt.Println(sum, diff)
}

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) (r int) {
	r = a - b
	return
}

func sumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func SumAndDiff(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}
