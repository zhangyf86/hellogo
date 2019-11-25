package sort

func InsertionSort(data InterFace) {
	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data.Less(j-1, j); j-- {
			data.Swap(j, j-1)
		}
	}
}
