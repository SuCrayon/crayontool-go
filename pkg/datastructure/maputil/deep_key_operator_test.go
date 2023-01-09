package maputil

import (
	"testing"
)

func Test_DeepKeyOperator(t *testing.T) {
	m, err := DeepKeyOperator.Gen("key", "hello")
	if err != nil {
		t.Error(err)
	}
	t.Log(m)
	if err := DeepKeyOperator.Set(m, "outerKey.innerKey.realKey", "hello"); err != nil {
		t.Error(err)
	}
	t.Log(m)
	if err := DeepKeyOperator.Set(m, "outerKey.realKey", "hello"); err != nil {
		t.Error(err)
	}
	t.Log(m)
}
