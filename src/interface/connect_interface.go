package main

import "fmt"

type myInt int

func (i myInt) print() {
	fmt.Println(i)
}

type Printer interface {
	print()
}

func main() {
	var i myInt = 5
	var myPrinter Printer

	myPrinter = i
	myPrinter.print()

}
