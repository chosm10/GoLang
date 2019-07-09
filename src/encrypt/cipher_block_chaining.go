package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encrypt(block cipher.Block, plaintext []byte) []byte {
	//암호화할 데이터의 길이를 블록 크기의 배수로 맞춰야함!
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		//부족한 만큼 채워줌
		padding := make([]byte, aes.BlockSize-mod)
		plaintext = append(plaintext, padding...)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	initVertor := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, initVertor); err != nil {
		fmt.Println(err)
		return nil
	}

	blockMode := cipher.NewCBCEncrypter(block, initVertor)
	blockMode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext
}

func decrypt(block cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("encrypted data's length must become block size multiple!")
		return nil
	}

	initVector := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	blockMode := cipher.NewCBCDecrypter(block, initVector)

	blockMode.CryptBlocks(plaintext, ciphertext)
	return plaintext
}

func main() {
	key := "Happy World 9876"
	str := `내가 가는 길이 험하고
	 멀지라도 그대와 함께라면 좋겠네.`

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := encrypt(block, []byte(str))
	fmt.Printf("%x\n", ciphertext)

	plaintext := decrypt(block, ciphertext)
	fmt.Println(string(plaintext))
}
