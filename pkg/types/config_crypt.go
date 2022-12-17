package types

// ConfigCryptConf 项目配置加解密配置
type ConfigCryptConf struct {
	Algorithm string
	Secret    Secret
	// CryptConfKeys []string
}

type SecretType string

type Secret struct {
	Type  SecretType
	Value string
}
