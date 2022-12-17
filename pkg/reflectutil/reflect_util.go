package reflectutil

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/datastructure/set"
	"crayontool-go/pkg/logger"
	"errors"
	"reflect"
)

var (
	TypeIsNotStructErr    = errors.New("input type is not a struct")
	TypeIsNotSliceErr     = errors.New("input type is not a slice")
	TypeCannotCallElemErr = errors.New("input type is not a kind which can call Elem()")
)

var (
	// only kind: Array, Chan, Map, Ptr, Slice can call Elem()
	canElemSet = set.NewSetWithCap(5).AddAll(
		reflect.Array,
		reflect.Chan,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
	)
)

/*func StructOriginalTypeOfFields(t interface{}) ([]reflect.StructField, error) {
	if !TypeIsStruct(t) {
		return nil, TypeIsNotStructErr
	}
	return StructOriginalTypeOfFieldByMatchFunc(
		func(field reflect.StructField) bool { return constant.True },
		t,
	)
}*/

func GetStructRealFieldByMatchFunc() {

}

func GetStructFieldByMatchFunc(match func(field reflect.StructField) bool, t interface{}) ([]reflect.StructField, error) {
	var ret []reflect.StructField
	if !RealTypeIsStruct(t) {
		return nil, TypeIsNotStructErr
	}
	_type := RealTypeOf(t)
	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i)
		if match(field) {
			ret = append(ret, field)
		}
	}
	return ret, nil
}

/*func StructOriginalTypeOf(t interface{}) (reflect.Type, error) {
	if !TypeIsStruct(t) {
		return nil, TypeIsNotStructErr
	}
	_type := reflect.TypeOf(t)
	if _type.Kind() == reflect.Ptr {
		_type = _type.Elem()
	}
	return _type, nil
}*/

/*func StructOriginalFieldValue(tv interface{}) ([]types.FieldValue, error) {
	if !TypeIsStruct(tv) {
		return nil, TypeIsNotStructErr
	}
	var ret []types.FieldValue
	_type, err := StructOriginalTypeOf(tv)
	if err != nil {
		return nil, err
	}
	value, err := StructOriginalValueOf(tv)
	if err != nil {
		return nil, err
	}
	for i := 0; i < _type.NumField(); i++ {
		elem := types.FieldValue{
			Name:  _type.Field(i).Name,
			Field: _type.Field(i),
			Value: value.Field(i),
		}

		ret = append(ret, elem)
	}
	return ret, nil
}

func recursiveParseFieldValueTree(tv interface{}, root types.FieldValueNode) error {
	if !TypeIsStruct(tv) &&
		!TypeIsSlice(tv) &&
		!TypeIsArray(tv) {
		// !TypeIsMap(tv)
		// 只有结构体、map、数组、切片才有继续解析的必要
		return nil
	}
	fieldValues, err := StructOriginalFieldValue(tv)
	if err != nil {
		return err
	}
	for i := range fieldValues {
		cur := types.FieldValueNode{
			Value: fieldValues[i],
		}
		// 加入子节点列表
		root.Children = append(root.Children, cur)
		// 循环递归
		recursiveParseFieldValueTree(fieldValues[i].Value.Interface(), cur)
	}
	return nil
}

func StructOriginalFieldValueTree(tv interface{}) (types.FieldValueNode, error) {
	var root types.FieldValueNode
	err := recursiveParseFieldValueTree(tv, root)
	return root, err
}*/

/*func StructOriginalValueOf(v interface{}) (*reflect.Value, error) {
	if !TypeIsStruct(v) {
		return nil, TypeIsNotStructErr
	}
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return &value, nil
}*/

func TypeIs(t interface{}, k reflect.Kind) bool {
	return reflect.TypeOf(t).Kind() == k
}

func ElemTypeIs(t interface{}, k reflect.Kind) (bool, error) {
	_type := reflect.TypeOf(t)
	if !canElemSet.Contains(_type.Kind()) {
		logger.Debugf("%s, inputType's kind: %v", TypeCannotCallElemErr.Error(), _type.Kind().String())
		return constant.False, TypeCannotCallElemErr
	}
	return _type.Elem().Kind() == k, nil
}

func TypeIsSlice(t interface{}) bool {
	return TypeIs(t, reflect.Slice)
}

func TypeIsString(t interface{}) bool {
	return TypeIs(t, reflect.String)
}

func TypeIsStruct(t interface{}) bool {
	return TypeIs(t, reflect.Struct)
}

func TypeIsMap(t interface{}) bool {
	return TypeIs(t, reflect.Map)
}

func TypeIsArray(t interface{}) bool {
	return TypeIs(t, reflect.Array)
}

func TypeIsPtr(t interface{}) bool {
	return TypeIs(t, reflect.Ptr)
}

func KindIs(kind1 reflect.Kind, kind2 reflect.Kind) bool {
	return kind1 == kind2
}

func KindIsString(kind reflect.Kind) bool {
	return kind == reflect.String
}

func KindIsStruct(kind reflect.Kind) bool {
	return kind == reflect.Struct
}

func KindIsCanElem(kind reflect.Kind) bool {
	return canElemSet.Contains(kind)
}

func RealTypeIs(t interface{}, k reflect.Kind) bool {
	if TypeIs(t, k) {
		return constant.True
	}
	ok, err := ElemTypeIs(t, k)
	if err != nil {
		return constant.False
	}
	return ok
}

func RealTypeIsStruct(t interface{}) bool {
	return RealTypeIs(t, reflect.Struct)
}

func RealTypeIsMap(t interface{}) bool {
	return RealTypeIs(t, reflect.Map)
}

func RealTypeIsArray(t interface{}) bool {
	return RealTypeIs(t, reflect.Array)
}

func RealTypeIsSlice(t interface{}) bool {
	return RealTypeIs(t, reflect.Slice)
}

func RealTypeOf(t interface{}) reflect.Type {
	_type := reflect.TypeOf(t)
	if canElemSet.Contains(_type.Kind()) {
		return _type.Elem()
	}
	return _type
}

func RealValueOf(v interface{}) reflect.Value {
	value := reflect.ValueOf(v)
	if canElemSet.Contains(value.Kind()) {
		return value.Elem()
	}
	return value
}

func doGetSliceElems(v interface{}, judge func() bool, valueOf func(interface{}) reflect.Value) ([]reflect.Value, error) {
	if !judge() {
		return nil, TypeIsNotSliceErr
	}
	var elems []reflect.Value
	value := valueOf(v)
	for i := 0; i < value.Len(); i++ {
		elems = append(elems, value.Index(i))
	}
	return elems, nil
}

func GetSliceElems(v interface{}) ([]reflect.Value, error) {
	return doGetSliceElems(
		v,
		func() bool {
			return TypeIsSlice(v)
		},
		reflect.ValueOf,
	)
}

func GetRealSliceElems(v interface{}) ([]reflect.Value, error) {
	return doGetSliceElems(
		v,
		func() bool {
			return RealTypeIsSlice(v)
		},
		RealValueOf,
	)
}
