package sliceutil

import (
	"testing"
)

func TestGenIndexOrderSlice(t *testing.T) {
	t.Log(GenIndexOrderSlice(10))
}

func TestGenIntOrderSlice(t *testing.T) {
	t.Log(GenIntOrderSlice(10))
}

func TestGenNegativeIndexOrderSlice(t *testing.T) {
	t.Log(GenNegativeIndexOrderSlice(10))
}

func TestGenNegativeIntOrderSlice(t *testing.T) {
	t.Log(GenNegativeIntOrderSlice(10))
}

func TestGenIndexReverseOrderSlice(t *testing.T) {
	t.Log(GenIndexReverseOrderSlice(10))
}

func TestGenIntReverseOrderSlice(t *testing.T) {
	t.Log(GenIntReverseOrderSlice(10))
}

func TestGenNegativeIndexReverseOrderSlice(t *testing.T) {
	t.Log(GenNegativeIndexReverseOrderSlice(10))
}

func TestGenNegativeIntReverseOrderSlice(t *testing.T) {
	t.Log(GenNegativeIntReverseOrderSlice(10))
}

func TestFibonacciGenFunc(t *testing.T) {
	t.Log(GenSliceWithGenFunc(10, FibonacciGenFunc))
}

func TestGenReverseFibonacciSlice(t *testing.T) {
	t.Log(GenFibonacciReverseSlice(10))
}

func TestGenFibonacciSlice(t *testing.T) {
	t.Log(GenFibonacciSlice(10))
}

func TestGenFactorialSlice(t *testing.T) {
	// 0! = 1
	// 1! = 1
	// 2! = 2
	// 3! = 6
	// 4! = 24
	t.Log(GenFactorialSlice(10))
}

func TestGenFactorialReverseSlice(t *testing.T) {
	t.Log(GenFactorialReverseSlice(10))
}
