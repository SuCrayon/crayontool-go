package types

import "reflect"

type FieldValue struct {
	Name  string
	Field reflect.StructField
	Value reflect.Value
}

type FieldTreeNode struct {
	FieldValue
	Children []FieldTreeNode
}
