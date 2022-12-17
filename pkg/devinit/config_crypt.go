package devinit

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/datastructure/set"
	"crayontool-go/pkg/logger"
	"crayontool-go/pkg/reflectutil"
	"crayontool-go/pkg/types"
	"os"
	"reflect"
)

type ConfigCryptProcessor interface {
	Crypt() error
	Decrypt() error
}

type configCryptProcessor struct {
	conf            types.ConfigCryptConf
	target          interface{}
	secret          string
	cryptConfKeySet set.Set
}

func NewConfigCryptProcessor(conf types.ConfigCryptConf, target interface{}) ConfigCryptProcessor {
	p := &configCryptProcessor{
		conf:   conf,
		target: target,
	}
	return p
}

/*func (p *configCryptProcessor) parseCryptConfKeys() bool {
	_set, ok := sliceutil.AnySlice2Set(p.conf.CryptConfKeys)
	if !ok {
		return constant.False
	}
	p.cryptConfKeySet = _set
	return constant.True
}*/

func (p *configCryptProcessor) parseSecret() bool {
	secret := ""
	defer func() {
		p.secret = secret
	}()
	switch p.conf.Secret.Type {
	case constant.SecretTypeEnv:
		secret = os.Getenv(p.conf.Secret.Value)
	case constant.SecretTypeText:
		secret = p.conf.Secret.Value
	case constant.SecretTypeFile:

	case constant.SecretTypeURL:
	default:

	}
	return constant.True
}

func (p *configCryptProcessor) ParseCryptConf() bool {
	if !p.parseSecret() {
		return constant.False
	}

	/*if !p.parseCryptConfKeys() {
		return constant.False
	}*/

	return constant.True
}

func (p *configCryptProcessor) DoCrypt(s string) string {
	return ""
}

func (p *configCryptProcessor) DoDecrypt(s string) string {
	return ""
}

func (p *configCryptProcessor) Crypt() error {
	return nil
}

func (p *configCryptProcessor) Decrypt() error {
	tFields, err := reflectutil.GetStructFieldByMatchFunc(
		func(field reflect.StructField) bool {
			_, ok := field.Tag.Lookup(constant.ConfigCryptTag)
			return ok
		},
		p.target,
	)
	if err != nil {
		return err
	}
	for i := range tFields {
		value := reflectutil.RealValueOf(p.target)
		fieldValue := value.FieldByName(tFields[i].Name)
		if reflectutil.KindIsCanElem(fieldValue.Kind()) {

		}
		if !reflectutil.KindIsString(fieldValue.Kind()) {
			logger.Debugf("kind is not a string, skip decrypt! kind: %v", fieldValue.Kind().String())
			continue
		}
		fieldValue.SetString(p.DoDecrypt(fieldValue.String()))
	}
	return nil
}
