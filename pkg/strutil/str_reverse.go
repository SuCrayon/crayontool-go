package strutil

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/datastructure/sliceutil"
)

// Reverse 字符串反转
/** eg: hello world -> dlrow olleh
 */
func Reverse(s string) string {
	bs := []byte(s)
	sliceutil.ByteSliceReverse(bs)
	return string(bs)
}

// ReverseWord 字符串单词间反转
/** eg: hello world -> world hello
 */
func ReverseWord(s string) string {
	bs := []byte(s)
	// 整体反转
	sliceutil.ByteSliceReverse(bs)
	leftIndex := 0
	for i := 0; i < len(bs); i++ {
		// 单词内再次反转
		if bs[i] == constant.SpaceSymbol {
			// meet space, word separator
			sliceutil.ByteSliceReverse(bs[leftIndex:i])
			leftIndex = i + 1
		}
	}
	sliceutil.ByteSliceReverse(bs[leftIndex:])
	return string(bs)
}

func ReverseRune(s string) string {
	rs := []rune(s)
	sliceutil.RuneSliceReverse(rs)
	return string(rs)
}

func ReverseRuneWord(s string) string {
	rs := []rune(s)
	// 整体反转
	sliceutil.RuneSliceReverse(rs)
	leftIndex := 0
	for i := 0; i < len(rs); i++ {
		// 单词内再次反转
		if rs[i] == constant.SpaceSymbol {
			// meet space, word separator
			sliceutil.RuneSliceReverse(rs[leftIndex:i])
			leftIndex = i + 1
		}
	}
	sliceutil.RuneSliceReverse(rs[leftIndex:])
	return string(rs)
}
