package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

const (
	zapLogStackKey = "stack"
)

type zapWriter struct {
	once    sync.Once
	rwMutex sync.RWMutex
	logger  *zap.Logger
}

func NewZapWriter(core zapcore.Core, opts ...zap.Option) (LogWriter, error) {
	logger := zap.New(core, opts...)
	return &zapWriter{
		logger: logger,
	}, nil
}

func ToZapFields(fields ...LogField) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}

func (w *zapWriter) Debug(v interface{}, fields ...LogField) {
	w.logger.Debug(fmt.Sprint(v), ToZapFields(fields...)...)
}

func (w *zapWriter) Info(v interface{}, fields ...LogField) {
	w.logger.Info(fmt.Sprint(v), ToZapFields(fields...)...)
}

func (w *zapWriter) Warn(v interface{}, fields ...LogField) {
	w.logger.Warn(fmt.Sprint(v), ToZapFields(fields...)...)
}

func (w *zapWriter) Error(v interface{}, fields ...LogField) {
	w.logger.Error(fmt.Sprint(v), ToZapFields(fields...)...)
}

func (w *zapWriter) Fatal(v interface{}, fields ...LogField) {
	w.logger.Fatal(fmt.Sprint(v), ToZapFields(fields...)...)
}

func (w *zapWriter) Close() error {
	return w.logger.Sync()
}
