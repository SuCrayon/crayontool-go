package timeutil

import "testing"

func Test_subscribe_time_jump(t *testing.T) {
	// Init()
	_, _ = TimeJumpSubject.Subscribe(SignalClockBack, func(signal TimeJumpSignal) {
		t.Log("clock back")
	})
	TimeJumpSubject.Run()
}
