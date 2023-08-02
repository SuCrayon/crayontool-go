package reflectutil

import (
	"reflect"
	"sort"
	"testing"
)

type GetTagNamesTestStruct struct {
	Name  string `json:"Name"`
	Email string `bson:"email"`
}

func TestGetTagNames(t *testing.T) {
	type args struct {
		structT interface{}
		opts    []GetTagNameOption
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "case01",
			args: args{
				structT: (*GetTagNamesTestStruct)(nil),
				opts:    nil,
			},
			want: []string{
				"name",
				"email",
			},
		},
		{
			name: "case02",
			args: args{
				structT: GetTagNamesTestStruct{},
				opts:    nil,
			},
			want: []string{
				"name",
				"email",
			},
		},
		{
			name: "case03",
			args: args{
				structT: GetTagNamesTestStruct{
					Name: "",
				},
				opts: nil,
			},
			want: []string{
				"name",
				"email",
			},
		},
		{
			name: "case04",
			args: args{
				structT: (*GetTagNamesTestStruct)(nil),
				opts: []GetTagNameOption{
					WithTagKey("json"),
				},
			},
			want: []string{
				"Name",
				"email",
			},
		},
		{
			name: "case05",
			args: args{
				structT: (*GetTagNamesTestStruct)(nil),
				opts: []GetTagNameOption{
					WithTagKey("json"),
					WithDefaultStrategy(&FieldNameStrategy{}),
				},
			},
			want: []string{
				"Name",
				"Email",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTagNames(tt.args.structT, tt.args.opts...)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTagNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
