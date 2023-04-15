package reflectutil

import "reflect"

func IsKindPtr(kind reflect.Kind) bool {
	return kind == reflect.Ptr
}

func IsKindMap(kind reflect.Kind) bool {
	return kind == reflect.Map
}

func IsKindArray(kind reflect.Kind) bool {
	return kind == reflect.Array
}

func IsKindSlice(kind reflect.Kind) bool {
	return kind == reflect.Slice
}

func IsKindOriginSequence(kind reflect.Kind) bool {
	return IsKindArray(kind) || IsKindSlice(kind)
}

func IsValueKindMap(v interface{}) bool {
	valueOf := reflect.ValueOf(v)
	return IsKindMap(valueOf.Kind())
}

func OriginKindOfValue(v interface{}) reflect.Kind {
	valueOf := reflect.ValueOf(v)
	if !IsKindPtr(valueOf.Kind()) {
		return valueOf.Kind()
	}
	return valueOf.Elem().Kind()
}

func KindOfValue(v interface{}) reflect.Kind {
	valueOf := reflect.ValueOf(v)
	return valueOf.Kind()
}

func ValueOf(v interface{}) reflect.Value {
	return reflect.ValueOf(v)
}

func OriginValueOf(v interface{}) reflect.Value {
	valueOf := ValueOf(v)
	if IsKindPtr(valueOf.Kind()) {
		return valueOf.Elem()
	}
	return valueOf
}

func TypeOf(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

func OriginTypeOf(v interface{}) reflect.Type {
	typeOf := TypeOf(v)
	if IsKindPtr(typeOf.Kind()) {
		return typeOf.Elem()
	}
	return typeOf
}
