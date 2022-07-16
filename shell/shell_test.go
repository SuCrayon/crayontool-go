package shell

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestReq_load(t *testing.T) {
	type fields struct {
		IType    InterceptorType
		Opts     []string
		Cmd      string
		fullCmd  string
		executor *exec.Cmd
		in       *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				IType: Sh,
				Cmd:   `echo "hello world"`,
			},
		},
		{
			name: "",
			fields: fields{
				Cmd:   `shell_test.sh Crayon`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReqByOption(
				WithIType(tt.fields.IType),
				WithCmd(tt.fields.Cmd),
			).FlatAppendOptions(tt.fields.Opts)
			ret, err := r.Do()
			if err != nil {
				t.Errorf("some errors occur, err: %v\n", err)
			}
			t.Logf("exec output: %s\n", string(ret))
		})
	}
}
