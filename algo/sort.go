package algo

// QuickSort 快排
func QuickSort(data []int, left, right int) {
	if left >= right {
		return
	}
	povit := Partition(data, left, right)
	QuickSort(data, left, povit-1)
	QuickSort(data, povit+1, right)
}

// Partition 快排分区函数
func Partition(data []int, left, right int) int {
	l := left
	r := right - 1
	partition := right
	v := data[partition]
	for l <= r {
		for (l <= r) && (data[l] < v) {
			l++
		}
		for (l <= r) && (data[r] > v) {
			r--
		}
		if l <= r {
			data[l], data[r] = data[r], data[l]
			l++
			r--
		}
	}
	data[l], data[partition] = data[partition], data[l]
	return l
}
