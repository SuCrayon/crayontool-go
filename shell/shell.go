package shell

import (
	"bytes"
	"crayontool-go/strutil"
	"fmt"
	"os/exec"
	"strings"
)

type InterceptorType string

const (
	Leave   = ""
	Bash    = "/bin/bash"
	Sh      = "/bin/sh"
	CSh     = "/usr/bin/csh"
	KSh     = "/usr/bin/ksh"
	Sh4Root = "/sbin/sh"
)

const (
	// 缓冲大小
	bufferSize = 1 << 5
	// 执行可执行文件时需添加前缀
	execPrefix = "./"
)

type Req struct {
	// shell解释器类型
	IType InterceptorType
	// 解释器参数
	Opts []string
	// shell命令
	Cmd string
	// 完整命令
	fullCmd  string
	executor *exec.Cmd
	in       *bytes.Buffer
}

type ReqOption func(req *Req)

func WithIType(it InterceptorType) ReqOption {
	return func(req *Req) {
		req.IType = it
	}
}

func WithCmd(cmd string) ReqOption {
	return func(req *Req) {
		req.Cmd = cmd
	}
}

func (t InterceptorType) ToString() string {
	return string(t)
}

func NewReqByOption(options ...ReqOption) *Req {
	req := Req{}
	for i := range options {
		options[i](&req)
	}
	return &req
}

func (r *Req) AppendOption(opt string) *Req {
	r.Opts = append(r.Opts, opt)
	return r
}

func (r *Req) FlatAppendOptions(opts []string) *Req {
	if len(opts) > 0 {
		for i := range opts {
			r.AppendOption(opts[i])
		}
	}
	return r
}

func (r *Req) loadFullCmd() {
	joinReq := strutil.NewJoinReq(
		[]string{
			r.IType.ToString(),
			strings.Join(r.Opts, " "),
			r.Cmd,
		},
	).SetSep(" ").SetOmitEmpty(true)
	if r.IType == Leave {
		joinReq.SetPrefix(execPrefix)
	}
	// strings.Join空安全的
	r.fullCmd = joinReq.Join()
	// r.fullCmd = strings.Join([]string{r.IType.ToString(), strings.Join(r.Opts, " "), r.Cmd}, " ")
	fmt.Printf("execute command: %s\n", r.fullCmd)
}

func (r *Req) load() {
	r.executor = exec.Command(string(r.IType))
	r.in = bytes.NewBuffer(make([]byte, 0, bufferSize))
	r.executor.Stdin = r.in

}

func (r *Req) Do() ([]byte, error) {
	r.load()
	r.in.WriteString(r.fullCmd)
	return r.executor.Output()
}

/*package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

type InterceptorType string

const (
	Bash    = "/bin/bash"
	Sh      = "/bin/sh"
	CSh     = "/usr/bin/csh"
	KSh     = "/usr/bin/ksh"
	Sh4Root = "/sbin/sh"
)

const (
	bufferSize        = 1 << 5
	tmpShPathTemplate = "/tmp/tmp.%s.sh"
	modeExec          = 0755
)

type Req struct {
	IType InterceptorType
	Cmd   string
	file  *os.File
}

func SimpleNewReq(cmd string) *Req {
	return &Req{
		IType: Sh,
		Cmd:   cmd,
	}
}

func (r *Req) createShFile() error {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	fName := fmt.Sprintf(tmpShPathTemplate, id)
	file, err := os.Create(fName)
	if err != nil {
		return err
	}
	if err := os.Chmod(fName, modeExec); err != nil {
		return err
	}
	r.file = file
	return nil
}

func (r *Req) writeCmd() error {
	_, err := r.file.WriteString(r.Cmd)
	return err
}

func (r *Req) deleteShFile() error {
	return os.Remove(r.file.Name())
}

func (r *Req) innerDo() ([]byte, error) {
	cmd := exec.Command(string(r.IType), r.file.Name())
	return cmd.Output()
}

func (r *Req) preprocess() error {
	if err := r.createShFile(); err != nil {
		return err
	}
	if err := r.writeCmd(); err != nil {
		return err
	}
	return nil
}

func (r *Req) postprocess() error {
	return r.deleteShFile()
}

func (r *Req) Do() ([]byte, error) {
	if err := r.preprocess(); err != nil {
		return nil, err
	}
	ret, err := r.innerDo()
	if err != nil {
		return nil, err
	}
	if err := r.postprocess(); err != nil {
		return nil, err
	}
	return ret, nil
}*/
