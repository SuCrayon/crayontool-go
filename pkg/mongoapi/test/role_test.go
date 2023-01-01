package mongoapi

import (
	"testing"
)

func Test_roleCommander_RolesInfo(t *testing.T) {
	ctl := globalCtl
	ctl.RoleCommander().RolesInfo().Do()
}
