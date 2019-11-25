package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	QuickSort(IntSlice(n))
	t.Log(n)
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(IntSlice(n))
	}
}
