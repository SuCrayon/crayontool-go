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
	initLogWriterErr = errors.New("init log writer err")
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
	Debug(v interface{}, fields ...LogField)
	Info(v interface{}, fields ...LogField)
	Warn(v interface{}, fields ...LogField)
	Error(v interface{}, fields ...LogField)
	Fatal(v interface{}, fields ...LogField)
	Close() error
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
		log.Fatalf("%s err: %+v\n", initLogWriterErr.Error(), err)
	}
	sugarLogWriter = logWriter
	log.Print(getBanner())
}

func Debug(v interface{}, fields ...LogField) {
	sugarLogWriter.Debug(v, fields...)
}

func Debugf(format string, v ...interface{}) {
	sugarLogWriter.Debug(fmt.Sprintf(format, v...))
}

func Info(v interface{}, fields ...LogField) {
	sugarLogWriter.Info(v, fields...)
}

func Infof(format string, v ...interface{}) {
	sugarLogWriter.Info(fmt.Sprintf(format, v...))
}

func Warn(v interface{}, fields ...LogField) {
	sugarLogWriter.Warn(v, fields...)
}

func Warnf(format string, v ...interface{}) {
	sugarLogWriter.Warn(fmt.Sprintf(format, v...))
}

func Error(v interface{}, fields ...LogField) {
	sugarLogWriter.Error(v, fields...)
}

func Errorf(format string, v ...interface{}) {
	sugarLogWriter.Error(fmt.Sprintf(format, v...))
}

func Fatal(v interface{}, fields ...LogField) {
	sugarLogWriter.Fatal(v, fields...)
}

func Fatalf(format string, v ...interface{}) {
	sugarLogWriter.Fatal(fmt.Sprintf(format, v...))
}

func Close() error {
	return sugarLogWriter.Close()
}
