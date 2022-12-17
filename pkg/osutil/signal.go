package osutil

import (
	"fmt"
	"os"
	"os/signal"
)

// ListenSignalAsync 异步监听操作系统事件
func ListenSignalAsync(cb func(), sig ...os.Signal) {
	go ListenSignal(cb, sig...)
}

// ListenSignal 监听操作系统事件
func ListenSignal(cb func(), sig ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig...)
	select {
	case s := <-c:
		fmt.Printf("got signal: %v\n", s)
		cb()
		break
	}
}
