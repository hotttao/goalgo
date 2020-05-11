package algo

// QuickSelect 求无序数组中的第 K 大元素
func QuickSelect() {

}

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

// MergeSort 归并排序
func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	merge := Merge(left, right)
	return merge
}

// Merge 归并排序的合并函数
func Merge(A, B []int) []int {
	C := make([]int, len(A)+len(B), len(A)+len(B))
	i, j := 0, 0
	for i+j < len(C) {
		if i > len(A)-1 || (j <= len(B)-1) && (A[i] > B[j]) {
			C[i+j] = B[j]
			j++
		} else {
			C[i+j] = A[i]
			i++
		}
	}
	return C
}

// DoubbleSort 冒泡排序
func DoubbleSort() {

}

// InsertSort 插入排序
func InsertSort() {

}

// SelectSort 选择排序
func SelectSort() {

}

// BucketSort 桶排序
func BucketSort() {

}

// CountSort 计数排序
func CountSort() {

}

// BaseSort 基数排序
func BaseSort() {

}
