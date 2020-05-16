/*
Hash 实现过程
1. 定义 Hash 接口包括 Get，Set，Del，Len 方法
2. 实现 Hashbase ，作为 HashMap 的抽象框架，定义包括
	- 需要子类实现的 bucketGet，bucketSet，bucketDel 方法
	- hash 哈希函数
	- Resize 动态扩缩容
3. 需要注意:
	- size 字段的更新，删除元素的时 -1，增加元素时 +1
4. 实现 ProbeHashMap 线性探测的 hashmap
	- 实现线性探测函数 findSlot，查找 key 位置
	- 实现  bucketGet，bucketSet，bucketDel，Iter 方法
	- 通过实例重新赋值，实现方法重载

难点: 如何使用 Go 实现模板模式
	- 抽象方法 eg: bucketGet 必须作为HashMapbase 的属性而不是方法存在
	- 通过对子类实例重新赋值，实现方法重载
*/

package algo

import "fmt"

var _delete HashItem

// Map 映射类型的接口规范
type Map interface {
	Get(int) string
	Set(int) string
	Del(int) bool
	Len() int
}

// Item MAP Iter 返回值接口
type Item interface {
	Key() int
	Value() string
}

// HashItem hash 使用 item，用于保存数据
type HashItem struct {
	key   int
	value string
}

// Key 方法
func (i *HashItem) Key() int {
	return i.key
}

// Value 方法
func (i *HashItem) Value() string {
	return i.value
}

// HashMapBase 散列表实现的样板
type HashMapBase struct {
	table     []Item
	size      int
	capbility int

	// 压缩函数 MAD 的参数
	// [(ai + b) mod p] mod N
	// - N: 散列表内部数组的大小
	// - p: 比 N 大的素数
	// - a，b 是从区间 [0, p-1] 任意选择的整数，并且 a > 0
	scale int // a
	shift int // b
	prime int // p

	bucketGet func(i int, k int) Item // hash 的子类需实现的方法
	bucketSet func(i int, k int, v string)
	bucketDel func(i int, k int) bool
	Iter      func() []Item
}

func (h *HashMapBase) hash(k int) int {
	return (k*h.scale + h.shift) % h.prime % h.capbility

}

// Resize hash 扩容
func (h *HashMapBase) Resize(cap int) {
	if cap < h.size {
		panic("resize to low")
	}
	tmp := h.Iter()
	table := make([]Item, cap, cap)
	// 移动数据
	h.capbility = cap
	h.table = table
	h.size = 0
	for _, i := range tmp {
		h.Set(i.Key(), i.Value())
	}
}

// Get HashMapBase 通过键返回值
func (h *HashMapBase) Get(k int) (string, bool) {
	i := h.hash(k)
	item := h.bucketGet(i, k)
	if item != nil {
		return item.Value(), true
	}
	return "", false
}

// Set HashMapBase 设置键值
func (h *HashMapBase) Set(k int, v string) {
	if h.size > h.capbility/2 {
		// 创建新的 table
		cap := h.capbility * 2
		h.Resize(cap)
	}
	i := h.hash(k)
	h.bucketSet(i, k, v)
	// h.size++  // h.size 是否增加取决于是否设置了新值
}

// Del HashMapBase 删除键
func (h *HashMapBase) Del(k int) bool {
	i := h.hash(k)
	return h.bucketDel(i, k)
	// h.size-- h.size 是否减-1取决于是否设置了新值
}

// Len HashMapBase 通过键返回值
func (h *HashMapBase) Len() int {
	return h.size
}

// ProbeHashMap 线性探测的散列表
type ProbeHashMap struct {
	HashMapBase
	_delete HashItem
}

// 线性探测
func (h *ProbeHashMap) findSlot(i, k int) (int, bool) {
	first := i
	for h.table[i] != nil {
		if h.table[i] == &h._delete {
			first = i
		}
		if h.table[i].Key() == k {
			return i, true
		}
		i = (i + 1) % h.capbility
	}
	first = i
	return first, false
}

// hash 的子类需实现的方法
func (h *ProbeHashMap) bucketGet(i int, k int) Item {
	// i 为 table 列表的索引
	// fmt.Println("ProbeHashMap bucketGet called")
	p, ok := h.findSlot(i, k)
	if ok {
		return h.table[p]
	}
	return nil
}

func (h *ProbeHashMap) bucketSet(i int, k int, v string) {
	// fmt.Println("ProbeHashMap bucketSet called")
	item := HashItem{key: k, value: v}
	p, ok := h.findSlot(i, k)
	h.table[p] = &item
	if !ok {
		h.size++
	}
}

func (h *ProbeHashMap) bucketDel(i int, k int) bool {
	// fmt.Println("ProbeHashMap bucketDel called")
	p, ok := h.findSlot(i, k)
	if ok {
		h.table[p] = &h._delete
		h.size--
	}
	return ok
}

// Iter 返回 hashmap 中有效的 Item
func (h *ProbeHashMap) Iter() []Item {
	c := make([]Item, 0, h.size)
	for _, item := range h.table {
		if item != nil && item != &h._delete {
			c = append(c, item)
		}
	}
	return c
}

// String 输出
func (h *ProbeHashMap) String() string {
	c := make([]int, 0, h.capbility)
	for _, item := range h.table {
		if item != nil && item != &h._delete {
			c = append(c, item.Key())
		} else {
			c = append(c, 0)
		}
	}
	return fmt.Sprintf("%v, %d,%d", c, h.size, h.capbility)
}

// NewProbeHashMap 创建 ProbeHashMap
func NewProbeHashMap(capbility int) *ProbeHashMap {
	s := &ProbeHashMap{
		HashMapBase: HashMapBase{
			capbility: capbility,
			scale:     782,
			shift:     38937,
			prime:     109345121,
			table:     make([]Item, capbility, capbility),
		},
	}
	// 通过赋值来实现抽象方法
	s.HashMapBase.Iter = s.Iter
	s.HashMapBase.bucketGet = s.bucketGet
	s.HashMapBase.bucketSet = s.bucketSet
	s.HashMapBase.bucketDel = s.bucketDel
	return s
}

// PointEqual 指针相等测试
func PointEqual() {
	var t *HashItem
	var s HashItem
	m := NewProbeHashMap(20)
	fmt.Printf("var *HashItem == nil: %t\n", t == nil)
	fmt.Printf("var *HashItem == _delete: %t\n", t == &m._delete)
	fmt.Printf("var HashItem == _delete: %t\n", &s == &m._delete)
}
