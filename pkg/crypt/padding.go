package crypt

import "bytes"

func NoPadding() {

}

func ZeroPadding() {

}

// PKCS7Padding PKCS7填充：缺几位补几个几
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	paddingData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, paddingData...)
}

func PKCS5Padding() {

}
