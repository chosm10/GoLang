package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	// fmt.Println(reflect.ValueOf(a).Int())  포인터 값을 int로 받으려 하므로 런타임 에러
	fmt.Println(reflect.ValueOf(a).Elem())
	fmt.Println(reflect.ValueOf(a).Elem().Int())

	var b interface{}
	b = 1

	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(b).Int())
	// fmt.Println(reflect.ValueOf(b).Elem())  포인터 형이 아닌데 elem으로 받아오려 하므로 런타임 에러
}
