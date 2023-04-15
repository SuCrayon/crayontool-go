package sliceutil

import (
	"github.com/SuCrayon/crayontool-go/pkg/reflectutil"
	"reflect"
)

/**
数组/切片去重
*/

func Distinct(v interface{}) interface{} {
	rawValueOf := reflectutil.ValueOf(v)
	valueOf := rawValueOf
	isPtr := reflectutil.IsKindPtr(rawValueOf.Kind())
	if isPtr {
		valueOf = reflectutil.OriginValueOf(v)
	}
	if !reflectutil.IsKindOriginSequence(valueOf.Kind()) {
		return v
	}
	m := make(map[interface{}]struct{}, valueOf.Len())
	ret := reflect.MakeSlice(valueOf.Type(), 0, valueOf.Len())
	for i := 0; i < valueOf.Len(); i++ {
		val := valueOf.Index(i).Interface()
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			ret = reflect.Append(ret, reflect.ValueOf(val))
		}
	}
	if isPtr {
		ptrValue := reflect.New(valueOf.Type())
		ptrValue.Elem().Set(ret)
		// 转为指针类型
		ret = ptrValue
	}
	return ret.Interface()
}
