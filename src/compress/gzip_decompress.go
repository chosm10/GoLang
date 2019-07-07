package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("hello.txt.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reader.Close()

	decompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(decompressed))
}
