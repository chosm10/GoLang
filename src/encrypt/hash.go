package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {

	str := "Find Yourself"
	hash1 := sha512.Sum512([]byte(str))
	fmt.Printf("%x\n", hash1)

	sha := sha512.New()
	sha.Write([]byte("Find "))
	sha.Write([]byte("Yourself"))

	hash2 := sha.Sum(nil)
	fmt.Printf("%x\n", hash2)
}
