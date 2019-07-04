package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	i := 4
	if i >= 5 {
		println("5이상")
	} else {
		println("5이하")
	}

	a := 6
	if a >= 10 {
		println("10이상")
	} else if a >= 5 && a < 10 {
		println("5이상 10미만")
	} else {
		println("5미만")
	}

	// var b []byte
	// var err error

	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b)
	} else {
		fmt.Println(err)
	}

	// println(b)
	// println(err)
}
