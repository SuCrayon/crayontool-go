package timeutil

import "testing"

func Test_subscribe_time_jump(t *testing.T) {
	// Init()
	_ = TimeJumpSubject.Subscribe(NewTimeJumpSubjectReq().
		SetSubscribeName("default").
		SetSignal(SignalClockBack).
		SetTrigger(
			func(signal TimeJumpSignal) {
				t.Log("clock back")
			},
		),
	)
	TimeJumpSubject.Run()
}
