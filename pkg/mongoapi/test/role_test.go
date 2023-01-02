package mongoapi

import (
	"crayontool-go/pkg/constant"
	"crayontool-go/pkg/mongoapi/typereq"
	"testing"
)

func Test_roleCommander_RolesInfo(t *testing.T) {
	ctl := globalCtl
	rolesInfo := typereq.NewRolesInfo(typereq.WithSelectAllRolesInfo())
	result, err := ctl.RoleCommander().RolesInfo().SetRolesInfo(rolesInfo).SetShowBuiltinRoles(constant.True).SetShowPrivileges(constant.True).SetShowAuthenticationRestrictions(constant.True).Do().GetResult()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
	for _, role := range result.Roles {
		t.Log("===================================")
		t.Logf("Role: %s", role.Role)
		t.Logf("DB: %s", role.DB)
		t.Logf("IsBuiltin: %v", role.IsBuiltin)
		t.Logf("Roles: %v", role.Roles)
		t.Logf("InheritedRoles: %v", role.InheritedRoles)
	}
}
