package typereq

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/mongoapi/bsonkey"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type IResource interface {
	ToBSON() bsonx.Val
}

type DBCollResource struct {
	DB         string
	Collection string
}

func (d *DBCollResource) ToBSON() bsonx.Val {
	doc := bsonx.Doc{}
	doc = doc.Set(bsonkey.KeyDB, bsonx.String(d.DB))
	doc = doc.Set(bsonkey.KeyCollection, bsonx.String(d.Collection))
	return bsonx.Document(doc)
}

type ClusterResource struct {
}

func (c *ClusterResource) ToBSON() bsonx.Val {
	doc := bsonx.Doc{}
	doc = doc.Set(bsonkey.KeyCluster, bsonx.Boolean(constant.True))
	return bsonx.Document(doc)
}

// AnyResource Do not use this resource, other than in exceptional circumstances.
type AnyResource struct {
}

func (a *AnyResource) ToBSON() bsonx.Val {
	doc := bsonx.Doc{}
	doc = doc.Set(bsonkey.KeyAnyResource, bsonx.Boolean(constant.True))
	return bsonx.Document(doc)
}
