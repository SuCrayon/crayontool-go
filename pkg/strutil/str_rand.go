package strutil

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func RandomCharFromStr(str string) uint8 {
	ran, _ := rand.Int(rand.Reader, big.NewInt(int64(len(str))))
	return str[ran.Int64()]
}

func RandomStrFromStr(num int, str string) string {
	var result = ""
	charArr := strings.Split(str, "")
	for i := 0; i < num; i++ {
		ran, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charArr))))
		result += charArr[ran.Int64()]
	}
	return result
}
