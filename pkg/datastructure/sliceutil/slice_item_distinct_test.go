package sliceutil

import (
	"testing"
)

func TestDistinct(t *testing.T) {
	intList := []int{1, 2, 2, 3, 4, 1}
	intDistinctList := Distinct(intList)
	t.Log(intDistinctList)
	t.Log(intDistinctList.([]int))

	intListPtr := &[]int{1, 2, 2, 3, 4, 1}
	intDistinctPtrList := Distinct(intListPtr)
	t.Log(intDistinctPtrList)
	t.Log(intDistinctPtrList.(*[]int))
}
