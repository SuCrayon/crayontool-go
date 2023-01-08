package stream

import (
	"crayontool-go/pkg/function"
)

type Stream interface {
	ForEach(consumer function.FConsumerFunction)
}

func Of(vs ...interface{}) Stream {
	return &stream{}
}

type stream struct {
	elems []IElement
}

func (s *stream) ForEach(consumer function.FConsumerFunction) {
	for i := range s.elems {
		consumer(s.elems[i])
	}
}
