package algo

import "fmt"

// Heap 大堆
type Heap struct {
	data     []int
	len      int
	capacity int
}

// NewHeap 新建 heap
func NewHeap(capacity int) *Heap {
	return &Heap{
		data:     make([]int, capacity+1, capacity+1), // capacity 要加 1
		len:      0,
		capacity: capacity,
	}
}

// Pop 弹出栈顶元素
func (h *Heap) Pop() int {
	if h.len <= 0 {
		panic("Heap null")
	}
	v := h.data[1]
	h.data[1] = h.data[h.len]
	h.len--
	h.DownHeap(1)
	return v
}

// Add 插入元素
func (h *Heap) Add(value int) {
	if h.len == h.capacity {
		h.Resize(2 * h.len)
	}
	h.len++
	h.data[h.len] = value
	h.UpHeap(h.len)
}

func (h *Heap) parent(i int) int {
	return i / 2
}

func (h *Heap) left(i int) int {
	return 2 * i
}

func (h *Heap) right(i int) int {
	return 2*i + 1
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// UpHeap 从下往上的堆化
func (h *Heap) UpHeap(i int) {
	for i/2 > 0 && (h.data[i] > h.data[i/2]) {
		h.swap(i, i/2)
		i = i / 2
	}
}

// Resize heap 扩容
func (h *Heap) Resize(size int) {
	if size < h.len {
		return
	}
	data := make([]int, size+1, size+1) // size 必须加 1
	for i := h.len; i >= 1; i-- {
		data[i] = h.data[i]
	}
	h.data = data
	h.capacity = size
}

// DownHeap 从上往下的堆化
func (h *Heap) DownHeap(i int) {
	for {
		maxPos := i
		if left := h.left(i); left <= h.len && h.data[i] < h.data[left] {
			h.swap(i, left)
			maxPos = left // 注意事项 1: 不能直接更新 i，因为 right 比较中使用了i
		}
		if right := h.right(i); right <= h.len && h.data[i] < h.data[right] {
			h.swap(i, right)
			maxPos = right
		}
		if maxPos == i {
			break
		}
		i = maxPos
	}
}

// BuildHeap 初始化一个堆
func BuildHeap(data []int) Heap {
	num := len(data)
	data = append(data, data[0])
	data[0] = -1
	h := Heap{
		data:     data,
		len:      num,
		capacity: num,
	}
	for i := h.parent(h.len); i >= 1; i-- {
		h.DownHeap(i)
	}
	return h
}

// HeapSort 堆排序
func HeapSort(data []int) []int {
	h := BuildHeap(data)
	fmt.Println(h.data)
	for i := h.len; i >= 1; i-- {
		h.swap(1, i)
		h.len--
		h.DownHeap(1)
	}
	return h.data[1:]
}
