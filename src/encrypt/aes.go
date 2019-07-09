package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := "bingo, key 12345"
	str := "plain, texts! 12"

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := make([]byte, len(str))
	block.Encrypt(ciphertext, []byte(str))
	fmt.Printf("%x\n", ciphertext)

	plaintext := make([]byte, len(str))
	block.Decrypt(plaintext, ciphertext)
	fmt.Println(string(plaintext))
}
