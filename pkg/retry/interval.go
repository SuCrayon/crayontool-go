package retry

import (
	"crayontool-go/pkg/datastructure/sliceutil"
	"crayontool-go/pkg/logger"
	"sync"
	"time"
)

type IntervalGenerator func(int) time.Duration

type IInterval interface {
	ConstantInterval(time.Duration) IntervalGenerator
	LinearInterval(time.Duration) IntervalGenerator
	FibonacciInterval(time.Duration) IntervalGenerator
}

type interval struct {
}

var (
	Interval = interval{}
)

var (
	fibonacciNumsCacheMutex = sync.RWMutex{}
	fibonacciNumsCache      = make([]int, 0)
)

func (i *interval) ConstantInterval(d time.Duration) IntervalGenerator {
	return constantInterval(d)
}

func (i *interval) LinearInterval(d time.Duration) IntervalGenerator {
	return linearInterval(d)
}

func (i *interval) FibonacciInterval(d time.Duration) IntervalGenerator {
	return fibonacciInterval(d)
}

// constantInterval 常量级别的时间间隔
func constantInterval(d time.Duration) IntervalGenerator {
	return func(times int) time.Duration {
		return d
	}
}

// linearInterval 线性增大的时间间隔
func linearInterval(d time.Duration) IntervalGenerator {
	return func(times int) time.Duration {
		return time.Duration(times) * d
	}
}

// fibonacciInterval 斐波那契规则的时间间隔
func fibonacciInterval(d time.Duration) IntervalGenerator {
	return func(times int) time.Duration {
		var ret time.Duration
		fibonacciNumsCacheMutex.RLock()
		logger.Debugf("fibonacciNumsCache: %v\n", fibonacciNumsCache)
		if times <= len(fibonacciNumsCache) {
			logger.Debugf("hit fibonacciNumsCache")
			ret = time.Duration(fibonacciNumsCache[times-1]) * d
			fibonacciNumsCacheMutex.RUnlock()
			return ret
		}
		fibonacciNumsCacheMutex.RUnlock()
		nums := sliceutil.GenFibonacciSlice(times)
		ret = time.Duration(nums[times-1]) * d
		fibonacciNumsCacheMutex.Lock()
		fibonacciNumsCache = append(fibonacciNumsCache, nums[len(fibonacciNumsCache):]...)
		fibonacciNumsCacheMutex.Unlock()
		return ret
	}
}
