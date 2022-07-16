package timeutil

import (
	"testing"
	"time"
)

func TestUnixMilli(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "", args: args{t: time.Now()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("UnixMilli: %d\n", UnixMilli(tt.args.t))
		})
	}
}
