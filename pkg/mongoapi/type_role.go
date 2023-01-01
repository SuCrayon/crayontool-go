package mongoapi

import "go.mongodb.org/mongo-driver/x/bsonx"

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
	doc = doc.Set(KeyRole, bsonx.String(d.Role))
	doc = doc.Set(KeyDB, bsonx.String(d.DB))
	return bsonx.Document(doc)
}

type RRole struct {
	Role string
}

func (r *RRole) ToBSON() bsonx.Val {
	return bsonx.String(r.Role)
}
