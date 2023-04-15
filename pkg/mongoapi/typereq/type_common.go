package typereq

import (
	"github.com/SuCrayon/crayontool-go/pkg/constant"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type IntOne struct {
}

func (i *IntOne) ToBSON() bsonx.Val {
	return bsonx.Int64(constant.IntOne)
}
