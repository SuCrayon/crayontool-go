package retry

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/logger"
	"time"
)

type v2 struct {
	v1
}

type FunctionV2 func() error

var (
	V2 = v2{}
)

func (v *v2) Do(function FunctionV2, opts ...Option) Errors {
	config := NewConfig(opts...)
	if err := config.Validate(); err != nil {
		return NewError(err)
	}

	var (
		// timer 超时时钟
		timer = new(time.Timer)
		// clock 间隔时钟，一开始不需要间隔等待所以设置为0
		clock = time.NewTimer(constant.IntZero)
		// errs 错误列表
		errs Errors
	)

	v.config = config
	if config.Timeout > 0 {
		timer = time.NewTimer(config.Timeout)
	}

LOOP:
	for {
		select {
		// 超时中断
		case <-timer.C:
			logger.Warn("arrive timeout line, no retry")
			break LOOP
		// 间隔时间到
		case <-clock.C:
			err := funcExecWrapper(function, config.RecoverPanic)
			if err == nil {
				return nil
			}
			errs = append(errs, err)
			logger.Warnf("meet error, err: %+v\n", err)
			v.increaseRetryNum()
			if v.isArriveTimesThreshold() {
				// 重试次数够了
				logger.Warn("arrive times threshold, no retry")
				break LOOP
			}
			// 时钟重置
			d := config.IntervalGenerator(v.retryNum)
			clock.Reset(d)
			logger.Debugf("sleep before next retry, interval: %v, retryNum: %d\n", d, v.retryNum)
		}
	}

	return errs
}
