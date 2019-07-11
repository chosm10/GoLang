package main

import (
	"container/ring"
	"fmt"
)

func main() {
	people := []string{"Minho", "John", "Paul", "Ather"}

	r := ring.New(len(people))
	for i := 0; i < r.Len(); i++ {
		r.Value = people[i]
		r = r.Next()
	}

	r.Do(func(s interface{}) {
		fmt.Println(s)
	})

	fmt.Println("Move forward")
	r = r.Move(1)

	fmt.Println("Curr: ", r.Value)
	fmt.Println("Next: ", r.Next().Value)
	fmt.Println("Prev: ", r.Prev().Value)

	fmt.Println("Move backward")
	r = r.Move(-1)
	fmt.Println("Curr: ", r.Value)
	fmt.Println("Next: ", r.Next().Value)
	fmt.Println("Prev: ", r.Prev().Value)
}
