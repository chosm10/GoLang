package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt.gz",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := gzip.NewWriter(file)
	defer writer.Close()

	str := "Happy Coding!"
	writer.Write([]byte(str))
	writer.Flush()
}
