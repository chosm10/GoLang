package main

import (
	"fmt"
	"reflect"
)

func target(args []reflect.Value) []reflect.Value {
	fmt.Println("method connecting!")

	return nil
}

func main() {
	var make_function func()

	function_info := reflect.ValueOf(&make_function).Elem()
	target_info := reflect.MakeFunc(function_info.Type(), target)

	function_info.Set(target_info)
	make_function()
}
