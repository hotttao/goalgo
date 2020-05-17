/*
Package algo skiplist
skiplist 跳表实现:
1. SkipNode 跳表中的节点:
	- next []*SkipNode 表示当前节点的层高，以及每一层指向的后续节点
2. SkipList
	- CreateNode 创建节点，决定新插入节点的层高
	- find: 值查找，返回每层遍历的最有一个节点，这些节点是每层插入的前驱节点
		- 从顶层往底层查找
		- 每一层从前往后查找
		- 每一层终止的条件是，下一个节点的 Key >= num(待查找值)，并将当前节点保存下来*/
package algo

import (
	"math"
	"math/rand"
)

// SkipNode 调表表节点
type SkipNode struct {
	key  int
	next []*SkipNode
}

// Skiplist 跳表
type Skiplist struct {
	level int // 跳表的层高
	head  *SkipNode
	size  int

	// 层高控制
	levelUpBaseNum int32 // 每多少个节点增加一层
}

// Constructor 创建 SkipList
func Constructor() Skiplist {
	level := 10
	return Skiplist{
		level:          level,
		levelUpBaseNum: 3,
		head: &SkipNode{key: math.MinInt64,
			next: make([]*SkipNode, level, level)},
	}
}

// CreateNode 创建 SkipNode
func (s *Skiplist) CreateNode(key int) *SkipNode {
	level := 1
	for i := 1; i < s.level; i++ { // i < s.level
		r := rand.Int31()
		if r%s.levelUpBaseNum == 1 {
			level++
		}
	}
	return &SkipNode{key: key,
		next: make([]*SkipNode, level, level)}
}

// Find 节点查找
func (s *Skiplist) find(num int) ([]*SkipNode, bool) {
	cur := s.head
	pre := make([]*SkipNode, s.level, s.level)
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil {
			key := cur.next[i].key
			if key >= num {
				pre[i] = cur
				break
			}
			cur = cur.next[i]
		}
		if cur.next[i] == nil {
			pre[i] = cur
		}
	}
	last := pre[0].next[0]
	ok := false
	if last != nil && last.key == num {
		ok = true
	}
	return pre, ok
}

// Search 查找值是否存在
func (s *Skiplist) Search(target int) bool {
	_, ok := s.find(target)
	return ok
}

// Add 添加
func (s *Skiplist) Add(num int) {
	pre, _ := s.find(num)
	node := s.CreateNode(num)
	for i := len(node.next) - 1; i >= 0; i-- {
		node.next[i] = pre[i].next[i]
		pre[i].next[i] = node
	}
}

// Erase 删除
func (s *Skiplist) Erase(num int) bool {
	pre, ok := s.find(num)
	if !ok {
		return ok
	}
	node := pre[0].next[0]
	for i := len(node.next) - 1; i >= 0; i-- {
		pre[i].next[i] = pre[i].next[i].next[i]
	}
	node.next = nil
	return ok
}
