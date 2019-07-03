package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main() {
	var number int = 1
	fmt.Println(reflect.TypeOf(number))

	var my_string string = "my go study!"
	fmt.Println(reflect.TypeOf(my_string))

	var my_float float32 = 1.3
	fmt.Println(reflect.TypeOf(my_float))

	var my_data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(my_data))
}
