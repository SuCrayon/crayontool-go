package strutil

import (
	"strings"
	"testing"
)

func TestJoinReq_Join(t *testing.T) {
	type fields struct {
		Elems     []string
		Sep       string
		Prefix    string
		Suffix    string
		OmitEmpty bool
		sb        *strings.Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				Elems:     nil,
				Sep:       "",
				Prefix:    "",
				Suffix:    "",
				OmitEmpty: false,
			},
			want: "",
		},
		{
			name: "",
			fields: fields{
				Elems:     []string{"1", "", "2", "3"},
				Sep:       "->",
				Prefix:    "",
				Suffix:    "",
				OmitEmpty: false,
			},
			want: "1->->2->3",
		},
		{
			name: "",
			fields: fields{
				Elems:     []string{"1", "", "2", "3"},
				Sep:       "->",
				Prefix:    "",
				Suffix:    "",
				OmitEmpty: false,
			},
			want: "1->->2->3",
		},
		{
			name: "",
			fields: fields{
				Elems:     []string{"1", "", "2", "3"},
				Sep:       "->",
				Prefix:    "",
				Suffix:    "",
				OmitEmpty: true,
			},
			want: "1->2->3",
		},
		{
			name: "",
			fields: fields{
				Elems:     []string{"1", "", "2", "3"},
				Sep:       "->",
				Prefix:    "head->",
				Suffix:    "->tail",
				OmitEmpty: true,
			},
			want: "head->1->2->3->tail",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &JoinReq{
				Elems:     tt.fields.Elems,
				Sep:       tt.fields.Sep,
				Prefix:    tt.fields.Prefix,
				Suffix:    tt.fields.Suffix,
				OmitEmpty: tt.fields.OmitEmpty,
				sb:        tt.fields.sb,
			}
			if got := r.Join(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
