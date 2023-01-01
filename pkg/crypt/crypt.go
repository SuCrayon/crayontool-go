package crypt

type EncryptReq interface {
	Encrypt() (EncryptResp, error)
}

type baseEncryptReq struct {
	SecretKey []byte
	Plaintext []byte
}

type EncryptResp interface {
}

type DecryptReq interface {
}

type DecryptResp interface {
}
