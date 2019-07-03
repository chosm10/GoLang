package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 3.2
	f_type := reflect.TypeOf(f)
	f_value := reflect.ValueOf(f)

	fmt.Println(f_type.Name())
	fmt.Println(f_type.Size())
	fmt.Println(f_type.Kind() == reflect.Float64)

	fmt.Println(f_type.Kind() == reflect.Int64)

	fmt.Println(f_value.Type())
	fmt.Println(f_value.Kind() == reflect.Float64)

	fmt.Println(f_value.Kind() == reflect.Int32)
	fmt.Println(f_value.Float())
}
