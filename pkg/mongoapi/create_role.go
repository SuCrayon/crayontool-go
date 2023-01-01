package mongoapi

type CreateRoleReq struct {
	iCommandReq
	RoleName                   string
	Privileges                 []*Privilege
	Roles                      []IRole
	AuthenticationRestrictions []*AuthenticationRestriction
}

type Privilege struct {
	Resource IResource
	Actions  []TypeAction
}

type AuthenticationRestriction struct {
	ClientSources []string
	ServerAddress []string
}

func (c *CreateRoleReq) Do() *CreateRoleReq {
	return c
}
