/*
二分查找实现注意事项:
1. 循环截止条件是 start<=end
2. mid = start + (end-start)/2 放置 start+end 溢出
3. start=mid+1 end=mid-1 特殊情况下如果不 +-1 会导致死循环
4. start，end 如何更新思考:
	- DATA[mid] > Value 说明列表值大，需要缩小范围，e=mid-1
	- DATA[mid] < Value 说明列表值小，需要扩大范围，s=mid+1
	- 始终使用 DATA[mid] 与 value 去比较，保持前后思考一致

*/

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
		if v > value {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return -1
}

// BinaryEqualFirst 查找第一个值等于给定值的元素
func BinaryEqualFirst(data []int, value int) int {
	s := 0
	e := len(data) - 1
	for s <= e {
		mid := (e-s)/2 + s
		v := data[mid]

		if v > value {
			e = mid - 1
		} else if v < value {
			s = mid + 1
		} else if mid == 0 || data[mid-1] != value {
			return mid
		} else {
			e = mid - 1
		}

	}
	return -1
}

// BinaryEqualEnd 查找最后一个值等于给定值的元素
func BinaryEqualEnd(data []int, value int) int {
	s := 0
	e := len(data) - 1
	for s <= e {
		mid := s + (e-s)/2
		v := data[mid]
		if v > value {
			e = mid - 1
		} else if v < mid {
			s = mid + 1
		} else if mid == len(data)-1 || data[mid+1] != value {
			return mid
		} else {
			s = mid + 1
		}
	}
	return -1
}

// BinaryGteFirst 查找第一个大于等于给定值的元素
func BinaryGteFirst(data []int, value int) int {
	s := 0
	e := len(data) - 1
	for s <= e {
		mid := s + (e-s)/2
		v := data[mid]
		if v < value {
			s = mid + 1
		} else if mid == 0 || data[mid-1] < value {
			return mid
		} else {
			e = mid - 1
		}
	}
	return -1
}

// BinaryLteEnd 查找最后一个小于等于给定值的元素
func BinaryLteEnd(data []int, value int) int {
	s := 0
	e := len(data) - 1
	for s <= e {
		mid := s + (e-s)/2
		v := data[mid]
		if v > value {
			e = mid - 1
		} else if mid == len(data)-1 || data[mid+1] > value {
			return mid
		} else {
			s = mid + 1
		}
	}
	return -1
}
