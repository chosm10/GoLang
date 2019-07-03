package main
import "fmt"

type Rectangle struct {
	width int
	height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func main() {

	rect := NewRectangle(30, 40)
	fmt.Println(rect)
}