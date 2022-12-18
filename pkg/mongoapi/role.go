package mongoapi

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
