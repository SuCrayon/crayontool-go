package set

import "testing"

func Test_set_Add(t *testing.T) {
	type fields struct {
		bucket map[interface{}]struct{}
	}
	type args struct {
		i interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{name: "exist elem", fields: fields{bucket: map[interface{}]struct{}{"a": {}}}, args: args{i: "a"}, want: false},
		{name: "not exist elem", fields: fields{bucket: map[interface{}]struct{}{}}, args: args{i: "a"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &set{
				bucket: tt.fields.bucket,
			}
			if got := s.Add(tt.args.i); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
