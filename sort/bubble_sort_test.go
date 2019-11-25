package sort

import (
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	n := make([]int, 20)
	for i := 0; i < len(n); i++ {
		n[i] = rand.Intn(20)
	}
	BubbleSort(IntSlice(n))
	t.Log(n)
}

func BenchmarkBubbleSort(b *testing.B) {
	n := make([]int, 10000)
	for i := 0; i < len(n); i++ {
		n[i] = rand.Intn(1000)
	}
	for i := 0; i < b.N; i++ {
		BubbleSort(IntSlice(n))
	}
}
