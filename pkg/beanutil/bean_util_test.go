package beanutil

import (
	"github.com/SuCrayon/crayontool-go/pkg/reflectutil"
	"testing"
)

func Test_typeOfField(t *testing.T) {
	type innerStruct struct {
		name string
	}
	type testStruct struct {
		name string
		Name string
		innerStruct
	}
	s := testStruct{}
	typeOf := reflectutil.OriginTypeOf(s)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		t.Log(field.Name)
		t.Log(field.IsExported())
		t.Log(field.Anonymous)
	}
}

func TestCopyProperties(t *testing.T) {
	type innerStruct struct {
		InnerName string
	}

	type InnerStruct struct {
		InnerName string
	}

	type srcStruct struct {
		name    string
		Name    string
		Integer int32
		innerStruct
		InnerStruct
	}

	type destStruct struct {
		name    string
		Name    string
		Integer int64
		innerStruct
		InnerStruct
	}

	src := srcStruct{
		name:    "srcStruct name",
		Name:    "srcStruct Name",
		Integer: 64,
		innerStruct: innerStruct{
			InnerName: "srcStruct innerStruct InnerName",
		},
		InnerStruct: InnerStruct{
			InnerName: "srcStruct InnerStruct InnerName",
		},
	}
	var dest destStruct
	CopyProperties(src, &dest)
	t.Log(dest)
}
