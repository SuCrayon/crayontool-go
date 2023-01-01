package mongoapi

import (
	"sync"
)

// Database User Roles
const (
	// RoleRead Provides the ability to read data on all non-system collections and the system.js collection.
	RoleRead = "read"
	// RoleReadWrite Provides all the privileges of the `read` role plus ability to modify data on all non-system collections and the system.js collection.
	RoleReadWrite = "readWrite"
)

// Database Administration Roles
const (
	RoleDBAdmin   = "dbAdmin"
	RoleDBOwner   = "dbOwner"
	RoleUserAdmin = "userAdmin"
)

// Cluster Administration Roles
const (
	RoleClusterAdmin   = "clusterAdmin"
	RoleClusterManager = "clusterManager"
	RoleClusterMonitor = "clusterMonitor"
	RoleHostManager    = "hostManager"
)

// Backup and Restoration Roles
const (
	RoleBackup  = "backup"
	RoleRestore = "restore"
)

// All-Database Roles
const (
	RoleReadAnyDatabase      = "readAnyDatabase"
	RoleReadWriteAnyDatabase = "readWriteAnyDatabase"
	RoleUserAdminAnyDatabase = "userAdminAnyDatabase"
	RoleDBAdminAnyDatabase   = "dbAdminAnyDatabase"
)

// Superuser Roles
const (
	// RoleRoot root
	RoleRoot = "root"
)

// Internal Role
const (
	// RoleSystem __system
	RoleSystem = "__system"
)

// Role Management Commands
const (
	CmdCreateRole = "createRole"
	CmdDropRole   = "dropRole"
	CmdRolesInfo  = "rolesInfo"
)

type RoleCommander interface {
	CreateRole() *CreateRoleReq
	RolesInfo() *RolesInfoReq
}

type roleCommander struct {
	commander
}

var (
	rcGetOnce sync.Once
	rc        RoleCommander
)

func GetRoleCommander(ctl MongoCtl) RoleCommander {
	if rc == nil {
		rcGetOnce.Do(func() {
			rc = &roleCommander{
				commander: commander{
					ctl: ctl,
				},
			}
		})
	}
	return rc
}

func (r *roleCommander) CreateRole() *CreateRoleReq {
	return &CreateRoleReq{
		iCommandReq: NewDefaultCommandReq(r.ctl).setCommandStr(CmdCreateRole),
	}
}

func (r *roleCommander) RolesInfo() *RolesInfoReq {
	return &RolesInfoReq{
		iCommandReq: NewDefaultCommandReq(r.ctl).setCommandStr(CmdRolesInfo),
	}
}
