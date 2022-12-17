package timeutil

import (
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

// Now 获取现在时间
func Now() *Time {
	return &Time{
		time.Now(),
	}
}

// UnixSec 根据秒级时间戳获取Time
func UnixSec(sec int64) *Time {
	return &Time{
		time.Unix(sec, 0),
	}
}

// UnixMilli 根据毫秒级时间戳获取Time
func UnixMilli(milli int64) *Time {
	return &Time{
		time.Unix(milli/1000, 0),
	}
}

// UnixMilli 获取毫秒时间戳
func (t *Time) UnixMilli() int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// UnixSec 获取秒级时间戳
func (t *Time) UnixSec() int64 {
	return t.UnixNano() / int64(time.Second)
}

// BeginOfSecond 获取秒的开始
func (t *Time) BeginOfSecond() *Time {
	t.Time = time.Unix(0, t.UnixNano()-int64(t.Nanosecond()))
	return t
}

// BeginOfMinute 获取分钟的开始
func (t *Time) BeginOfMinute() *Time {
	t.BeginOfSecond()
	t.Time = time.Unix(t.Unix()-int64(t.Second()), 0)
	return t
}

// BeginOfHour 获取小时的开始
func (t *Time) BeginOfHour() *Time {
	t.BeginOfMinute()
	m := t.Minute()
	duration, _ := time.ParseDuration(fmt.Sprintf("-%dm", m))
	t.Time = t.Add(duration)
	return t
}

// BeginOfDay 获取天的开始
func (t *Time) BeginOfDay() *Time {
	t.BeginOfHour()
	h := t.Hour()
	duration, _ := time.ParseDuration(fmt.Sprintf("-%dh", h))
	t.Time = t.Add(duration)
	return t
}

// OffsetDays 偏移days天
func (t *Time) OffsetDays(days int) *Time {
	t.Time = t.AddDate(0, 0, days)
	return t
}

// OffsetMonths 偏移months月
func (t *Time) OffsetMonths(months int) *Time {
	t.Time = t.AddDate(0, months, 0)
	return t
}

// OffsetYears 偏移years年
func (t *Time) OffsetYears(years int) *Time {
	t.Time = t.AddDate(years, 0, 0)
	return t
}

// Tomorrow 明天
func (t *Time) Tomorrow() *Time {
	return t.OffsetDays(1)
}

// AfterTomorrow 后天
func (t *Time) AfterTomorrow() *Time {
	return t.OffsetDays(2)
}

// Yesterday 昨天
func (t *Time) Yesterday() *Time {
	return t.OffsetDays(-1)
}

// BeforeYesterday 前天
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

/*
DiffHumanDay 相差自然日
1. 秒数只差小于一天（绝对相差不到一天），那么只需要判断是否同一天即可
2. 对于绝对相差大于一天的情况，减掉绝对相差的秒数转到情况1即可
*/
func DiffHumanDay(t1, t2 *Time) int64 {
	ts1, ts2 := t1.UnixNano(), t2.UnixNano()
	if ts1 > ts2 {
		ts1, ts2 = ts2, ts1
		t1, t2 = t2, t1
	}
	absDay := (ts2 - ts1) / int64(24*time.Hour)
	tt := t2
	if absDay >= 1 {
		tt = &Time{
			time.Unix(0, ts2-absDay*int64(24*time.Hour)),
		}
	}
	return absDay + int64(tt.Day()-t1.Day())
}
