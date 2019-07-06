package main

import "fmt"

type myInt int
type Rectangle struct {
	width  int
	height int
}

func (i myInt) print() {
	fmt.Println(i)
}

func (r Rectangle) print() {
	fmt.Println(r.height, r.width)
}

type Printer interface {
	print()
}

func main() {
	var i myInt = 5
	var r Rectangle = Rectangle{10, 20}

	var p Printer
	p = i
	p.print()

	p = r
	p.print()

	p1 := Printer(i)
	p1.print()

	p1 = Printer(r)
	p1.print()

	pArr := []Printer{i, r}
	for index, _ := range pArr {
		pArr[index].print()
	}
	for _, value := range pArr {
		value.print()
	}
}
