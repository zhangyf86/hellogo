package sort

import "fmt"

func getMax(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func BucketSort(a []int) {
	num := len(a)
	if num <= 1 {
		return
	}
	max := getMax(a)
	buckets := make([][]int, num) // 二维切片

	index := 0
	for i := 0; i < num; i++ {
		index = a[i] * (num - 1) / max                // 桶序号
		buckets[index] = append(buckets[index], a[i]) // 加入对应的桶中
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(IntSlice(buckets[i])) // 桶内做快速排序
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}

// 计数排序
func CountingSort(source []int) {
	if len(source) <= 1 {
		return
	}
	array := make([]int, getMax(source)+1)
	for i := 0; i < len(source); i++ {
		array[source[i]]++
	}
	c := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for array[i] != 0 {
			c = append(c, i)
			array[i]--
		}
	}
	copy(source, c)
}

func getNumLen(num int) int {
	if num == 0 {
		return 1
	}
	len := 0
	for i := num; i > 0; i /= 10 {
		len++
	}
	return len
}

// 基数排序
func RadixSort(data []int) {
	maxDigit := getNumLen(getMax(data))
	mod, dev := 10, 1
	fmt.Println(maxDigit)
	for i := 0; i < maxDigit; i++ {
		counter := make([][]int, mod)
		for j := 0; j < len(data); j++ {
			bucket := (data[j] % mod) / dev
			counter[bucket] = append(counter[bucket], data[j])
		}
		pos := 0
		for _, bucket := range counter {
			for _, value := range bucket {
				data[pos] = value
				pos++
			}
		}
		mod *= 10
		dev *= 10
	}
}
