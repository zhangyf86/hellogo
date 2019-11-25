package sort

import "testing"

func TestBucketSort(t *testing.T) {
	t.Log(n)
	BucketSort(n)
	t.Log(n)
}

func TestCountingSort(t *testing.T) {
	CountingSort(n)
	t.Log(n)
}

func TestRadixSort(t *testing.T) {
	RadixSort(n)
	t.Log(n)
}

func BenchmarkBucketSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BucketSort(n)
	}
}

func BenchmarkCountingSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountingSort(n)
	}
}
