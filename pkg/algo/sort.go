package algo

// QuickSelect 求无序数组中的第 K 大元素
func QuickSelect(data []int, n int) int {
	s := 0
	e := len(data) - 1
	l := Partition(data, s, e)
	for l != n-1 {
		if n > l+1 { // 使用 n 与 l+1 比较更加直观
			s = l + 1
		} else {
			e = l - 1
		}
		l = Partition(data, s, e)
	}
	return data[l]
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
func DoubbleSort(data []int) []int {
	for i := len(data) - 1; i >= 2; i-- {
		for j := i; j >= 1; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return data
}

// InsertSort 插入排序
func InsertSort(data []int) {
	for i := 1; i <= len(data)-1; i++ {
		tmp := data[i]
		j := i - 1
		for ; j >= 0; j-- {
			if data[j] > tmp { // 注意不能使用 data[i]，因为其已经被覆盖
				data[j+1] = data[j]
			} else {
				break
			}
		}
		data[j+1] = tmp
	}
}

// SelectSort 选择排序
func SelectSort(data []int) {
	for i := 0; i <= len(data)-1; i++ {
		m := data[i]
		n := i
		for j := i + 1; j <= len(data)-1; j++ {
			if data[j] > m {
				m = data[j]
				n = j
			}
		}
		data[i], data[n] = data[n], data[i]
	}
}

// BucketSort 桶排序
func BucketSort() {

}

// CountSort 计数排序
func CountSort(data []int) []int {
	m := data[0]
	for _, d := range data {
		if d > m {
			m = d
		}
	}
	bucket := make([]int, m+1, m+1)
	for _, d := range data {
		bucket[d]++
	}
	for i := 1; i <= m; i++ {
		bucket[i] += bucket[i-1]
	}
	// fmt.Println(bucket)
	tmp := make([]int, len(data), len(data))
	for i := len(data) - 1; i >= 0; i-- {
		bucket[data[i]]-- // 必须减去 1 才是对应索引位置
		tmp[bucket[data[i]]] = data[i]
	}
	return tmp
}

// BaseSort 基数排序
func BaseSort(data []int) []int {
	return data
}
