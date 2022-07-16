package timeutil

import (
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

// UnixMilli 毫秒时间戳
func UnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func Now() *Time {
	return &Time{
		Time: time.Now(),
	}
}

func (t *Time) BeginOfSecond() *Time {
	t.Time = time.Unix(0, t.UnixNano()-int64(t.Nanosecond()))
	return t
}

func (t *Time) BeginOfMinute() *Time {
	t.BeginOfSecond()
	t.Time = time.Unix(t.Unix()-int64(t.Second()), 0)
	return t
}

func (t *Time) BeginOfHour() *Time {
	t.BeginOfMinute()
	m := t.Minute()
	duration, _ := time.ParseDuration(fmt.Sprintf("-%dm", m))
	t.Time = t.Add(duration)
	return t
}

func (t *Time) BeginOfDay() *Time {
	t.BeginOfHour()
	h := t.Hour()
	duration, _ := time.ParseDuration(fmt.Sprintf("-%dh", h))
	t.Time = t.Add(duration)
	return t
}

func (t *Time) OffsetDays(days int) *Time {
	t.Time = t.AddDate(0, 0, days)
	return t
}

func (t *Time) OffsetMonths(months int) *Time {
	t.Time = t.AddDate(0, 1, 0)
	return t
}

func (t *Time) OffsetYears(years int) *Time {
	t.Time = t.AddDate(1, 0, 0)
	return t
}

func (t *Time) Tomorrow() *Time {
	return t.OffsetDays(1)
}

func (t *Time) AfterTomorrow() *Time {
	return t.OffsetDays(2)
}

func (t *Time) Yesterday() *Time {
	return t.OffsetDays(-1)
}

func (t *Time) BeforeYesterday() *Time {
	return t.OffsetDays(-2)
}

// DiffDay 相差多少天
func DiffDay(t1, t2 *Time) int64 {
	ts1, ts2 := t1.UnixNano(), t2.UnixNano()
	if ts1 > ts2 {
		ts1, ts2 = ts2, ts1
	}
	return (ts2 - ts1) / int64(24*time.Hour)
}

// DiffHumanDay 相差自然日
func DiffHumanDay(t1, t2 *Time) int64 {
	return 0
}
