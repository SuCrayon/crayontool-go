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
	LackITypeErr = errors.New("ErrLackIType: InterceptorType is not specify")
	CmdEmptyErr  = errors.New("ErrCmdEmpty: CmdList is empty")
)

type Req struct {
	// shell解释器类型
	IType InterceptorType
	// shell命令
	CmdList []string
	// 退出码白名单
	exitCodeWhiteSet set.Set
	executor         *exec.Cmd
	in               *bytes.Buffer
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
	req.in = bytes.NewBuffer(make([]byte, 0, defaultBufferSize))
	req.executor.Stdin = req.in
	return &req
}

func (r *Req) validate() error {
	if r.IType == "" {
		return LackITypeErr
	}
	if len(r.CmdList) == 0 {
		return CmdEmptyErr
	}
	return nil
}

func (r *Req) load() error {
	if err := r.validate(); err != nil {
		return err
	}
	cmdStr := strutil.NewJoinReq(r.CmdList).SetSep("\n").Join()
	_, err := r.in.WriteString(cmdStr)
	return err
}

func (r *Req) isWhiteExitCode(err error) bool {
	if ee, ok := err.(*exec.ExitError); ok {
		return r.exitCodeWhiteSet.Contains(ee.ExitCode())
	}
	return false
}

func (r *Req) Do() ([]byte, error) {
	if err := r.load(); err != nil {
		return nil, err
	}
	output, err := r.executor.CombinedOutput()
	if r.isWhiteExitCode(err) {
		err = nil
	}
	return output, err
}
