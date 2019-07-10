package main

import (
	"fmt"
	"sort"
)

func main() {
	intArr := []int{9, 7, 5, 3, 1}
	floatArr := []float64{3.2, 7.8, 5.2, 10.8, 9.2}
	stringArr := []string{"Maria", "Destiny", "define", "console", "zebra"}

	sort.Sort(sort.IntSlice(intArr))
	fmt.Println(intArr)
	sort.Sort(sort.Reverse(sort.IntSlice(intArr)))
	fmt.Println(intArr)

	sort.Sort(sort.Float64Slice(floatArr))
	fmt.Println(floatArr)
	sort.Sort(sort.Reverse(sort.Float64Slice(floatArr)))
	fmt.Println(floatArr)

	//대문자 A-Z, 소문자 a-z 순으로 오름차순 정렬됨
	sort.Sort(sort.StringSlice(stringArr))
	fmt.Println(stringArr)
	sort.Sort(sort.Reverse(sort.StringSlice(stringArr)))
	fmt.Println(stringArr)
}
