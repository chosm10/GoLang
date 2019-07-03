package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case Person:
		p := arg.(Person)
		return "name : " + p.name + " age : " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return "name : " + p.name + " age : " + strconv.Itoa(p.age)
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(Person{"Maria", 20}))
	fmt.Println(formatString(&Person{"Clone", 30}))

	var p *Person = new(Person)
	p.name = "bolton"
	p.age = 15
	fmt.Println(p)

	var i interface{}
	i = Person{"Dingo", 25}

	if value, ok := i.(Person); ok {
		fmt.Println(value, ok)
	}
}
