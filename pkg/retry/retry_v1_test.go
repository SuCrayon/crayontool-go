package retry

import (
	"errors"
	"testing"
	"time"
)

func Test_v1_Do_0(t *testing.T) {
	err := V1.Do(func() error {
		return errors.New("meet error, need retry")
	}, WithIntervalGenerator(Interval.ConstantInterval(time.Second)))
	if err.HasError() {
		t.Log(err)
	}
}

func Test_v1_Do_1(t *testing.T) {
	err := V1.Do(func() error {
		return errors.New("meet error, need retry")
	}, WithIntervalGenerator(Interval.LinearInterval(time.Second)), WithTimes(1), WithNoTimeout())
	if err.HasError() {
		t.Log(err)
	}
}

func Test_v1_Do_2(t *testing.T) {
	err := V1.Do(func() error {
		return errors.New("meet error, need retry")
	}, WithIntervalGenerator(Interval.LinearInterval(time.Second)), WithNoTimes())
	if err.HasError() {
		t.Log(err)
	}
}

func Test_v1_Do_3(t *testing.T) {
	err := V1.Do(func() error {
		return errors.New("meet error, need retry")
	}, WithIntervalGenerator(Interval.LinearInterval(time.Second)), WithNoTimeout(), WithNoTimes())
	if err.HasError() {
		t.Log(err)
	}
}
