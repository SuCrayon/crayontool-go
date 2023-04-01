package typereq

import (
	"github.com/SuCrayon/crayontool-go/pkg/mongoapi/bsonkey"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type IRole interface {
	ToBSON() bsonx.Val
}

/*
{ role: "<role>", db: "<database>" } | "<role>"
*/

type DBRole struct {
	Role string
	DB   string
}

func (d *DBRole) ToBSON() bsonx.Val {
	doc := bsonx.Doc{}
	doc = doc.Set(bsonkey.KeyRole, bsonx.String(d.Role))
	doc = doc.Set(bsonkey.KeyDB, bsonx.String(d.DB))
	return bsonx.Document(doc)
}

type StrRole string

func (s *StrRole) ToBSON() bsonx.Val {
	return bsonx.String(string(*s))
}
