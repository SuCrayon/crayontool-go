package shell

import (
	"testing"
)

func Test(t *testing.T) {
	type fields struct {
		IType InterceptorType
		Cmd   string
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
				Cmd: `./shell_test.sh Crayon`,
			},
		},
		{
			name: "",
			fields: fields{
				Cmd: `echo now: "date +%F"
                     echo my name is Crayon
                     ./shell_test.sh Crayon
                     bash -c 'echo hello'`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := NewReq()
			if tt.fields.IType != "" {
				req.SetIType(tt.fields.IType)
			}
			ret, err := req.AddCmd(tt.fields.Cmd).Do()
			if err != nil {
				t.Errorf("some errors occur, err: %v\n", err)
			}
			t.Logf("exec output: %s\n", string(ret))
		})
	}
}
