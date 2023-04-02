package pkg

import (
	"github.com/SuCrayon/crayontool-go/pkg/timeutil"
	"sync"
)

var (
	once sync.Once
)

type Config struct {
	TimeUtilConfig timeutil.Config `json:"timeUtilConfig" yaml:"timeUtilConfig"`
}

func Init(c Config) {
	once.Do(func() {
		timeutil.Init(&c.TimeUtilConfig)
	})
}
