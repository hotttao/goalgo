package algo

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	maxLevel           = 25
	defaultProbability = 1 / math.E
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// SkipElement 存入跳表中的值必须满足的接口
type SkipElement interface {
	Key() string
	String() string
}

// SkipNode 调表表节点
type SkipNode struct {
	level int
	key   float64
	value SkipElement
	next  []*SkipNode
}

// SkipList 跳表
type SkipList struct {
	max    int // 调表的最高层高
	level  int // 调表的实际层高
	head   *SkipNode
	lenght int

	levelUpBaseNum int32 // 每多少个节点增加一层
}

// NewSkipList 创建 SkipList
func NewSkipList(level int, basenum int32) *SkipList {
	return &SkipList{
		max:            level,
		level:          1,
		lenght:         0,
		levelUpBaseNum: basenum,
		head: &SkipNode{level: level, key: math.MinInt32, value: nil,
			next: make([]*SkipNode, level, level)},
	}
}

func (s *SkipList) String() {
	for i := s.max - 1; i >= 0; i-- {
		fmt.Printf("head%d", i)
		c := "-"
		l0 := s.head
		cur := s.head
		for l0.next[0] != nil {
			// fmt.Println(l0.key)
			l0 = l0.next[0]
			if l0 != cur.next[i] {
				fmt.Printf("--")
			} else {
				cur = cur.next[i]
				fmt.Printf("%s%d", c, int(cur.key))
			}
		}
		fmt.Println()
	}
}

// CreateNode 创建 SkipNode
func (s *SkipList) CreateNode(key float64, value SkipElement) *SkipNode {
	level := 1
	for i := 1; i <= s.max; i++ {
		r := rand.Int31()
		fmt.Printf("random:  %d, %d\n", r, r%s.levelUpBaseNum)
		if r%s.levelUpBaseNum == 1 {
			level++
		}
	}
	fmt.Printf("New Node level: %d, Skip Level: %d\n", level, s.level)
	return &SkipNode{key: key, value: value, level: level,
		next: make([]*SkipNode, level, level)}
}

// Find 节点查找
func (s *SkipList) find(key float64) (pre []*SkipNode, ok bool) {
	return path, false
}

// Find 查找
func (s *SkipList) Find(key float64) (SkipElement, bool) {
	// 不能使用 find() 方法，因为无法知道是在第几层查找到的值
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil {
			k := cur.next[i].key
			v := cur.next[i].value
			if k == key {
				return v, true
			}
			if k > key {
				break
			}
			cur = cur.next[i]
		}
	}
	return nil, false
}

// Insert 插入
func (s *SkipList) Insert(key float64, value SkipElement) {

}

// Delete 删除节点
func (s *SkipList) Delete(key float64) (SkipElement, bool) {

}
