package beanutil

import (
	"github.com/SuCrayon/crayontool-go/pkg/logger"
	"github.com/SuCrayon/crayontool-go/pkg/reflectutil"
	"reflect"
)

func GetFieldMap(typeOf reflect.Type) map[string]reflect.StructField {
	ret := make(map[string]reflect.StructField, typeOf.NumField())
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		ret[field.Name] = field
	}
	return ret
}

func CopyProperties(src, dest interface{}) {
	if !reflectutil.IsKindPtr(reflectutil.TypeOf(dest).Kind()) {
		logger.Warn("dest cannot be set, only kind of ptr is valid")
		return
	}
	srcTypeOf := reflectutil.OriginTypeOf(src)
	destTypeOf := reflectutil.OriginTypeOf(dest)
	srcValueOf := reflectutil.OriginValueOf(src)
	destValueOf := reflectutil.OriginValueOf(dest)

	srcFieldMap := GetFieldMap(srcTypeOf)
	for i := 0; i < destTypeOf.NumField(); i++ {
		destFieldType := destTypeOf.Field(i)
		srcFieldType, ok := srcFieldMap[destFieldType.Name]
		if !ok {
			// 源结构体中并不存在对应的字段
			continue
		}
		if destFieldType.Anonymous != srcFieldType.Anonymous {
			// 源结构体中对应的字段匿名程度不同
			continue
		}
		if destFieldType.Type.Kind() != srcFieldType.Type.Kind() {
			// 源结构体中对应的字段类型不同
			continue
		}

		destFieldValue := destValueOf.FieldByName(destFieldType.Name)
		srcFieldValue := srcValueOf.FieldByName(srcFieldType.Name)
		if !destFieldValue.CanSet() {
			logger.Warnf("dest field value cannot be set, destFieldName: %s\n", destFieldType.Name)
			continue
		}
		destFieldValue.Set(srcFieldValue)
	}
}
