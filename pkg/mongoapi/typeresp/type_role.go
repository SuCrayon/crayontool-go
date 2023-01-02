package typeresp

type Role struct {
}

type TRole struct {
	Role                string      `bson:"role" json:"role" yaml:"role" xml:"role"`
	DB                  string      `bson:"db" json:"db" yaml:"db" xml:"db"`
	IsBuiltin           bool        `bson:"isBuiltin" json:"isBuiltin" yaml:"isBuiltin" xml:"isBuiltin"`
	Roles               []Role      `bson:"roles" json:"roles" yaml:"roles" xml:"roles"`
	InheritedRoles      []Role      `bson:"inheritedRoles" json:"inheritedRoles" yaml:"inheritedRoles" xml:"inheritedRoles"`
	Privileges          []Privilege `bson:"privileges" json:"privileges" yaml:"privileges" xml:"privileges"`
	InheritedPrivileges []Privilege `bson:"InheritedPrivileges" json:"InheritedPrivileges" yaml:"InheritedPrivileges" xml:"InheritedPrivileges"`
}
