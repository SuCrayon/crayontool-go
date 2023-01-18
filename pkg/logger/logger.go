package logger

import (
	"crayontool-go/pkg/strutil"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"sync"
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

const (
	currentInitFuncRegistryKey = "current"
	zapInitFuncRegistryKey     = "zap"
)

const (
	callerSkipOffset = 2
)

var (
	initLogWriterErr     = errors.New("init log writer err")
	initFuncNotExistsErr = errors.New("logger init func may not exists")
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

type InitFunc func() (LogWriter, error)

var (
	initOnce         sync.Once
	mutex            sync.RWMutex
	initFuncRegistry map[string]InitFunc
	sugarLogWriter   LogWriter
)

func init() {
	// 初始化函数注册表
	initFuncRegistry = map[string]InitFunc{
		currentInitFuncRegistryKey: defaultInitZapLogger,
		zapInitFuncRegistryKey:     defaultInitZapLogger,
	}
}

func execInitFunc() (LogWriter, error) {
	mutex.RLock()
	initFunc, ok := initFuncRegistry[currentInitFuncRegistryKey]
	mutex.RUnlock()
	if !ok {
		return nil, initFuncNotExistsErr
	}
	return initFunc()
}

func getBanner() string {
	returnStr := strutil.GetLineSep()
	return strutil.SprintfRepeatedTimes(bannerTemplate, returnStr, 7)
}

func defaultInitZapLogger() (LogWriter, error) {
	var opts []zap.Option
	opts = append(opts, zap.AddCallerSkip(callerSkipOffset))
	logger, err := zap.NewDevelopment(opts...)
	if err != nil {
		return nil, err
	}
	return &zapWriter{
		logger: logger,
	}, nil
}

func initLogWriter() {
	// 默认初始化zap库
	logWriter, err := execInitFunc()
	if err != nil {
		log.Fatalf("%s err: %+v\n", initLogWriterErr.Error(), err)
	}
	sugarLogWriter = logWriter
	log.Print(getBanner())
}

func InitLogWriter() {
	// 只执行一次
	initOnce.Do(initLogWriter)
}

func Debug(v interface{}, fields ...LogField) {
	InitLogWriter()
	sugarLogWriter.Debug(v, fields...)
}

func Debugf(format string, v ...interface{}) {
	InitLogWriter()
	sugarLogWriter.Debug(fmt.Sprintf(format, v...))
}

func Info(v interface{}, fields ...LogField) {
	InitLogWriter()
	sugarLogWriter.Info(v, fields...)
}

func Infof(format string, v ...interface{}) {
	InitLogWriter()
	sugarLogWriter.Info(fmt.Sprintf(format, v...))
}

func Warn(v interface{}, fields ...LogField) {
	InitLogWriter()
	sugarLogWriter.Warn(v, fields...)
}

func Warnf(format string, v ...interface{}) {
	InitLogWriter()
	sugarLogWriter.Warn(fmt.Sprintf(format, v...))
}

func Error(v interface{}, fields ...LogField) {
	InitLogWriter()
	sugarLogWriter.Error(v, fields...)
}

func Errorf(format string, v ...interface{}) {
	InitLogWriter()
	sugarLogWriter.Error(fmt.Sprintf(format, v...))
}

func Fatal(v interface{}, fields ...LogField) {
	InitLogWriter()
	sugarLogWriter.Fatal(v, fields...)
}

func Fatalf(format string, v ...interface{}) {
	InitLogWriter()
	sugarLogWriter.Fatal(fmt.Sprintf(format, v...))
}

func Close() error {
	InitLogWriter()
	return sugarLogWriter.Close()
}
