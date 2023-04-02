package timeutil

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
	initTimeJumpSubject(&c.TimeJumpSubjectConfig)
}
