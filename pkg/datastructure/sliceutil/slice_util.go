package sliceutil

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/datastructure/set"
)

/*func OptionalParams2Set(vs ...interface{}) set.Set {
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
}*/

func Slice2Set(slice []interface{}) set.Set {
	_set := set.NewSet()
	for i := range slice {
		_set.Add(slice[i])
	}
	return _set
}

func Equals(slice1, slice2 []interface{}) bool {
	if len(slice1) != len(slice2) {
		return constant.False
	}
	for i := range slice1 {
		elem1 := slice1[i]
		elem2 := slice2[i]

		if elem1 != elem2 {
			return constant.False
		}
	}
	return constant.True
}
