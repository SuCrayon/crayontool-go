package test

import (
	"crayontool-go/pkg/password"
	"testing"
)

func TestSugar_GenStrongPassword(t *testing.T) {
	var lens = []int{1, 2, 3, 5, 10, 32, 64, 0}
	for _, length := range lens {
		pwd := password.Sugar.GenStrongPassword(length)
		t.Logf("pwd: %s, len: %d, wantLen: %d\n", pwd, len(pwd), length)
		if len(pwd) != length {
			t.Fail()
		}
	}
}
