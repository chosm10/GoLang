package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("hello1.txt")
	defer file1.Close()
	fmt.Fprint(file1, 1, 1.1, "Hello, World!")

	file2, _ := os.Create("hello2.txt")
	defer file2.Close()
	fmt.Fprintln(file2, 2, 2.2, "Nice Day!")

	file3, _ := os.Create("hello3.txt")
	defer file3.Close()
	fmt.Fprintf(file3, "&d,&f,&s", 3, 3.3, "groovy everywhere!")
}
