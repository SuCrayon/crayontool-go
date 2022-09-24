package timeutil

import (
	"testing"
	"time"
)

func TestDiffHumanDay(t *testing.T) {
	type args struct {
		t1 *Time
		t2 *Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: struct {
				t1 *Time
				t2 *Time
			}{
				t1: &Time{
					// 2022-09-24 16:00:00
					time.Unix(1664006400, 0),
				},
				t2: &Time{
					// 2022-09-25 01:00:00
					time.Unix(1664038800, 0),
				},
			},
			want: 1,
		},
		{
			name: "",
			args: struct {
				t1 *Time
				t2 *Time
			}{
				t1: &Time{
					// 2022-09-24 23:00:00
					time.Unix(1664031600, 0),
				},
				t2: &Time{
					// 2022-09-25 01:00:00
					time.Unix(1664038800, 0),
				},
			},
			want: 1,
		},
		{
			name: "",
			args: struct {
				t1 *Time
				t2 *Time
			}{
				t1: &Time{
					// 2022-09-24 12:00:00
					time.Unix(1663992000, 0),
				},
				t2: &Time{
					// 2022-09-24 23:00:00
					time.Unix(1664031600, 0),
				},
			},
			want: 0,
		},
		{
			name: "",
			args: struct {
				t1 *Time
				t2 *Time
			}{
				t1: &Time{
					// 2022-09-25 23:30:00
					time.Unix(1664119800, 0),
				},
				t2: &Time{
					// 2022-09-24 23:00:00
					time.Unix(1664031600, 0),
				},
			},
			want: 1,
		},
		{
			name: "",
			args: struct {
				t1 *Time
				t2 *Time
			}{
				t1: &Time{
					// 2022-09-26 01:00:00
					time.Unix(1664125200, 0),
				},
				t2: &Time{
					// 2022-09-24 23:00:00
					time.Unix(1664031600, 0),
				},
			},
			want: 2,
		},
		{
			name: "",
			args: struct {
				t1 *Time
				t2 *Time
			}{
				t1: &Time{
					// 2022-09-28 23:00:00
					time.Unix(1664377200, 0),
				},
				t2: &Time{
					// 2022-09-24 01:00:00
					time.Unix(1664031600, 0),
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffHumanDay(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("DiffHumanDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffDay(t *testing.T) {
	type args struct {
		t1 *Time
		t2 *Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				t1: &Time{
					// 2022-09-26 01:00:00
					time.Unix(1664125200, 0),
				},
				t2: &Time{
					// 2022-09-25 01:00:00
					time.Unix(1664038800, 0),
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				t1: &Time{
					// 2022-09-26 01:00:00
					time.Unix(1664125200, 0),
				},
				t2: &Time{
					// 2022-09-24 01:00:00
					time.Unix(1663952400, 0),
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				t1: &Time{
					// 2022-09-26 01:00:00
					time.Unix(1664125200, 0),
				},
				t2: &Time{
					// 2022-09-25 23:00:00
					time.Unix(1664118000, 0),
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffDay(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("DiffDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
