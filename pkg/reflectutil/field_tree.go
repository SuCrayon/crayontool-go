package reflectutil

/*import (
	"crayontool-go/pkg/types"
	"errors"
)

var (
	TypeNotSupportErr = errors.New("type not support")
)

type FieldTreeParser interface {
	Parse(v interface{}) (*types.FieldTreeNode, error)
	Judge(v interface{}) bool
}

type fieldTreeParser struct {
}

func (p *fieldTreeParser) Judge(v interface{}) bool {
	return RealTypeIsStruct(v) || RealTypeIsSlice(v) || RealTypeIsMap(v) || RealTypeIsArray(v)
}

func (p *fieldTreeParser) Parse(v interface{}) (*types.FieldTreeNode, error) {
	var root types.FieldTreeNode
	err := p.doParse(v, root)
	return &root, err
}

func (p *fieldTreeParser) doParse(v interface{}, node types.FieldTreeNode) error {
	if !p.Judge(v) {
		return TypeNotSupportErr
	}
	_type := RealTypeOf(v)
	value := RealValueOf(v)
}*/
