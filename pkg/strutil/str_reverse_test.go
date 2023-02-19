package strutil

import "testing"

func TestReverseWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "hello world",
			},
			want: "world hello",
		},
		{
			name: "",
			args: args{
				s: "123 hello world 321",
			},
			want: "321 world hello 123",
		},
		{
			name: "",
			args: args{
				s: " hello world",
			},
			want: "world hello ",
		},
		{
			name: "",
			args: args{
				s: "hello world ",
			},
			want: " world hello",
		},
		{
			name: "",
			args: args{
				s: " hello world ",
			},
			want: " world hello ",
		},
		{
			name: "",
			args: args{
				s: "  hello world ",
			},
			want: " world hello  ",
		},
		{
			name: "",
			args: args{
				s: "hello",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWord(tt.args.s); got != tt.want {
				t.Errorf("ReverseWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "1234 5 6789",
			},
			want: "9876 5 4321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseRune(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "湘北篮球部",
			},
			want: "部球篮北湘",
		},
		{
			name: "",
			args: args{
				s: "上海自来水来自海上",
			},
			want: "上海自来水来自海上",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseRune(tt.args.s); got != tt.want {
				t.Errorf("ReverseRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseRuneWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "湘北篮球部",
			},
			want: "湘北篮球部",
		},
		{
			name: "",
			args: args{
				s: "上海 自来水 来自 海上",
			},
			want: "海上 来自 自来水 上海",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseRuneWord(tt.args.s); got != tt.want {
				t.Errorf("ReverseRuneWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
