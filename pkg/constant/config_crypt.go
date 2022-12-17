package constant

import "crayontool-go/pkg/types"

const (
	ConfigCryptTag = "conf$crypt"
)

const (
	SecretTypeText types.SecretType = "text"
	SecretTypeEnv  types.SecretType = "env"
	SecretTypeFile types.SecretType = "file"
	SecretTypeURL  types.SecretType = "url"
)
