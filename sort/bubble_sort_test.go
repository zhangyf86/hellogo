package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	BubbleSort(IntSlice(n))
	t.Log(n)
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(IntSlice(n))
	}
}
