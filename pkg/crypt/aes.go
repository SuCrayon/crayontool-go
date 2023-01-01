package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

type encryptReq struct {
	baseEncryptReq
}

type encryptResp struct {
}

func (r *encryptReq) Encrypt() (EncryptResp, error) {
	block, err := aes.NewCipher(r.SecretKey)
	if err != nil {
		return nil, err
	}
	// 判断加密块大小
	blockSize := block.BlockSize()
	// 填充
	encryptBytes := PKCS7Padding(r.Plaintext, blockSize)
	ret := make([]byte, len(encryptBytes))
	encrypter := cipher.NewCBCEncrypter(block, r.SecretKey[:blockSize])
	encrypter.CryptBlocks(ret, encryptBytes)
	fmt.Println(base64.StdEncoding.EncodeToString(ret))
	return nil, nil
}

func Decrypt() {

}
