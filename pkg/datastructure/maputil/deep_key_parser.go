package maputil

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/strutil"
	"strings"
	"sync"
)

const (
	escape       = constant.Escape
	escapeSymbol = constant.EscapeSymbol
	keySep       = constant.Dot
	keySepSymbol = constant.DotSymbol
)

type IDeepKeyParser interface {
	SetEscapeSymbol(uint8) IDeepKeyParser
	SetKeySepSymbol(uint8) IDeepKeyParser
	Parse(deepKey string) ([]string, error)
}

type deepKeyParser struct {
	mutex        sync.RWMutex
	escapeSymbol uint8
	keySepSymbol uint8
}

type deepKeyParserV2 struct {
	deepKeyParser
}

var (
	DeepKeyParser = deepKeyParser{
		escapeSymbol: escapeSymbol,
		keySepSymbol: keySepSymbol,
	}
)

func (d *deepKeyParser) SetEscapeSymbol(symbol uint8) IDeepKeyParser {
	d.escapeSymbol = symbol
	return d
}

func (d *deepKeyParser) SetKeySepSymbol(symbol uint8) IDeepKeyParser {
	d.keySepSymbol = symbol
	return d
}

func (d *deepKeyParser) Parse(deepKey string) ([]string, error) {
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
		index := strings.Index(curStr, strutil.Symbol2Str(d.keySepSymbol))
		if index == -1 {
			// 没有找到分隔符
			appendFunc(len(curStr))
			break
		}
		if index >= 1 && curStr[index-1] == d.escapeSymbol {
			// 前面一个字符是转义字符
			if index >= 2 && curStr[index-2] == d.escapeSymbol {
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

func (d *deepKeyParserV2) Parse(deepKey string) ([]string, error) {
	var (
		ret       = make([]string, 0)
		char      uint8
		leftIndex uint64
	)

	for i := 0; i < len(deepKey); i++ {
		char = deepKey[i]
		if char == d.keySepSymbol {
			/*// 找到分隔符
			  if i >= 1 && deepKey[i-1] == escapeSymbol {
			  	if i >= 2 && deepKey[i-2] == escapeSymbol {
			  		// 再前面一个是转义字符，则转义不生效，正常解析
			  		outerKey\\.innerKey

			  		append(ret, str[leftIndex:i])
			  		leftIndex = i + 1
			  	} else {
			  		// 前一个字符是转义字符，转义
			  		outerKey\.innerKey


			  	}
			  } else {
			  	// 没有遇到转义，正常解析
			  	outerKey.innerKey

			  	append(ret, str[leftIndex:i])
			  	leftIndex = i + 1
			  }*/
			if i >= 1 && deepKey[i-1] == d.escapeSymbol {
				if !(i >= 2 && deepKey[i-2] == d.escapeSymbol) {
					// 被转义了
					// 把转义用的转义字符去除
					// TODO: [OPT] 循环内修改字符串，待优化
					deepKey = deepKey[:i-1] + deepKey[i:]
					continue
				}
			}
			ret = append(ret, deepKey[leftIndex:i])
			leftIndex = uint64(i + 1)
		}
	}

	if leftIndex != uint64(len(deepKey)) {
		// 最后的状态应该是leftIndex等于deepKey索引+1，即等于len(deepKey)，不等则需要把leftIndex开始的字符串加入结果中
		ret = append(ret, deepKey[leftIndex:])
	}
	return ret, nil
}
