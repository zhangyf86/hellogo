package sort

func ShellSort(data InterFace) {
	n := data.Len()
	for step := n / 2; step > 0; step /= 2 {
		for i := step; i < n; i++ {
			for j := i - step; j >= 0 && data.Less(j, j+step); j -= step {
				data.Swap(j, j+step)
			}
		}
	}
}
