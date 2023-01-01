package mongoapi

import (
	"testing"
)

func init() {
	ctlInit()
}

func Test_diagnosticCommander_Ping(t *testing.T) {
	ctl := globalCtl
	result, err := ctl.DiagnosticCommander().Ping().Do().GetResult()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
