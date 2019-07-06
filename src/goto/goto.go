package main

import "fmt"

func main() {
	// var a int = 1

	// if a == 1 {
	// 	goto ERROR1
	// } else if a == 2 {
	// 	goto ERROR2
	// } else if a == 3 {
	// 	goto ERROR1
	// }

	// fmt.Println(a)
	// fmt.Println("Success")

	var b int = 1
	if b == 1 {
		goto ERROR
		c := 1
		fmt.Println(c)
	}

	return
ERROR:
	fmt.Println("ERROR")
	// ERROR1:
	// 	fmt.Println("ERROR1")
	// 	return

	// ERROR2:
	// 	fmt.Println("ERROR2 발생")
	// 	return

}
