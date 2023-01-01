package crypt

import (
	"testing"
)

func Test_encryptReq_Encrypt(t *testing.T) {
	req := encryptReq{baseEncryptReq{
		SecretKey: []byte("1234567890123456"),
		Plaintext: []byte("123456"),
	}}
	req.Encrypt()
}
