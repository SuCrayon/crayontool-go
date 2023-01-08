package retry

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/logger"
	"time"
)

type v1 struct {
	config *Config
	// startTime 开始时间
	startTime time.Time
	// TimeoutLine 超时时间线
	timeoutLine time.Time
	// retryNum 重试次数
	retryNum int
}

type FunctionV1 func() error

var (
	V1 = v1{}
)

func (v *v1) isTimeout() bool {
	if v.config.Timeout < constant.IntZero {
		// 小于0说明没有超时限制
		return constant.False
	}
	// 超时时间线在当前时间之前了，说明超时了
	return v.timeoutLine.Before(time.Now())
}

func (v *v1) isArriveTimesThreshold() bool {
	if v.config.Times < constant.IntZero {
		// 小于0说明没有重试次数的限制
		return constant.False
	}
	return v.retryNum > v.config.Times
}

func (v *v1) increaseRetryNum() int {
	v.retryNum += 1
	return v.retryNum
}

func (v *v1) Do(function FunctionV1, opts ...Option) Error {
	config := NewConfig(opts...)
	if err := config.Validate(); err != nil {
		return NewError(err)
	}
	var errs Error
	v.config = config
	v.startTime = time.Now()
	v.timeoutLine = time.Now().Add(config.Timeout)
	for {
		err := function()
		if err == nil {
			return nil
		}
		errs = append(errs, err)
		logger.Warnf("meet error, err: %+v\n", err)

		v.increaseRetryNum()
		if v.isArriveTimesThreshold() {
			// 重试次数够了
			logger.Warn("arrive times threshold, no retry")
			break
		}
		if v.isTimeout() {
			// 超时时间线比当前时间靠前，说明已经超时了，直接退出
			logger.Warn("arrive timeout line, no retry")
			break
		}
		d := config.IntervalGenerator(v.retryNum)
		logger.Infof("sleep before next retry, interval: %v, retryNum: %d\n", d, v.retryNum)
		time.Sleep(d)
	}

	return errs
}
