package strutil

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/datastructure/set"
	"runtime"
)

var (
	windowsReturnOSSet = set.NewSetWithCap(1).AddAll(constant.WindowsOSName)
	linuxReturnOSSet   = set.NewSetWithCap(1).AddAll(constant.LinuxOSName)
	macReturnOSSet     = set.NewSetWithCap(1).AddAll(constant.MacOSName)
)

func GetLineSep() string {
	// 换行风格默认为win
	returnStr := constant.WindowsLineSep
	osName := runtime.GOOS
	switch {
	case windowsReturnOSSet.Contains(osName):
		{
			returnStr = constant.WindowsLineSep
		}
	case linuxReturnOSSet.Contains(osName):
		{
			returnStr = constant.UnixLineSep
		}
	case macReturnOSSet.Contains(osName):
		{
			returnStr = constant.MacOSLineSep
		}
	}
	return returnStr
}

func Symbol2Str(symbol uint8) string {
	return string(symbol)
}
