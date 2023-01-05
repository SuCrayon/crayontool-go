package maputil

import (
	"strings"
)

const (
	escape       = "\\"
	escapeSymbol = '\\'
	keySep       = "."
)

func ParseDeepKey(deepKey string) ([]string, error) {
	var ret = make([]string, 0)
	// 最大循环次数
	maxLoopCount := len(deepKey)
	sb := strings.Builder{}
	curStr := deepKey
	appendFunc := func(index int) {
		s := ""
		if sb.Len() > 0 {
			s += sb.String()
			sb.Reset()
		}
		ret = append(ret, s+curStr[:index])
		/*if index == len(curStr)-1 {
			// 最后一个字符
			return nil, deepKeySyntaxErr
		}*/
	}
	for i := 0; i < maxLoopCount; i++ {
		index := strings.Index(curStr, keySep)
		if index == -1 {
			// 没有找到分隔符
			appendFunc(len(curStr))
			break
		}
		if index >= 1 && curStr[index-1] == escapeSymbol {
			// 前面一个字符是转义字符
			if index >= 2 && curStr[index-2] == escapeSymbol {
				// 再前面一个字符是转义字符，则转义不生效，正常解析的同时去除
				// 丢弃分隔符
				appendFunc(index)
			} else {
				// 跳过转义，包含分隔符
				sb.WriteString(curStr[:index-1])
				sb.WriteString(curStr[index : index+1])
				if index == len(curStr)-1 {
					// 最后一个字符
					break
				}
			}
		} else {
			// 正常解析
			// 丢弃分隔符
			appendFunc(index)
		}
		curStr = curStr[index+1:]
	}

	if sb.Len() > 0 {
		ret = append(ret, sb.String())
	}
	return ret, nil
}
