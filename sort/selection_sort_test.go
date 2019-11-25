package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	SelectionSort(IntSlice(n))
	t.Log(n)
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(IntSlice(n))
	}
}
