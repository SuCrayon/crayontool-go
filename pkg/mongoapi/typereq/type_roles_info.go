package typereq

import (
	"github.com/SuCrayon/crayontool-go/pkg/constant"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type IRolesInfo interface {
	ToBSON() bsonx.Val
}

type RolesInfo []IRole

func (r *RolesInfo) ToBSON() bsonx.Val {
	if len(*r) == 0 {
		return bsonx.Document(bsonx.Doc{})
	}
	if len(*r) == 1 {
		return (*r)[0].ToBSON()
	}
	arr := bsonx.Arr{}
	for _, v := range *r {
		arr = append(arr, v.ToBSON())
	}
	return bsonx.Array(arr)
}

type RolesInfoOption func(r *RolesInfo)

func WithSelectAllRolesInfo() RolesInfoOption {
	return func(r *RolesInfo) {
		*r = make([]IRole, constant.IntOne)
		(*r)[0] = &IntOne{}
	}
}

func WithDBRole(role DBRole) RolesInfoOption {
	return func(r *RolesInfo) {
		*r = append(*r, &role)
	}
}

func WithStrRole(role StrRole) RolesInfoOption {
	return func(r *RolesInfo) {
		*r = append(*r, &role)
	}
}

func NewRolesInfo(options ...RolesInfoOption) IRolesInfo {
	rolesInfo := &RolesInfo{}
	for _, f := range options {
		f(rolesInfo)
	}
	return rolesInfo
}
