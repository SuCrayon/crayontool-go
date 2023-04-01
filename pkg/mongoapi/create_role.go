package mongoapi

import "github.com/SuCrayon/crayontool-go/pkg/mongoapi/typereq"

type CreateRoleReq struct {
	iCommandReq
	RoleName                   string
	Privileges                 []*Privilege
	Roles                      []typereq.IRole
	AuthenticationRestrictions []*AuthenticationRestriction
}

type Privilege struct {
	Resource typereq.IResource
	Actions  []typereq.TypeAction
}

type AuthenticationRestriction struct {
	ClientSources []string
	ServerAddress []string
}

func (c *CreateRoleReq) Do() *CreateRoleReq {
	return c
}
