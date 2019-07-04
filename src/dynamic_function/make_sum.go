package main

import (
	"fmt"
	"reflect"
)

func sum(args []reflect.Value) []reflect.Value {
	first, second := args[0], args[1]

	if first.Kind() != second.Kind() {
		fmt.Println("Please check variables type!")
		return nil
	}

	switch first.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(first.Int() + second.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(first.Uint() + second.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(first.Float() + second.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(first.String() + second.String())}
	default:
		return []reflect.Value{}
	}
}

func make_sum(func_pointer interface{}) {
	make_function := reflect.ValueOf(func_pointer).Elem()
	target := reflect.MakeFunc(make_function.Type(), sum)
	make_function.Set(target)
}

func main() {
	var int_sum func(int, int) int64
	var uint_sum func(uint, uint) uint64
	var float_sum func(float32, float32) float64
	var string_sum func(string, string) string

	make_sum(&int_sum)
	make_sum(&uint_sum)
	make_sum(&float_sum)
	make_sum(&string_sum)

	fmt.Println(int_sum(2, 3))
	fmt.Println(uint_sum(2, 4))
	fmt.Println(float_sum(2.1, 3.6))
	fmt.Println(string_sum("Happy ", "NewYear"))
}
