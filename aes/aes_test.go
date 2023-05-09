package utils

import (
    "testing"
    "fmt"
    "strings"
    "encoding/hex"
)

func TestAes(t *testing.T) {
	text := "hello"
	fmt.Printf("原文:%v\n",text)
	key := []byte{'a', 's', 'd', 'f', '+', '-', '*', '/', 'h', 'j', 'k', 'm', 5, 6, 7, 8}
	// 初始向量
	iv := []byte{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', 'g', 'b', 51, 53, 55, 59}

	// 加密
	result := CBCEncrypter([]byte(text), key, iv)
	ctext := strings.ToUpper(hex.EncodeToString(result)) 
	fmt.Printf("加密后:%v\n",ctext)

	// 解密
	b, _ := hex.DecodeString(ctext)
	fmt.Printf("解密后:%v\n",string(CBCDecrypter(b, key, iv)))
}
