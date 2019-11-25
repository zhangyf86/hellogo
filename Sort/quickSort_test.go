package Sort

import (
	"math/rand"
	"testing"
)

var n []int

func init() {
	n = make([]int, 100000)
	for i := 0; i < len(n); i++ {
		n[i] = rand.Intn(10000)
	}
}
func Test_quickSort(t *testing.T) {
	QuickSort(IntSlice(n))
	t.Log(n)
}

func Benchmark_quickSort(t *testing.B) {
	QuickSort(IntSlice(n))
}
