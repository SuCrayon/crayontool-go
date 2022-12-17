package logger

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/datastructure/set"
	"crayontool-go/pkg/strutil"
	"errors"
	"fmt"
	"log"
)

var (
	initLogWriterError = errors.New("init log writer err")
)

const (
	bannerTemplate = "%s" +
		"                                     __              __      __                           %s" +
		"  ______________ ___  ______  ____  / /_____  ____  / /     / /___  ____ _____ ____  _____%s" +
		" / ___/ ___/ __ `/ / / / __ \\/ __ \\/ __/ __ \\/ __ \\/ /_____/ / __ \\/ __ `/ __ `/ _ \\/ ___/%s" +
		"/ /__/ /  / /_/ / /_/ / /_/ / / / / /_/ /_/ / /_/ / /_____/ / /_/ / /_/ / /_/ /  __/ /    %s" +
		"\\___/_/   \\__,_/\\__, /\\____/_/ /_/\\__/\\____/\\____/_/     /_/\\____/\\__, /\\__, /\\___/_/     %s" +
		"               /____/                                            /____//____/             %s"
)

var (
	windowsReturnOSSet = set.NewSetWithCap(1).AddAll(constant.WindowsOSName)
	linuxReturnOSSet   = set.NewSetWithCap(1).AddAll(constant.LinuxOSName)
	macReturnOSSet     = set.NewSetWithCap(1).AddAll(constant.MacOSName)
)

type LogField struct {
	Key   string
	Value interface{}
}

type LogWriter interface {
	Alert(v interface{})
	Close() error
	Debug(v interface{}, fields ...LogField)
	Error(v interface{}, fields ...LogField)
	Info(v interface{}, fields ...LogField)
	Severe(v interface{})
	Slow(v interface{}, fields ...LogField)
	Stack(v interface{})
	Stat(v interface{}, fields ...LogField)
}

var (
	sugarLogWriter LogWriter
)

func getBanner() string {
	returnStr := strutil.GetLineSep()
	return strutil.SprintfRepeatedTimes(bannerTemplate, returnStr, 7)
}

func init() {
	// 默认初始化zap库
	logWriter, err := NewZapWriter()
	if err != nil {
		log.Fatalf("%s err: %+v\n", initLogWriterError.Error(), err)
	}
	sugarLogWriter = logWriter
	log.Print(getBanner())
}

func Alert(v interface{}) {
	sugarLogWriter.Alert(v)
}

func Close() error {
	return sugarLogWriter.Close()
}

func Debug(v interface{}) {
	sugarLogWriter.Debug(v)
}

func Debugf(format string, v ...interface{}) {
	sugarLogWriter.Debug(fmt.Sprintf(format, v...))
}

func Error(v interface{}) {
	sugarLogWriter.Error(v)
}

func Errorf(format string, v ...interface{}) {
	sugarLogWriter.Error(fmt.Sprintf(format, v...))
}

func Info(v interface{}) {
	sugarLogWriter.Info(v)
}

func Infof(format string, v ...interface{}) {
	sugarLogWriter.Info(fmt.Sprintf(format, v...))
}

func Severe(v interface{}) {
	sugarLogWriter.Severe(v)
}

func Slow(v interface{}) {
	sugarLogWriter.Slow(v)
}

func Stack(v interface{}) {
	sugarLogWriter.Stack(v)
}

func Stat(v interface{}) {
	sugarLogWriter.Stat(v)
}
