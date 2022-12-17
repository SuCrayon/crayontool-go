package sliceutil

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/datastructure/set"
	"crayontool-go/pkg/reflectutil"
)

func OptionalParams2Set(vs ...interface{}) set.Set {
	var slice []interface{}
	for i := range vs {
		slice = append(slice, vs[i])
	}
	return Slice2Set(slice)
}

func AnySlice2Set(slice interface{}) (set.Set, bool) {
	if !reflectutil.TypeIsSlice(slice) {
		return nil, constant.False
	}
	elems, err := reflectutil.GetSliceElems(slice)
	if err != nil {
		return nil, constant.False
	}
	var tempSlice []interface{}
	for i := range elems {
		tempSlice = append(tempSlice, elems[i].Interface())
	}
	return Slice2Set(tempSlice), constant.True
}

func Slice2Set(slice []interface{}) set.Set {
	_set := set.NewSet()
	for i := range slice {
		_set.Add(slice[i])
	}
	return _set
}
