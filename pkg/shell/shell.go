package shell

import (
	"bytes"
	"errors"
	"github.com/SuCrayon/crayontool-go/pkg/datastructure/set"
	"github.com/SuCrayon/crayontool-go/pkg/strutil"
	"os/exec"
)

type InterceptorType string

// 解释器类型
const (
	defaultIType = Sh
	Sh           = "/bin/sh"
	Sh4Root      = "/sbin/sh"
	Bash         = "/bin/bash"
	CSh          = "/usr/bin/csh"
	KSh          = "/usr/bin/ksh"
)

// 缓冲大小
const (
	// defaultBufferSize 默认缓冲大小
	defaultBufferSize  = 1 << 5
	defaultCmdListSize = 1 << 2
)

var (
	ErrLackIType = errors.New("ErrLackIType: InterceptorType is not specify")
	ErrCmdEmpty  = errors.New("ErrCmdEmpty: CmdList is empty")
)

type Req struct {
	// shell解释器类型
	IType InterceptorType
	// shell命令
	CmdList []string
	// 退出码白名单
	exitCodeWhiteSet set.Set
	executor         *exec.Cmd
	stdin            *bytes.Buffer
	stdout           *bytes.Buffer
	stderr           *bytes.Buffer
}

type Resp struct {
	req *Req
}

func (r *Req) SetIType(it InterceptorType) *Req {
	r.IType = it
	return r
}

func (r *Req) AddCmd(cmd ...string) *Req {
	r.CmdList = append(r.CmdList, cmd...)
	return r
}

func (r *Req) AddWhiteExitCode(codes ...int) *Req {
	if r.exitCodeWhiteSet == nil {
		r.exitCodeWhiteSet = set.NewSetWithCap(len(codes))
	}
	for _, code := range codes {
		r.exitCodeWhiteSet.Add(code)
	}
	return r
}

func (t InterceptorType) ToString() string {
	return string(t)
}

func NewReq() *Req {
	req := Req{
		IType:            defaultIType,
		CmdList:          make([]string, 0, defaultCmdListSize),
		exitCodeWhiteSet: set.NewSet(),
		executor:         exec.Command(defaultIType),
	}
	req.stdin = bytes.NewBuffer(make([]byte, 0, defaultBufferSize))
	req.stdout = bytes.NewBuffer(make([]byte, 0, defaultBufferSize))
	req.stderr = bytes.NewBuffer(make([]byte, 0, defaultBufferSize))
	req.executor.Stdin = req.stdin
	req.executor.Stdout = req.stdout
	req.executor.Stderr = req.stderr
	return &req
}

func (r *Req) validate() error {
	if r.IType == "" {
		return ErrLackIType
	}
	if len(r.CmdList) == 0 {
		return ErrCmdEmpty
	}
	return nil
}

func (r *Req) load() error {
	if err := r.validate(); err != nil {
		return err
	}
	cmdStr := strutil.NewJoinReq(r.CmdList).SetSep("\n").Join()
	_, err := r.stdin.WriteString(cmdStr)
	return err
}

func (r *Req) isWhiteExitCode() bool {
	return r.exitCodeWhiteSet.Contains(r.executor.ProcessState.ExitCode())
}

func (r *Req) Do() (*Resp, error) {
	resp := Resp{
		req: r,
	}
	if err := r.load(); err != nil {
		return nil, err
	}
	err := r.executor.Run()
	if r.isWhiteExitCode() {
		err = nil
	}
	return &resp, err
}

func (r *Resp) ExitCode() int {
	return r.req.executor.ProcessState.ExitCode()
}

func (r *Resp) Stdout() *bytes.Buffer {
	return r.req.stdout
}

func (r *Resp) Stderr() *bytes.Buffer {
	return r.req.stderr
}
