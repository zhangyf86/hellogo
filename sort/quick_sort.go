package sort

func QuickSort(data InterFace) {
	separateSort(data, 0, data.Len()-1)
}

func separateSort(data InterFace, start, end int) {
	if start >= end {
		return
	}
	i := partition(data, start, end)
	separateSort(data, start, i-1)
	separateSort(data, i+1, end)
}

func partition(data InterFace, start, end int) int {
	i := start
	for j := start; j < end; j++ {
		if data.Less(end, j) {
			if j != i {
				data.Swap(i, j)
			}
			i++
		}
	}
	data.Swap(end, i)
	return i
}
