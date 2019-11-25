package sort

import "math/rand"

var n []int

func init() {
	n = make([]int, 20)
	for i := 0; i < len(n); i++ {
		n[i] = rand.Intn(20)
	}
}
