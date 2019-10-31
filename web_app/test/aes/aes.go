/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 17:04 2019-10-29
 */
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Encrypt string to base64 crypto using AES
func AESEncrypt(key []byte, text string) (string, bool) {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", false
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", false
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), true
}

// Decrypt from base64 to decrypted string
func AESDecrypt(key []byte, cryptoText string) (string, bool) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", false
	}

	if len(ciphertext) < aes.BlockSize {
		return "", false
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), true
}

func main() {
	var key = []byte("6aaf508205a7e5cca6b03ee5747a92f3")

	// 每次得到的结果都不同，但是都可以解密
	msg, ok := AESEncrypt(key, "abcd===a")
	if !ok {
		fmt.Println("encrypt is failed.")
	}
	fmt.Println("msg=", msg)

	text, ok := AESDecrypt(key, msg)
	if !ok {
		fmt.Println("decrypt is failed.")
	}

	fmt.Println("text=", text)
}
