package std

import (
	"fmt"
	"github.com/SuCrayon/crayontool-go/pkg/shell"
	"os"
)

func Errorf(format string, vs ...interface{}) {
	ErrorfFullParam(format, shell.OperationNotPermitted, vs...)
}

func Error(vs ...interface{}) {
	ErrorFullParam(shell.OperationNotPermitted, vs...)
}

func ErrorfFullParam(format string, exitCode int, vs ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, vs...)
	os.Exit(exitCode)
}

func ErrorFullParam(exitCode int, vs ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, vs...)
	os.Exit(exitCode)
}
