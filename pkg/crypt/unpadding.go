package crypt

import (
	"errors"
)

var (
	DataLengthIsZeroErr = errors.New("data length is zero")
)

func NoUnPadding() {

}

func ZeroUnPadding() {

}

// PKCS7UnPadding PKCS7填充：缺几位补几个几
func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, DataLengthIsZeroErr
	}
	// 获取填充的个数
	paddingLen := int(data[length-1])
	return data[:(length - paddingLen)], nil
}

func PKCS5UnPadding() {

}
