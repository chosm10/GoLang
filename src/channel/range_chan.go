package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

	channel := make(chan int, 1)
	go func() {
		channel <- 1
	}()

	a, ok := <-channel
	fmt.Println(a, ok)

	close(channel)

	a, ok = <-channel
	fmt.Println(a, ok)

}
