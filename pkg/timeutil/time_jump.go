package timeutil

/**
检测时间跳变
    作用在于检测墙上时钟（客户端时间改变），并发出事件，
    提供一个发现的机制，让程序灵活应变这种情况，提高程序的健壮性
    eg: 有些定时任务依赖墙上时钟，所以当客户端修改了时间，会导致定时任务异常
*/

import (
	"github.com/SuCrayon/crayontool-go/pkg/constant"
	"github.com/SuCrayon/crayontool-go/pkg/logger"
	"sync"
	"time"
)

var (
	startOnce       sync.Once
	TimeJumpSubject ITimeJumpSubject
)

type TimeJumpSignal int8

const (
	SignalClockBack    TimeJumpSignal = -1
	SignalClockForward TimeJumpSignal = 1
)

const (
	defaultObserveInterval = 5
)

type TimeJumpTrigger = func(TimeJumpSignal)

type ITimeJumpSubject interface {
	Subscribe(*TimeJumpSubjectReq) error
	UnSubscribe(string) error
	Start()
	Run()
	doRun()
	notify(TimeJumpSignal)
	asyncNotify(TimeJumpSignal)
}

type TimeJumpSubjectConfig struct {
	Enable             bool `json:"enable" yaml:"enable"`
	ObserveInterval    int  `json:"observeInterval" yaml:"observeInterval"`
	ObserveImmediately bool `json:"observeImmediately" yaml:"observeImmediately"`
}

type baseTimeJumpSubject struct {
	config          TimeJumpSubjectConfig
	lastTime        time.Time
	triggerRegistry sync.Map
}

type defaultTimeJumpSubject struct {
	baseTimeJumpSubject
}

type TimeJumpSubjectReq struct {
	subscribeName string
	signal        TimeJumpSignal
	trigger       TimeJumpTrigger
}

func NewTimeJumpSubjectReq() *TimeJumpSubjectReq {
	return &TimeJumpSubjectReq{}
}

func (r *TimeJumpSubjectReq) SetSubscribeName(subscribeName string) *TimeJumpSubjectReq {
	r.subscribeName = subscribeName
	return r
}

func (r *TimeJumpSubjectReq) SetSignal(signal TimeJumpSignal) *TimeJumpSubjectReq {
	r.signal = signal
	return r
}

func (r *TimeJumpSubjectReq) SetTrigger(trigger TimeJumpTrigger) *TimeJumpSubjectReq {
	r.trigger = trigger
	return r
}

func (d *defaultTimeJumpSubject) Subscribe(req *TimeJumpSubjectReq) error {
	tTrigger := func(sSignal TimeJumpSignal) {
		if sSignal == req.signal {
			req.trigger(sSignal)
		}
	}
	d.triggerRegistry.Store(req.subscribeName, tTrigger)
	return nil
}

func (d *defaultTimeJumpSubject) UnSubscribe(subscriptionName string) error {
	d.triggerRegistry.Delete(subscriptionName)
	return nil
}

func (d *defaultTimeJumpSubject) notify(signal TimeJumpSignal) {
	d.triggerRegistry.Range(func(key, value interface{}) bool {
		trigger, ok := value.(TimeJumpTrigger)
		if ok {
			trigger(signal)
		} else {
			logger.Warnf("triggers range asset value to TimeJumpTrigger failed, value: %v\n", value)
		}
		return constant.True
	})
}

func (d *defaultTimeJumpSubject) asyncNotify(signal TimeJumpSignal) {
	go d.notify(signal)
}

func (d *defaultTimeJumpSubject) doRun() {
	currentTime := time.Now()
	logger.Debugf("defaultTimeJumpSubject get currentTime: %s\n", currentTime)
	defer func() {
		d.lastTime = currentTime
	}()
	if currentTime.Before(d.lastTime) {
		// 当前时间比上一次的时间要小了，说明时间改小了，时间跳变
		d.asyncNotify(SignalClockBack)
	} else if currentTime.After(d.lastTime) {
		d.asyncNotify(SignalClockForward)
	}
}

func (d *defaultTimeJumpSubject) Run() {
	if d.config.Enable == constant.False {
		// 没有开启
		logger.Info("time jump subject is not enabled")
		return
	}
	if d.config.ObserveImmediately {
		d.doRun()
	}
	duration := time.Duration(d.config.ObserveInterval) * time.Second
	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			d.doRun()
			timer.Reset(duration)
		}
	}
}

func (d *defaultTimeJumpSubject) Start() {
	startOnce.Do(func() {
		go d.Run()
	})
}

func getFinalConfig(c *TimeJumpSubjectConfig) *TimeJumpSubjectConfig {
	ret := TimeJumpSubjectConfig{
		ObserveInterval:    defaultObserveInterval,
		ObserveImmediately: constant.False,
	}
	// 默认配置
	if c == nil {
		return &ret
	}
	if c.ObserveInterval <= 0 {
		logger.Warn("c.ObserveInterval <= 0, will replace by default config")
		c.ObserveInterval = ret.ObserveInterval
	}
	return c
}

func newDefaultTimeJumpSubject(c *TimeJumpSubjectConfig) *defaultTimeJumpSubject {
	config := getFinalConfig(c)
	return &defaultTimeJumpSubject{
		baseTimeJumpSubject{
			config:   *config,
			lastTime: time.Now(),
		},
	}
}
