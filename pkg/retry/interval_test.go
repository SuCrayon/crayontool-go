package retry

import (
	"testing"
	"time"
)

func Test_fibonacciInterval(t *testing.T) {
	generator := fibonacciInterval(time.Second)
	generator(1)
	generator(1)
	generator(5)
	generator(3)
}
