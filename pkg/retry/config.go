package retry

import (
	"crayontool-go/pkg/constant"
	"errors"
	"time"
)

const (
	defaultTimes        = 5
	defaultTimeout      = 30 * time.Second
	defaultRecoverPanic = constant.True
)

var (
	defaultIntervalGenerator = Interval.ConstantInterval(time.Second)
)

type Option func(*Config)

/*
1. 有重试次数和超时时间限制，重试指定次数，如果中途检测到超时了就退出
2. 只有重试次数限制，则重试指定次数即可
3. 只有超时时间限制，则一直重试直到超时为止
错误配置
1. 没有重试次数和超时时间限制，这样会导致死循环
*/

type Config struct {
	Times             int
	Timeout           time.Duration
	RecoverPanic      bool
	IntervalGenerator IntervalGenerator
}

var (
	DeadLoopRetryErr = errors.New("this configuration may cause dead loop retry")
)

func NewDefaultConfig() *Config {
	return &Config{
		Times:             defaultTimes,
		Timeout:           defaultTimeout,
		RecoverPanic:      defaultRecoverPanic,
		IntervalGenerator: defaultIntervalGenerator,
	}
}

func NewConfig(opts ...Option) *Config {
	config := NewDefaultConfig()
	for _, opt := range opts {
		opt(config)
	}
	return config
}

func WithTimes(times int) Option {
	return func(c *Config) {
		c.Times = times
	}
}

func WithNoTimes() Option {
	return func(c *Config) {
		c.Times = constant.IntNegativeOne
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

func WithNoTimeout() Option {
	return func(c *Config) {
		c.Timeout = constant.IntNegativeOne
	}
}

func WithRecoverPanic(recoverPanic bool) Option {
	return func(c *Config) {
		c.RecoverPanic = recoverPanic
	}
}

func WithIntervalGenerator(intervalGenerator IntervalGenerator) Option {
	return func(c *Config) {
		c.IntervalGenerator = intervalGenerator
	}
}

func (c *Config) Validate() error {
	if c.Times < 0 && c.Timeout < 0 {
		return DeadLoopRetryErr
	}

	return nil
}
