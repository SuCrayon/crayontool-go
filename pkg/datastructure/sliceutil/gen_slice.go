package sliceutil

func IndexOrderGenFunc(arr []int, i int) int {
	return i
}

func IntOrderGenFunc(arr []int, i int) int {
	return i + 1
}

func NegativeIndexOrderGenFunc(arr []int, i int) int {
	return -i
}

func NegativeIntOrderGenFunc(arr []int, i int) int {
	return -i - 1
}

func FibonacciGenFunc(arr []int, i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return arr[i-2] + arr[i-1]
}

func ReverseFibonacciGenFunc(arr []int, i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	// return arr[len(arr)-1-i+2] + arr[len(arr)-1-i+1]
	return arr[len(arr)-i+1] + arr[len(arr)-i]
}

func FactorialGenFunc(arr []int, i int) int {
	if i == 0 {
		return 1
	}
	return arr[i-1] * i
}

func ReverseFactorialGenFunc(arr []int, i int) int {
	if i == 0 {
		return 1
	}
	return arr[len(arr)-i] * i
}

func GenSliceWithGenFunc(cap int, genFunc func(arr []int, i int) int) []int {
	var ret = make([]int, cap)
	for i := 0; i < len(ret); i++ {
		ret[i] = genFunc(ret, i)
	}
	return ret
}

func GenReverseSliceWithGenFunc(cap int, genFunc func(arr []int, i int) int) []int {
	var ret = make([]int, cap)
	for i := 0; i < len(ret); i++ {
		ret[len(ret)-1-i] = genFunc(ret, i)
	}
	return ret
}

func GenIndexOrderSlice(cap int) []int {
	return GenSliceWithGenFunc(cap, IndexOrderGenFunc)
}

func GenIntOrderSlice(cap int) []int {
	return GenSliceWithGenFunc(cap, IntOrderGenFunc)
}

func GenNegativeIndexOrderSlice(cap int) []int {
	return GenSliceWithGenFunc(cap, NegativeIndexOrderGenFunc)
}

func GenNegativeIntOrderSlice(cap int) []int {
	return GenSliceWithGenFunc(cap, NegativeIntOrderGenFunc)
}

func GenIndexReverseOrderSlice(cap int) []int {
	return GenReverseSliceWithGenFunc(cap, IndexOrderGenFunc)
}

func GenIntReverseOrderSlice(cap int) []int {
	return GenReverseSliceWithGenFunc(cap, IntOrderGenFunc)
}

func GenNegativeIndexReverseOrderSlice(cap int) []int {
	return GenReverseSliceWithGenFunc(cap, NegativeIndexOrderGenFunc)
}

func GenNegativeIntReverseOrderSlice(cap int) []int {
	return GenReverseSliceWithGenFunc(cap, NegativeIntOrderGenFunc)
}

func GenFibonacciSlice(cap int) []int {
	return GenSliceWithGenFunc(cap, FibonacciGenFunc)
}

func GenFibonacciReverseSlice(cap int) []int {
	return GenReverseSliceWithGenFunc(cap, ReverseFibonacciGenFunc)
}

func GenFactorialSlice(cap int) []int {
	return GenSliceWithGenFunc(cap, FactorialGenFunc)
}

func GenFactorialReverseSlice(cap int) []int {
	return GenReverseSliceWithGenFunc(cap, ReverseFactorialGenFunc)
}
