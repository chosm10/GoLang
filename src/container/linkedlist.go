package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()
	myList.PushBack(10)
	myList.PushBack(20)
	myList.PushBack(30)

	fmt.Println("Front: ", myList.Front().Value)
	fmt.Println("Back: ", myList.Back().Value)

	for node := myList.Front(); node != nil; node = node.Next() {
		fmt.Println(node.Value)
	}
}
