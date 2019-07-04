package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello() {
	fmt.Println("Hello World")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	once := new(sync.Once)

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("Go Routine : ", n)
			once.Do(hello)
		}(i)
	}

	fmt.Scanln()
}
