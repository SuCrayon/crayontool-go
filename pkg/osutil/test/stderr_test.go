package test

import (
	"github.com/SuCrayon/crayontool-go/pkg/osutil/std"
	"testing"
)

func TestErrorf(t *testing.T) {
	std.Errorf("some errors occur when execute command, command: %s\n", "ls -l")
}

func TestError(t *testing.T) {
	std.Error("some errors occur when execute command, command: ", "ls -l", "\n")
}
