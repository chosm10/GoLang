package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func main() {
	var rect1 Rectangle = Rectangle{10, 20}
	rect2 := Rectangle{45, 30}
	rect3 := Rectangle{height: 30, width: 45}

	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)

	var rect4 Rectangle
	rect4.height = 4
	rect4.width = 12
	fmt.Println(rect4)

	var rect5 *Rectangle = new(Rectangle)
	rect5.height = 12
	rect5.width = 15
	fmt.Println(rect5)
}
