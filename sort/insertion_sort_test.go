package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	InsertionSort(IntSlice(n))
	t.Log(n)
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(IntSlice(n))
	}
}
