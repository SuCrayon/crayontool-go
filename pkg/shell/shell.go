package shell

import (
	"bytes"
	"crayontool-go/pkg/strutil"
	"errors"
	"fmt"
	"io/ioutil"
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
	CmdList  []string
	executor *exec.Cmd
	in       *bytes.Buffer
	err      *bytes.Buffer
}

func (r *Req) SetIType(it InterceptorType) *Req {
	r.IType = it
	return r
}

func (r *Req) AddCmd(cmd ...string) *Req {
	r.CmdList = append(r.CmdList, cmd...)
	return r
}

func (t InterceptorType) ToString() string {
	return string(t)
}

func NewReq() *Req {
	req := Req{
		IType:    defaultIType,
		CmdList:  make([]string, 0, defaultCmdListSize),
		executor: exec.Command(defaultIType),
	}
	req.in = bytes.NewBuffer(make([]byte, 0, defaultBufferSize))
	req.executor.Stdin = req.in
	req.err = bytes.NewBuffer(make([]byte, 0, defaultBufferSize))
	req.executor.Stderr = req.err
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
	_, err := r.in.WriteString(cmdStr)
	return err
}

func (r *Req) readStderr() []byte {
	stderr, err := ioutil.ReadAll(r.err)
	if err != nil {
		fmt.Printf("some errors occur when read stderr, err: %v\n", err)
		return nil
	}
	return stderr
}

func (r *Req) Do() ([]byte, error) {
	if err := r.load(); err != nil {
		return nil, err
	}
	output, err := r.executor.Output()
	if err != nil {
		stderr := r.readStderr()
		err = fmt.Errorf("%w\ndetail: %s\n", err, string(stderr))
	}
	return output, err
}
