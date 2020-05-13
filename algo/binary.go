package algo

// BinarySearch 精确二分查找
func BinarySearch(data []int, value int) int {
	s := 0
	e := len(data) - 1
	for s <= e {
		mid := s + (e-s)/2
		v := data[mid]
		if v == value {
			return mid
		}
		if value > v {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return -1
}

// BinaryEqualFirst 查找第一个值等于给定值的元素
func BinaryEqualFirst() {

}

// BinaryEqualEnd 查找最后一个值等于给定值的元素
func BinaryEqualEnd() {

}

// BinaryGteFirst 查找第一个大于等于给定值的元素
func BinaryGteFirst() {

}

// BinaryLteEnd 查找最后一个小于等于给定值的元素
func BinaryLteEnd() {

}
