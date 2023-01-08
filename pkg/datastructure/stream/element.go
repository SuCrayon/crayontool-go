package stream

import "reflect"

type IElement interface {
	SetKind(reflect.Kind) IElement
	SetValue(interface{}) IElement
	Kind() reflect.Kind
	Value() interface{}
}

type Element struct {
	kind  reflect.Kind
	value interface{}
}

func (e *Element) SetKind(kind reflect.Kind) IElement {
	e.kind = kind
	return e
}

func (e *Element) SetValue(value interface{}) IElement {
	e.value = value
	return e
}

func (e *Element) Kind() reflect.Kind {
	return e.kind
}

func (e *Element) Value() interface{} {
	return e.value
}
