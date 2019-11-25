package sort

func SelectionSort(data InterFace) {
	n := data.Len()
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if data.Less(minIndex, j) {
				minIndex = j
			}
		}
		data.Swap(i, minIndex)
	}
}
