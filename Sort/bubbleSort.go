package Sort

func BubbleSort(data InterFace) {
	n := data.Len()
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		flag := false
		for j := i; j < n; j++ {
			if data.Less(i, j) {
				data.Swap(i, j)
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
