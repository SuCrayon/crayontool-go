package password

import (
	"github.com/SuCrayon/crayontool-go/pkg/datastructure/sliceutil"
	"github.com/SuCrayon/crayontool-go/pkg/strutil"
	"strings"
)

const (
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	UppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers          = "0123456789"
)

type Generator interface {
	OriginStrings() []string
	GenStrongPassword(length int) string
	// TODO: GenPassword(length int) string
}

type sugar struct {
}

var (
	Sugar = sugar{}
)

func (s *sugar) OriginStrings() []string {
	return []string{
		LowercaseLetters,
		UppercaseLetters,
		Numbers,
	}
}

func (s *sugar) GenStrongPassword(length int) string {
	var (
		originStringsLen = len(s.OriginStrings())
		counts           = sliceutil.GenSliceWithGenFunc(originStringsLen, func(arr []int, i int) int {
			return 0
		})
	)

	if length%originStringsLen == 0 {
		// 整除
		sliceutil.IntSliceFill(counts, length/originStringsLen)
	} else {
		remainder := length % originStringsLen
		count := (length - remainder) / originStringsLen
		sliceutil.IntSliceFill(counts[1:], count)
		counts[0] = remainder + count
	}

	sb := strings.Builder{}
	for i := 0; i < originStringsLen; i++ {
		sb.WriteString(strutil.RandomStrFromStr(counts[i], s.OriginStrings()[i]))
	}
	tmp := sb.String()
	indexArr := sliceutil.GenIndexShuffleOrderSlice(sb.Len())

	sb.Reset()
	for i := range indexArr {
		sb.WriteByte(tmp[indexArr[i]])
	}

	return sb.String()
}

func (s *sugar) GenPassword(length int) string {
	return ""
}
