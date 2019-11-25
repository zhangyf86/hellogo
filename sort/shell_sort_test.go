package sort

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	ShellSort(IntSlice(n))
	t.Log(n)
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort(IntSlice(n))
	}
}
