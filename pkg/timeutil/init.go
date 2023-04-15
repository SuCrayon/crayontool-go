package timeutil

import "sync"

var (
	initOnce sync.Once
)

type Config struct {
	TimeJumpSubjectConfig TimeJumpSubjectConfig `json:"timeJumpSubjectConfig" yaml:"timeJumpSubjectConfig"`
}

func initTimeJumpSubject(c *TimeJumpSubjectConfig) {
	// 初始化TimeJumpSubject
	TimeJumpSubject = newDefaultTimeJumpSubject(c)
	// 启动
	TimeJumpSubject.Start()
}

func Init(c *Config) {
	initOnce.Do(func() {
		initTimeJumpSubject(&c.TimeJumpSubjectConfig)
	})
}
