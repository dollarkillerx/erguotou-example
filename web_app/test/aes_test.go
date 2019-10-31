/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:57 2019-10-29
 */
package test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/dollarkillerx/erguotou/util"
	"io"
	"log"
	"testing"
)

// 对称加密 AES 高级标准加密

// Encrypt string to base64 crypto using AES
func AESEncrypt(key []byte, plaintext []byte) ([]byte, bool) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, false
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, false
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, true
}

// Decrypt from base64 to decrypted string
func AESDecrypt(key []byte, ciphertext []byte) ([]byte, bool) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, false
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, false
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, true
}

func TestAes(t *testing.T) {
	key := []byte("ff6f2ee1d83b3f8b235da9355aee44ac")
	yao := "你好啊"
	s, b := AESEncrypt(key, []byte(yao))
	if !b {
		panic("error")
	}
	bytes, i := AESDecrypt(key, s)
	if !i {
		panic("ss")
	}

	log.Println(string(bytes))
}
func TestNu(t *testing.T) {
	da := "6aaf508205a7e5cca6b03ee5747a92f3"
	fmt.Println(len(da))

	superRand := util.SuperRand()
	log.Println(len(superRand[:32]))
	log.Println(superRand[:32])
}

// 对称加密 AES 高级标准加密
func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func AESEncode(key []byte, src []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func AESDecode(key []byte, src []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}

func TestNu2(t *testing.T) {
	key := util.SuperRand()[:32]
	encode, e := AESEncrypt([]byte(key), []byte("this is data"))
	if !e {
		panic(e)
	}
	decode, e := AESDecrypt([]byte(key), encode)
	log.Println(string(decode))
}
