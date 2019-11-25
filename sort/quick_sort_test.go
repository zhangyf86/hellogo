package sort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	n := make([]int, 20)
	for i := 0; i < len(n); i++ {
		n[i] = rand.Intn(20)
	}
	QuickSort(IntSlice(n))
	t.Log(n)
}

func BenchmarkQuickSort(t *testing.B) {
	n := make([]int, 20)
	for i := 0; i < len(n); i++ {
		n[i] = rand.Intn(20)
	}
	QuickSort(IntSlice(n))
}
